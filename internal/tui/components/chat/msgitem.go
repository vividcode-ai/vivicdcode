package chat

import (
	"context"
	"encoding/json"
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/vividcode-ai/vividcode/internal/config"
	"github.com/vividcode-ai/vividcode/internal/diff"
	"github.com/vividcode-ai/vividcode/internal/llm/agent"
	"github.com/vividcode-ai/vividcode/internal/llm/models"
	"github.com/vividcode-ai/vividcode/internal/llm/tools"
	"github.com/vividcode-ai/vividcode/internal/message"
	"github.com/vividcode-ai/vividcode/internal/tui/styles"
	"github.com/vividcode-ai/vividcode/internal/tui/theme"
)

const maxResultHeight = 10

type msgItem struct {
	msg             message.Message
	allMessages     []message.Message
	messagesService message.Service
	cache           map[int]*msgCache
}

type msgCache struct {
	content    string
	height     int
	multiItems []Item
}

func (m *msgItem) ID() string {
	return m.msg.ID
}

func (m *msgItem) getCache(width int) (string, int, []Item, bool) {
	if m.cache == nil {
		return "", 0, nil, false
	}
	if cr, ok := m.cache[width]; ok {
		if cr.multiItems != nil {
			return "", cr.height, cr.multiItems, true
		}
		return cr.content, cr.height, nil, true
	}
	return "", 0, nil, false
}

func (m *msgItem) setCacheMulti(width int, height int, items []Item) {
	if m.cache == nil {
		m.cache = make(map[int]*msgCache)
	}
	m.cache[width] = &msgCache{height: height, multiItems: items}
}

func (m *msgItem) setCache(width int, content string, height int) {
	if m.cache == nil {
		m.cache = make(map[int]*msgCache)
	}
	m.cache[width] = &msgCache{content: content, height: height}
}

func (m *msgItem) Render(width int) string {
	if content, _, items, ok := m.getCache(width); ok {
		if content != "" {
			return content
		}
		if len(items) > 0 {
			var parts []string
			for _, item := range items {
				parts = append(parts, item.Render(width))
			}
			return strings.Join(parts, "\n")
		}
		return ""
	}
	items := m.doRenderMulti(width)
	var rendered string
	if len(items) == 1 {
		rendered = items[0].Render(width)
	} else {
		var parts []string
		for _, item := range items {
			parts = append(parts, item.Render(width))
		}
		rendered = strings.Join(parts, "\n")
	}
	h := strings.Count(rendered, "\n") + 1
	m.setCache(width, rendered, h)
	return rendered
}

func (m *msgItem) Height(width int) int {
	if _, height, _, ok := m.getCache(width); ok {
		return height
	}
	items := m.doRenderMulti(width)
	var totalHeight int
	for _, item := range items {
		totalHeight += item.Height(width)
	}
	m.setCacheMulti(width, totalHeight, items)
	return totalHeight
}

func (m *msgItem) RenderMulti(width int) []Item {
	if _, _, items, ok := m.getCache(width); ok {
		return items
	}
	return m.doRenderMulti(width)
}

func (m *msgItem) doRender(width int) string {
	items := m.doRenderMulti(width)
	if len(items) == 1 {
		return items[0].Render(width)
	}
	var parts []string
	for _, item := range items {
		parts = append(parts, item.Render(width))
	}
	return strings.Join(parts, "\n")
}

func (m *msgItem) doRenderMulti(width int) []Item {
	t := theme.CurrentTheme()

	switch m.msg.Role {
	case message.User:
		return m.renderUserMsgMulti(width, t)
	case message.Assistant:
		return m.renderAssistantMsgMulti(width, t)
	}
	return nil
}

