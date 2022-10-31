package syntax

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeclarationIdentifierVisitor_VisitDeclarationIdentifier(t *testing.T) {
	const identifier = `var type = "testdata"`

	decl := ParseVariableDecl(t, identifier)
	assert.Equal(t, "type", decl.Name)
	assert.Equal(t, "testdata", decl.GetInitialValue().GetStringLiteralExpr().GetValue())
}
