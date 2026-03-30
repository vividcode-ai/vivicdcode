package theme

const (
	flexokiPaper   = "#FFFCF0"
	flexokiBase50  = "#F2F0E5"
	flexokiBase100 = "#E6E4D9"
	flexokiBase150 = "#DAD8CE"
	flexokiBase200 = "#CECDC3"
	flexokiBase300 = "#B7B5AC"
	flexokiBase500 = "#878580"
	flexokiBase600 = "#6F6E69"
	flexokiBase700 = "#575653"
	flexokiBase800 = "#403E3C"
	flexokiBase850 = "#343331"
	flexokiBase900 = "#282726"
	flexokiBase950 = "#1C1B1A"
	flexokiBlack   = "#100F0F"

	flexokiRed600     = "#AF3029"
	flexokiOrange600  = "#BC5215"
	flexokiYellow600  = "#AD8301"
	flexokiGreen600   = "#66800B"
	flexokiCyan600    = "#24837B"
	flexokiBlue600    = "#205EA6"
	flexokiPurple600  = "#5E409D"
	flexokiMagenta600 = "#A02F6F"

	flexokiRed400     = "#D14D41"
	flexokiOrange400  = "#DA702C"
	flexokiYellow400  = "#D0A215"
	flexokiGreen400   = "#879A39"
	flexokiCyan400    = "#3AA99F"
	flexokiBlue400    = "#4385BE"
	flexokiPurple400  = "#8B7EC8"
	flexokiMagenta400 = "#CE5D97"
)

// FlexokiTheme implements the Theme interface with Flexoki colors.
type FlexokiTheme struct {
	BaseTheme
}

// NewFlexokiTheme creates a new instance of the Flexoki theme.
func NewFlexokiTheme() *FlexokiTheme {
	theme := &FlexokiTheme{}

	// Base colors
	theme.PrimaryColor = flexokiBlue400
	theme.SecondaryColor = flexokiPurple400
	theme.AccentColor = flexokiOrange400

	// Status colors
	theme.ErrorColor = flexokiRed400
	theme.WarningColor = flexokiYellow400
	theme.SuccessColor = flexokiGreen400
	theme.InfoColor = flexokiCyan400

	// Text colors
	theme.TextColor = flexokiBase300
	theme.TextMutedColor = flexokiBase700
	theme.TextEmphasizedColor = flexokiYellow400

	// Background colors
	theme.BackgroundColor = flexokiBlack
	theme.BackgroundSecondaryColor = flexokiBase950
	theme.BackgroundDarkerColor = flexokiBase900

	// Border colors
	theme.BorderNormalColor = flexokiBase900
	theme.BorderFocusedColor = flexokiBlue400
	theme.BorderDimColor = flexokiBase850

	// Diff view colors
	theme.DiffAddedColor = flexokiGreen400
	theme.DiffRemovedColor = flexokiRed400
	theme.DiffContextColor = flexokiBase700
	theme.DiffHunkHeaderColor = flexokiBase700
	theme.DiffHighlightAddedColor = flexokiGreen400
	theme.DiffHighlightRemovedColor = flexokiRed400
	theme.DiffAddedBgColor = "#1D2419"
	theme.DiffRemovedBgColor = "#241919"
	theme.DiffContextBgColor = flexokiBlack
	theme.DiffLineNumberColor = flexokiBase700
	theme.DiffAddedLineNumberBgColor = "#1A2017"
	theme.DiffRemovedLineNumberBgColor = "#201717"

	// Markdown colors
	theme.MarkdownTextColor = flexokiBase300
	theme.MarkdownHeadingColor = flexokiYellow400
	theme.MarkdownLinkColor = flexokiCyan400
	theme.MarkdownLinkTextColor = flexokiMagenta400
	theme.MarkdownCodeColor = flexokiGreen400
	theme.MarkdownBlockQuoteColor = flexokiCyan400
	theme.MarkdownEmphColor = flexokiYellow400
	theme.MarkdownStrongColor = flexokiOrange400
	theme.MarkdownHorizontalRuleColor = flexokiBase800
	theme.MarkdownListItemColor = flexokiBlue400
	theme.MarkdownListEnumerationColor = flexokiBlue400
	theme.MarkdownImageColor = flexokiPurple400
	theme.MarkdownImageTextColor = flexokiMagenta400
	theme.MarkdownCodeBlockColor = flexokiBase300

	// Syntax highlighting colors
	theme.SyntaxCommentColor = flexokiBase700
	theme.SyntaxKeywordColor = flexokiGreen400
	theme.SyntaxFunctionColor = flexokiOrange400
	theme.SyntaxVariableColor = flexokiBlue400
	theme.SyntaxStringColor = flexokiCyan400
	theme.SyntaxNumberColor = flexokiPurple400
	theme.SyntaxTypeColor = flexokiYellow400
	theme.SyntaxOperatorColor = flexokiBase500
	theme.SyntaxPunctuationColor = flexokiBase500

	return theme
}

func init() {
	// Register the Flexoki theme with the theme manager
	RegisterTheme("flexoki", NewFlexokiTheme())
}
