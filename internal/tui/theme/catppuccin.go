package theme

import (
	catppuccin "github.com/catppuccin/go"
)

// CatppuccinTheme implements the Theme interface with Catppuccin colors.
// It provides both dark (Mocha) and light (Latte) variants.
type CatppuccinTheme struct {
	BaseTheme
}

// NewCatppuccinTheme creates a new instance of the Catppuccin theme.
func NewCatppuccinTheme() *CatppuccinTheme {
	mocha := catppuccin.Mocha

	theme := &CatppuccinTheme{}

	// Base colors
	theme.PrimaryColor = mocha.Blue().Hex
	theme.SecondaryColor = mocha.Mauve().Hex
	theme.AccentColor = mocha.Peach().Hex

	// Status colors
	theme.ErrorColor = mocha.Red().Hex
	theme.WarningColor = mocha.Peach().Hex
	theme.SuccessColor = mocha.Green().Hex
	theme.InfoColor = mocha.Blue().Hex

	// Text colors
	theme.TextColor = mocha.Text().Hex
	theme.TextMutedColor = mocha.Subtext0().Hex
	theme.TextEmphasizedColor = mocha.Lavender().Hex

	// Background colors
	theme.BackgroundColor = "#212121"
	theme.BackgroundSecondaryColor = "#2c2c2c"
	theme.BackgroundDarkerColor = "#181818"

	// Border colors
	theme.BorderNormalColor = "#4b4c5c"
	theme.BorderFocusedColor = mocha.Blue().Hex
	theme.BorderDimColor = mocha.Surface0().Hex

	// Diff view colors
	theme.DiffAddedColor = "#478247"
	theme.DiffRemovedColor = "#7C4444"
	theme.DiffContextColor = "#a0a0a0"
	theme.DiffHunkHeaderColor = "#a0a0a0"
	theme.DiffHighlightAddedColor = "#DAFADA"
	theme.DiffHighlightRemovedColor = "#FADADD"
	theme.DiffAddedBgColor = "#303A30"
	theme.DiffRemovedBgColor = "#3A3030"
	theme.DiffContextBgColor = "#212121"
	theme.DiffLineNumberColor = "#888888"
	theme.DiffAddedLineNumberBgColor = "#293229"
	theme.DiffRemovedLineNumberBgColor = "#332929"

	// Markdown colors
	theme.MarkdownTextColor = mocha.Text().Hex
	theme.MarkdownHeadingColor = mocha.Mauve().Hex
	theme.MarkdownLinkColor = mocha.Sky().Hex
	theme.MarkdownLinkTextColor = mocha.Pink().Hex
	theme.MarkdownCodeColor = mocha.Green().Hex
	theme.MarkdownBlockQuoteColor = mocha.Yellow().Hex
	theme.MarkdownEmphColor = mocha.Yellow().Hex
	theme.MarkdownStrongColor = mocha.Peach().Hex
	theme.MarkdownHorizontalRuleColor = mocha.Overlay0().Hex
	theme.MarkdownListItemColor = mocha.Blue().Hex
	theme.MarkdownListEnumerationColor = mocha.Sky().Hex
	theme.MarkdownImageColor = mocha.Sapphire().Hex
	theme.MarkdownImageTextColor = mocha.Pink().Hex
	theme.MarkdownCodeBlockColor = mocha.Text().Hex

	// Syntax highlighting colors
	theme.SyntaxCommentColor = mocha.Overlay1().Hex
	theme.SyntaxKeywordColor = mocha.Pink().Hex
	theme.SyntaxFunctionColor = mocha.Green().Hex
	theme.SyntaxVariableColor = mocha.Sky().Hex
	theme.SyntaxStringColor = mocha.Yellow().Hex
	theme.SyntaxNumberColor = mocha.Teal().Hex
	theme.SyntaxTypeColor = mocha.Sky().Hex
	theme.SyntaxOperatorColor = mocha.Pink().Hex
	theme.SyntaxPunctuationColor = mocha.Text().Hex

	return theme
}

func init() {
	// Register the Catppuccin theme with the theme manager
	RegisterTheme("catppuccin", NewCatppuccinTheme())
}
