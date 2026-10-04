package document

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewText(t *testing.T) {
	for _, text := range []string{"", "hello", "你好\nworld"} {
		assert.Equal(t, text, NewText(text).Val)
	}
}
