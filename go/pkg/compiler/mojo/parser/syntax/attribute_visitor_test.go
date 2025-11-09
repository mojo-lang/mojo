package syntax

import (
	"testing"

	"github.com/mojo-lang/mojo/go/pkg/mojo/lang"
	"github.com/stretchr/testify/assert"

	"github.com/mojo-lang/mojo/go/pkg/compiler/context"
)

func TestAttributeVisitor_VisitAttribute_Number(t *testing.T) {
	const typeAttribute = `@1 type Mailbox{}`

	parser := &Parser{}
	file, err := parser.ParseString(context.Empty(), typeAttribute)

	assert.NoError(t, err)
	attribute := getAttribute(file)
	assert.NotNil(t, attribute)

	assert.Equal(t, uint64(1), attribute.Arguments[0].GetIntegerLiteralExpr().Value)
}

func TestAttributeVisitor_VisitAttribute_Object(t *testing.T) {
	const typeAttribute = `@t(foo: 11, bar: 22) type Mailbox{}`

	parser := &Parser{}
	file, err := parser.ParseString(context.Empty(), typeAttribute)

	assert.NoError(t, err)
	attribute := getAttribute(file)
	assert.NotNil(t, attribute)

	assert.Equal(t, "t", attribute.Name)

	assert.Equal(t, "foo", attribute.Arguments[0].Label)
	assert.Equal(t, uint64(11), attribute.Arguments[0].GetIntegerLiteralExpr().Value)

	assert.Equal(t, "bar", attribute.Arguments[1].Label)
	assert.Equal(t, uint64(22), attribute.Arguments[1].GetIntegerLiteralExpr().Value)
}

func TestAttributeVisitor_VisitAttribute_Array(t *testing.T) {
	const typeAttribute = `@foo("test1", "test2", "test3") type Mailbox{}`

	parser := &Parser{}
	file, err := parser.ParseString(context.Empty(), typeAttribute)

	assert.NoError(t, err)
	attribute := getAttribute(file)
	assert.NotNil(t, attribute)

	assert.Equal(t, "foo", attribute.Name)
	assert.Equal(t, "test1", attribute.Arguments[0].GetStringLiteralExpr().Value)
	assert.Equal(t, "test2", attribute.Arguments[1].GetStringLiteralExpr().Value)
	assert.Equal(t, "test3", attribute.Arguments[2].GetStringLiteralExpr().Value)
}

func TestAttributeVisitor_VisitAttribute_FieldNumber(t *testing.T) {
	const typeAttribute = `type Mailbox { user : String @1 }`

	parser := &Parser{}
	file, err := parser.ParseString(context.Empty(), typeAttribute)

	assert.NoError(t, err)
	attribute := getFieldAttribute(file, 0)
	assert.NotNil(t, attribute)

	assert.Equal(t, uint64(1), attribute.Arguments[0].GetIntegerLiteralExpr().Value)
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