func (m *msgItem) renderUserMsgMulti(width int, t theme.Theme) []Item {
	items := []Item{}
	baseStyle := styles.BaseStyle()

	var styledAttachments []string
	attachmentStyles := baseStyle.
		MarginLeft(1).
		Background(lipgloss.Color(t.TextMuted())).
		Foreground(lipgloss.Color(t.Text()))
	for _, attachment := range m.msg.BinaryContent() {
		file := filepath.Base(attachment.Path)
		var filename string
		if len(file) > 10 {
			filename = fmt.Sprintf(" %s %s...", styles.DocumentIcon, file[0:7])
		} else {
			filename = fmt.Sprintf(" %s %s", styles.DocumentIcon, file)
		}
		styledAttachments = append(styledAttachments, attachmentStyles.Render(filename))
	}

	content := m.msg.Content().String()
	var info []string
	if len(styledAttachments) > 0 {
		attachmentContent := baseStyle.Width(width).Render(lipgloss.JoinHorizontal(lipgloss.Left, styledAttachments...))
		info = append(info, attachmentContent)
	}

	items = append(items, &simpleItem{
		id:       m.msg.ID + "-content",
		rendered: renderUserMsgUI(content, width, true, false, t.Text(), info...),
	})

	return items
}

func (m *msgItem) renderAssistantMsgMulti(width int, t theme.Theme) []Item {
	items := []Item{}
	baseStyle := styles.BaseStyle()

	content := m.msg.Content().String()
	thinking := m.msg.IsThinking()
	thinkingContent := m.msg.ReasoningContent().Thinking
	finished := m.msg.IsFinished()
	finishData := m.msg.FinishPart()
	info := []string{}

	if finished {
		switch finishData.Reason {
		case message.FinishReasonEndTurn:
			took := fmtTimeDiff(m.msg.CreatedAt, finishData.Time)
			info = append(info, baseStyle.
				Width(width-1).
				Foreground(lipgloss.Color(t.TextMuted())).
				Render(fmt.Sprintf(" %s (%s)", models.SupportedModels[m.msg.Model].Name, took)))
		case message.FinishReasonCanceled:
			info = append(info, baseStyle.
				Width(width-1).
				Foreground(lipgloss.Color(t.TextMuted())).
				Render(fmt.Sprintf(" %s (%s)", models.SupportedModels[m.msg.Model].Name, "canceled")))
		case message.FinishReasonError:
			info = append(info, baseStyle.
				Width(width-1).
				Foreground(lipgloss.Color(t.Error())).
				Render(fmt.Sprintf(" %s (%s)", models.SupportedModels[m.msg.Model].Name, "error")))
		case message.FinishReasonPermissionDenied:
			info = append(info, baseStyle.
				Width(width-1).
				Foreground(lipgloss.Color(t.Error())).
				Render(fmt.Sprintf(" %s (%s)", models.SupportedModels[m.msg.Model].Name, "permission denied")))
		}
	}

	if thinking && thinkingContent != "" {
		thinkingPrefix := "Thinking: "
		items = append(items, &simpleItem{
			id:       m.msg.ID + "-thinking",
			rendered: renderMsgUI(thinkingPrefix+thinkingContent, width, false, true, t.TextMuted()),
		})
	} else if content != "" || (finished && finishData.Reason == message.FinishReasonEndTurn) {
		if content == "" {
			content = "*Finished without output*"
		}
		items = append(items, &simpleItem{
			id:       m.msg.ID + "-content",
			rendered: renderMsgUI(content, width, false, false, t.Text(), info...),
		})
	}

	for _, toolCall := range m.msg.ToolCalls() {
		toolItem := m.renderToolMessageItem(toolCall, width)
		items = append(items, toolItem)
	}

	return items
}

