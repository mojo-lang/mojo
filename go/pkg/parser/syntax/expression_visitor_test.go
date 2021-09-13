package syntax

import (
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExpressionVisitor_VisitExpression_False(t *testing.T) {
	const typeAttribute = `
@default(false)
type Mailbox{}`

	parser := &Parser{}
	file, err := parser.ParseString(typeAttribute)

	assert.NoError(t, err)
	expr := getExpression(file)
	assert.NotNil(t, expr)
	assert.NotNil(t, expr.GetBoolLiteralExpr())
	assert.Equal(t, false, expr.GetBoolLiteralExpr().Value)
}

func TestExpressionVisitor_VisitExpression_True(t *testing.T) {
	const typeAttribute = `@default(true) type Mailbox{}`

	parser := &Parser{}
	file, err := parser.ParseString(typeAttribute)

	assert.NoError(t, err)
	expr := getExpression(file)
	assert.NotNil(t, expr)
	assert.NotNil(t, expr.GetBoolLiteralExpr())
	assert.Equal(t, true, expr.GetBoolLiteralExpr().Value)
}

func getExpression(file *lang.SourceFile) *lang.Expression {
	if len(file.Statements) > 0 {
		statement := file.Statements[0]
		if decl := statement.GetDeclaration(); decl != nil {
			attributes := decl.GetStructDecl().Attributes
			if len(attributes) > 0 {
				expressions := attributes[0].GetArguments()
				if len(expressions) > 0 {
					return expressions[0].Value
				}
			}
		}
	}
	return nil
}
