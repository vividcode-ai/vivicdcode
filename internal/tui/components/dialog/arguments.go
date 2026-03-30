package dialog

import (
	"charm.land/bubbles/v2/key"
	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"fmt"
	uv "github.com/charmbracelet/ultraviolet"

	"github.com/vividcode-ai/vividcode/internal/tui/render"
	"github.com/vividcode-ai/vividcode/internal/tui/styles"
	"github.com/vividcode-ai/vividcode/internal/tui/theme"
	"github.com/vividcode-ai/vividcode/internal/tui/util"
)

type argumentsDialogKeyMap struct {
	Enter  key.Binding
	Escape key.Binding
}

// ShortHelp implements key.Map.
func (k argumentsDialogKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "confirm"),
		),
		key.NewBinding(
			key.WithKeys("esc"),
			key.WithHelp("esc", "cancel"),
		),
	}
}

// FullHelp implements key.Map.
func (k argumentsDialogKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{k.ShortHelp()}
}

// ShowMultiArgumentsDialogMsg is a message that is sent to show the multi-arguments dialog.
type ShowMultiArgumentsDialogMsg struct {
	CommandID string
	Content   string
	ArgNames  []string
}

// CloseMultiArgumentsDialogMsg is a message that is sent when the multi-arguments dialog is closed.
type CloseMultiArgumentsDialogMsg struct {
	Submit    bool
	CommandID string
	Content   string
	Args      map[string]string
}

// MultiArgumentsDialogCmp is a component that asks the user for multiple command arguments.
type MultiArgumentsDialogCmp struct {
	width, height int
	inputs        []textinput.Model
	focusIndex    int
	keys          argumentsDialogKeyMap
	commandID     string
	content       string
	argNames      []string
}

func applyInputStyles(ti *textinput.Model, t theme.Theme, focused bool) {
	s := ti.Styles()
	bg := lipgloss.Color(t.Background())
	fg := lipgloss.Color(t.Primary())

	s.Blurred.Placeholder = s.Blurred.Placeholder.Background(bg)
	s.Blurred.Prompt = s.Blurred.Prompt.Background(bg)
	s.Blurred.Text = s.Blurred.Text.Background(bg)

	if focused {
		s.Focused.Prompt = s.Focused.Prompt.Background(bg).Foreground(fg)
		s.Focused.Text = s.Focused.Text.Background(bg).Foreground(fg)
	} else {
		s.Focused.Prompt = s.Focused.Prompt.Background(bg)
		s.Focused.Text = s.Focused.Text.Background(bg)
	}

	ti.SetStyles(s)
}

// NewMultiArgumentsDialogCmp creates a new MultiArgumentsDialogCmp.
func NewMultiArgumentsDialogCmp(commandID, content string, argNames []string) *MultiArgumentsDialogCmp {
	t := theme.CurrentTheme()
	inputs := make([]textinput.Model, len(argNames))

	for i, name := range argNames {
		ti := textinput.New()
		ti.Placeholder = fmt.Sprintf("Enter value for %s...", name)
		ti.SetWidth(40)
		ti.Prompt = ""
		applyInputStyles(&ti, t, i == 0)

		if i == 0 {
			ti.Focus()
		} else {
			ti.Blur()
		}

		inputs[i] = ti
	}

	return &MultiArgumentsDialogCmp{
		inputs:     inputs,
		keys:       argumentsDialogKeyMap{},
		commandID:  commandID,
		content:    content,
		argNames:   argNames,
		focusIndex: 0,
	}
}

// Init implements tea.Model.
func (m *MultiArgumentsDialogCmp) Init() tea.Cmd {
	for i := range m.inputs {
		if i == 0 {
			m.inputs[i].Focus()
		} else {
			m.inputs[i].Blur()
		}
	}

	return textinput.Blink
}

