package syntax

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPackageIdentifierVisitor_VisitPackageIdentifier(t *testing.T) {
	const typeDecl = `type Mailbox : ab.cd.ef.Box`

	decl := parseStructDecl(t, typeDecl)
	assert.Equal(t, "ab.cd.ef", decl.Type.Inherits[0].PackageName)
}
