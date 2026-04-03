package layout

import (
	"strings"

	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	uv "github.com/charmbracelet/ultraviolet"
	"github.com/vividcode-ai/vividcode/internal/tui/theme"
)

type SplitPaneLayout interface {
	tea.Model
	Sizeable
	Bindings
	SetLeftPanel(panel Container) tea.Cmd
	SetRightPanel(panel Container) tea.Cmd
	SetBottomPanel(panel Container) tea.Cmd

	ClearLeftPanel() tea.Cmd
	ClearRightPanel() tea.Cmd
	ClearBottomPanel() tea.Cmd

	ScrollBy(lines int) tea.Cmd

	Draw(scr uv.Screen, area uv.Rectangle) *tea.Cursor
}

type splitPaneLayout struct {
	width         int
	height        int
	ratio         float64
	verticalRatio float64

	rightPanel             Container
	leftPanel              Container
	bottomPanel            Container
	bottomPanelFixedHeight int
}

type SplitPaneOption func(*splitPaneLayout)

func (s *splitPaneLayout) Init() tea.Cmd {
	var cmds []tea.Cmd

	if s.leftPanel != nil {
		cmds = append(cmds, s.leftPanel.Init())
	}

	if s.rightPanel != nil {
		cmds = append(cmds, s.rightPanel.Init())
	}

	if s.bottomPanel != nil {
		cmds = append(cmds, s.bottomPanel.Init())
	}

	return tea.Batch(cmds...)
}

func (s *splitPaneLayout) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		return s, s.SetSize(msg.Width, msg.Height)
	}

	if s.rightPanel != nil {
		u, cmd := s.rightPanel.Update(msg)
		s.rightPanel = u.(Container)
		if cmd != nil {
			cmds = append(cmds, cmd)
		}
	}

	if s.leftPanel != nil {
		u, cmd := s.leftPanel.Update(msg)
		s.leftPanel = u.(Container)
		if cmd != nil {
			cmds = append(cmds, cmd)
		}
	}

	if s.bottomPanel != nil {
		u, cmd := s.bottomPanel.Update(msg)
		s.bottomPanel = u.(Container)
		if cmd != nil {
			cmds = append(cmds, cmd)
		}
	}

	return s, tea.Batch(cmds...)
}

func (s *splitPaneLayout) View() tea.View {
	t := theme.CurrentTheme()

	var topSection string
	if s.leftPanel != nil && s.rightPanel != nil {
		leftView := s.leftPanel.View()
		rightView := s.rightPanel.View()
		topSection = lipgloss.JoinHorizontal(lipgloss.Top, leftView.Content, rightView.Content)
	} else if s.leftPanel != nil {
		topSection = s.leftPanel.View().Content
	} else if s.rightPanel != nil {
		topSection = s.rightPanel.View().Content
	} else {
		topSection = ""
	}

	var finalView string
	if s.bottomPanel != nil && topSection != "" {
		bottomView := s.bottomPanel.View()
		finalView = lipgloss.JoinVertical(lipgloss.Left, topSection, bottomView.Content)
	} else if s.bottomPanel != nil {
		finalView = s.bottomPanel.View().Content
	} else {
		finalView = topSection
	}

	if finalView == "" {
		finalView = strings.Repeat(" ", s.width)
	}

	currentHeight := lipgloss.Height(finalView)
	if currentHeight < s.height {
		fillLine := strings.Repeat(" ", s.width)
		for i := currentHeight; i < s.height; i++ {
			finalView += "\n" + fillLine
		}
	}

	style := lipgloss.NewStyle().
		Width(s.width).
		Height(s.height).
		Background(lipgloss.Color(t.Background()))

	return tea.View{Content: style.Render(finalView)}
}

func (s *splitPaneLayout) SetSize(width, height int) tea.Cmd {
	s.width = width
	s.height = height

	var topHeight, bottomHeight int
	if s.bottomPanel != nil {
		if s.bottomPanelFixedHeight > 0 {
			bottomHeight = s.bottomPanelFixedHeight
			topHeight = height - bottomHeight
		} else {
			topHeight = int(float64(height) * s.verticalRatio)
			bottomHeight = height - topHeight
		}
	} else {
		topHeight = height
		bottomHeight = 0
	}

	var leftWidth, rightWidth int
	if s.leftPanel != nil && s.rightPanel != nil {
		leftWidth = int(float64(width) * s.ratio)
		rightWidth = width - leftWidth
	} else if s.leftPanel != nil {
		leftWidth = width
		rightWidth = 0
	} else if s.rightPanel != nil {
		leftWidth = 0
		rightWidth = width
	}

	var cmds []tea.Cmd
	if s.leftPanel != nil {
		cmd := s.leftPanel.SetSize(leftWidth, topHeight)
		cmds = append(cmds, cmd)
	}

	if s.rightPanel != nil {
		cmd := s.rightPanel.SetSize(rightWidth, topHeight)
		cmds = append(cmds, cmd)
	}

	if s.bottomPanel != nil {
		cmd := s.bottomPanel.SetSize(width, bottomHeight)
		cmds = append(cmds, cmd)
	}
	return tea.Batch(cmds...)
}

