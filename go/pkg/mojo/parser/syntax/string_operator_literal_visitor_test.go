package syntax

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringOperatorLiteralVisitor_VisitStringOperatorLiteral(t *testing.T) {
	const StringOperatorExpression = `r"^ab$"`

	parser := &Parser{}
	file, err := parser.ParseString(StringOperatorExpression)

	assert.NoError(t, err)
	expr := getExpression(file)
	assert.NotNil(t, expr)

	stringExpr := expr.GetStringLiteralUnaryExpr()
	assert.NotNil(t, stringExpr)
	assert.Equal(t, "r", stringExpr.Operator)
	assert.Equal(t, "^ab$", stringExpr.Expression.GetStringLiteralExpr().Value)
}
