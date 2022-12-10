package printer

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mojo-lang/mojo/go/pkg/context"
)

func TestPrinter_PrintNominalType(t *testing.T) {
	const typeDecl = `type Mailbox: Int32 | String`
	const expect = `Int32 | String`

	decl := parseStructDecl(t, typeDecl)
	assert.NotNil(t, decl.Type)
	assert.NotEmptyf(t, decl.Type.Inherits, "the type inherits is empty")

	p := New(&Config{}).PrintNominalType(context.Empty(), decl.Type.Inherits[0])
	assert.Equal(t, expect, p.Buffer.String())
}
