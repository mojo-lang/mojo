package printer

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstLineNoneBreaker_Break(t *testing.T) {
	buffer := bytes.NewBuffer(nil)
	p := New(&Config{}, buffer)

	p.PrintLine("foo")
	breaker := NewFirstLineNoneBreaker(p)
	breaker.Break()
	assert.Equal(t, "foo", buffer.String())

	p.PrintLine("bar")
	breaker.Break()
	assert.Equal(t, "foo\nbar\n", buffer.String())
}
