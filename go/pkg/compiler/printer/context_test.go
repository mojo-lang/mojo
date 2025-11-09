package printer

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mojo-lang/mojo/go/pkg/compiler/context"
)

func TestContextColumns(t *testing.T) {
	columns := Columns{0, 6}
	ctx := WithColumns(context.Empty(), columns)

	out := ContextColumns(ctx)
	assert.Equal(t, Columns{0, 6}, out)
}
