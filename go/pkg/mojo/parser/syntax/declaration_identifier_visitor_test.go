package syntax

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeclarationIdentifierVisitor_VisitDeclarationIdentifier(t *testing.T) {
	const identifier = `var type = "test-data"`

	decl := ParseVariableDecl(t, identifier)
	assert.Equal(t, "type", decl.Name)
	assert.Equal(t, "test-data", decl.GetInitialValue().GetStringLiteralExpr().GetValue())
}
