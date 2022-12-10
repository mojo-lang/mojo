package syntax

import (
	"testing"

	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/stretchr/testify/assert"

	"github.com/mojo-lang/mojo/go/pkg/context"
)

func parseEnumDecl(t *testing.T, decl string) *lang.EnumDecl {
	parser := &Parser{}
	file, err := parser.ParseString(context.Empty(), decl)
	assert.NoError(t, err)

	if len(file.Statements) > 0 {
		statement := file.Statements[0]
		if d := statement.GetDeclaration(); d != nil {
			return d.GetEnumDecl()
		}
	}

	assert.NotNil(t, nil, "failed to parse the enum decl")
	return nil
}

func TestEnumDeclarationVisitor_VisitEnumDeclaration(t *testing.T) {
	const typeDecl = `
@foo('bar')
enum Box {
	black
	white
}
`
	decl := parseEnumDecl(t, typeDecl)
	assert.Equal(t, "Box", decl.Name)
	assert.Equal(t, 2, len(decl.Type.Enumerators))
}
