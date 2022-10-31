package format

import (
	"bytes"
	"testing"

	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/stretchr/testify/assert"

	"github.com/mojo-lang/mojo/go/pkg/mojo/context"
	"github.com/mojo-lang/mojo/go/pkg/mojo/parser/syntax"
)

func getTypeAliasDecl(file *lang.SourceFile) *lang.TypeAliasDecl {
	if len(file.Statements) > 0 {
		statement := file.Statements[0]
		if decl := statement.GetDeclaration(); decl != nil {
			return decl.GetTypeAliasDecl()
		}
	}
	return nil
}

func parseTypeAliasDecl(t *testing.T, decl string) *lang.TypeAliasDecl {
	parser := &syntax.Parser{}
	file, err := parser.ParseString(decl)
	assert.NoError(t, err)

	alias := getTypeAliasDecl(file)
	assert.NotNil(t, alias)
	return alias
}

func TestPrinter_PrintTypeAliasDecl(t *testing.T) {
	const typeDecl = `
// comment1
// comment2

// comment3
type Mailbox=Box  //< following document - 1
//< following document - 2
`
	const expect = `// comment1
// comment2

// comment3
type Mailbox = Box //< following document - 1
                   //< following document - 2
`

	decl := parseTypeAliasDecl(t, typeDecl)
	buffer := bytes.NewBuffer(nil)
	NewPrinter(Config{}, buffer).PrintTypeAliasDecl(context.Empty(), decl)
	assert.Equal(t, expect, buffer.String())
}
