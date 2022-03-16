package syntax

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestNumericOperatorLiteralVisitor_VisitNumericOperatorLiteral(t *testing.T) {
    const NumericOperatorExpression = `12s`
    expr := parseExpression(t, NumericOperatorExpression)

    unaryExpr := expr.GetNumericLiteralUnaryExpr()
    assert.NotNil(t, unaryExpr)
    assert.Equal(t, "s", unaryExpr.Operator.Symbol)
    assert.Equal(t, uint64(12), unaryExpr.Argument.GetIntegerLiteralExpr().Value)
}

func TestNumericOperatorLiteralVisitor_VisitNumericOperatorLiteral2(t *testing.T) {
    const NumericOperatorExpression = `12.122s`
    expr := parseExpression(t, NumericOperatorExpression)

    unaryExpr := expr.GetNumericLiteralUnaryExpr()
    assert.NotNil(t, unaryExpr)
    assert.Equal(t, "s", unaryExpr.Operator.Symbol)
    assert.Equal(t, 12.122, unaryExpr.Argument.GetFloatLiteralExpr().Value)
}
