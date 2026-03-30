package chat

import (
	"fmt"
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/vividcode-ai/vividcode/internal/llm/models"
	"github.com/vividcode-ai/vividcode/internal/message"
	"github.com/vividcode-ai/vividcode/internal/tui/styles"
	"github.com/vividcode-ai/vividcode/internal/tui/theme"
)

type msgItem struct {
	msg   message.Message
	cache map[int]*msgCache
}

type msgCache struct {
	content string
	height  int
}

func (m *msgItem) ID() string {
	return m.msg.ID
}

func (m *msgItem) getCache(width int) (string, int, bool) {
	if m.cache == nil {
		return "", 0, false
	}
	if cr, ok := m.cache[width]; ok {
		return cr.content, cr.height, true
	}
	return "", 0, false
}

func (m *msgItem) setCache(content string, width int, height int) {
	if m.cache == nil {
		m.cache = make(map[int]*msgCache)
	}
	m.cache[width] = &msgCache{content: content, height: height}
}

func (m *msgItem) Render(width int) string {
	if content, _, ok := m.getCache(width); ok {
		return content
	}
	rendered := m.doRender(width)
	height := strings.Count(rendered, "\n") + 1
	m.setCache(rendered, width, height)
	return rendered
}

func (m *msgItem) Height(width int) int {
	if _, height, ok := m.getCache(width); ok {
		return height
	}
	rendered := m.doRender(width)
	height := strings.Count(rendered, "\n") + 1
	m.setCache(rendered, width, height)
	return height
}

func (m *msgItem) doRender(width int) string {
	t := theme.CurrentTheme()
	baseStyle := styles.BaseStyle()

	switch m.msg.Role {
	case message.User:
		return m.renderUserMsg(width, baseStyle, t)
	case message.Assistant:
		return m.renderAssistantMsg(width, baseStyle, t)
	}
	return ""
}

func (m *msgItem) renderUserMsg(width int, baseStyle lipgloss.Style, t theme.Theme) string {
	var styledAttachments []string
	attachmentStyles := baseStyle.
		MarginLeft(1).
		Background(lipgloss.Color(t.TextMuted())).
		Foreground(lipgloss.Color(t.Text()))
	for _, attachment := range m.msg.BinaryContent() {
		filename := fmt.Sprintf(" %s %s", styles.DocumentIcon, attachment.Path)
		styledAttachments = append(styledAttachments, attachmentStyles.Render(filename))
	}
	content := m.msg.Content().String()
	var info []string
	if len(styledAttachments) > 0 {
		attachmentContent := baseStyle.Width(width).Render(lipgloss.JoinHorizontal(lipgloss.Left, styledAttachments...))
		info = append(info, attachmentContent)
	}
	return renderMsgUI(content, width, true, info...)
}

func (m *msgItem) renderAssistantMsg(width int, baseStyle lipgloss.Style, t theme.Theme) string {
	content := m.msg.Content().String()
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

	if content != "" || (finished && finishData.Reason == message.FinishReasonEndTurn) {
		if content == "" {
			content = "*Finished without output*"
		}
		return renderMsgUI(content, width, false, info...)
	}
	return ""
}

func renderMsgUI(msg string, width int, isUser bool, info ...string) string {
	t := theme.CurrentTheme()

	style := styles.BaseStyle().
		Width(width - 1).
		BorderLeft(true).
		Foreground(lipgloss.Color(t.TextMuted())).
		BorderForeground(lipgloss.Color(t.Primary())).
		BorderStyle(lipgloss.ThickBorder())

	if isUser {
		style = style.BorderForeground(lipgloss.Color(t.Secondary()))
	}

	parts := []string{
		styles.ForceReplaceBackgroundWithLipgloss(toMd(msg, width), lipgloss.Color(t.Background())),
	}
	parts[0] = strings.TrimSuffix(parts[0], "\n")
	if len(info) > 0 {
		parts = append(parts, info...)
	}

	return style.Render(lipgloss.JoinVertical(lipgloss.Left, parts...))
}

func toMd(content string, width int) string {
	r := styles.GetMarkdownRenderer(width)
	rendered, _ := r.Render(content)
	return rendered
}

func messagesToItems(messages []message.Message) []Item {
	items := make([]Item, 0, len(messages))
	for _, msg := range messages {
		items = append(items, newMsgItem(msg))
	}
	return items
}

func newMsgItem(msg message.Message) Item {
	return &msgItem{msg: msg, cache: make(map[int]*msgCache)}
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
