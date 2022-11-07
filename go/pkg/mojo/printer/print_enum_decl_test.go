package printer

import (
	"bytes"
	"testing"

	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/stretchr/testify/assert"

	"github.com/mojo-lang/mojo/go/pkg/context"
	"github.com/mojo-lang/mojo/go/pkg/mojo/parser/syntax"
)

func getEnumDecl(file *lang.SourceFile) *lang.EnumDecl {
	if len(file.Statements) > 0 {
		statement := file.Statements[0]
		if decl := statement.GetDeclaration(); decl != nil {
			return decl.GetEnumDecl()
		}
	}
	return nil
}

func parseEnumDecl(t *testing.T, str string) *lang.EnumDecl {
	parser := &syntax.Parser{}
	file, err := parser.ParseString(str)
	assert.NoError(t, err)

	decl := getEnumDecl(file)
	assert.NotNil(t, decl)
	return decl
}

func TestPrinter_PrintEnumDecl(t *testing.T) {
	const typeDecl = `
// comment1
// comment2

// comment3
enum Mailbox {

none  @1 //< following document - 1
//< following document - 2

box @2
}
`
	const expect = `// comment1
// comment2

// comment3
enum Mailbox {
    none @1 //< following document - 1
            //< following document - 2
    box  @2
}`

	decl := parseEnumDecl(t, typeDecl)
	buffer := bytes.NewBuffer(nil)
	New(Config{}, buffer).PrintEnumDecl(context.Empty(), decl)
	assert.Equal(t, expect, buffer.String())
}