func (m *msgItem) renderToolMessageItem(toolCall message.ToolCall, width int) *simpleItem {
	t := theme.CurrentTheme()
	baseStyle := styles.BaseStyle()

	toolNameText := toolName(toolCall.Name)

	if !toolCall.Finished {
		toolAction := getToolAction(toolCall.Name)
		content := baseStyle.
			Width(width - 1).
			BorderLeft(true).
			BorderStyle(lipgloss.ThickBorder()).
			PaddingLeft(1).
			BorderForeground(lipgloss.Color(t.TextMuted())).
			Foreground(lipgloss.Color(t.TextMuted())).
			Render(fmt.Sprintf("%s: %s", toolNameText, toolAction))
		return &simpleItem{
			id:       m.msg.ID + "-tool-" + toolCall.ID,
			rendered: content,
		}
	}

	params := renderToolParams(width-2-lipgloss.Width(toolNameText+": "), toolCall)

	responseContent := ""
	response := findToolResponse(toolCall.ID, m.allMessages)
	if response != nil {
		responseContent = renderToolResponse(toolCall, *response, width-2)
		responseContent = strings.TrimSuffix(responseContent, "\n")
	} else {
		responseContent = baseStyle.
			Italic(true).
			Width(width - 2).
			Foreground(lipgloss.Color(t.TextMuted())).
			Render("Waiting for response...")
	}

	style := baseStyle.
		Width(width - 1).
		BorderLeft(true).
		BorderStyle(lipgloss.ThickBorder()).
		PaddingLeft(1).
		BorderForeground(lipgloss.Color(t.TextMuted()))

	formattedParams := baseStyle.
		Width(width - 2 - lipgloss.Width(toolNameText+": ")).
		Foreground(lipgloss.Color(t.TextMuted())).
		Render(params)

	parts := []string{}
	parts = append(parts, fmt.Sprintf("%s: %s", toolNameText, formattedParams))

	if toolCall.Name == agent.AgentToolName && m.messagesService != nil {
		taskMessages, _ := m.messagesService.List(context.Background(), toolCall.ID)
		toolCalls := []message.ToolCall{}
		for _, v := range taskMessages {
			toolCalls = append(toolCalls, v.ToolCalls()...)
		}
		for _, call := range toolCalls {
			nestedItem := m.renderToolMessageItem(call, width-3)
			parts = append(parts, " └ "+strings.TrimPrefix(nestedItem.Render(width-3), " "))
		}
	}

	if responseContent != "" {
		parts = append(parts, responseContent)
	}

	content := style.Render(lipgloss.JoinVertical(lipgloss.Left, parts...))

	return &simpleItem{
		id:       m.msg.ID + "-tool-" + toolCall.ID,
		rendered: content,
	}
}

type simpleItem struct {
	id       string
	rendered string
}

func (s *simpleItem) ID() string {
	return s.id
}

func (s *simpleItem) Render(width int) string {
	return s.rendered
}

func (s *simpleItem) Height(width int) int {
	return strings.Count(s.rendered, "\n") + 1
}

func (s *simpleItem) RenderMulti(width int) []Item {
	return []Item{s}
}

func renderMsgUI(msg string, width int, isUser bool, isThinking bool, fgColor string, info ...string) string {
	t := theme.CurrentTheme()
	bgColor := t.BackgroundSecondary()

	style := styles.BaseStyle().
		Width(width-1).
		Background(lipgloss.Color(bgColor)).
		Padding(1, 1, 1, 2).
		MarginBottom(1).
		BorderLeft(true).
		BorderForeground(lipgloss.Color(t.Primary())).
		BorderStyle(lipgloss.ThickBorder())

	if isThinking {
		style = style.Foreground(lipgloss.Color(t.TextMuted()))
	} else {
		style = style.Foreground(lipgloss.Color(t.Text()))
	}

	if isUser {
		style = style.BorderForeground(lipgloss.Color(t.Secondary()))
	}

	parts := []string{
		styles.ForceReplaceForegroundAndBackgroundWithLipgloss(toMd(msg, width), lipgloss.Color(fgColor), lipgloss.Color(bgColor)),
	}

	parts[0] = strings.TrimSuffix(parts[0], "\n")
	if len(info) > 0 {
		parts = append(parts, info...)
	}
	content := lipgloss.JoinVertical(lipgloss.Left, parts...)

	return style.Render(content)
}
func renderUserMsgUI(msg string, width int, isUser bool, isThinking bool, fgColor string, info ...string) string {
	t := theme.CurrentTheme()
	bgColor := t.BackgroundSecondary()

	style := styles.BaseStyle().
		Width(width-1).
		Background(lipgloss.Color(bgColor)).
		Padding(1, 1, 1, 2).
		MarginBottom(1).
		BorderLeft(true).
		BorderForeground(lipgloss.Color(t.Secondary())).
		BorderStyle(lipgloss.ThickBorder())

	parts := []string{
		msg,
	}

	parts[0] = strings.TrimSuffix(parts[0], "\n")
	if len(info) > 0 {
		parts = append(parts, info...)
	}
	content := lipgloss.JoinVertical(lipgloss.Left, parts...)

	return style.Render(content)
}

