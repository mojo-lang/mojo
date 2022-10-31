package format

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mojo-lang/mojo/go/pkg/mojo/context"
)

func TestPrinter_PrintNominalType(t *testing.T) {
	const typeDecl = `type Mailbox: Int32 | String`
	const expect = `Int32 | String`

	decl := parseStructDecl(t, typeDecl)
	assert.NotNil(t, decl.Type)
	assert.NotEmptyf(t, decl.Type.Inherits, "the type inherits is empty")

	buffer := bytes.NewBuffer(nil)
	NewPrinter(Config{}, buffer).PrintNominalType(context.Empty(), decl.Type.Inherits[0])
	assert.Equal(t, expect, buffer.String())
}
