package syntax

import (
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExpressionVisitor_VisitExpression_False(t *testing.T) {
	const falseLiteral = `false`

	parser := &Parser{}
	file, err := parser.ParseString(falseLiteral)

	assert.NoError(t, err)
	expr := getExpression(file)
	assert.NotNil(t, expr)
	assert.NotNil(t, expr.GetBoolLiteralExpr())
	assert.Equal(t, false, expr.GetBoolLiteralExpr().Value)
}

func TestExpressionVisitor_VisitExpression_True(t *testing.T) {
	const trueLiteral = `true`

	parser := &Parser{}
	file, err := parser.ParseString(trueLiteral)

	assert.NoError(t, err)
	expr := getExpression(file)
	assert.NotNil(t, expr)
	assert.NotNil(t, expr.GetBoolLiteralExpr())
	assert.Equal(t, true, expr.GetBoolLiteralExpr().Value)
}

func TestExpressionVisitor_VisitExpression_BinaryExpression(t *testing.T) {
	const binaryExpression = `a+b+c*d*e`

	parser := &Parser{}
	file, err := parser.ParseString(binaryExpression)

	assert.NoError(t, err)
	expr := getExpression(file)
	assert.NotNil(t, expr)

	binaryExpr := expr.GetBinaryExpr()
	assert.NotNil(t, binaryExpr)
	assert.Equal(t, "+", binaryExpr.Operator.Symbol)
	assert.NotNil(t, binaryExpr.LeftHandArgument.GetBinaryExpr())
	assert.Equal(t, "+", binaryExpr.LeftHandArgument.GetBinaryExpr().Operator.Symbol)
}

func TestExpressionVisitor_VisitExpression_BinaryExpressionError(t *testing.T) {
	const expression = `a - - - b`

	parser := &Parser{}
	_, err := parser.ParseString(expression)
	assert.Error(t, err)
}

func getExpression(file *lang.SourceFile) *lang.Expression {
	if len(file.Statements) > 0 {
		statement := file.Statements[0]
		return statement.GetExpression()
	}
	return nil
}
