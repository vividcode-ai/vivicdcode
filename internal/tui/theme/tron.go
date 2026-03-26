package theme

// TronTheme implements the Theme interface with Tron-inspired colors.
type TronTheme struct {
	BaseTheme
}

// NewTronTheme creates a new instance of the Tron theme.
func NewTronTheme() *TronTheme {
	darkBackground := "#0c141f"
	darkCurrentLine := "#1a2633"
	darkSelection := "#1a2633"
	darkForeground := "#caf0ff"
	darkComment := "#4d6b87"
	darkCyan := "#00d9ff"
	darkBlue := "#007fff"
	darkOrange := "#ff9000"
	darkPink := "#ff00a0"
	darkPurple := "#b73fff"
	darkRed := "#ff3333"
	darkYellow := "#ffcc00"
	darkGreen := "#00ff8f"
	darkBorder := "#1a2633"

	theme := &TronTheme{}

	// Base colors
	theme.PrimaryColor = darkCyan
	theme.SecondaryColor = darkBlue
	theme.AccentColor = darkOrange

	// Status colors
	theme.ErrorColor = darkRed
	theme.WarningColor = darkOrange
	theme.SuccessColor = darkGreen
	theme.InfoColor = darkCyan

	// Text colors
	theme.TextColor = darkForeground
	theme.TextMutedColor = darkComment
	theme.TextEmphasizedColor = darkYellow

	// Background colors
	theme.BackgroundColor = darkBackground
	theme.BackgroundSecondaryColor = darkCurrentLine
	theme.BackgroundDarkerColor = "#070d14"

	// Border colors
	theme.BorderNormalColor = darkBorder
	theme.BorderFocusedColor = darkCyan
	theme.BorderDimColor = darkSelection

	// Diff view colors
	theme.DiffAddedColor = darkGreen
	theme.DiffRemovedColor = darkRed
	theme.DiffContextColor = darkComment
	theme.DiffHunkHeaderColor = darkBlue
	theme.DiffHighlightAddedColor = "#00ff8f"
	theme.DiffHighlightRemovedColor = "#ff3333"
	theme.DiffAddedBgColor = "#0a2a1a"
	theme.DiffRemovedBgColor = "#2a0a0a"
	theme.DiffContextBgColor = darkBackground
	theme.DiffLineNumberColor = darkComment
	theme.DiffAddedLineNumberBgColor = "#082015"
	theme.DiffRemovedLineNumberBgColor = "#200808"

	// Markdown colors
	theme.MarkdownTextColor = darkForeground
	theme.MarkdownHeadingColor = darkCyan
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
	theme.SyntaxKeywordColor = darkCyan
	theme.SyntaxFunctionColor = darkGreen
	theme.SyntaxVariableColor = darkOrange
	theme.SyntaxStringColor = darkYellow
	theme.SyntaxNumberColor = darkBlue
	theme.SyntaxTypeColor = darkPurple
	theme.SyntaxOperatorColor = darkPink
	theme.SyntaxPunctuationColor = darkForeground

	return theme
}

func init() {
	// Register the Tron theme with the theme manager
	RegisterTheme("tron", NewTronTheme())
}
