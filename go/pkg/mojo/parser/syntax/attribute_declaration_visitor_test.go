package syntax

import (
	"testing"

	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/stretchr/testify/assert"

	"github.com/mojo-lang/mojo/go/pkg/context"
)

func TestAttributeDeclarationVisitor_VisitAttributeDeclaration(t *testing.T) {
	const attributeDecl = `@foo('bar') attribute mailbox: String`

	decl := parseAttributeDecl(t, attributeDecl)

	assert.Equal(t, "mailbox", decl.Name)
	assert.Equal(t, "String", decl.GetNominalType().GetName())
	assert.Equal(t, 1, len(decl.Attributes))
	assert.Equal(t, "foo", decl.Attributes[0].Name)
}

func getAttributeDecl(file *lang.SourceFile) *lang.AttributeDecl {
	if len(file.Statements) > 0 {
		statement := file.Statements[0]
		if decl := statement.GetDeclaration(); decl != nil {
			return decl.GetAttributeDecl()
		}
	}
	return nil
}

func parseAttributeDecl(t *testing.T, str string) *lang.AttributeDecl {
	parser := &Parser{}
	file, err := parser.ParseString(context.Empty(), str)
	assert.NoError(t, err)

	decl := getAttributeDecl(file)
	assert.NotNil(t, decl)
	return decl
}