// Update implements tea.Model.
func (m *MultiArgumentsDialogCmp) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	t := theme.CurrentTheme()

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, key.NewBinding(key.WithKeys("esc"))):
			return m, util.CmdHandler(CloseMultiArgumentsDialogMsg{
				Submit:    false,
				CommandID: m.commandID,
				Content:   m.content,
				Args:      nil,
			})
		case key.Matches(msg, key.NewBinding(key.WithKeys("enter"))):
			if m.focusIndex == len(m.inputs)-1 {
				args := make(map[string]string)
				for i, name := range m.argNames {
					args[name] = m.inputs[i].Value()
				}
				return m, util.CmdHandler(CloseMultiArgumentsDialogMsg{
					Submit:    true,
					CommandID: m.commandID,
					Content:   m.content,
					Args:      args,
				})
			}
			m.inputs[m.focusIndex].Blur()
			applyInputStyles(&m.inputs[m.focusIndex], t, false)
			m.focusIndex++
			m.inputs[m.focusIndex].Focus()
			applyInputStyles(&m.inputs[m.focusIndex], t, true)
		case key.Matches(msg, key.NewBinding(key.WithKeys("tab"))):
			m.inputs[m.focusIndex].Blur()
			applyInputStyles(&m.inputs[m.focusIndex], t, false)
			m.focusIndex = (m.focusIndex + 1) % len(m.inputs)
			m.inputs[m.focusIndex].Focus()
			applyInputStyles(&m.inputs[m.focusIndex], t, true)
		case key.Matches(msg, key.NewBinding(key.WithKeys("shift+tab"))):
			m.inputs[m.focusIndex].Blur()
			applyInputStyles(&m.inputs[m.focusIndex], t, false)
			m.focusIndex = (m.focusIndex - 1 + len(m.inputs)) % len(m.inputs)
			m.inputs[m.focusIndex].Focus()
			applyInputStyles(&m.inputs[m.focusIndex], t, true)
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}

	var cmd tea.Cmd
	m.inputs[m.focusIndex], cmd = m.inputs[m.focusIndex].Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

// View implements tea.Model.
func (m *MultiArgumentsDialogCmp) View() tea.View {
	t := theme.CurrentTheme()
	baseStyle := styles.BaseStyle()

	maxWidth := 60

	title := lipgloss.NewStyle().
		Foreground(lipgloss.Color(t.Primary())).
		Bold(true).
		Width(maxWidth).
		Padding(0, 1).
		Background(lipgloss.Color(t.Background())).
		Render("Command Arguments")

	explanation := lipgloss.NewStyle().
		Foreground(lipgloss.Color(t.Text())).
		Width(maxWidth).
		Padding(0, 1).
		Background(lipgloss.Color(t.Background())).
		Render("This command requires multiple arguments. Please enter values for each:")

	inputFields := make([]string, len(m.inputs))
	for i, input := range m.inputs {
		labelStyle := lipgloss.NewStyle().
			Width(maxWidth).
			Padding(1, 1, 0, 1).
			Background(lipgloss.Color(t.Background()))

		if i == m.focusIndex {
			labelStyle = labelStyle.Foreground(lipgloss.Color(t.Primary())).Bold(true)
		} else {
			labelStyle = labelStyle.Foreground(lipgloss.Color(t.TextMuted()))
		}

		label := labelStyle.Render(m.argNames[i] + ":")

		field := lipgloss.NewStyle().
			Foreground(lipgloss.Color(t.Text())).
			Width(maxWidth).
			Padding(0, 1).
			Background(lipgloss.Color(t.Background())).
			Render(input.View())

		inputFields[i] = lipgloss.JoinVertical(lipgloss.Left, label, field)
	}

	maxWidth = min(maxWidth, m.width-10)

	elements := []string{title, explanation}
	elements = append(elements, inputFields...)

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		elements...,
	)

	return tea.View{Content: baseStyle.Padding(1, 2).
		Border(lipgloss.RoundedBorder()).
		BorderBackground(lipgloss.Color(t.Background())).
		BorderForeground(lipgloss.Color(t.TextMuted())).
		Background(lipgloss.Color(t.Background())).
		Width(lipgloss.Width(content) + 4).
		Render(content)}
}

func (m *MultiArgumentsDialogCmp) Draw(scr uv.Screen, area uv.Rectangle) *tea.Cursor {
	view := m.View().Content
	render.DrawCenter(scr, area, view)
	return nil
}

// SetSize sets the size of the component.
func (m *MultiArgumentsDialogCmp) SetSize(width, height int) {
	m.width = width
	m.height = height
}

// Bindings implements layout.Bindings.
func (m *MultiArgumentsDialogCmp) Bindings() []key.Binding {
	return m.keys.ShortHelp()
}
