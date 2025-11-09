package syntax

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrayLiteralVisitor_VisitArrayLiteral(t *testing.T) {
	const arrayLiteral = `[1, 2, 3]`
	expr := parseExpression(t, arrayLiteral)

	arrayExpr := expr.GetArrayLiteralExpr()
	assert.NotNil(t, arrayExpr)
	assert.Equal(t, 3, len(arrayExpr.Elements))
	assert.Equal(t, uint64(1), arrayExpr.Elements[0].GetIntegerLiteralExpr().GetValue())
}
