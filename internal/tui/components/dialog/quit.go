package dialog

import (
	"strings"

	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	uv "github.com/charmbracelet/ultraviolet"
	"github.com/vividcode-ai/vividcode/internal/tui/layout"
	"github.com/vividcode-ai/vividcode/internal/tui/render"
	"github.com/vividcode-ai/vividcode/internal/tui/styles"
	"github.com/vividcode-ai/vividcode/internal/tui/theme"
	"github.com/vividcode-ai/vividcode/internal/tui/util"
)

const question = "Are you sure you want to quit?"

type CloseQuitMsg struct{}

type QuitDialog interface {
	tea.Model
	layout.Bindings
	Draw(scr uv.Screen, area uv.Rectangle) *tea.Cursor
}

type quitDialogCmp struct {
	selectedNo bool
}

type helpMapping struct {
	LeftRight  key.Binding
	EnterSpace key.Binding
	Yes        key.Binding
	No         key.Binding
	Tab        key.Binding
}

var helpKeys = helpMapping{
	LeftRight: key.NewBinding(
		key.WithKeys("left", "right"),
		key.WithHelp("←/→", "switch options"),
	),
	EnterSpace: key.NewBinding(
		key.WithKeys("enter", " "),
		key.WithHelp("enter/space", "confirm"),
	),
	Yes: key.NewBinding(
		key.WithKeys("y", "Y"),
		key.WithHelp("y/Y", "yes"),
	),
	No: key.NewBinding(
		key.WithKeys("n", "N"),
		key.WithHelp("n/N", "no"),
	),
	Tab: key.NewBinding(
		key.WithKeys("tab"),
		key.WithHelp("tab", "switch options"),
	),
}

func (q *quitDialogCmp) Init() tea.Cmd {
	return nil
}

func (q *quitDialogCmp) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, helpKeys.LeftRight) || key.Matches(msg, helpKeys.Tab):
			q.selectedNo = !q.selectedNo
			return q, nil
		case key.Matches(msg, helpKeys.EnterSpace):
			if !q.selectedNo {
				return q, tea.Quit
			}
			return q, util.CmdHandler(CloseQuitMsg{})
		case key.Matches(msg, helpKeys.Yes):
			return q, tea.Quit
		case key.Matches(msg, helpKeys.No):
			return q, util.CmdHandler(CloseQuitMsg{})
		}
	}
	return q, nil
}

func (q *quitDialogCmp) View() tea.View {
	t := theme.CurrentTheme()
	baseStyle := styles.BaseStyle()

	yesStyle := baseStyle
	noStyle := baseStyle
	spacerStyle := baseStyle.Background(lipgloss.Color(t.Background()))

	if q.selectedNo {
		noStyle = noStyle.Background(lipgloss.Color(t.Primary())).Foreground(lipgloss.Color(t.Background()))
		yesStyle = yesStyle.Background(lipgloss.Color(t.Background())).Foreground(lipgloss.Color(t.Primary()))
	} else {
		yesStyle = yesStyle.Background(lipgloss.Color(t.Primary())).Foreground(lipgloss.Color(t.Background()))
		noStyle = noStyle.Background(lipgloss.Color(t.Background())).Foreground(lipgloss.Color(t.Primary()))
	}

	yesButton := yesStyle.Padding(0, 1).Render("Yes")
	noButton := noStyle.Padding(0, 1).Render("No")

	buttons := lipgloss.JoinHorizontal(lipgloss.Left, yesButton, spacerStyle.Render("  "), noButton)

	width := lipgloss.Width(question)
	remainingWidth := width - lipgloss.Width(buttons)
	if remainingWidth > 0 {
		buttons = spacerStyle.Render(strings.Repeat(" ", remainingWidth)) + buttons
	}

	content := baseStyle.Render(
		lipgloss.JoinVertical(
			lipgloss.Center,
			question,
			"",
			buttons,
		),
	)

	return tea.View{Content: baseStyle.Padding(1, 2).
		Border(lipgloss.RoundedBorder()).
		BorderBackground(lipgloss.Color(t.Background())).
		BorderForeground(lipgloss.Color(t.TextMuted())).
		Width(lipgloss.Width(content) + 6).
		Render(content)}
}

func (q *quitDialogCmp) Draw(scr uv.Screen, area uv.Rectangle) *tea.Cursor {
	view := q.View().Content
	render.DrawCenter(scr, area, view)
	return nil
}

func (q *quitDialogCmp) BindingKeys() []key.Binding {
	return layout.KeyMapToSlice(helpKeys)
}

func NewQuitCmp() QuitDialog {
	return &quitDialogCmp{
		selectedNo: true,
	}
}
