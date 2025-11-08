package lang

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewArguments(t *testing.T) {
	arguments := NewArguments(NewStringLiteralExpressionFrom("foo"), NewStringLiteralExpressionFrom("bar"))
	assert.Equal(t, 2, len(arguments))
}
