package theme

// DraculaTheme implements the Theme interface with Dracula colors.
// It provides both dark and light variants, though Dracula is primarily a dark theme.
type DraculaTheme struct {
	BaseTheme
}

// NewDraculaTheme creates a new instance of the Dracula theme.
func NewDraculaTheme() *DraculaTheme {
	darkBackground := "#282a36"
	darkCurrentLine := "#44475a"
	darkSelection := "#44475a"
	darkForeground := "#f8f8f2"
	darkComment := "#6272a4"
	darkCyan := "#8be9fd"
	darkGreen := "#50fa7b"
	darkOrange := "#ffb86c"
	darkPink := "#ff79c6"
	darkPurple := "#bd93f9"
	darkRed := "#ff5555"
	darkYellow := "#f1fa8c"
	darkBorder := "#44475a"

	theme := &DraculaTheme{}

	// Base colors
	theme.PrimaryColor = darkPurple
	theme.SecondaryColor = darkPink
	theme.AccentColor = darkCyan

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
	theme.BackgroundDarkerColor = "#21222c"

	// Border colors
	theme.BorderNormalColor = darkBorder
	theme.BorderFocusedColor = darkPurple
	theme.BorderDimColor = darkSelection

	// Diff view colors
	theme.DiffAddedColor = darkGreen
	theme.DiffRemovedColor = darkRed
	theme.DiffContextColor = darkComment
	theme.DiffHunkHeaderColor = darkPurple
	theme.DiffHighlightAddedColor = "#50fa7b"
	theme.DiffHighlightRemovedColor = "#ff5555"
	theme.DiffAddedBgColor = "#2c3b2c"
	theme.DiffRemovedBgColor = "#3b2c2c"
	theme.DiffContextBgColor = darkBackground
	theme.DiffLineNumberColor = darkComment
	theme.DiffAddedLineNumberBgColor = "#253025"
	theme.DiffRemovedLineNumberBgColor = "#302525"

	// Markdown colors
	theme.MarkdownTextColor = darkForeground
	theme.MarkdownHeadingColor = darkPink
	theme.MarkdownLinkColor = darkPurple
	theme.MarkdownLinkTextColor = darkCyan
	theme.MarkdownCodeColor = darkGreen
	theme.MarkdownBlockQuoteColor = darkYellow
	theme.MarkdownEmphColor = darkYellow
	theme.MarkdownStrongColor = darkOrange
	theme.MarkdownHorizontalRuleColor = darkComment
	theme.MarkdownListItemColor = darkPurple
	theme.MarkdownListEnumerationColor = darkCyan
	theme.MarkdownImageColor = darkPurple
	theme.MarkdownImageTextColor = darkCyan
	theme.MarkdownCodeBlockColor = darkForeground

	// Syntax highlighting colors
	theme.SyntaxCommentColor = darkComment
	theme.SyntaxKeywordColor = darkPink
	theme.SyntaxFunctionColor = darkGreen
	theme.SyntaxVariableColor = darkOrange
	theme.SyntaxStringColor = darkYellow
	theme.SyntaxNumberColor = darkPurple
	theme.SyntaxTypeColor = darkCyan
	theme.SyntaxOperatorColor = darkPink
	theme.SyntaxPunctuationColor = darkForeground

	return theme
}

func init() {
	// Register the Dracula theme with the theme manager
	RegisterTheme("dracula", NewDraculaTheme())
}