func toMd(content string, width int) string {
	r := styles.GetMarkdownRenderer(width)
	rendered, _ := r.Render(content)
	return rendered
}

func findToolResponse(toolCallID string, futureMessages []message.Message) *message.ToolResult {
	for _, msg := range futureMessages {
		for _, result := range msg.ToolResults() {
			if result.ToolCallID == toolCallID {
				return &result
			}
		}
	}
	return nil
}

func toolName(name string) string {
	switch name {
	case agent.AgentToolName:
		return "Task"
	case tools.BashToolName:
		return "Bash"
	case tools.EditToolName:
		return "Edit"
	case tools.FetchToolName:
		return "Fetch"
	case tools.GlobToolName:
		return "Glob"
	case tools.GrepToolName:
		return "Grep"
	case tools.LSToolName:
		return "List"
	case tools.SourcegraphToolName:
		return "Sourcegraph"
	case tools.ViewToolName:
		return "View"
	case tools.WriteToolName:
		return "Write"
	case tools.PatchToolName:
		return "Patch"
	}
	return name
}

func getToolAction(name string) string {
	switch name {
	case agent.AgentToolName:
		return "Preparing prompt..."
	case tools.BashToolName:
		return "Building command..."
	case tools.EditToolName:
		return "Preparing edit..."
	case tools.FetchToolName:
		return "Writing fetch..."
	case tools.GlobToolName:
		return "Finding files..."
	case tools.GrepToolName:
		return "Searching content..."
	case tools.LSToolName:
		return "Listing directory..."
	case tools.SourcegraphToolName:
		return "Searching code..."
	case tools.ViewToolName:
		return "Reading file..."
	case tools.WriteToolName:
		return "Preparing write..."
	case tools.PatchToolName:
		return "Preparing patch..."
	}
	return "Working..."
}

func renderParams(paramsWidth int, params ...string) string {
	if len(params) == 0 {
		return ""
	}
	mainParam := params[0]
	if len(mainParam) > paramsWidth {
		mainParam = mainParam[:paramsWidth-3] + "..."
	}

	if len(params) == 1 {
		return mainParam
	}
	otherParams := params[1:]
	if len(otherParams)%2 != 0 {
		otherParams = append(otherParams, "")
	}
	parts := make([]string, 0, len(otherParams)/2)
	for i := 0; i < len(otherParams); i += 2 {
		key := otherParams[i]
		value := otherParams[i+1]
		if value == "" {
			continue
		}
		parts = append(parts, fmt.Sprintf("%s=%s", key, value))
	}

	partsRendered := strings.Join(parts, ", ")
	remainingWidth := paramsWidth - lipgloss.Width(partsRendered) - 5
	if remainingWidth < 30 {
		return mainParam
	}

	if len(parts) > 0 {
		mainParam = fmt.Sprintf("%s (%s)", mainParam, strings.Join(parts, ", "))
	}

	return ansi.Truncate(mainParam, paramsWidth, "...")
}

func removeWorkingDirPrefix(path string) string {
	wd := config.WorkingDirectory()
	if strings.HasPrefix(path, wd) {
		path = strings.TrimPrefix(path, wd)
	}
	if strings.HasPrefix(path, "/") {
		path = strings.TrimPrefix(path, "/")
	}
	if strings.HasPrefix(path, "./") {
		path = strings.TrimPrefix(path, "./")
	}
	if strings.HasPrefix(path, "../") {
		path = strings.TrimPrefix(path, "../")
	}
	return path
}

