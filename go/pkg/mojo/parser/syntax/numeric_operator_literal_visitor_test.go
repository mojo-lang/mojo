package syntax

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mojo-lang/mojo/go/pkg/context"
)

func TestNumericOperatorLiteralVisitor_VisitNumericOperatorLiteral(t *testing.T) {
	const NumericOperatorExpression = `12s`
	expr := parseExpression(t, NumericOperatorExpression)

	unaryExpr := expr.GetNumericSuffixLiteralExpr()
	if assert.NotNil(t, unaryExpr) {
		assert.Equal(t, "s", unaryExpr.GetOperator().Symbol)
		assert.Equal(t, uint64(12), unaryExpr.Argument.GetIntegerLiteralExpr().Value)
	}
}

func TestNumericOperatorLiteralVisitor_VisitNumericOperatorLiteral2(t *testing.T) {
	const NumericOperatorExpression = `12.122s`
	expr := parseExpression(t, NumericOperatorExpression)

	unaryExpr := expr.GetNumericSuffixLiteralExpr()
	if assert.NotNil(t, unaryExpr) {
		assert.Equal(t, "s", unaryExpr.GetOperator().Symbol)
		assert.Equal(t, 12.122, unaryExpr.Argument.GetFloatLiteralExpr().Value)
	}
}

func TestNumericOperatorLiteralVisitor_VisitNumericOperatorLiteral3(t *testing.T) {
	const NumericOperatorExpression = `12 s`
	parser := &Parser{}
	_, err := parser.ParseString(context.Empty(), NumericOperatorExpression)
	assert.Error(t, err)
}
