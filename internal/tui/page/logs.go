package page

import (
	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	uv "github.com/charmbracelet/ultraviolet"
	"github.com/vividcode-ai/vividcode/internal/tui/components/logs"
	"github.com/vividcode-ai/vividcode/internal/tui/layout"
	"github.com/vividcode-ai/vividcode/internal/tui/styles"
)

var LogsPage PageID = "logs"

type LogPage interface {
	tea.Model
	layout.Sizeable
	layout.Bindings
}
type logsPage struct {
	width, height int
	table         layout.Container
	details       layout.Container
}

func (p *logsPage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		p.width = msg.Width
		p.height = msg.Height
		return p, p.SetSize(msg.Width, msg.Height)
	}

	table, cmd := p.table.Update(msg)
	cmds = append(cmds, cmd)
	p.table = table.(layout.Container)
	details, cmd := p.details.Update(msg)
	cmds = append(cmds, cmd)
	p.details = details.(layout.Container)

	return p, tea.Batch(cmds...)
}

func (p *logsPage) View() tea.View {
	style := styles.BaseStyle().Width(p.width).Height(p.height)
	return tea.View{Content: style.Render(lipgloss.JoinVertical(lipgloss.Top,
		p.table.View().Content,
		p.details.View().Content,
	))}
}

func (p *logsPage) BindingKeys() []key.Binding {
	return p.table.BindingKeys()
}

func (p *logsPage) Draw(scr uv.Screen, area uv.Rectangle) *tea.Cursor {
	view := p.View().Content
	uv.NewStyledString(view).Draw(scr, area)
	return nil
}

// GetSize implements LogPage.
func (p *logsPage) GetSize() (int, int) {
	return p.width, p.height
}

// SetSize implements LogPage.
func (p *logsPage) SetSize(width int, height int) tea.Cmd {
	p.width = width
	p.height = height
	return tea.Batch(
		p.table.SetSize(width, height/2),
		p.details.SetSize(width, height/2),
	)
}

func (p *logsPage) Init() tea.Cmd {
	return tea.Batch(
		p.table.Init(),
		p.details.Init(),
	)
}

func NewLogsPage() LogPage {
	return &logsPage{
		table:   layout.NewContainer(logs.NewLogsTable(), layout.WithBorderAll()),
		details: layout.NewContainer(logs.NewLogsDetails(), layout.WithBorderAll()),
	}
}
