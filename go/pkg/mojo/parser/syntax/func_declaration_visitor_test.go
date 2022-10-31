package syntax

import (
	"testing"

	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/stretchr/testify/assert"
)

func parseFunctionDecl(t *testing.T, decl string) *lang.FunctionDecl {
	parser := &Parser{}
	file, err := parser.ParseString(decl)
	assert.NoError(t, err)

	if len(file.Statements) > 0 {
		statement := file.Statements[0]
		if d := statement.GetDeclaration(); d != nil {
			return d.GetFunctionDecl()
		}
	}

	assert.NotNil(t, nil, "failed to parse the function decl")
	return nil
}

func TestFuncDeclarationVisitor_VisitFunctionDeclaration(t *testing.T) {
	const typeDecl = `
@foo('bar')
func abc(a:Int) -> Int
`
	decl := parseFunctionDecl(t, typeDecl)
	assert.Equal(t, "abc", decl.Name)
	assert.Equal(t, 1, len(decl.Signature.GetParameters()))
	assert.Equal(t, "Int", decl.Signature.GetResultType().Name)
}
