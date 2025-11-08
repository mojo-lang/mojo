package lang

import (
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
)

func TestStatementCodec_Encode(t *testing.T) {
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
	stmt := NewDeclarationStatement(decl)

	str, err := jsoniter.ConfigDefault.Marshal(stmt)
	assert.NoError(t, err)
	assert.Equal(t, EnumDeclStr, string(str))
}

func TestStatementCodec_Decode(t *testing.T) {
	stmt := &Statement{}
	err := jsoniter.ConfigDefault.Unmarshal([]byte(EnumDeclStr), stmt)
	assert.NoError(t, err)
	assert.Equal(t, "Foo", stmt.GetDeclaration().GetEnumDecl().Name)
}
