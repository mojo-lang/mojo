package syntax

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommentParser_Parse(t *testing.T) {
	const typeDecl = `
// comment1
// comment2
type Mailbox {
	// free floating document
	
	address: String
	/* block comment
	*/
}
`

	decl := parseStructDecl(t, typeDecl)
	assert.NotNil(t, decl)
}
