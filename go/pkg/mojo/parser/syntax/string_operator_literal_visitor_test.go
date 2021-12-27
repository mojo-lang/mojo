package syntax

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringOperatorLiteralVisitor_VisitStringOperatorLiteral(t *testing.T) {
	const StringOperatorExpression = `r"^ab$"`
	expr := parseExpression(t, StringOperatorExpression)

	stringExpr := expr.GetStringLiteralUnaryExpr()
	assert.NotNil(t, stringExpr)
	assert.Equal(t, "r", stringExpr.Operator.Symbol)
	assert.Equal(t, "^ab$", stringExpr.Argument.GetStringLiteralExpr().Value)
}
