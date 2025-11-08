package lang

import (
	"testing"

	"github.com/mojo-lang/mojo/packages/core/go/pkg/mojo/core"
	"github.com/stretchr/testify/assert"
)

func TestNewNullLiteralExpressionFrom(t *testing.T) {
	expr := NewNullLiteralExpressionFrom()
	assert.NotNil(t, expr)
	assert.NotNil(t, expr.GetNullLiteralExpr())
}

func TestNewNullLiteralExpression(t *testing.T) {
	expr := NewNullLiteralExpression(&NullLiteralExpr{})
	assert.NotNil(t, expr)
	assert.NotNil(t, expr.GetNullLiteralExpr())
}

func TestNewBoolLiteralExpressionFrom(t *testing.T) {
	expr := NewBoolLiteralExpressionFrom(true)
	assert.NotNil(t, expr)

	v, err := expr.EvalBoolLiteral()
	assert.NoError(t, err)
	assert.True(t, v)
}

func TestNewBoolLiteralExpression(t *testing.T) {
	expr := NewBoolLiteralExpression(NewBoolLiteralExpr(true))
	assert.NotNil(t, expr)

	v, err := expr.EvalBoolLiteral()
	assert.NoError(t, err)
	assert.True(t, v)
}

func TestNewExpression(t *testing.T) {
	expr := NewIntegerLiteralExpressionFrom(100)
	assert.NotNil(t, expr)
	assert.NotNil(t, expr.GetIntegerLiteralExpr())
	assert.Equal(t, int64(100), expr.GetIntegerLiteralExpr().EvalValue())

	expr = NewIntegerLiteralExpression(NewIntegerLiteralExpr(100))
	assert.NotNil(t, expr)
	assert.NotNil(t, expr.GetIntegerLiteralExpr())
	assert.Equal(t, int64(100), expr.GetIntegerLiteralExpr().EvalValue())

	expr = NewFloatLiteralExpressionFrom(10.1)
	assert.NotNil(t, expr)
	assert.NotNil(t, expr.GetFloatLiteralExpr())
	assert.Equal(t, 10.1, expr.GetFloatLiteralExpr().EvalValue())

	expr = NewFloatLiteralExpression(NewFloatLiteralExpr(10.1))
	assert.NotNil(t, expr)
	assert.NotNil(t, expr.GetFloatLiteralExpr())
	assert.Equal(t, 10.1, expr.GetFloatLiteralExpr().EvalValue())

	expr = NewStringLiteralExpressionFrom("foo")
	assert.NotNil(t, expr)
	assert.NotNil(t, expr.GetStringLiteralExpr())
	assert.Equal(t, "foo", expr.GetStringLiteralExpr().EvalValue())

	expr = NewStringLiteralExpression(NewStringLiteralExpr("foo"))
	assert.NotNil(t, expr)
	assert.NotNil(t, expr.GetStringLiteralExpr())
	assert.Equal(t, "foo", expr.GetStringLiteralExpr().EvalValue())

	expr = NewRangeLiteralExpressionFrom(&core.IntRange{Start: 100, End: 200})
	assert.NotNil(t, expr)
	assert.NotNil(t, expr.GetRangeLiteralExpr())
	assert.Equal(t, int64(100), expr.GetRangeLiteralExpr().EvalValue().Start)

	expr = NewRangeLiteralExpression(NewRangeLiteralExpr(100, 200))
	assert.NotNil(t, expr)
	assert.NotNil(t, expr.GetRangeLiteralExpr())
	assert.Equal(t, int64(100), expr.GetRangeLiteralExpr().EvalValue().Start)
}