func (s *splitPaneLayout) GetSize() (int, int) {
	return s.width, s.height
}

func (s *splitPaneLayout) SetLeftPanel(panel Container) tea.Cmd {
	s.leftPanel = panel
	if s.width > 0 && s.height > 0 {
		return s.SetSize(s.width, s.height)
	}
	return nil
}

func (s *splitPaneLayout) SetRightPanel(panel Container) tea.Cmd {
	s.rightPanel = panel
	if s.width > 0 && s.height > 0 {
		return s.SetSize(s.width, s.height)
	}
	return nil
}

func (s *splitPaneLayout) SetBottomPanel(panel Container) tea.Cmd {
	s.bottomPanel = panel
	if s.width > 0 && s.height > 0 {
		return s.SetSize(s.width, s.height)
	}
	return nil
}

func (s *splitPaneLayout) ClearLeftPanel() tea.Cmd {
	s.leftPanel = nil
	if s.width > 0 && s.height > 0 {
		return s.SetSize(s.width, s.height)
	}
	return nil
}

func (s *splitPaneLayout) ClearRightPanel() tea.Cmd {
	s.rightPanel = nil
	if s.width > 0 && s.height > 0 {
		return s.SetSize(s.width, s.height)
	}
	return nil
}

func (s *splitPaneLayout) ClearBottomPanel() tea.Cmd {
	s.bottomPanel = nil
	if s.width > 0 && s.height > 0 {
		return s.SetSize(s.width, s.height)
	}
	return nil
}

func (s *splitPaneLayout) BindingKeys() []key.Binding {
	keys := []key.Binding{}
	if s.leftPanel != nil {
		if b, ok := s.leftPanel.(Bindings); ok {
			keys = append(keys, b.BindingKeys()...)
		}
	}
	if s.rightPanel != nil {
		if b, ok := s.rightPanel.(Bindings); ok {
			keys = append(keys, b.BindingKeys()...)
		}
	}
	if s.bottomPanel != nil {
		if b, ok := s.bottomPanel.(Bindings); ok {
			keys = append(keys, b.BindingKeys()...)
		}
	}
	return keys
}

func (s *splitPaneLayout) Draw(scr uv.Screen, area uv.Rectangle) *tea.Cursor {
	view := s.View().Content
	uv.NewStyledString(view).Draw(scr, area)

	// Debug: Check actual drawing area
	_ = area // area is the rect we're drawing into
	return nil
}

func NewSplitPane(options ...SplitPaneOption) SplitPaneLayout {

	layout := &splitPaneLayout{
		ratio:         0.8,
		verticalRatio: 0.9, // Default 90% for top section, 10% for bottom
	}
	for _, option := range options {
		option(layout)
	}
	return layout
}

func WithLeftPanel(panel Container) SplitPaneOption {
	return func(s *splitPaneLayout) {
		s.leftPanel = panel
	}
}

func WithRightPanel(panel Container) SplitPaneOption {
	return func(s *splitPaneLayout) {
		s.rightPanel = panel
	}
}

func WithRatio(ratio float64) SplitPaneOption {
	return func(s *splitPaneLayout) {
		s.ratio = ratio
	}
}

func WithBottomPanel(panel Container) SplitPaneOption {
	return func(s *splitPaneLayout) {
		s.bottomPanel = panel
	}
}

func WithVerticalRatio(ratio float64) SplitPaneOption {
	return func(s *splitPaneLayout) {
		s.verticalRatio = ratio
	}
}

func WithBottomPanelFixed(panel Container, fixedHeight int) SplitPaneOption {
	return func(s *splitPaneLayout) {
		s.bottomPanel = panel
		s.bottomPanelFixedHeight = fixedHeight
	}
}

func (s *splitPaneLayout) ScrollBy(lines int) tea.Cmd {
	if s.leftPanel != nil {
		if scrollable, ok := s.leftPanel.(interface{ ScrollBy(lines int) tea.Cmd }); ok {
			return scrollable.ScrollBy(lines)
		}
	}
	return nil
}
