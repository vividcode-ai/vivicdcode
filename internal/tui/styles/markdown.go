package styles

import (
	"charm.land/glamour/v2"
	"charm.land/glamour/v2/ansi"
	"github.com/vividcode-ai/vividcode/internal/tui/theme"
)

const defaultMargin = 1

func boolPtr(b bool) *bool       { return &b }
func stringPtr(s string) *string { return &s }
func uintPtr(u uint) *uint       { return &u }

func GetMarkdownRenderer(width int) *glamour.TermRenderer {
	r, _ := glamour.NewTermRenderer(
		glamour.WithStyles(generateMarkdownStyleConfig()),
		glamour.WithWordWrap(width),
	)
	return r
}

func generateMarkdownStyleConfig() ansi.StyleConfig {
	t := theme.CurrentTheme()

	return ansi.StyleConfig{
		Document: ansi.StyleBlock{
			StylePrimitive: ansi.StylePrimitive{
				BlockPrefix: "",
				BlockSuffix: "",
				Color:       stringPtr(t.MarkdownText()),
			},
			Margin: uintPtr(defaultMargin),
		},
		BlockQuote: ansi.StyleBlock{
			StylePrimitive: ansi.StylePrimitive{
				Color:  stringPtr(t.MarkdownBlockQuote()),
				Italic: boolPtr(true),
				Prefix: "┃ ",
			},
			Indent:      uintPtr(1),
			IndentToken: stringPtr(BaseStyle().Render(" ")),
		},
		List: ansi.StyleList{
			LevelIndent: defaultMargin,
			StyleBlock: ansi.StyleBlock{
				IndentToken: stringPtr(BaseStyle().Render(" ")),
				StylePrimitive: ansi.StylePrimitive{
					Color: stringPtr(t.MarkdownText()),
				},
			},
		},
		Heading: ansi.StyleBlock{
			StylePrimitive: ansi.StylePrimitive{
				BlockSuffix: "\n",
				Color:       stringPtr(t.MarkdownHeading()),
				Bold:        boolPtr(true),
			},
		},
		H1: ansi.StyleBlock{
			StylePrimitive: ansi.StylePrimitive{
				Prefix: "# ",
				Color:  stringPtr(t.MarkdownHeading()),
				Bold:   boolPtr(true),
			},
		},
		H2: ansi.StyleBlock{
			StylePrimitive: ansi.StylePrimitive{
				Prefix: "## ",
				Color:  stringPtr(t.MarkdownHeading()),
				Bold:   boolPtr(true),
			},
		},
		H3: ansi.StyleBlock{
			StylePrimitive: ansi.StylePrimitive{
				Prefix: "### ",
				Color:  stringPtr(t.MarkdownHeading()),
				Bold:   boolPtr(true),
			},
		},
		H4: ansi.StyleBlock{
			StylePrimitive: ansi.StylePrimitive{
				Prefix: "#### ",
				Color:  stringPtr(t.MarkdownHeading()),
				Bold:   boolPtr(true),
			},
		},
		H5: ansi.StyleBlock{
			StylePrimitive: ansi.StylePrimitive{
				Prefix: "##### ",
				Color:  stringPtr(t.MarkdownHeading()),
				Bold:   boolPtr(true),
			},
		},
		H6: ansi.StyleBlock{
			StylePrimitive: ansi.StylePrimitive{
				Prefix: "###### ",
				Color:  stringPtr(t.MarkdownHeading()),
				Bold:   boolPtr(true),
			},
		},
		Strikethrough: ansi.StylePrimitive{
			CrossedOut: boolPtr(true),
			Color:      stringPtr(t.TextMuted()),
		},
		Emph: ansi.StylePrimitive{
			Color:  stringPtr(t.MarkdownEmph()),
			Italic: boolPtr(true),
		},
		Strong: ansi.StylePrimitive{
			Bold:  boolPtr(true),
			Color: stringPtr(t.MarkdownStrong()),
		},
		HorizontalRule: ansi.StylePrimitive{
			Color:  stringPtr(t.MarkdownHorizontalRule()),
			Format: "\n─────────────────────────────────────────\n",
		},
		Item: ansi.StylePrimitive{
			BlockPrefix: "• ",
			Color:       stringPtr(t.MarkdownListItem()),
		},
		Enumeration: ansi.StylePrimitive{
			BlockPrefix: ". ",
			Color:       stringPtr(t.MarkdownListEnumeration()),
		},
		Task: ansi.StyleTask{
			StylePrimitive: ansi.StylePrimitive{},
			Ticked:         "[✓] ",
			Unticked:       "[ ] ",
		},
		Link: ansi.StylePrimitive{
			Color:     stringPtr(t.MarkdownLink()),
			Underline: boolPtr(true),
		},
		LinkText: ansi.StylePrimitive{
			Color: stringPtr(t.MarkdownLinkText()),
			Bold:  boolPtr(true),
		},
		Image: ansi.StylePrimitive{
			Color:     stringPtr(t.MarkdownImage()),
			Underline: boolPtr(true),
			Format:    "🖼 {{.text}}",
		},
		ImageText: ansi.StylePrimitive{
			Color:  stringPtr(t.MarkdownImageText()),
			Format: "{{.text}}",
		},
		Code: ansi.StyleBlock{
			StylePrimitive: ansi.StylePrimitive{
				Color:  stringPtr(t.MarkdownCode()),
				Prefix: "",
				Suffix: "",
			},
		},
		CodeBlock: ansi.StyleCodeBlock{
			StyleBlock: ansi.StyleBlock{
				StylePrimitive: ansi.StylePrimitive{
					Prefix: " ",
					Color:  stringPtr(t.MarkdownCodeBlock()),
				},
				Margin: uintPtr(defaultMargin),
			},
			Chroma: &ansi.Chroma{
				Text: ansi.StylePrimitive{
					Color: stringPtr(t.MarkdownText()),
				},
				Error: ansi.StylePrimitive{
					Color: stringPtr(t.Error()),
				},
				Comment: ansi.StylePrimitive{
					Color: stringPtr(t.SyntaxComment()),
				},
				CommentPreproc: ansi.StylePrimitive{
					Color: stringPtr(t.SyntaxKeyword()),
				},
				Keyword: ansi.StylePrimitive{
					Color: stringPtr(t.SyntaxKeyword()),
				},
				KeywordReserved: ansi.StylePrimitive{
					Color: stringPtr(t.SyntaxKeyword()),
				},
				KeywordNamespace: ansi.StylePrimitive{
					Color: stringPtr(t.SyntaxKeyword()),
				},
				KeywordType: ansi.StylePrimitive{
					Color: stringPtr(t.SyntaxType()),
				},
				Operator: ansi.StylePrimitive{
					Color: stringPtr(t.SyntaxOperator()),
				},
				Punctuation: ansi.StylePrimitive{
					Color: stringPtr(t.SyntaxPunctuation()),
				},
				Name: ansi.StylePrimitive{
					Color: stringPtr(t.SyntaxVariable()),
				},
				NameBuiltin: ansi.StylePrimitive{
					Color: stringPtr(t.SyntaxVariable()),
				},
				NameTag: ansi.StylePrimitive{
					Color: stringPtr(t.SyntaxKeyword()),
				},
				NameAttribute: ansi.StylePrimitive{
					Color: stringPtr(t.SyntaxFunction()),
				},
				NameClass: ansi.StylePrimitive{
					Color: stringPtr(t.SyntaxType()),
				},
				NameConstant: ansi.StylePrimitive{
					Color: stringPtr(t.SyntaxVariable()),
				},
				NameDecorator: ansi.StylePrimitive{
					Color: stringPtr(t.SyntaxFunction()),
				},
				NameFunction: ansi.StylePrimitive{
					Color: stringPtr(t.SyntaxFunction()),
				},
				LiteralNumber: ansi.StylePrimitive{
					Color: stringPtr(t.SyntaxNumber()),
				},
				LiteralString: ansi.StylePrimitive{
					Color: stringPtr(t.SyntaxString()),
				},
				LiteralStringEscape: ansi.StylePrimitive{
					Color: stringPtr(t.SyntaxKeyword()),
				},
				GenericDeleted: ansi.StylePrimitive{
					Color: stringPtr(t.DiffRemoved()),
				},
				GenericEmph: ansi.StylePrimitive{
					Color:  stringPtr(t.MarkdownEmph()),
					Italic: boolPtr(true),
				},
				GenericInserted: ansi.StylePrimitive{
					Color: stringPtr(t.DiffAdded()),
				},
				GenericStrong: ansi.StylePrimitive{
					Color: stringPtr(t.MarkdownStrong()),
					Bold:  boolPtr(true),
				},
				GenericSubheading: ansi.StylePrimitive{
					Color: stringPtr(t.MarkdownHeading()),
				},
			},
		},
		Table: ansi.StyleTable{
			StyleBlock: ansi.StyleBlock{
				StylePrimitive: ansi.StylePrimitive{
					BlockPrefix: "\n",
					BlockSuffix: "\n",
				},
			},
			CenterSeparator: stringPtr("┼"),
			ColumnSeparator: stringPtr("│"),
			RowSeparator:    stringPtr("─"),
		},
		DefinitionDescription: ansi.StylePrimitive{
			BlockPrefix: "\n ❯ ",
			Color:       stringPtr(t.MarkdownLinkText()),
		},
		Text: ansi.StylePrimitive{
			Color: stringPtr(t.MarkdownText()),
		},
		Paragraph: ansi.StyleBlock{
			StylePrimitive: ansi.StylePrimitive{
				Color: stringPtr(t.MarkdownText()),
			},
		},
	}
}
