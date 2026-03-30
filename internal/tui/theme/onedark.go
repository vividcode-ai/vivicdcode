package theme

// OneDarkTheme implements the Theme interface with Atom's One Dark colors.
type OneDarkTheme struct {
	BaseTheme
}

// NewOneDarkTheme creates a new instance of the One Dark theme.
func NewOneDarkTheme() *OneDarkTheme {
	darkBackground := "#282c34"
	darkCurrentLine := "#2c313c"
	darkSelection := "#3e4451"
	darkForeground := "#abb2bf"
	darkComment := "#5c6370"
	darkRed := "#e06c75"
	darkOrange := "#d19a66"
	darkYellow := "#e5c07b"
	darkGreen := "#98c379"
	darkCyan := "#56b6c2"
	darkBlue := "#61afef"
	darkPurple := "#c678dd"
	darkBorder := "#3b4048"

	theme := &OneDarkTheme{}

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
	theme.BackgroundDarkerColor = "#21252b"

	// Border colors
	theme.BorderNormalColor = darkBorder
	theme.BorderFocusedColor = darkBlue
	theme.BorderDimColor = darkSelection

	// Diff view colors
	theme.DiffAddedColor = "#478247"
	theme.DiffRemovedColor = "#7C4444"
	theme.DiffContextColor = "#a0a0a0"
	theme.DiffHunkHeaderColor = "#a0a0a0"
	theme.DiffHighlightAddedColor = "#DAFADA"
	theme.DiffHighlightRemovedColor = "#FADADD"
	theme.DiffAddedBgColor = "#303A30"
	theme.DiffRemovedBgColor = "#3A3030"
	theme.DiffContextBgColor = darkBackground
	theme.DiffLineNumberColor = "#888888"
	theme.DiffAddedLineNumberBgColor = "#293229"
	theme.DiffRemovedLineNumberBgColor = "#332929"

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
	// Register the One Dark theme with the theme manager
	RegisterTheme("onedark", NewOneDarkTheme())
}
