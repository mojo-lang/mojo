package lang

import (
	"testing"

	"github.com/mojo-lang/mojo/packages/core/go/pkg/mojo/core"
	"github.com/stretchr/testify/assert"
)

func TestNewRangeLiteralExpr(t *testing.T) {
	r := NewRangeLiteralExpr(0, 100)
	assert.NotNil(t, r)
	assert.Equal(t, int64(0), r.Value.Start)
	assert.Equal(t, int64(100), r.Value.End)

	r = NewRangeLiteralExprFrom(&core.IntRange{Start: 0, End: 100})
	assert.Equal(t, int64(0), r.Value.Start)
	assert.Equal(t, int64(100), r.Value.End)
}
