package syntax

import (
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTypeAliasDeclarationVisitor_VisitTypeAliasDeclaration(t *testing.T) {
	const typeAlias = `@foo('bar') type Mailbox = String`

	parser := &Parser{}
	file, err := parser.ParseString(typeAlias)

	assert.NoError(t, err)
	decl := getDecl(file)
	assert.NotNil(t, decl)
	assert.Equal(t, "Mailbox", decl.Name)
	assert.Equal(t, "String", decl.Type.Name)
	assert.Equal(t, 1, len(decl.Attributes))
	assert.Equal(t, "foo", decl.Attributes[0].Name)
}

func getDecl(file *lang.SourceFile) *lang.TypeAliasDecl {
	if len(file.Statements) > 0 {
		statement := file.Statements[0]
		if decl := statement.GetDeclaration(); decl != nil {
			return decl.GetTypeAliasDecl()
		}
	}
	return nil
}
