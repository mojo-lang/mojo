package document

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewTableCell(t *testing.T) {
	blocks := []*Block{NewTextPlainBlock("first"), nil, NewHeaderBlock(NewHeader(2))}
	assert.Equal(t, blocks, NewTableCell(blocks...).Vals)
	assert.Empty(t, NewTableCell().Vals)
}

func TestNewTextTableCell(t *testing.T) {
	for _, text := range []string{"", "cell", "你好\nworld"} {
		cell := NewTextTableCell(text)
		require.Len(t, cell.Vals, 1)
		plain := cell.Vals[0].GetPlain()
		require.NotNil(t, plain)
		require.Len(t, plain.Inlines, 1)
		assert.Equal(t, text, plain.Inlines[0].GetText().GetVal())
	}
}
