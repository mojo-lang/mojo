package syntax

import (
	"testing"

	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/stretchr/testify/assert"
)

func TestExpressionVisitor_VisitExpression_Null(t *testing.T) {
	const nullLiteral = `null`
	expr := parseExpression(t, nullLiteral)

	assert.NotNil(t, expr.GetNullLiteralExpr())
}

func TestExpressionVisitor_VisitExpression_False(t *testing.T) {
	const falseLiteral = `false`
	expr := parseExpression(t, falseLiteral)

	assert.NotNil(t, expr.GetBoolLiteralExpr())
	assert.Equal(t, false, expr.GetBoolLiteralExpr().Value)
}

func TestExpressionVisitor_VisitExpression_True(t *testing.T) {
	const trueLiteral = `true`
	expr := parseExpression(t, trueLiteral)

	assert.NotNil(t, expr.GetBoolLiteralExpr())
	assert.Equal(t, true, expr.GetBoolLiteralExpr().Value)
}

func TestExpressionVisitor_VisitExpression_Wildcard(t *testing.T) {
	const wildcard = `_`
	expr := parseExpression(t, wildcard)

	assert.NotNil(t, expr.GetWildcardExpr())
}

func TestExpressionVisitor_VisitExpression_String(t *testing.T) {
	stringCases := map[string]string{
		`"testdata"`: `testdata`,
		`'testdata'`: `testdata`,
	}

	for k, v := range stringCases {
		expr := parseExpression(t, k)

		assert.NotNil(t, expr.GetStringLiteralExpr())
		assert.Equal(t, v, expr.GetStringLiteralExpr().Value)
	}
}

func TestExpressionVisitor_VisitExpression_Integer(t *testing.T) {
	integerCases := map[string]struct {
		IsNegative bool
		Value      uint64
	}{
		"1234":  {IsNegative: false, Value: 1234},
		"-1234": {IsNegative: true, Value: 1234},
	}

	for k, v := range integerCases {
		expr := parseExpression(t, k)

		integerExpr := expr.GetIntegerLiteralExpr()
		assert.NotNil(t, integerExpr)
		assert.Equal(t, v.IsNegative, integerExpr.IsNegative)
		assert.Equal(t, v.Value, integerExpr.Value)
	}
}

func TestExpressionVisitor_VisitExpression_Float(t *testing.T) {
	floatCases := map[string]struct {
		IsNegative bool
		Value      float64
	}{
		"0.1234":  {IsNegative: false, Value: 0.1234},
		"-0.1234": {IsNegative: true, Value: 0.1234},
	}

	for k, v := range floatCases {
		expr := parseExpression(t, k)

		floatExpr := expr.GetFloatLiteralExpr()
		assert.NotNil(t, floatExpr)
		assert.Equal(t, v.IsNegative, floatExpr.IsNegative)
		assert.Equal(t, v.Value, floatExpr.Value)
	}
}

func TestExpressionVisitor_VisitExpression_BinaryExpression(t *testing.T) {
	const binaryExpression = `a+b+c*d*e`
	expr := parseExpression(t, binaryExpression)

	binaryExpr := expr.GetBinaryExpr()
	assert.NotNil(t, binaryExpr)
	assert.Equal(t, "+", binaryExpr.GetOperator().Symbol)
	assert.NotNil(t, binaryExpr.LeftArgument.GetBinaryExpr())
	assert.Equal(t, "+", binaryExpr.LeftArgument.GetBinaryExpr().GetOperator().Symbol)
}

func TestExpressionVisitor_VisitExpression_ConditionalExpr(t *testing.T) {
	const conditionalExpression = `a ? b : c`
	expr := parseExpression(t, conditionalExpression)

	conditionalExpr := expr.GetConditionalExpr()
	assert.NotNil(t, conditionalExpr)
	assert.Equal(t, "a", conditionalExpr.Condition.GetIdentifierExpr().GetName())
	assert.Equal(t, "b", conditionalExpr.ThenBranch.GetIdentifierExpr().GetName())
	assert.Equal(t, "c", conditionalExpr.ElseBranch.GetIdentifierExpr().GetName())
}

func TestExpressionVisitor_VisitExpression_BinaryExpressionError(t *testing.T) {
	const expression = `a - - - b`

	parser := &Parser{}
	_, err := parser.ParseString(expression)
	assert.Error(t, err)
}

func parseExpression(t *testing.T, str string) (expr *lang.Expression) {
	parser := &Parser{}
	file, err := parser.ParseString(str)
	assert.NoError(t, err)

	if file != nil && len(file.Statements) > 0 {
		statement := file.Statements[0]
		expr = statement.GetExpression()
	}
	assert.NotNil(t, expr)
	return expr
}
