package syntax

import (
	"testing"

	"github.com/mojo-lang/lang/go/pkg/mojo/lang"

	"github.com/stretchr/testify/assert"

	"github.com/mojo-lang/mojo/go/pkg/context"
)

func parseInterfaceDecl(t *testing.T, decl string) *lang.InterfaceDecl {
	parser := &Parser{}
	file, err := parser.ParseString(context.Empty(), decl)
	assert.NoError(t, err)

	if len(file.Statements) > 0 {
		statement := file.Statements[0]
		if d := statement.GetDeclaration(); d != nil {
			return d.GetInterfaceDecl()
		}
	}
	assert.NotNil(t, nil, "failed to parse the interface decl")
	return nil
}

func TestInterfaceDeclarationVisitor_VisitInterfaceDeclaration(t *testing.T) {
	const typeDecl = `
@foo('bar')
interface Mailbox {
}
`
	decl := parseInterfaceDecl(t, typeDecl)
	assert.Equal(t, "Mailbox", decl.Name)
}
