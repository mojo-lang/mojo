package document

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLine(t *testing.T) {
	values := []*Inline{NewTextInline("first"), nil, NewCodeInlineFrom("second")}
	assert.Equal(t, values, NewLine(values...).Vals)
	assert.Empty(t, NewLine().Vals)
}
