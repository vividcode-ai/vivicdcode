package chat

import (
	"context"
	"fmt"
	"strings"

	"charm.land/bubbles/v2/key"
	"charm.land/bubbles/v2/spinner"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	uv "github.com/charmbracelet/ultraviolet"
	"github.com/vividcode-ai/vividcode/internal/app"
	"github.com/vividcode-ai/vividcode/internal/logging"
	"github.com/vividcode-ai/vividcode/internal/message"
	"github.com/vividcode-ai/vividcode/internal/pubsub"
	"github.com/vividcode-ai/vividcode/internal/session"
	"github.com/vividcode-ai/vividcode/internal/tui/components/dialog"
	"github.com/vividcode-ai/vividcode/internal/tui/styles"
	"github.com/vividcode-ai/vividcode/internal/tui/theme"
	"github.com/vividcode-ai/vividcode/internal/tui/util"
)

type messagesCmp struct {
	app           *app.App
	width, height int
	vlist         *VirtualList
	session       session.Session
	messages      []message.Message
	currentMsgID  string
	spinner       spinner.Model
	rendering     bool
	follow        bool
	idIndexMap    map[string]int
}

type renderFinishedMsg struct{}

type MessageKeys struct {
	PageDown     key.Binding
	PageUp       key.Binding
	HalfPageUp   key.Binding
	HalfPageDown key.Binding
	ScrollUp     key.Binding
	ScrollDown   key.Binding
}

var messageKeys = MessageKeys{
	PageDown: key.NewBinding(
		key.WithKeys("pgdown", "f"),
		key.WithHelp("f/pgdn", "page down"),
	),
	PageUp: key.NewBinding(
		key.WithKeys("pgup", "b"),
		key.WithHelp("b/pgup", "page up"),
	),
	HalfPageUp: key.NewBinding(
		key.WithKeys("ctrl+u"),
		key.WithHelp("ctrl+u", "½ page up"),
	),
	HalfPageDown: key.NewBinding(
		key.WithKeys("ctrl+d"),
		key.WithHelp("ctrl+d", "½ page down"),
	),
	ScrollUp: key.NewBinding(
		key.WithKeys("up", "k"),
		key.WithHelp("↑/k", "scroll up"),
	),
	ScrollDown: key.NewBinding(
		key.WithKeys("down", "j"),
		key.WithHelp("↓/j", "scroll down"),
	),
}

func (m *messagesCmp) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m *messagesCmp) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case dialog.ThemeChangedMsg:
		m.rerender()
		return m, nil

	case SessionSelectedMsg:
		if msg.ID != m.session.ID {
			cmd := m.SetSession(msg)
			return m, cmd
		}
		m.follow = true
		return m, nil

	case SessionClearedMsg:
		m.session = session.Session{}
		m.messages = make([]message.Message, 0)
		m.currentMsgID = ""
		m.rendering = false
		m.idIndexMap = make(map[string]int)
		m.vlist.SetItems(nil)
		m.follow = false
		return m, nil

	case tea.KeyMsg:
		if key.Matches(msg, messageKeys.PageDown) {
			m.ScrollBy(m.height)
		} else if key.Matches(msg, messageKeys.PageUp) {
			m.ScrollBy(-m.height)
		} else if key.Matches(msg, messageKeys.HalfPageUp) {
			m.ScrollBy(-m.height / 2)
		} else if key.Matches(msg, messageKeys.HalfPageDown) {
			m.ScrollBy(m.height / 2)
		} else if key.Matches(msg, messageKeys.ScrollUp) {
			m.ScrollBy(-1)
		} else if key.Matches(msg, messageKeys.ScrollDown) {
			m.ScrollBy(1)
		}

	case renderFinishedMsg:
		m.rendering = false
		m.vlist.ScrollToBottom()

	case ForceFollowMsg:
		logging.Debug("ForceFollowMsg received, setting follow = true")
		m.follow = true
		return m, nil

	case pubsub.Event[session.Session]:
		if msg.Type == pubsub.UpdatedEvent && msg.Payload.ID == m.session.ID {
			m.session = msg.Payload
			if m.session.SummaryMessageID == m.currentMsgID {
				m.rerender()
			}
		}

	case pubsub.Event[message.Message]:
		if msg.Payload.SessionID != m.session.ID {
			return m, nil
		}
		logging.Debug("message event received", "type", msg.Type, "messageID", msg.Payload.ID)
		if msg.Type == pubsub.CreatedEvent {
			if _, exists := m.idIndexMap[msg.Payload.ID]; !exists {
				logging.Debug("new message, follow", "follow", m.follow)
				m.messages = append(m.messages, msg.Payload)
				m.idIndexMap[msg.Payload.ID] = len(m.messages) - 1
				m.currentMsgID = msg.Payload.ID
				m.vlist.AppendItems(newMsgItem(msg.Payload))
				if m.follow {
					logging.Debug("scrolling to bottom")
					m.vlist.ScrollToBottom()
				}
			}
		} else if msg.Type == pubsub.UpdatedEvent {
			if idx, exists := m.idIndexMap[msg.Payload.ID]; exists {
				m.messages[idx] = msg.Payload
				m.vlist.UpdateItem(idx, newMsgItem(msg.Payload))
				if m.follow {
					m.vlist.ScrollToBottom()
				}
			}
		}
	}
	return m, nil
}

