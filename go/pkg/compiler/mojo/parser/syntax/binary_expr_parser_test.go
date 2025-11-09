package syntax

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinaryExprParser_Parse(t *testing.T) {
	const exprStr = `2 * 4 + 7 * 8`

	expr := parseExpression(t, exprStr)
	binary := expr.GetBinaryExpr()
	assert.Equal(t, "+", binary.GetOperator().Symbol)
}
