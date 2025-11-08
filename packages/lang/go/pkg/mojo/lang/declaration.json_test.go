package lang

import (
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
)

const EnumDeclStr = `{"@type":"EnumDecl","packageName":"mojo.lang.test","name":"Foo","type":{"enumerators":[{"name":"a"}]}}`

func TestDeclarationCodec_Encode(t *testing.T) {
	decl := NewEnumDeclaration(&EnumDecl{
		Name:        "Foo",
		PackageName: "mojo.lang.test",
		Type: &EnumType{
			Enumerators: []*ValueDecl{
				{
					Name: "a",
				},
			},
		},
	})

	str, err := jsoniter.ConfigDefault.Marshal(decl)
	assert.NoError(t, err)
	assert.Equal(t, EnumDeclStr, string(str))
}

func TestDeclarationCodec_Decode(t *testing.T) {
	decl := &Declaration{}
	err := jsoniter.ConfigDefault.Unmarshal([]byte(EnumDeclStr), decl)
	assert.NoError(t, err)
	assert.Equal(t, "Foo", decl.GetEnumDecl().Name)
}
