package syntax

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDocumentVisitor_VisitDocument(t *testing.T) {
	const typeDecl = `
/// document1
/// document2
type Mailbox {
	/// free floating document
	address: String
}
`

	decl := parseStructDecl(t, typeDecl)

	assert.Equal(t, 2, len(decl.Document.Lines))
}
