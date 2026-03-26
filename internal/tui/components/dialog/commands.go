package dialog

import (
	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	uv "github.com/charmbracelet/ultraviolet"
	utilComponents "github.com/vividcode-ai/vividcode/internal/tui/components/util"
	"github.com/vividcode-ai/vividcode/internal/tui/layout"
	"github.com/vividcode-ai/vividcode/internal/tui/render"
	"github.com/vividcode-ai/vividcode/internal/tui/styles"
	"github.com/vividcode-ai/vividcode/internal/tui/theme"
	"github.com/vividcode-ai/vividcode/internal/tui/util"
)

// Command represents a command that can be executed
type Command struct {
	ID          string
	Title       string
	Description string
	Handler     func(cmd Command) tea.Cmd
}

func (ci Command) Render(selected bool, width int) string {
	t := theme.CurrentTheme()
	baseStyle := styles.BaseStyle()

	descStyle := baseStyle.Width(width).Foreground(lipgloss.Color(t.TextMuted()))
	itemStyle := baseStyle.Width(width).
		Foreground(lipgloss.Color(t.Text())).
		Background(lipgloss.Color(t.Background()))

	if selected {
		itemStyle = itemStyle.
			Background(lipgloss.Color(t.Primary())).
			Foreground(lipgloss.Color(t.Background())).
			Bold(true)
		descStyle = descStyle.
			Background(lipgloss.Color(t.Primary())).
			Foreground(lipgloss.Color(t.Background()))
	}

	title := itemStyle.Padding(0, 1).Render(ci.Title)
	if ci.Description != "" {
		description := descStyle.Padding(0, 1).Render(ci.Description)
		return lipgloss.JoinVertical(lipgloss.Left, title, description)
	}
	return title
}

// CommandSelectedMsg is sent when a command is selected
type CommandSelectedMsg struct {
	Command Command
}

// CloseCommandDialogMsg is sent when the command dialog is closed
type CloseCommandDialogMsg struct{}

// CommandDialog interface for the command selection dialog
type CommandDialog interface {
	tea.Model
	layout.Bindings
	SetCommands(commands []Command)
	Draw(scr uv.Screen, area uv.Rectangle) *tea.Cursor
}

type commandDialogCmp struct {
	listView utilComponents.SimpleList[Command]
	width    int
	height   int
}

type commandKeyMap struct {
	Enter  key.Binding
	Escape key.Binding
}

var commandKeys = commandKeyMap{
	Enter: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "select command"),
	),
	Escape: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "close"),
	),
}

func (c *commandDialogCmp) Init() tea.Cmd {
	return c.listView.Init()
}

func (c *commandDialogCmp) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, commandKeys.Enter):
			selectedItem, idx := c.listView.GetSelectedItem()
			if idx != -1 {
				return c, util.CmdHandler(CommandSelectedMsg{
					Command: selectedItem,
				})
			}
		case key.Matches(msg, commandKeys.Escape):
			return c, util.CmdHandler(CloseCommandDialogMsg{})
		}
	case tea.WindowSizeMsg:
		c.width = msg.Width
		c.height = msg.Height
	}

	u, cmd := c.listView.Update(msg)
	c.listView = u.(utilComponents.SimpleList[Command])
	cmds = append(cmds, cmd)

	return c, tea.Batch(cmds...)
}

func (c *commandDialogCmp) View() tea.View {
	t := theme.CurrentTheme()
	baseStyle := styles.BaseStyle()

	maxWidth := 40

	commands := c.listView.GetItems()

	for _, cmd := range commands {
		if len(cmd.Title) > maxWidth-4 {
			maxWidth = len(cmd.Title) + 4
		}
		if cmd.Description != "" {
			if len(cmd.Description) > maxWidth-4 {
				maxWidth = len(cmd.Description) + 4
			}
		}
	}

	c.listView.SetMaxWidth(maxWidth)

	title := baseStyle.
		Foreground(lipgloss.Color(t.Primary())).
		Bold(true).
		Width(maxWidth).
		Padding(0, 1).
		Render("Commands")

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		baseStyle.Width(maxWidth).Render(""),
		baseStyle.Width(maxWidth).Render(c.listView.View().Content),
		baseStyle.Width(maxWidth).Render(""),
	)

	return tea.View{Content: baseStyle.Padding(1, 2).
		Border(lipgloss.RoundedBorder()).
		BorderBackground(lipgloss.Color(t.Background())).
		BorderForeground(lipgloss.Color(t.TextMuted())).
		Width(lipgloss.Width(content) + 4).
		Render(content)}
}

func (c *commandDialogCmp) Draw(scr uv.Screen, area uv.Rectangle) *tea.Cursor {
	view := c.View().Content
	render.DrawCenter(scr, area, view)
	return nil
}

func (c *commandDialogCmp) BindingKeys() []key.Binding {
	return layout.KeyMapToSlice(commandKeys)
}

func (c *commandDialogCmp) SetCommands(commands []Command) {
	c.listView.SetItems(commands)
}

// NewCommandDialogCmp creates a new command selection dialog
func NewCommandDialogCmp() CommandDialog {
	listView := utilComponents.NewSimpleList[Command](
		[]Command{},
		10,
		"No commands available",
		true,
	)
	return &commandDialogCmp{
		listView: listView,
	}
}
