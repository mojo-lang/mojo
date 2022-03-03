package syntax

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAttributesVisitor_VisitAttributes(t *testing.T) {
	const typeDecl = `
@foo('bar') @bar @foobar(1234)
type Mailbox {
	address: String
}
`
	decl := parseStructDecl(t, typeDecl)
	assert.Equal(t, "Mailbox", decl.Name)
	assert.Equal(t, 3, len(decl.Attributes))
}

func TestAttributesVisitor_VisitAttributes_BankLine(t *testing.T) {
	const typeDecl = `
@foo('bar',
	 "car")
@bar
@foobar(1234) @same_line(true)
type Mailbox {
	address: String
}
`

	decl := parseStructDecl(t, typeDecl)
	assert.Equal(t, "Mailbox", decl.Name)
	assert.Equal(t, 4, len(decl.Attributes))
}