func (m *messagesCmp) IsAgentWorking() bool {
	return m.app.CoderAgent.IsSessionBusy(m.session.ID)
}

func (m *messagesCmp) updateList() {
	items := messagesToItems(m.messages)
	m.vlist.SetItems(items)
}

func (m *messagesCmp) View() tea.View {
	baseStyle := styles.BaseStyle()

	var content string
	if m.rendering {
		content = baseStyle.
			Width(m.width).
			Render(
				lipgloss.JoinVertical(
					lipgloss.Top,
					"Loading...",
					m.working(),
					m.help(),
				),
			)
	} else if len(m.messages) == 0 {
		innerContent := baseStyle.
			Width(m.width).
			Height(m.height - 1).
			Render(
				m.initialScreen(),
			)
		content = baseStyle.
			Width(m.width).
			Render(
				lipgloss.JoinVertical(
					lipgloss.Top,
					innerContent,
					"",
					m.help(),
				),
			)
	} else {
		listContent := m.vlist.Render()
		content = baseStyle.
			Width(m.width).
			Render(
				lipgloss.JoinVertical(
					lipgloss.Top,
					listContent,
					m.working(),
					m.help(),
				),
			)
	}

	actualHeight := lipgloss.Height(content)
	if actualHeight < m.height {
		fillLine := strings.Repeat(" ", m.width)
		for i := actualHeight; i < m.height; i++ {
			content += "\n" + fillLine
		}
	}

	return tea.View{Content: content}
}

func (m *messagesCmp) Draw(scr uv.Screen, area uv.Rectangle) {
	content := m.View().Content
	uv.NewStyledString(content).Draw(scr, area)
}

func hasToolsWithoutResponse(messages []message.Message) bool {
	toolCalls := make([]message.ToolCall, 0)
	toolResults := make([]message.ToolResult, 0)
	for _, m := range messages {
		toolCalls = append(toolCalls, m.ToolCalls()...)
		toolResults = append(toolResults, m.ToolResults()...)
	}
	for _, v := range toolCalls {
		found := false
		for _, r := range toolResults {
			if v.ID == r.ToolCallID {
				found = true
				break
			}
		}
		if !found && v.Finished {
			return true
		}
	}
	return false
}

func hasUnfinishedToolCalls(messages []message.Message) bool {
	toolCalls := make([]message.ToolCall, 0)
	for _, m := range messages {
		toolCalls = append(toolCalls, m.ToolCalls()...)
	}
	for _, v := range toolCalls {
		if !v.Finished {
			return true
		}
	}
	return false
}

func (m *messagesCmp) working() string {
	text := ""
	if m.IsAgentWorking() && len(m.messages) > 0 {
		t := theme.CurrentTheme()
		baseStyle := styles.BaseStyle()

		task := "Thinking..."
		lastMessage := m.messages[len(m.messages)-1]
		if hasToolsWithoutResponse(m.messages) {
			task = "Waiting for tool response..."
		} else if hasUnfinishedToolCalls(m.messages) {
			task = "Building tool call..."
		} else if !lastMessage.IsFinished() {
			task = "Generating..."
		}
		if task != "" {
			text += baseStyle.
				Width(m.width).
				Foreground(lipgloss.Color(t.Primary())).
				Bold(true).
				Render(fmt.Sprintf("%s %s ", m.spinner.View(), task))
		}
	}
	return text
}

