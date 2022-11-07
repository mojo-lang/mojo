package printer

import (
	"bytes"
	"testing"

	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
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
	file, err := parser.ParseString(str)
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
	buffer := bytes.NewBuffer(nil)
	New(Config{}, buffer).PrintAttributeDecl(context.Empty(), decl)
	assert.Equal(t, expect, buffer.String())
}
