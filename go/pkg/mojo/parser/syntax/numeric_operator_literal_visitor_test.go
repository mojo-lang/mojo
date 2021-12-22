package syntax

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumericOperatorLiteralVisitor_VisitNumericOperatorLiteral(t *testing.T) {
	const NumericOperatorExpression = `12s`

	parser := &Parser{}
	file, err := parser.ParseString(NumericOperatorExpression)

	assert.NoError(t, err)
	expr := getExpression(file)
	assert.NotNil(t, expr)

	unaryExpr := expr.GetNumericLiteralUnaryExpr()
	assert.NotNil(t, unaryExpr)
	assert.Equal(t, "s", unaryExpr.Operator)
	assert.Equal(t, int64(12), unaryExpr.Expression.GetIntegerLiteralExpr().Value)
}

func TestNumericOperatorLiteralVisitor_VisitNumericOperatorLiteral2(t *testing.T) {
	const NumericOperatorExpression = `12.122s`

	parser := &Parser{}
	file, err := parser.ParseString(NumericOperatorExpression)

	assert.NoError(t, err)
	expr := getExpression(file)
	assert.NotNil(t, expr)

	unaryExpr := expr.GetNumericLiteralUnaryExpr()
	assert.NotNil(t, unaryExpr)
	assert.Equal(t, "s", unaryExpr.Operator)
	assert.Equal(t, 12.122, unaryExpr.Expression.GetFloatLiteralExpr().Value)
}
