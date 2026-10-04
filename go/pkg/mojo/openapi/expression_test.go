package openapi

import (
	"testing"

	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExpressionFromValue(t *testing.T) {
	expression, err := ExpressionFromValue(core.NewStringValue("$request.query.id"))
	require.NoError(t, err)
	assert.Equal(t, "$request.query.id", expression.ToString())
	expression, err = ExpressionFromValue(core.NewStringValue("invalid"))
	assert.Error(t, err)
	assert.Nil(t, expression)
	for _, value := range []*core.Value{nil, core.NewIntValue(1), core.NewBoolValue(true), core.NewNullValue()} {
		expression, err = ExpressionFromValue(value)
		require.NoError(t, err)
		assert.Equal(t, Expression_TYPE_UNSPECIFIED, expression.Type)
		old := &Expression{Type: Expression_TYPE_URL}
		require.NoError(t, old.FromValue(value))
		assert.Equal(t, Expression_TYPE_URL, old.Type)
	}
	assert.NoError(t, (*Expression)(nil).FromValue(core.NewStringValue("$url")))
}
