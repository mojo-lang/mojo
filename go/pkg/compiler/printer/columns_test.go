package printer

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestColumns_PrintTo(t *testing.T) {
	columns := Columns{0, 5}
	buffer := bytes.NewBuffer(nil)
	p := New(&Config{}, buffer)

	p.PrintRaw("foo")
	columns.PrintTo(1, p).PrintRaw("bar")

	assert.Equal(t, "foo  bar", buffer.String())
}

func TestColumns_Column(t *testing.T) {
	columns := Columns{0, 5}
	assert.Equal(t, 0, columns.Column(-1))
	assert.Equal(t, 5, columns.Column(3))
}
