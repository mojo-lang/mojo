package format

import (
    "bytes"
    "github.com/mojo-lang/lang/go/pkg/mojo/lang"
    "github.com/mojo-lang/mojo/go/pkg/mojo/context"
    "github.com/mojo-lang/mojo/go/pkg/mojo/parser/syntax"
    "github.com/stretchr/testify/assert"
    "testing"
)

func getStructDecl(file *lang.SourceFile) *lang.StructDecl {
    if len(file.Statements) > 0 {
        statement := file.Statements[0]
        if decl := statement.GetDeclaration(); decl != nil {
            return decl.GetStructDecl()
        }
    }
    return nil
}

func parseStructDecl(t *testing.T, decl string) *lang.StructDecl {
    parser := &syntax.Parser{}
    file, err := parser.ParseString(decl)
    assert.NoError(t, err)

    structDecl := getStructDecl(file)
    assert.NotNil(t, structDecl)

    return structDecl
}

func TestPrinter_PrintStructDecl(t *testing.T) {
    const typeDecl = `
// comment1
// comment2

// comment3
type Mailbox {
    // free floating comment
    
    address: String
    /* block comment
    *///comment4
// comment5

	following: Bool //< following document - 1
//< following document - 2
}
`
    const expect = `// comment1
// comment2

// comment3
type Mailbox {
    // free floating comment
    address: String

    /* block comment
    */

    //comment4
    // comment5
    following: Bool //< following document - 1
                    //< following document - 2
}`

    decl := parseStructDecl(t, typeDecl)
    buffer := bytes.NewBuffer(nil)
    NewPrinter(Config{}, buffer).PrintStructDecl(context.Empty(), decl)
    assert.Equal(t, expect, buffer.String())
}
