package document

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewEmptyBlocks(t *testing.T) {
	// Empty and nil inputs must still select the requested union branch.
	for _, tt := range []struct {
		name    string
		block   *Block
		variant interface{}
	}{
		{"plain", NewPainBlock(), &Block_Plain{}},
		{"header", NewHeaderBlock(nil), &Block_Header{}},
		{"quote", NewQuoteBlockBlock(), &Block_QuoteBlock{}},
		{"code", NewCodeBlockBlock(nil), &Block_CodeBlock{}},
		{"ordered list", NewOrderedListBlock(nil), &Block_OrderedList{}},
		{"bullet list", NewBulletListBlock(nil), &Block_BulletList{}},
		{"paragraph", NewParagraphBlock(), &Block_Paragraph{}},
		{"line block", NewLineBlockBlock(), &Block_LineBlock{}},
		{"division", NewDivisionBlock(nil), &Block_Division{}},
		{"table", NewTableBlock(nil), &Block_Table{}},
	} {
		t.Run(tt.name, func(t *testing.T) { assert.IsType(t, tt.variant, tt.block.Block) })
	}
	assert.Empty(t, NewPainBlock().GetPlain().Inlines)
	assert.Empty(t, NewParagraphBlock().GetParagraph().Inlines)
	assert.Empty(t, NewQuoteBlockBlock().GetQuoteBlock().Blocks)
	assert.Empty(t, NewLineBlockBlock().GetLineBlock().Lines)
}

func TestNewTextPlainBlock(t *testing.T) {
	for _, text := range []string{"", "plain text", "你好\nworld"} {
		block := NewTextPlainBlock(text)
		require.NotNil(t, block.GetPlain())
		require.Len(t, block.GetPlain().Inlines, 1)
		assert.Equal(t, text, block.GetPlain().Inlines[0].GetText().GetVal())
	}
}
