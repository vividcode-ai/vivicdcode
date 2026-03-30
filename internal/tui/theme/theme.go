package theme

// Theme defines the interface for all UI themes in the application.
// All colors are defined as hex strings for compatibility with various
// terminal rendering libraries.
type Theme interface {
	// Base colors
	Primary() string
	Secondary() string
	Accent() string

	// Status colors
	Error() string
	Warning() string
	Success() string
	Info() string

	// Text colors
	Text() string
	TextMuted() string
	TextEmphasized() string

	// Background colors
	Background() string
	BackgroundSecondary() string
	BackgroundDarker() string

	// Border colors
	BorderNormal() string
	BorderFocused() string
	BorderDim() string

	// Diff view colors
	DiffAdded() string
	DiffRemoved() string
	DiffContext() string
	DiffHunkHeader() string
	DiffHighlightAdded() string
	DiffHighlightRemoved() string
	DiffAddedBg() string
	DiffRemovedBg() string
	DiffContextBg() string
	DiffLineNumber() string
	DiffAddedLineNumberBg() string
	DiffRemovedLineNumberBg() string

	// Markdown colors
	MarkdownText() string
	MarkdownHeading() string
	MarkdownLink() string
	MarkdownLinkText() string
	MarkdownCode() string
	MarkdownBlockQuote() string
	MarkdownEmph() string
	MarkdownStrong() string
	MarkdownHorizontalRule() string
	MarkdownListItem() string
	MarkdownListEnumeration() string
	MarkdownImage() string
	MarkdownImageText() string
	MarkdownCodeBlock() string

	// Syntax highlighting colors
	SyntaxComment() string
	SyntaxKeyword() string
	SyntaxFunction() string
	SyntaxVariable() string
	SyntaxString() string
	SyntaxNumber() string
	SyntaxType() string
	SyntaxOperator() string
	SyntaxPunctuation() string
}

// BaseTheme provides a default implementation of the Theme interface
// that can be embedded in concrete theme implementations.
type BaseTheme struct {
	// Base colors
	PrimaryColor   string
	SecondaryColor string
	AccentColor    string

	// Status colors
	ErrorColor   string
	WarningColor string
	SuccessColor string
	InfoColor    string

	// Text colors
	TextColor           string
	TextMutedColor      string
	TextEmphasizedColor string

	// Background colors
	BackgroundColor          string
	BackgroundSecondaryColor string
	BackgroundDarkerColor    string

	// Border colors
	BorderNormalColor  string
	BorderFocusedColor string
	BorderDimColor     string

	// Diff view colors
	DiffAddedColor               string
	DiffRemovedColor             string
	DiffContextColor             string
	DiffHunkHeaderColor          string
	DiffHighlightAddedColor      string
	DiffHighlightRemovedColor    string
	DiffAddedBgColor             string
	DiffRemovedBgColor           string
	DiffContextBgColor           string
	DiffLineNumberColor          string
	DiffAddedLineNumberBgColor   string
	DiffRemovedLineNumberBgColor string

	// Markdown colors
	MarkdownTextColor            string
	MarkdownHeadingColor         string
	MarkdownLinkColor            string
	MarkdownLinkTextColor        string
	MarkdownCodeColor            string
	MarkdownBlockQuoteColor      string
	MarkdownEmphColor            string
	MarkdownStrongColor          string
	MarkdownHorizontalRuleColor  string
	MarkdownListItemColor        string
	MarkdownListEnumerationColor string
	MarkdownImageColor           string
	MarkdownImageTextColor       string
	MarkdownCodeBlockColor       string

	// Syntax highlighting colors
	SyntaxCommentColor     string
	SyntaxKeywordColor     string
	SyntaxFunctionColor    string
	SyntaxVariableColor    string
	SyntaxStringColor      string
	SyntaxNumberColor      string
	SyntaxTypeColor        string
	SyntaxOperatorColor    string
	SyntaxPunctuationColor string
}

// Implement the Theme interface for BaseTheme
func (t *BaseTheme) Primary() string   { return t.PrimaryColor }
func (t *BaseTheme) Secondary() string { return t.SecondaryColor }
func (t *BaseTheme) Accent() string    { return t.AccentColor }

func (t *BaseTheme) Error() string   { return t.ErrorColor }
func (t *BaseTheme) Warning() string { return t.WarningColor }
func (t *BaseTheme) Success() string { return t.SuccessColor }
func (t *BaseTheme) Info() string    { return t.InfoColor }