func (m *messagesCmp) help() string {
	t := theme.CurrentTheme()
	baseStyle := styles.BaseStyle()

	text := ""

	if m.app.CoderAgent.IsBusy() {
		text += lipgloss.JoinHorizontal(
			lipgloss.Left,
			baseStyle.Foreground(lipgloss.Color(t.TextMuted())).Bold(true).Render("press "),
			baseStyle.Foreground(lipgloss.Color(t.Text())).Bold(true).Render("esc"),
			baseStyle.Foreground(lipgloss.Color(t.TextMuted())).Bold(true).Render(" to exit cancel"),
		)
	} else {
		text += lipgloss.JoinHorizontal(
			lipgloss.Left,
			baseStyle.Foreground(lipgloss.Color(t.TextMuted())).Bold(true).Render("press "),
			baseStyle.Foreground(lipgloss.Color(t.Text())).Bold(true).Render("enter"),
			baseStyle.Foreground(lipgloss.Color(t.TextMuted())).Bold(true).Render(" to send the message,"),
			baseStyle.Foreground(lipgloss.Color(t.TextMuted())).Bold(true).Render(" write"),
			baseStyle.Foreground(lipgloss.Color(t.Text())).Bold(true).Render(" \\"),
			baseStyle.Foreground(lipgloss.Color(t.TextMuted())).Bold(true).Render(" and enter to add a new line"),
		)
	}
	return baseStyle.
		Width(m.width).
		Render(text)
}

func (m *messagesCmp) initialScreen() string {
	baseStyle := styles.BaseStyle()

	return baseStyle.Width(m.width).Render(
		lipgloss.JoinVertical(
			lipgloss.Top,
			Header(m.width),
			"",
			lspsConfigured(m.width),
		),
	)
}

func (m *messagesCmp) rerender() {
	m.vlist.SetItems(messagesToItems(m.messages))
}

func (m *messagesCmp) SetSize(width, height int) tea.Cmd {
	if m.width == width && m.height == height {
		return nil
	}
	m.width = width
	m.height = height
	m.vlist.SetSize(width, height-2)
	m.rerender()
	return nil
}

func (m *messagesCmp) GetSize() (int, int) {
	return m.width, m.height
}

func (m *messagesCmp) SetSession(s session.Session) tea.Cmd {
	if m.session.ID == s.ID {
		return nil
	}
	m.session = s
	messages, err := m.app.Messages.List(context.Background(), s.ID)
	if err != nil {
		return func() tea.Msg { return util.ReportError(err) }
	}
	m.messages = messages
	m.idIndexMap = make(map[string]int)
	for i, msg := range messages {
		m.idIndexMap[msg.ID] = i
	}
	if len(m.messages) > 0 {
		m.currentMsgID = m.messages[len(m.messages)-1].ID
	}
	m.rendering = true
	m.follow = true
	m.updateList()
	return func() tea.Msg { return renderFinishedMsg{} }
}

func (m *messagesCmp) ScrollBy(lines int) tea.Cmd {
	cmd := m.vlist.ScrollBy(lines)
	if lines > 0 {
		m.follow = m.vlist.AtBottom()
	} else {
		m.follow = false
	}
	return cmd
}

func (m *messagesCmp) BindingKeys() []key.Binding {
	return []key.Binding{
		messageKeys.PageDown,
		messageKeys.PageUp,
		messageKeys.HalfPageUp,
		messageKeys.HalfPageDown,
		messageKeys.ScrollUp,
		messageKeys.ScrollDown,
	}
}

func NewMessagesCmp(app *app.App) tea.Model {
	s := spinner.New()
	s.Spinner = spinner.Pulse
	vlist := NewVirtualList()
	return &messagesCmp{
		app:        app,
		spinner:    s,
		vlist:      vlist,
		idIndexMap: make(map[string]int),
		follow:     true,
	}
}