func renderToolParams(paramWidth int, toolCall message.ToolCall) string {
	params := ""
	switch toolCall.Name {
	case agent.AgentToolName:
		var p agent.AgentParams
		json.Unmarshal([]byte(toolCall.Input), &p)
		prompt := strings.ReplaceAll(p.Prompt, "\n", " ")
		return renderParams(paramWidth, prompt)
	case tools.BashToolName:
		var p tools.BashParams
		json.Unmarshal([]byte(toolCall.Input), &p)
		command := strings.ReplaceAll(p.Command, "\n", " ")
		return renderParams(paramWidth, command)
	case tools.EditToolName:
		var p tools.EditParams
		json.Unmarshal([]byte(toolCall.Input), &p)
		filePath := removeWorkingDirPrefix(p.FilePath)
		return renderParams(paramWidth, filePath)
	case tools.FetchToolName:
		var p tools.FetchParams
		json.Unmarshal([]byte(toolCall.Input), &p)
		toolParams := []string{p.URL}
		if p.Format != "" {
			toolParams = append(toolParams, "format", p.Format)
		}
		if p.Timeout != 0 {
			toolParams = append(toolParams, "timeout", (time.Duration(p.Timeout) * time.Second).String())
		}
		return renderParams(paramWidth, toolParams...)
	case tools.GlobToolName:
		var p tools.GlobParams
		json.Unmarshal([]byte(toolCall.Input), &p)
		toolParams := []string{p.Pattern}
		if p.Path != "" {
			toolParams = append(toolParams, "path", p.Path)
		}
		return renderParams(paramWidth, toolParams...)
	case tools.GrepToolName:
		var p tools.GrepParams
		json.Unmarshal([]byte(toolCall.Input), &p)
		toolParams := []string{p.Pattern}
		if p.Path != "" {
			toolParams = append(toolParams, "path", p.Path)
		}
		if p.Include != "" {
			toolParams = append(toolParams, "include", p.Include)
		}
		if p.LiteralText {
			toolParams = append(toolParams, "literal", "true")
		}
		return renderParams(paramWidth, toolParams...)
	case tools.LSToolName:
		var p tools.LSParams
		json.Unmarshal([]byte(toolCall.Input), &p)
		path := p.Path
		if path == "" {
			path = "."
		}
		return renderParams(paramWidth, path)
	case tools.SourcegraphToolName:
		var p tools.SourcegraphParams
		json.Unmarshal([]byte(toolCall.Input), &p)
		return renderParams(paramWidth, p.Query)
	case tools.ViewToolName:
		var p tools.ViewParams
		json.Unmarshal([]byte(toolCall.Input), &p)
		filePath := removeWorkingDirPrefix(p.FilePath)
		toolParams := []string{filePath}
		if p.Limit != 0 {
			toolParams = append(toolParams, "limit", fmt.Sprintf("%d", p.Limit))
		}
		if p.Offset != 0 {
			toolParams = append(toolParams, "offset", fmt.Sprintf("%d", p.Offset))
		}
		return renderParams(paramWidth, toolParams...)
	case tools.WriteToolName:
		var p tools.WriteParams
		json.Unmarshal([]byte(toolCall.Input), &p)
		filePath := removeWorkingDirPrefix(p.FilePath)
		return renderParams(paramWidth, filePath)
	default:
		input := strings.ReplaceAll(toolCall.Input, "\n", " ")
		params = renderParams(paramWidth, input)
	}
	return params
}

func truncateHeight(content string, height int) string {
	lines := strings.Split(content, "\n")
	if len(lines) > height {
		return strings.Join(lines[:height], "\n")
	}
	return content
}