func (t *BaseTheme) Text() string           { return t.TextColor }
func (t *BaseTheme) TextMuted() string      { return t.TextMutedColor }
func (t *BaseTheme) TextEmphasized() string { return t.TextEmphasizedColor }

func (t *BaseTheme) Background() string          { return t.BackgroundColor }
func (t *BaseTheme) BackgroundSecondary() string { return t.BackgroundSecondaryColor }
func (t *BaseTheme) BackgroundDarker() string    { return t.BackgroundDarkerColor }

func (t *BaseTheme) BorderNormal() string  { return t.BorderNormalColor }
func (t *BaseTheme) BorderFocused() string { return t.BorderFocusedColor }
func (t *BaseTheme) BorderDim() string     { return t.BorderDimColor }

func (t *BaseTheme) DiffAdded() string            { return t.DiffAddedColor }
func (t *BaseTheme) DiffRemoved() string          { return t.DiffRemovedColor }
func (t *BaseTheme) DiffContext() string          { return t.DiffContextColor }
func (t *BaseTheme) DiffHunkHeader() string       { return t.DiffHunkHeaderColor }
func (t *BaseTheme) DiffHighlightAdded() string   { return t.DiffHighlightAddedColor }
func (t *BaseTheme) DiffHighlightRemoved() string { return t.DiffHighlightRemovedColor }
func (t *BaseTheme) DiffAddedBg() string          { return t.DiffAddedBgColor }
func (t *BaseTheme) DiffRemovedBg() string        { return t.DiffRemovedBgColor }
func (t *BaseTheme) DiffContextBg() string        { return t.DiffContextBgColor }
func (t *BaseTheme) DiffLineNumber() string       { return t.DiffLineNumberColor }
func (t *BaseTheme) DiffAddedLineNumberBg() string {
	return t.DiffAddedLineNumberBgColor
}
func (t *BaseTheme) DiffRemovedLineNumberBg() string {
	return t.DiffRemovedLineNumberBgColor
}

func (t *BaseTheme) MarkdownText() string       { return t.MarkdownTextColor }
func (t *BaseTheme) MarkdownHeading() string    { return t.MarkdownHeadingColor }
func (t *BaseTheme) MarkdownLink() string       { return t.MarkdownLinkColor }
func (t *BaseTheme) MarkdownLinkText() string   { return t.MarkdownLinkTextColor }
func (t *BaseTheme) MarkdownCode() string       { return t.MarkdownCodeColor }
func (t *BaseTheme) MarkdownBlockQuote() string { return t.MarkdownBlockQuoteColor }
func (t *BaseTheme) MarkdownEmph() string       { return t.MarkdownEmphColor }
func (t *BaseTheme) MarkdownStrong() string     { return t.MarkdownStrongColor }
func (t *BaseTheme) MarkdownHorizontalRule() string {
	return t.MarkdownHorizontalRuleColor
}
func (t *BaseTheme) MarkdownListItem() string        { return t.MarkdownListItemColor }
func (t *BaseTheme) MarkdownListEnumeration() string { return t.MarkdownListEnumerationColor }
func (t *BaseTheme) MarkdownImage() string           { return t.MarkdownImageColor }
func (t *BaseTheme) MarkdownImageText() string       { return t.MarkdownImageTextColor }
func (t *BaseTheme) MarkdownCodeBlock() string       { return t.MarkdownCodeBlockColor }

func (t *BaseTheme) SyntaxComment() string     { return t.SyntaxCommentColor }
func (t *BaseTheme) SyntaxKeyword() string     { return t.SyntaxKeywordColor }
func (t *BaseTheme) SyntaxFunction() string    { return t.SyntaxFunctionColor }
func (t *BaseTheme) SyntaxVariable() string    { return t.SyntaxVariableColor }
func (t *BaseTheme) SyntaxString() string      { return t.SyntaxStringColor }
func (t *BaseTheme) SyntaxNumber() string      { return t.SyntaxNumberColor }
func (t *BaseTheme) SyntaxType() string        { return t.SyntaxTypeColor }
func (t *BaseTheme) SyntaxOperator() string    { return t.SyntaxOperatorColor }
func (t *BaseTheme) SyntaxPunctuation() string { return t.SyntaxPunctuationColor }
