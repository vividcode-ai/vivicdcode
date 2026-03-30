package theme

const (
	gruvboxDarkBg0          = "#282828"
	gruvboxDarkBg0Soft      = "#32302f"
	gruvboxDarkBg1          = "#3c3836"
	gruvboxDarkBg2          = "#504945"
	gruvboxDarkBg3          = "#665c54"
	gruvboxDarkBg4          = "#7c6f64"
	gruvboxDarkFg0          = "#fbf1c7"
	gruvboxDarkFg1          = "#ebdbb2"
	gruvboxDarkFg2          = "#d5c4a1"
	gruvboxDarkFg3          = "#bdae93"
	gruvboxDarkFg4          = "#a89984"
	gruvboxDarkGray         = "#928374"
	gruvboxDarkRed          = "#cc241d"
	gruvboxDarkRedBright    = "#fb4934"
	gruvboxDarkGreen        = "#98971a"
	gruvboxDarkGreenBright  = "#b8bb26"
	gruvboxDarkYellow       = "#d79921"
	gruvboxDarkYellowBright = "#fabd2f"
	gruvboxDarkBlue         = "#458588"
	gruvboxDarkBlueBright   = "#83a598"
	gruvboxDarkPurple       = "#b16286"
	gruvboxDarkPurpleBright = "#d3869b"
	gruvboxDarkAqua         = "#689d6a"
	gruvboxDarkAquaBright   = "#8ec07c"
	gruvboxDarkOrange       = "#d65d0e"
	gruvboxDarkOrangeBright = "#fe8019"

	gruvboxLightBg0          = "#fbf1c7"
	gruvboxLightBg0Soft      = "#f2e5bc"
	gruvboxLightBg1          = "#ebdbb2"
	gruvboxLightBg2          = "#d5c4a1"
	gruvboxLightBg3          = "#bdae93"
	gruvboxLightBg4          = "#a89984"
	gruvboxLightFg0          = "#282828"
	gruvboxLightFg1          = "#3c3836"
	gruvboxLightFg2          = "#504945"
	gruvboxLightFg3          = "#665c54"
	gruvboxLightFg4          = "#7c6f64"
	gruvboxLightGray         = "#928374"
	gruvboxLightRed          = "#9d0006"
	gruvboxLightRedBright    = "#cc241d"
	gruvboxLightGreen        = "#79740e"
	gruvboxLightGreenBright  = "#98971a"
	gruvboxLightYellow       = "#b57614"
	gruvboxLightYellowBright = "#d79921"
	gruvboxLightBlue         = "#076678"
	gruvboxLightBlueBright   = "#458588"
	gruvboxLightPurple       = "#8f3f71"
	gruvboxLightPurpleBright = "#b16286"
	gruvboxLightAqua         = "#427b58"
	gruvboxLightAquaBright   = "#689d6a"
	gruvboxLightOrange       = "#af3a03"
	gruvboxLightOrangeBright = "#d65d0e"
)

// GruvboxTheme implements the Theme interface with Gruvbox colors.
type GruvboxTheme struct {
	BaseTheme
}

// NewGruvboxTheme creates a new instance of the Gruvbox theme.
func NewGruvboxTheme() *GruvboxTheme {
	theme := &GruvboxTheme{}

	// Base colors
	theme.PrimaryColor = gruvboxDarkBlueBright
	theme.SecondaryColor = gruvboxDarkPurpleBright
	theme.AccentColor = gruvboxDarkOrangeBright

	// Status colors
	theme.ErrorColor = gruvboxDarkRedBright
	theme.WarningColor = gruvboxDarkYellowBright
	theme.SuccessColor = gruvboxDarkGreenBright
	theme.InfoColor = gruvboxDarkBlueBright

	// Text colors
	theme.TextColor = gruvboxDarkFg1
	theme.TextMutedColor = gruvboxDarkFg4
	theme.TextEmphasizedColor = gruvboxDarkYellowBright

	// Background colors
	theme.BackgroundColor = gruvboxDarkBg0
	theme.BackgroundSecondaryColor = gruvboxDarkBg1
	theme.BackgroundDarkerColor = gruvboxDarkBg0Soft

	// Border colors
	theme.BorderNormalColor = gruvboxDarkBg2
	theme.BorderFocusedColor = gruvboxDarkBlueBright
	theme.BorderDimColor = gruvboxDarkBg1

	// Diff view colors
	theme.DiffAddedColor = gruvboxDarkGreenBright
	theme.DiffRemovedColor = gruvboxDarkRedBright
	theme.DiffContextColor = gruvboxDarkFg4
	theme.DiffHunkHeaderColor = gruvboxDarkFg3
	theme.DiffHighlightAddedColor = gruvboxDarkGreenBright
	theme.DiffHighlightRemovedColor = gruvboxDarkRedBright
	theme.DiffAddedBgColor = "#3C4C3C"
	theme.DiffRemovedBgColor = "#4C3C3C"
	theme.DiffContextBgColor = gruvboxDarkBg0
	theme.DiffLineNumberColor = gruvboxDarkFg4
	theme.DiffAddedLineNumberBgColor = "#32432F"
	theme.DiffRemovedLineNumberBgColor = "#43322F"

	// Markdown colors
	theme.MarkdownTextColor = gruvboxDarkFg1
	theme.MarkdownHeadingColor = gruvboxDarkYellowBright
	theme.MarkdownLinkColor = gruvboxDarkBlueBright
	theme.MarkdownLinkTextColor = gruvboxDarkAquaBright
	theme.MarkdownCodeColor = gruvboxDarkGreenBright
	theme.MarkdownBlockQuoteColor = gruvboxDarkAquaBright
	theme.MarkdownEmphColor = gruvboxDarkYellowBright
	theme.MarkdownStrongColor = gruvboxDarkOrangeBright
	theme.MarkdownHorizontalRuleColor = gruvboxDarkBg3
	theme.MarkdownListItemColor = gruvboxDarkBlueBright
	theme.MarkdownListEnumerationColor = gruvboxDarkBlueBright
	theme.MarkdownImageColor = gruvboxDarkPurpleBright
	theme.MarkdownImageTextColor = gruvboxDarkAquaBright
	theme.MarkdownCodeBlockColor = gruvboxDarkFg1

	// Syntax highlighting colors
	theme.SyntaxCommentColor = gruvboxDarkGray
	theme.SyntaxKeywordColor = gruvboxDarkRedBright
	theme.SyntaxFunctionColor = gruvboxDarkGreenBright
	theme.SyntaxVariableColor = gruvboxDarkBlueBright
	theme.SyntaxStringColor = gruvboxDarkYellowBright
	theme.SyntaxNumberColor = gruvboxDarkPurpleBright
	theme.SyntaxTypeColor = gruvboxDarkYellow
	theme.SyntaxOperatorColor = gruvboxDarkAquaBright
	theme.SyntaxPunctuationColor = gruvboxDarkFg1

	return theme
}

func init() {
	// Register the Gruvbox theme with the theme manager
	RegisterTheme("gruvbox", NewGruvboxTheme())
}
