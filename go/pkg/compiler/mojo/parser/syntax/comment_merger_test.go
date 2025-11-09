package syntax

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommentMerger_Merge(t *testing.T) {
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

	following: Bool // following comment
}
`

	decl := parseStructDecl(t, typeDecl)
	assert.NotNil(t, decl)
	assert.Equal(t, 2, len(decl.StartPosition.LeadingComments))
	assert.Equal(t, 1, len(decl.StartPosition.LeadingComments[1].GetMultiLineComment().GetLines()))
	assert.Equal(t, 1, len(decl.Type.Fields[0].StartPosition.LeadingComments))
	assert.Equal(t, 2, len(decl.Type.Fields[1].StartPosition.LeadingComments))
	assert.Equal(t, 1, len(decl.Type.Fields[1].EndPosition.TailingComments))
}
