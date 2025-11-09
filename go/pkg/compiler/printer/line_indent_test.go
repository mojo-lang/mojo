package printer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLineIndent_IndentSpace(t *testing.T) {
	space := (&lineIndent{Width: 4}).Outdent().Indent().Indent().Outdent().IndentSpace()
	assert.Equal(t, []byte("    "), space)
}
