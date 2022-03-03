package syntax

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetPosition(t *testing.T) {
	const typeDecl = `@foo('bar') type Mailbox {}`

	decl := parseStructDecl(t, typeDecl)
	assert.Equal(t, int64(1), decl.StartPosition.Line)
	assert.Equal(t, int64(1), decl.StartPosition.Column)
	assert.Equal(t, int64(1), decl.EndPosition.Line)
	assert.Equal(t, int64(26), decl.EndPosition.Column)
}
