package document

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDocumentAppendBlocks(t *testing.T) {
	var doc Document
	first, second, third := NewTextPlainBlock("first"), NewTextPlainBlock("second"), NewTextPlainBlock("third")
	doc.AppendBlock(first)
	doc.AppendBlocks(second, third)
	doc.AppendBlocks()
	require.Len(t, doc.Blocks, 3)
	assert.Same(t, first, doc.Blocks[0])
	assert.Same(t, second, doc.Blocks[1])
	assert.Same(t, third, doc.Blocks[2])
	assert.NotPanics(t, func() {
		(*Document)(nil).AppendBlock(first)
		(*Document)(nil).AppendBlocks(second, third)
	})
}

func TestDocumentAppendContent(t *testing.T) {
	text := NewTextInline("content")
	header := NewHeader(2, text)
	quoted := NewTextPlainBlock("quoted")
	code := NewCodeBlock("go", "first", "second")
	ordered, bullet := &OrderedList{}, &BulletList{}
	line, division, table := NewLine(text), NewTocDivision(), &Table{Header: NewTextTableHeader("column")}
	for _, tt := range []struct {
		name   string
		append func(*Document)
		check  func(*testing.T, *Block)
	}{
		{"plain", func(d *Document) { d.AppendPain(text) }, func(t *testing.T, b *Block) {
			require.NotNil(t, b.GetPlain())
			assert.Equal(t, []*Inline{text}, b.GetPlain().Inlines)
		}},
		{"header", func(d *Document) { d.AppendHeader(header) }, func(t *testing.T, b *Block) { assert.Same(t, header, b.GetHeader()) }},
		{"header from inlines", func(d *Document) { d.AppendHeaderFrom(3, text) }, func(t *testing.T, b *Block) {
			require.NotNil(t, b.GetHeader())
			assert.Equal(t, int64(3), b.GetHeader().Level)
			assert.Equal(t, []*Inline{text}, b.GetHeader().Text)
		}},
		{"header from text", func(d *Document) { d.AppendHeaderFromText(4, "title") }, func(t *testing.T, b *Block) {
			h := b.GetHeader()
			require.NotNil(t, h)
			assert.Equal(t, int64(4), h.Level)
			require.Len(t, h.Text, 1)
			assert.Equal(t, "title", h.Text[0].GetText().GetVal())
		}},
		{"quote", func(d *Document) { d.AppendBlockQuote(quoted) }, func(t *testing.T, b *Block) {
			require.NotNil(t, b.GetQuoteBlock())
			assert.Equal(t, []*Block{quoted}, b.GetQuoteBlock().Blocks)
		}},
		{"code", func(d *Document) { d.AppendCodeBlock(code) }, func(t *testing.T, b *Block) { assert.Same(t, code, b.GetCodeBlock()) }},
		{"code from lines", func(d *Document) { d.AppendCodeBlockFrom("go", "first\nsecond") }, func(t *testing.T, b *Block) {
			require.NotNil(t, b.GetCodeBlock())
			assert.Equal(t, "go", b.GetCodeBlock().Language)
			assert.Equal(t, "first\nsecond", string(b.GetCodeBlock().GetCode()))
		}},
		{"ordered list", func(d *Document) { d.AppendOrderedList(ordered) }, func(t *testing.T, b *Block) { assert.Same(t, ordered, b.GetOrderedList()) }},
		{"bullet list", func(d *Document) { d.AppendBulletList(bullet) }, func(t *testing.T, b *Block) { assert.Same(t, bullet, b.GetBulletList()) }},
		{"paragraph", func(d *Document) { d.AppendParagraph(text) }, func(t *testing.T, b *Block) {
			require.NotNil(t, b.GetParagraph())
			assert.Equal(t, []*Inline{text}, b.GetParagraph().Inlines)
		}},
		{"line block", func(d *Document) { d.AppendLineBlock(line) }, func(t *testing.T, b *Block) {
			require.NotNil(t, b.GetLineBlock())
			assert.Equal(t, []*Line{line}, b.GetLineBlock().Lines)
		}},
		{"division", func(d *Document) { d.AppendDivision(division) }, func(t *testing.T, b *Block) { assert.Same(t, division, b.GetDivision()) }},
		{"table", func(d *Document) { d.AppendTable(table) }, func(t *testing.T, b *Block) { assert.Same(t, table, b.GetTable()) }},
	} {
		t.Run(tt.name, func(t *testing.T) {
			prefix := NewTextPlainBlock("prefix")
			doc := &Document{Blocks: []*Block{prefix}}
			tt.append(doc)
			require.Len(t, doc.Blocks, 2)
			assert.Same(t, prefix, doc.Blocks[0])
			tt.check(t, doc.Blocks[1])
			assert.NotPanics(t, func() { tt.append(nil) })
		})
	}
}
