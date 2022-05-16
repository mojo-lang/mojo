package format

import (
    "bytes"
    "github.com/mojo-lang/mojo/go/pkg/mojo/context"
    "github.com/stretchr/testify/assert"
    "testing"
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
