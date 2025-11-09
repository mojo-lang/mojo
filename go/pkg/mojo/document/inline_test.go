package document

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInline_SetText(t *testing.T) {
	foo := NewTextInline("").SetText("foo")
	assert.Equal(t, "foo", foo.GetText().Val)
}
