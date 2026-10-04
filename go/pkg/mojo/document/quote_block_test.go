package document

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewQuoteBlock(t *testing.T) {
	blocks := []*Block{NewTextPlainBlock("first"), nil, NewParagraphBlock(NewTextInline("second"))}
	assert.Equal(t, blocks, NewQuoteBlock(blocks...).Blocks)
	assert.Empty(t, NewQuoteBlock().Blocks)
}
