package syntax

import (
	"testing"

	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/stretchr/testify/assert"

	"github.com/mojo-lang/mojo/go/pkg/context"
)

func ParseVariableDecl(t *testing.T, decl string) *lang.VariableDecl {
	parser := &Parser{}
	file, err := parser.ParseString(context.Empty(), decl)

	if assert.NoError(t, err) {
		if len(file.Statements) > 0 {
			statement := file.Statements[0]
			if d := statement.GetDeclaration(); d != nil {
				return d.GetVariableDecl()
			}
		}
	}

	assert.NotNil(t, nil, "failed to parse the variable value decl")
	return nil
}

func ParseConstantDecl(t *testing.T, decl string) *lang.ConstantDecl {
	parser := &Parser{}
	file, err := parser.ParseString(context.Empty(), decl)

	assert.NoError(t, err)
	if len(file.Statements) > 0 {
		statement := file.Statements[0]
		if d := statement.GetDeclaration(); d != nil {
			return d.GetConstantDecl()
		}
	}

	assert.NotNil(t, nil, "failed to parse the constant value decl")
	return nil
}

func TestValueDeclarationVisitor_VisitVariableDeclaration(t *testing.T) {

}
