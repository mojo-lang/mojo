package markdown

import (
	"bytes"
	"strings"
	"testing"

	"github.com/mojo-lang/mojo/go/pkg/mojo/document"
	"github.com/stretchr/testify/require"
	"github.com/yuin/goldmark/ast"
	extast "github.com/yuin/goldmark/extension/ast"
	"github.com/yuin/goldmark/text"
)

func TestRenderTableMultilineCells(t *testing.T) {
	for _, source := range []string{
		"Before\n\n```text\nfirst | second\n  `<tag>` & value\nliteral \\| pipe\n\nlast\n```\n\nAfter",
		"Before\n\n    first | second\n      `<tag>` & value\n\nAfter",
		"- First item\n  continuation\n- Second item\n\n  ```text\n  a | b\n  c\n  ```",
	} {
		t.Run(source, func(t *testing.T) {
			md := New()
			description, err := md.Parse(source)
			require.NoError(t, err)
			doc := &document.Document{Blocks: []*document.Block{document.NewTableBlock(&document.Table{
				Header: document.NewTextTableHeader("field", "description | example"),
				Rows: []*document.Table_Row{
					{Vals: []*document.Table_Cell{document.NewTextTableCell("value"), document.NewTableCell(description.Blocks...)}},
					{Vals: []*document.Table_Cell{document.NewTextTableCell("next"), document.NewTextTableCell("intact")}},
				},
			})}}
			rendered, err := md.RenderToString(doc)
			require.NoError(t, err)
			require.Len(t, strings.Split(strings.TrimSuffix(rendered, "\n"), "\n"), 4, rendered)
			require.Contains(t, rendered, "<br>")
			parsed := md.MdEngine.Parser().Parse(text.NewReader([]byte(rendered)))
			require.NotNil(t, parsed.FirstChild())
			require.Nil(t, parsed.FirstChild().NextSibling(), rendered)
			table, ok := parsed.FirstChild().(*extast.Table)
			require.True(t, ok, rendered)
			require.Equal(t, 3, table.ChildCount(), rendered)
			for row := table.FirstChild(); row != nil; row = row.NextSibling() {
				require.Equal(t, 2, row.ChildCount(), rendered)
			}
			var html bytes.Buffer
			require.NoError(t, md.ConvertToHtml([]byte(rendered), &html))
			if strings.Contains(source, "<tag>") {
				require.Contains(t, html.String(), "<code>first | second</code>")
				require.Contains(t, html.String(), "`&lt;tag&gt;` &amp; value")
				if strings.Contains(source, "literal") {
					require.Contains(t, html.String(), "<code>literal \\| pipe</code>")
				}
			}
			// Cell rendering must not change ordinary fenced code output or retain state.
			outside, err := md.RenderToString(&document.Document{Blocks: []*document.Block{
				document.NewCodeBlockBlock(document.NewCodeBlock("text", "outside | table")),
			}})
			require.NoError(t, err)
			require.Equal(t, "```text\noutside | table\n```\n", outside)
		})
	}
}

func TestParseCodeBlocksPreservesLines(t *testing.T) {
	for _, source := range []string{"```text\nfirst\n  second\n```", "    first\n      second\n"} {
		doc, err := ParseMarkdown([]byte(source))
		require.NoError(t, err)
		require.Len(t, doc.Blocks, 1)
		require.Equal(t, "first\n  second", string(doc.Blocks[0].GetCodeBlock().GetCode()))
	}
}

func TestTableEscapedPipesAndLineEndings(t *testing.T) {
	doc := &document.Document{Blocks: []*document.Block{document.NewTableBlock(&document.Table{
		Header: document.NewTextTableHeader("description"),
		Rows:   []*document.Table_Row{{Vals: []*document.Table_Cell{document.NewTextTableCell("a | b \\| c\r\nd\re\nf")}}},
	})}}
	rendered, err := RenderToString(doc)
	require.NoError(t, err)
	require.Contains(t, rendered, "a \\| b \\| c<br>d<br>e<br>f")
	parsed := NewParser().MdEngine.Parser().Parse(text.NewReader([]byte(rendered)))
	require.NoError(t, ast.Walk(parsed, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
		if entering && n.Kind() == extast.KindTableRow {
			require.Equal(t, 1, n.ChildCount())
		}
		return ast.WalkContinue, nil
	}))
}
