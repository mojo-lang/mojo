package syntax

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPackageIdentifierVisitor_VisitPackageIdentifier(t *testing.T) {
	const typeDecl = `type Mailbox : ab.cd.ef.Box`

	decl := parseStructDecl(t, typeDecl)
	assert.Equal(t, "ab.cd.ef", decl.Type.Inherits[0].PackageName)
}
