package printer

import (
	"testing"

	"github.com/mojo-lang/mojo/packages/lang/go/pkg/mojo/lang"
	"github.com/stretchr/testify/assert"

	"github.com/mojo-lang/mojo/go/pkg/context"
	"github.com/mojo-lang/mojo/go/pkg/mojo/parser/syntax"
)

func getAttributeDecl(file *lang.SourceFile) *lang.AttributeDecl {
	if len(file.Statements) > 0 {
		statement := file.Statements[0]
		if decl := statement.GetDeclaration(); decl != nil {
			return decl.GetAttributeDecl()
		}
	}
	return nil
}

func parseAttributeDecl(t *testing.T, str string) *lang.AttributeDecl {
	parser := &syntax.Parser{}
	file, err := parser.ParseString(context.Empty(), str)
	assert.NoError(t, err)

	decl := getAttributeDecl(file)
	assert.NotNil(t, decl)
	return decl
}

func TestPrinter_PrintAttributeDecl(t *testing.T) {
	const typeDecl = `
// comment1
// comment2

// comment3
attribute mailbox: Box  //< following document - 1
//< following document - 2
`
	const expect = `// comment1
// comment2

// comment3
attribute mailbox: Box //< following document - 1
                       //< following document - 2
`

	decl := parseAttributeDecl(t, typeDecl)
	p := New(&Config{}).PrintAttributeDecl(context.Empty(), decl)
	assert.Equal(t, expect, p.Buffer.String())
}
