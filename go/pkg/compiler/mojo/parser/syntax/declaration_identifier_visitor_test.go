package syntax

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeclarationIdentifierVisitor_VisitDeclarationIdentifier(t *testing.T) {
	const identifier = `var foo = "testdata"`

	decl := ParseVariableDecl(t, identifier)
	if assert.NotNil(t, decl) {
		assert.Equal(t, "foo", decl.Name)
		assert.Equal(t, "testdata", decl.GetInitialValue().GetStringLiteralExpr().GetValue())
	}
}
