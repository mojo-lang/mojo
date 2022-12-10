package syntax

import (
	"testing"

	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/stretchr/testify/assert"

	"github.com/mojo-lang/mojo/go/pkg/context"
)

func TestTypeAliasDeclarationVisitor_VisitTypeAliasDeclaration(t *testing.T) {
	const typeAlias = `@foo('bar') type Mailbox = String`

	decl := parseTypeAlias(t, typeAlias)

	assert.Equal(t, "Mailbox", decl.Name)
	assert.Equal(t, "String", decl.Type.Name)
	assert.Equal(t, 1, len(decl.Attributes))
	assert.Equal(t, "foo", decl.Attributes[0].Name)
}

func parseTypeAlias(t *testing.T, decl string) *lang.TypeAliasDecl {
	parser := &Parser{}
	file, err := parser.ParseString(context.Empty(), decl)

	assert.NoError(t, err)
	aliasDecl := getAliasDecl(file)
	assert.NotNil(t, aliasDecl)

	return aliasDecl
}

func getAliasDecl(file *lang.SourceFile) *lang.TypeAliasDecl {
	if len(file.Statements) > 0 {
		statement := file.Statements[0]
		if decl := statement.GetDeclaration(); decl != nil {
			return decl.GetTypeAliasDecl()
		}
	}
	return nil
}
