package syntax

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTypeIdentifier(t *testing.T) {
	const typeDecl = `type Mailbox : Box<Mail<T>>`

	decl := parseStructDecl(t, typeDecl)
	assert.Equal(t, "Box", decl.Type.Inherits[0].Name)
}

func TestGetTypeIdentifier_Package(t *testing.T) {
	const typeDecl = `type Mailbox : some.pkg.Box<Mail<T>>`

	decl := parseStructDecl(t, typeDecl)

	assert.Equal(t, "Box", decl.Type.Inherits[0].Name)
	assert.Equal(t, "some.pkg", decl.Type.Inherits[0].PackageName)
}

func TestGetTypeIdentifier_Package_Enclosing(t *testing.T) {
	const typeDecl = `type Mailbox : some.pkg.Box<T>.Mail<T>`

	decl := parseStructDecl(t, typeDecl)

	assert.Equal(t, "Mail", decl.Type.Inherits[0].Name)
	assert.Equal(t, "T", decl.Type.Inherits[0].GenericArguments[0].Name)
	assert.Equal(t, "some.pkg", decl.Type.Inherits[0].PackageName)
	assert.Equal(t, "Box", decl.Type.Inherits[0].Enclosing.Name)
	assert.Equal(t, "T", decl.Type.Inherits[0].Enclosing.GenericArguments[0].Name)

	assert.Equal(t, "some.pkg.Box<T>.Mail<T>", decl.Type.Inherits[0].GetGenericFullName())
}
