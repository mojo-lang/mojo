package lang

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVariableDecl_ToValueDecl(t *testing.T) {
	variable := &VariableDecl{
		Implicit: true,
		Document: &Document{
			Private: true,
			Lines:   []*Document_Line{{Content: "document"}},
		},
		PackageName:    "foo",
		SourceFileName: "bar.mojo",
		Name:           "bar",
		Type:           &NominalType{Name: "Foobar"},
	}

	value := variable.ToValueDecl()
	assert.NotNil(t, value)

	assert.Equal(t, variable.Implicit, value.Implicit)
	assert.Equal(t, variable.Document.Private, value.Document.Private)
	assert.Equal(t, variable.Document.Lines[0].Content, value.Document.Lines[0].Content)
	assert.Equal(t, variable.PackageName, value.PackageName)
	assert.Equal(t, variable.SourceFileName, value.SourceFileName)
	assert.Equal(t, variable.Name, value.Name)
	assert.Equal(t, variable.Type.Name, value.Type.Name)
}