func renderToolResponse(toolCall message.ToolCall, response message.ToolResult, width int) string {
	t := theme.CurrentTheme()
	baseStyle := styles.BaseStyle()

	if response.IsError {
		errContent := fmt.Sprintf("Error: %s", strings.ReplaceAll(response.Content, "\n", " "))
		errContent = ansi.Truncate(errContent, width-1, "...")
		return baseStyle.
			Width(width).
			Foreground(lipgloss.Color(t.Error())).
			Render(errContent)
	}

	resultContent := truncateHeight(response.Content, maxResultHeight)
	switch toolCall.Name {
	case agent.AgentToolName:
		return styles.ForceReplaceBackgroundWithLipgloss(
			toMd(resultContent, width),
			lipgloss.Color(t.Background()),
		)
	case tools.BashToolName:
		resultContent = fmt.Sprintf("```bash\n%s\n```", resultContent)
		return styles.ForceReplaceBackgroundWithLipgloss(
			toMd(resultContent, width),
			lipgloss.Color(t.Background()),
		)
	case tools.EditToolName:
		metadata := tools.EditResponseMetadata{}
		json.Unmarshal([]byte(response.Metadata), &metadata)
		truncDiff := truncateHeight(metadata.Diff, maxResultHeight)
		formattedDiff, _ := diff.FormatDiff(truncDiff, diff.WithTotalWidth(width))
		return formattedDiff
	case tools.FetchToolName:
		var params tools.FetchParams
		json.Unmarshal([]byte(toolCall.Input), &params)
		mdFormat := "markdown"
		switch params.Format {
		case "text":
			mdFormat = "text"
		case "html":
			mdFormat = "html"
		}
		resultContent = fmt.Sprintf("```%s\n%s\n```", mdFormat, resultContent)
		return styles.ForceReplaceBackgroundWithLipgloss(
			toMd(resultContent, width),
			lipgloss.Color(t.Background()),
		)
	case tools.GlobToolName:
		return baseStyle.Width(width).Foreground(lipgloss.Color(t.TextMuted())).Render(resultContent)
	case tools.GrepToolName:
		return baseStyle.Width(width).Foreground(lipgloss.Color(t.TextMuted())).Render(resultContent)
	case tools.LSToolName:
		return baseStyle.Width(width).Foreground(lipgloss.Color(t.TextMuted())).Render(resultContent)
	case tools.SourcegraphToolName:
		return baseStyle.Width(width).Foreground(lipgloss.Color(t.TextMuted())).Render(resultContent)
	case tools.ViewToolName:
		metadata := tools.ViewResponseMetadata{}
		json.Unmarshal([]byte(response.Metadata), &metadata)
		ext := filepath.Ext(metadata.FilePath)
		if ext != "" {
			ext = strings.ToLower(ext[1:])
		}
		resultContent = fmt.Sprintf("```%s\n%s\n```", ext, truncateHeight(metadata.Content, maxResultHeight))
		return styles.ForceReplaceBackgroundWithLipgloss(
			toMd(resultContent, width),
			lipgloss.Color(t.Background()),
		)
	case tools.WriteToolName:
		params := tools.WriteParams{}
		json.Unmarshal([]byte(toolCall.Input), &params)
		ext := filepath.Ext(params.FilePath)
		if ext != "" {
			ext = strings.ToLower(ext[1:])
		}
		resultContent = fmt.Sprintf("```%s\n%s\n```", ext, truncateHeight(params.Content, maxResultHeight))
		return styles.ForceReplaceBackgroundWithLipgloss(
			toMd(resultContent, width),
			lipgloss.Color(t.Background()),
		)
	default:
		resultContent = fmt.Sprintf("```text\n%s\n```", resultContent)
		return styles.ForceReplaceBackgroundWithLipgloss(
			toMd(resultContent, width),
			lipgloss.Color(t.Background()),
		)
	}
}

func messagesToItems(messages []message.Message, messagesService message.Service) []Item {
	items := make([]Item, 0, len(messages))
	for _, msg := range messages {
		items = append(items, newMsgItem(msg, messages, messagesService))
	}
	return items
}

func newMsgItem(msg message.Message, allMessages []message.Message, messagesService message.Service) Item {
	return &msgItem{
		msg:             msg,
		allMessages:     allMessages,
		messagesService: messagesService,
		cache:           make(map[int]*msgCache),
	}
}

func fmtTimeDiff(start, end int64) string {
	diffSeconds := float64(end-start) / 1000.0
	if diffSeconds < 1 {
		return fmt.Sprintf("%dms", int(diffSeconds*1000))
	}
	if diffSeconds < 60 {
		return fmt.Sprintf("%.1fs", diffSeconds)
	}
	return fmt.Sprintf("%.1fm", diffSeconds/60)
}

func toMarkdown(content string, focused bool, width int) string {
	r := styles.GetMarkdownRenderer(width)
	rendered, _ := r.Render(content)
	return rendered
}
