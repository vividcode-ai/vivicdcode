package theme

// TokyoNightTheme implements the Theme interface with Tokyo Night colors.
type TokyoNightTheme struct {
	BaseTheme
}

// NewTokyoNightTheme creates a new instance of the Tokyo Night theme.
func NewTokyoNightTheme() *TokyoNightTheme {
	darkBackground := "#222436"
	darkCurrentLine := "#1e2030"
	darkSelection := "#2f334d"
	darkForeground := "#c8d3f5"
	darkComment := "#636da6"
	darkRed := "#ff757f"
	darkOrange := "#ff966c"
	darkYellow := "#ffc777"
	darkGreen := "#c3e88d"
	darkCyan := "#86e1fc"
	darkBlue := "#82aaff"
	darkPurple := "#c099ff"
	darkBorder := "#3b4261"

	theme := &TokyoNightTheme{}

	// Base colors
	theme.PrimaryColor = darkBlue
	theme.SecondaryColor = darkPurple
	theme.AccentColor = darkOrange

	// Status colors
	theme.ErrorColor = darkRed
	theme.WarningColor = darkOrange
	theme.SuccessColor = darkGreen
	theme.InfoColor = darkBlue

	// Text colors
	theme.TextColor = darkForeground
	theme.TextMutedColor = darkComment
	theme.TextEmphasizedColor = darkYellow

	// Background colors
	theme.BackgroundColor = darkBackground
	theme.BackgroundSecondaryColor = darkCurrentLine
	theme.BackgroundDarkerColor = "#191B29"

	// Border colors
	theme.BorderNormalColor = darkBorder
	theme.BorderFocusedColor = darkBlue
	theme.BorderDimColor = darkSelection

	// Diff view colors
	theme.DiffAddedColor = "#4fd6be"
	theme.DiffRemovedColor = "#c53b53"
	theme.DiffContextColor = "#828bb8"
	theme.DiffHunkHeaderColor = "#828bb8"
	theme.DiffHighlightAddedColor = "#b8db87"
	theme.DiffHighlightRemovedColor = "#e26a75"
	theme.DiffAddedBgColor = "#20303b"
	theme.DiffRemovedBgColor = "#37222c"
	theme.DiffContextBgColor = darkBackground
	theme.DiffLineNumberColor = "#545c7e"
	theme.DiffAddedLineNumberBgColor = "#1b2b34"
	theme.DiffRemovedLineNumberBgColor = "#2d1f26"

	// Markdown colors
	theme.MarkdownTextColor = darkForeground
	theme.MarkdownHeadingColor = darkPurple
	theme.MarkdownLinkColor = darkBlue
	theme.MarkdownLinkTextColor = darkCyan
	theme.MarkdownCodeColor = darkGreen
	theme.MarkdownBlockQuoteColor = darkYellow
	theme.MarkdownEmphColor = darkYellow
	theme.MarkdownStrongColor = darkOrange
	theme.MarkdownHorizontalRuleColor = darkComment
	theme.MarkdownListItemColor = darkBlue
	theme.MarkdownListEnumerationColor = darkCyan
	theme.MarkdownImageColor = darkBlue
	theme.MarkdownImageTextColor = darkCyan
	theme.MarkdownCodeBlockColor = darkForeground

	// Syntax highlighting colors
	theme.SyntaxCommentColor = darkComment
	theme.SyntaxKeywordColor = darkPurple
	theme.SyntaxFunctionColor = darkBlue
	theme.SyntaxVariableColor = darkRed
	theme.SyntaxStringColor = darkGreen
	theme.SyntaxNumberColor = darkOrange
	theme.SyntaxTypeColor = darkYellow
	theme.SyntaxOperatorColor = darkCyan
	theme.SyntaxPunctuationColor = darkForeground

	return theme
}

func init() {
	// Register the Tokyo Night theme with the theme manager
	RegisterTheme("tokyonight", NewTokyoNightTheme())
}
