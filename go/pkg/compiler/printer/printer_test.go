package printer

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrinter_PrintIndent(t *testing.T) {
	buffer := bytes.NewBuffer(nil)
	p := New(&Config{}, buffer)

	p.Indent().Indent().PrintIndent()
	assert.Equal(t, "        ", buffer.String())
}

func TestPrinter_PrintBlankLine(t *testing.T) {
	buffer := bytes.NewBuffer(nil)
	p := New(&Config{}, buffer)

	p.PrintRaw("foo").PrintBlankLine()
	assert.Equal(t, "foo\n\n", buffer.String())

	buffer2 := bytes.NewBuffer(nil)
	p.Reset(buffer2)
	p.PrintRaw("foo").BreakLine().BreakLine()
	assert.Equal(t, "foo\n\n", buffer.String())
}

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

func TestPrinter_PrintLine(t *testing.T) {
	buffer := bytes.NewBuffer(nil)
	p := New(&Config{}, buffer)

	p.PrintLine("foo")
	p.PrintLine("bar")
	assert.Equal(t, "foo\nbar", buffer.String())
}

func TestNew(t *testing.T) {
	buffer := bytes.NewBuffer(nil)
	p := New(&Config{IndentWidth: 2, Indent: 1}, buffer)

	p.PrintLine("foo")
	assert.Equal(t, "  foo", buffer.String())
}

func TestPrinter_PrintTo(t *testing.T) {
	buffer := bytes.NewBuffer(nil)
	p := New(&Config{IndentWidth: 2, Indent: 1}, buffer)

	p.PrintLine("foo")
	p.PrintTo(Cursor{
		Line:   2,
		Column: 4,
	}).PrintRaw("bar")

	assert.Equal(t, "  foo\n\n    bar", buffer.String())
}
