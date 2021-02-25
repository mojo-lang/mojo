package syntax

import (
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAttributeVisitor_VisitAttribute_Number(t *testing.T) {
	const typeAttribute = `@1 type Mailbox{}`

	parser := &Parser{}
	file, err := parser.ParseString(typeAttribute)

	assert.NoError(t, err)
	attribute := getAttribute(file)
	assert.NotNil(t, attribute)

	assert.Equal(t, int64(1), attribute.Arguments[0].GetIntegerLiteralExpr().Value)
}

func TestAttributeVisitor_VisitAttribute_FieldNumber(t *testing.T) {
	const typeAttribute = `type Mailbox { user : String @1 }`

	parser := &Parser{}
	file, err := parser.ParseString(typeAttribute)

	assert.NoError(t, err)
	attribute := getFieldAttribute(file, 0)
	assert.NotNil(t, attribute)

	assert.Equal(t, int64(1), attribute.Arguments[0].GetIntegerLiteralExpr().Value)
}

func getAttribute(file *lang.SourceFile) *lang.Attribute {
	if len(file.Statements) > 0 {
		statement := file.Statements[0]
		if decl := statement.GetDeclaration(); decl != nil {
			attributes := decl.GetStructDecl().Attributes
			if len(attributes) > 0 {
				return attributes[0]
			}
		}
	}
	return nil
}

func getFieldAttribute(file *lang.SourceFile, index int) *lang.Attribute {
	if len(file.Statements) > 0 {
		statement := file.Statements[0]
		if decl := statement.GetDeclaration(); decl != nil {
			structDecl := decl.GetStructDecl()
			if structDecl == nil || structDecl.Type == nil || len(structDecl.Type.Fields) == 0 {
				return nil
			}
			attributes := structDecl.Type.Fields[0].Type.Attributes
			if len(attributes) > index {
				return attributes[index]
			}
		}
	}
	return nil
}
