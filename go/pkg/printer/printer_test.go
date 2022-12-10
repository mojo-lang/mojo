package printer

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrinter_PrintRaw1(t *testing.T) {
	buffer := bytes.NewBuffer(nil)
	p := New(&Config{}, buffer)

	v := "foo"
	p.PrintRaw(&v)

	out := buffer.String()
	assert.Equal(t, "foo", out)
}

func TestPrinter_PrintRaw2(t *testing.T) {
	buffer := bytes.NewBuffer(nil)
	p := New(&Config{}, buffer)

	sv := "foo"
	iv := 123
	p.PrintRaw(&sv, " ", &iv, " ", 456, " ", true)

	out := buffer.String()
	assert.Equal(t, "foo 123 456 true", out)
}
