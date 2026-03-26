package dialog

import (
	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	uv "github.com/charmbracelet/ultraviolet"

	"github.com/vividcode-ai/vividcode/internal/tui/render"
	"github.com/vividcode-ai/vividcode/internal/tui/styles"
	"github.com/vividcode-ai/vividcode/internal/tui/theme"
	"github.com/vividcode-ai/vividcode/internal/tui/util"
)

// InitDialogCmp is a component that asks the user if they want to initialize the project.
type InitDialogCmp struct {
	width, height int
	selected      int
	keys          initDialogKeyMap
}

// NewInitDialogCmp creates a new InitDialogCmp.
func NewInitDialogCmp() InitDialogCmp {
	return InitDialogCmp{
		selected: 0,
		keys:     initDialogKeyMap{},
	}
}

type initDialogKeyMap struct {
	Tab    key.Binding
	Left   key.Binding
	Right  key.Binding
	Enter  key.Binding
	Escape key.Binding
	Y      key.Binding
	N      key.Binding
}

// ShortHelp implements key.Map.
func (k initDialogKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		key.NewBinding(
			key.WithKeys("tab", "left", "right"),
			key.WithHelp("tab/←/→", "toggle selection"),
		),
		key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "confirm"),
		),
		key.NewBinding(
			key.WithKeys("esc", "q"),
			key.WithHelp("esc/q", "cancel"),
		),
		key.NewBinding(
			key.WithKeys("y", "n"),
			key.WithHelp("y/n", "yes/no"),
		),
	}
}

// FullHelp implements key.Map.
func (k initDialogKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{k.ShortHelp()}
}

// Init implements tea.Model.
func (m InitDialogCmp) Init() tea.Cmd {
	return nil
}

// Update implements tea.Model.
func (m InitDialogCmp) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, key.NewBinding(key.WithKeys("esc"))):
			return m, util.CmdHandler(CloseInitDialogMsg{Initialize: false})
		case key.Matches(msg, key.NewBinding(key.WithKeys("tab", "left", "right", "h", "l"))):
			m.selected = (m.selected + 1) % 2
			return m, nil
		case key.Matches(msg, key.NewBinding(key.WithKeys("enter"))):
			return m, util.CmdHandler(CloseInitDialogMsg{Initialize: m.selected == 0})
		case key.Matches(msg, key.NewBinding(key.WithKeys("y"))):
			return m, util.CmdHandler(CloseInitDialogMsg{Initialize: true})
		case key.Matches(msg, key.NewBinding(key.WithKeys("n"))):
			return m, util.CmdHandler(CloseInitDialogMsg{Initialize: false})
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}
	return m, nil
}

// View implements tea.Model.
func (m InitDialogCmp) View() tea.View {
	t := theme.CurrentTheme()
	baseStyle := styles.BaseStyle()

	// Calculate width needed for content
	maxWidth := 60 // Width for explanation text

	title := baseStyle.
		Foreground(lipgloss.Color(t.Primary())).
		Bold(true).
		Width(maxWidth).
		Padding(0, 1).
		Render("Initialize Project")

	explanation := baseStyle.
		Foreground(lipgloss.Color(t.Text())).
		Width(maxWidth).
		Padding(0, 1).
		Render("Initialization generates a new VividCode.md file that contains information about your codebase, this file serves as memory for each project, you can freely add to it to help the agents be better at their job.")

	question := baseStyle.
		Foreground(lipgloss.Color(t.Text())).
		Width(maxWidth).
		Padding(1, 1).
		Render("Would you like to initialize this project?")

	maxWidth = min(maxWidth, m.width-10)
	yesStyle := baseStyle
	noStyle := baseStyle

	if m.selected == 0 {
		yesStyle = yesStyle.
			Background(lipgloss.Color(t.Primary())).
			Foreground(lipgloss.Color(t.Background())).
			Bold(true)
		noStyle = noStyle.
			Background(lipgloss.Color(t.Background())).
			Foreground(lipgloss.Color(t.Primary()))
	} else {
		noStyle = noStyle.
			Background(lipgloss.Color(t.Primary())).
			Foreground(lipgloss.Color(t.Background())).
			Bold(true)
		yesStyle = yesStyle.
			Background(lipgloss.Color(t.Background())).
			Foreground(lipgloss.Color(t.Primary()))
	}

	yes := yesStyle.Padding(0, 3).Render("Yes")
	no := noStyle.Padding(0, 3).Render("No")

	buttons := lipgloss.JoinHorizontal(lipgloss.Center, yes, baseStyle.Render("  "), no)
	buttons = baseStyle.
		Width(maxWidth).
		Padding(1, 0).
		Render(buttons)

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		baseStyle.Width(maxWidth).Render(""),
		explanation,
		question,
		buttons,
		baseStyle.Width(maxWidth).Render(""),
	)

	return tea.View{Content: baseStyle.Padding(1, 2).
		Border(lipgloss.RoundedBorder()).
		BorderBackground(lipgloss.Color(t.Background())).
		BorderForeground(lipgloss.Color(t.TextMuted())).
		Width(lipgloss.Width(content) + 4).
		Render(content)}
}

func (m InitDialogCmp) Draw(scr uv.Screen, area uv.Rectangle) *tea.Cursor {
	view := m.View().Content
	render.DrawCenter(scr, area, view)
	return nil
}

// SetSize sets the size of the component.
func (m *InitDialogCmp) SetSize(width, height int) {
	m.width = width
	m.height = height
}

// Bindings implements layout.Bindings.
func (m InitDialogCmp) Bindings() []key.Binding {
	return m.keys.ShortHelp()
}

// CloseInitDialogMsg is a message that is sent when the init dialog is closed.
type CloseInitDialogMsg struct {
	Initialize bool
}

// ShowInitDialogMsg is a message that is sent to show the init dialog.
type ShowInitDialogMsg struct {
	Show bool
}
