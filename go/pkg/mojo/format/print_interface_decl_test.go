package format

import (
	"bytes"
	"testing"

	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/stretchr/testify/assert"

	"github.com/mojo-lang/mojo/go/pkg/mojo/context"
	"github.com/mojo-lang/mojo/go/pkg/mojo/parser/syntax"
)

func getInterfaceDecl(file *lang.SourceFile) *lang.InterfaceDecl {
	if len(file.Statements) > 0 {
		statement := file.Statements[0]
		if decl := statement.GetDeclaration(); decl != nil {
			return decl.GetInterfaceDecl()
		}
	}
	return nil
}

func parseInterfaceDecl(t *testing.T, str string) *lang.InterfaceDecl {
	parser := &syntax.Parser{}
	file, err := parser.ParseString(str)
	assert.NoError(t, err)

	decl := getInterfaceDecl(file)
	assert.NotNil(t, decl)
	return decl
}

func TestPrinter_PrintInterfaceDecl(t *testing.T) {
	const typeDecl = `
// comment1
// comment2

// comment3
interface Mailbox {

@http.get("/mailbox/v1/mails/{id}")
get_mail(id: String @1)//< following document - 1
//< following document - 2
-> Mail //< the mail



}
`
	const expect = `// comment1
// comment2

// comment3
interface Mailbox {

    @http.get("/mailbox/v1/mails/{id}")
    get_mail(id String @1) //< following document - 1
                           //< following document - 2
             -> Mail //< the mail
}`

	decl := parseInterfaceDecl(t, typeDecl)
	buffer := bytes.NewBuffer(nil)
	NewPrinter(Config{}, buffer).PrintInterfaceDecl(context.Empty(), decl)
	assert.Equal(t, expect, buffer.String())
}
