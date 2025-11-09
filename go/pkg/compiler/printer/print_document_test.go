package printer

import (
	"bytes"
	"testing"

	"github.com/mojo-lang/mojo/go/pkg/mojo/lang"
	"github.com/stretchr/testify/assert"

	"github.com/mojo-lang/mojo/go/pkg/compiler/context"
)

func TestPrinter_PrintDocument(t *testing.T) {
	buffer := bytes.NewBuffer(nil)
	p := New(&Config{}, buffer)

	p.PrintDocument(context.Empty(), &lang.Document{Lines: []*lang.Document_Line{{Content: "foo"}}})
	assert.Equal(t, "/// foo\n", buffer.String())
}
