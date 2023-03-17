package printer

import (
	"bytes"
	"testing"

	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/stretchr/testify/assert"

	"github.com/mojo-lang/mojo/go/pkg/context"
)

func TestPrinter_PrintComments(t *testing.T) {
	buffer := bytes.NewBuffer(nil)
	p := New(&Config{}, buffer)

	p.PrintComments(context.Empty(), &lang.Comment{
		Comment: &lang.Comment_MultiLineComment{
			MultiLineComment: &lang.MultiLineComment{
				Lines: []*lang.LineComment{
					{Content: "foo"},
				},
			},
		},
	})
	assert.Equal(t, "// foo\n", buffer.String())
}
