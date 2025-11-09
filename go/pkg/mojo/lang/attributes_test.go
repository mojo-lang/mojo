package lang

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAttributes(t *testing.T) {
	decl := NewStructDeclStatement(&StructDecl{Attributes: []*Attribute{{Name: "test"}}})
	attributes := GetAttributes(decl)

	assert.NotEmpty(t, attributes)
	assert.NotNil(t, attributes.GetAttribute("test"))
}
