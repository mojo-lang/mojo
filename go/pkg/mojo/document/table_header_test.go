package document

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewTextTableHeader(t *testing.T) {
	assert.Empty(t, NewTextTableHeader().Vals)
	texts := []string{"first", "", "第三列"}
	header := NewTextTableHeader(texts...)
	require.Len(t, header.Vals, len(texts))
	for i, cell := range header.Vals {
		require.Len(t, cell.Vals, 1)
		plain := cell.Vals[0].GetPlain()
		require.NotNil(t, plain)
		require.Len(t, plain.Inlines, 1)
		assert.Equal(t, texts[i], plain.Inlines[0].GetText().GetVal())
	}
	header.Vals[0].Vals[0].GetPlain().Inlines[0].SetText("changed")
	assert.Equal(t, "", header.Vals[1].Vals[0].GetPlain().Inlines[0].GetText().GetVal())
}
