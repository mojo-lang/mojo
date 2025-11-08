package printer

import (
	"testing"

	"github.com/mojo-lang/mojo/packages/lang/go/pkg/mojo/lang"
	"github.com/stretchr/testify/assert"

	"github.com/mojo-lang/mojo/go/pkg/context"
	"github.com/mojo-lang/mojo/go/pkg/mojo/parser/syntax"
)

func getFunctionDecl(file *lang.SourceFile) *lang.FunctionDecl {
	if len(file.Statements) > 0 {
		statement := file.Statements[0]
		if decl := statement.GetDeclaration(); decl != nil {
			return decl.GetFunctionDecl()
		}
	}
	return nil
}

func parseFunctionDecl(t *testing.T, str string) *lang.FunctionDecl {
	parser := &syntax.Parser{}
	file, err := parser.ParseString(context.Empty(), str)
	assert.NoError(t, err)

	decl := getFunctionDecl(file)
	assert.NotNil(t, decl)
	return decl
}

func TestPrinter_PrintFunctionDecl(t *testing.T) {
	const typeDecl = `
// comment1
// comment2

// comment3
func mailbox(box: Box @1  //< following document - 1
//< following document - 2
mail: Mail @2 @must) -> Void //< nothing
`
	const expect = `// comment1
// comment2

// comment3
func mailbox(box Box @1 //< following document - 1
                        //< following document - 2
             mail Mail @2 @must)
             -> Void //< nothing
`

	decl := parseFunctionDecl(t, typeDecl)
	p := New(&Config{}).PrintFunctionDecl(context.Empty(), decl)
	assert.Equal(t, expect, p.Buffer.String())
}
