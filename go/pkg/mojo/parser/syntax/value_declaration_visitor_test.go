package syntax

import (
	"testing"

	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/stretchr/testify/assert"
)

func ParseVariableDecl(t *testing.T, decl string) *lang.VariableDecl {
	parser := &Parser{}
	file, err := parser.ParseString(decl)

	assert.NoError(t, err)
	if len(file.Statements) > 0 {
		statement := file.Statements[0]
		if d := statement.GetDeclaration(); d != nil {
			return d.GetVariableDecl()
		}
	}

	assert.NotNil(t, nil, "failed to parse the variable value decl")
	return nil
}

func ParseConstantDecl(t *testing.T, decl string) *lang.ConstantDecl {
	parser := &Parser{}
	file, err := parser.ParseString(decl)

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
