package theme

// MonokaiProTheme implements the Theme interface with Monokai Pro colors.
type MonokaiProTheme struct {
	BaseTheme
}

// NewMonokaiProTheme creates a new instance of the Monokai Pro theme.
func NewMonokaiProTheme() *MonokaiProTheme {
	darkBackground := "#2d2a2e"
	darkCurrentLine := "#403e41"
	darkSelection := "#5b595c"
	darkForeground := "#fcfcfa"
	darkComment := "#727072"
	darkRed := "#ff6188"
	darkOrange := "#fc9867"
	darkYellow := "#ffd866"
	darkGreen := "#a9dc76"
	darkCyan := "#78dce8"
	darkBlue := "#ab9df2"
	darkPurple := "#ab9df2"
	darkBorder := "#403e41"

	theme := &MonokaiProTheme{}

	// Base colors
	theme.PrimaryColor = darkCyan
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
	theme.BackgroundDarkerColor = "#221f22"

	// Border colors
	theme.BorderNormalColor = darkBorder
	theme.BorderFocusedColor = darkCyan
	theme.BorderDimColor = darkSelection

	// Diff view colors
	theme.DiffAddedColor = "#a9dc76"
	theme.DiffRemovedColor = "#ff6188"
	theme.DiffContextColor = "#a0a0a0"
	theme.DiffHunkHeaderColor = "#a0a0a0"
	theme.DiffHighlightAddedColor = "#c2e7a9"
	theme.DiffHighlightRemovedColor = "#ff8ca6"
	theme.DiffAddedBgColor = "#3a4a35"
	theme.DiffRemovedBgColor = "#4a3439"
	theme.DiffContextBgColor = darkBackground
	theme.DiffLineNumberColor = "#888888"
	theme.DiffAddedLineNumberBgColor = "#2d3a28"
	theme.DiffRemovedLineNumberBgColor = "#3d2a2e"

	// Markdown colors
	theme.MarkdownTextColor = darkForeground
	theme.MarkdownHeadingColor = darkPurple
	theme.MarkdownLinkColor = darkCyan
	theme.MarkdownLinkTextColor = darkBlue
	theme.MarkdownCodeColor = darkGreen
	theme.MarkdownBlockQuoteColor = darkYellow
	theme.MarkdownEmphColor = darkYellow
	theme.MarkdownStrongColor = darkOrange
	theme.MarkdownHorizontalRuleColor = darkComment
	theme.MarkdownListItemColor = darkCyan
	theme.MarkdownListEnumerationColor = darkBlue
	theme.MarkdownImageColor = darkCyan
	theme.MarkdownImageTextColor = darkBlue
	theme.MarkdownCodeBlockColor = darkForeground

	// Syntax highlighting colors
	theme.SyntaxCommentColor = darkComment
	theme.SyntaxKeywordColor = darkRed
	theme.SyntaxFunctionColor = darkGreen
	theme.SyntaxVariableColor = darkForeground
	theme.SyntaxStringColor = darkYellow
	theme.SyntaxNumberColor = darkPurple
	theme.SyntaxTypeColor = darkBlue
	theme.SyntaxOperatorColor = darkCyan
	theme.SyntaxPunctuationColor = darkForeground

	return theme
}

func init() {
	// Register the Monokai Pro theme with the theme manager
	RegisterTheme("monokai", NewMonokaiProTheme())
}
