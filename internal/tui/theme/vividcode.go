package theme

// VividCodeTheme implements the Theme interface with VividCode brand colors.
type VividCodeTheme struct {
	BaseTheme
}

// NewVividCodeTheme creates a new instance of the VividCode theme.
func NewVividCodeTheme() *VividCodeTheme {
	darkBackground := "#212121"
	darkCurrentLine := "#252525"
	darkSelection := "#303030"
	darkForeground := "#e0e0e0"
	darkComment := "#6a6a6a"
	darkPrimary := "#fab283"
	darkSecondary := "#5c9cf5"
	darkAccent := "#9d7cd8"
	darkRed := "#e06c75"
	darkOrange := "#f5a742"
	darkGreen := "#7fd88f"
	darkCyan := "#56b6c2"
	darkYellow := "#e5c07b"
	darkBorder := "#4b4c5c"

	theme := &VividCodeTheme{}

	// Base colors
	theme.PrimaryColor = darkPrimary
	theme.SecondaryColor = darkSecondary
	theme.AccentColor = darkAccent

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
	theme.BackgroundDarkerColor = "#121212"

	// Border colors
	theme.BorderNormalColor = darkBorder
	theme.BorderFocusedColor = darkPrimary
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
	theme.MarkdownHeadingColor = darkSecondary
	theme.MarkdownLinkColor = darkPrimary
	theme.MarkdownLinkTextColor = darkCyan
	theme.MarkdownCodeColor = darkGreen
	theme.MarkdownBlockQuoteColor = darkYellow
	theme.MarkdownEmphColor = darkYellow
	theme.MarkdownStrongColor = darkAccent
	theme.MarkdownHorizontalRuleColor = darkComment
	theme.MarkdownListItemColor = darkPrimary
	theme.MarkdownListEnumerationColor = darkCyan
	theme.MarkdownImageColor = darkPrimary
	theme.MarkdownImageTextColor = darkCyan
	theme.MarkdownCodeBlockColor = darkForeground

	// Syntax highlighting colors
	theme.SyntaxCommentColor = darkComment
	theme.SyntaxKeywordColor = darkSecondary
	theme.SyntaxFunctionColor = darkPrimary
	theme.SyntaxVariableColor = darkRed
	theme.SyntaxStringColor = darkGreen
	theme.SyntaxNumberColor = darkAccent
	theme.SyntaxTypeColor = darkYellow
	theme.SyntaxOperatorColor = darkCyan
	theme.SyntaxPunctuationColor = darkForeground

	return theme
}

func init() {
	// Register the VividCode theme with the theme manager
	RegisterTheme("vividcode", NewVividCodeTheme())
}
