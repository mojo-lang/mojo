package syntax

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClosureExprVisitor_VisitClosureExpression(t *testing.T) {
	const exprStr = `{x,y -> x + y}`

	expr := parseExpression(t, exprStr)
	closure := expr.GetClosureExpr()
	assert.Equal(t, "x", closure.Signature.GetParameters()[0].Name)
	assert.Equal(t, "y", closure.Signature.GetParameters()[1].Name)
	assert.Equal(t, 1, len(closure.Body.Statements))

	closureExpr := closure.Body.Statements[0].GetExpression().GetBinaryExpr()
	assert.Equal(t, "+", closureExpr.GetOperator().Symbol)
}
