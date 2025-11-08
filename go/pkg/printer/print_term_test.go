package printer

import (
	"bytes"
	"testing"

	"github.com/mojo-lang/mojo/packages/lang/go/pkg/mojo/lang"
	"github.com/stretchr/testify/assert"

	"github.com/mojo-lang/mojo/go/pkg/context"
)

func TestPrinter_PrintTerm(t *testing.T) {
	buffer := bytes.NewBuffer(nil)
	p := New(&Config{}, buffer)

	p.PrintTerm(context.Empty(), &lang.Term{Value: "type Int32"})
	assert.Equal(t, "type Int32", buffer.String())
}
