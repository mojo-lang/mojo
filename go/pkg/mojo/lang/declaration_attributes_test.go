package lang

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type attributeOwner interface {
	GetAttribute(string) *Attribute
	HasAttribute(string) bool
	GetAttributeArgument(string) (*Argument, error)
	GetAttributeArguments(string) ([]*Argument, error)
	GetBoolAttribute(string) (bool, error)
	GetIntegerAttribute(string) (int64, error)
	GetStringAttribute(string) (string, error)
	SetImplicitBoolAttribute(string, bool) *Attribute
	SetImplicitIntegerAttribute(string, int64) *Attribute
	SetImplicitStringAttribute(string, string) *Attribute
	RemoveAttribute(string)
}

func TestDeclarationAttributesUpdateTheRequestedAttribute(t *testing.T) {
	for _, tt := range []struct {
		name  string
		owner attributeOwner
	}{
		{"struct", &StructDecl{}}, {"alias", &TypeAliasDecl{}}, {"function", &FunctionDecl{}},
		{"value", &ValueDecl{Type: &NominalType{}}}, {"nominal", &NominalType{}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			x := tt.owner
			x.SetImplicitBoolAttribute("enabled", true)
			x.SetImplicitIntegerAttribute("number", 1)
			x.SetImplicitStringAttribute("name", "old")
			for _, name := range []string{"enabled", "number", "name"} {
				x.GetAttribute(name).Implicit = false
			}
			updated := x.SetImplicitBoolAttribute("enabled", false)
			require.Same(t, x.GetAttribute("enabled"), updated)
			require.True(t, updated.Implicit)
			require.False(t, x.GetAttribute("name").Implicit, "updating one attribute must not mark another implicit")
			require.Same(t, x.GetAttribute("number"), x.SetImplicitIntegerAttribute("number", -7))
			require.Same(t, x.GetAttribute("name"), x.SetImplicitStringAttribute("name", "new"))
			b, err := x.GetBoolAttribute("enabled")
			require.NoError(t, err)
			require.False(t, b)
			n, err := x.GetIntegerAttribute("number")
			require.NoError(t, err)
			require.Equal(t, int64(-7), n)
			s, err := x.GetStringAttribute("name")
			require.NoError(t, err)
			require.Equal(t, "new", s)
			require.True(t, x.HasAttribute("name"))
			args, err := x.GetAttributeArguments("name")
			require.NoError(t, err)
			require.Len(t, args, 1)
			arg, err := x.GetAttributeArgument("name")
			require.NoError(t, err)
			require.Same(t, args[0], arg)
			_, err = x.GetIntegerAttribute("name")
			require.Error(t, err)
			for _, name := range []string{"enabled", "number", "name"} {
				x.RemoveAttribute(name)
				require.False(t, x.HasAttribute(name))
			}
			_, err = x.GetBoolAttribute("enabled")
			require.Error(t, err)
			_, err = x.GetIntegerAttribute("number")
			require.Error(t, err)
			_, err = x.GetStringAttribute("name")
			require.Error(t, err)
		})
	}
}

func TestDeclarationAttributeDefaultsAndNilOwners(t *testing.T) {
	marker := &Attribute{Name: "flag"}
	for _, owner := range []attributeOwner{
		&StructDecl{Attributes: []*Attribute{marker}}, &TypeAliasDecl{Attributes: []*Attribute{marker}},
		&FunctionDecl{Attributes: []*Attribute{marker}}, &ValueDecl{Attributes: []*Attribute{marker}}, &NominalType{Attributes: []*Attribute{marker}},
	} {
		b, err := owner.GetBoolAttribute("flag")
		require.NoError(t, err)
		require.True(t, b)
		n, err := owner.GetIntegerAttribute("flag")
		require.NoError(t, err)
		require.Zero(t, n)
		s, err := owner.GetStringAttribute("flag")
		require.NoError(t, err)
		require.Empty(t, s)
	}
	for _, owner := range []attributeOwner{(*StructDecl)(nil), (*TypeAliasDecl)(nil), (*FunctionDecl)(nil), (*ValueDecl)(nil), (*NominalType)(nil)} {
		require.Nil(t, owner.GetAttribute("flag"))
		require.False(t, owner.HasAttribute("flag"))
		_, err := owner.GetAttributeArguments("flag")
		require.Error(t, err)
		_, err = owner.GetAttributeArgument("flag")
		require.Error(t, err)
		require.Nil(t, owner.SetImplicitBoolAttribute("flag", true))
		require.Nil(t, owner.SetImplicitIntegerAttribute("number", 1))
		require.Nil(t, owner.SetImplicitStringAttribute("name", "x"))
		require.NotPanics(t, func() { owner.RemoveAttribute("flag") })
	}
}

func TestValueDeclAttributesPreferDeclarationAndFallBackToType(t *testing.T) {
	field := &ValueDecl{Attributes: []*Attribute{NewStringAttribute("", "alias", "field")}, Type: &NominalType{Attributes: []*Attribute{NewStringAttribute("", "alias", "type"), NewIntegerAttribute("", "number", 9)}}}
	value, err := field.GetStringAttribute("alias")
	require.NoError(t, err)
	require.Equal(t, "field", value)
	valueArg, err := field.GetAttributeArgument("number")
	require.NoError(t, err)
	require.Same(t, field.Type.Attributes[1].Arguments[0], valueArg)
	args, err := field.GetAttributeArguments("number")
	require.NoError(t, err)
	require.Len(t, args, 1)
	require.True(t, field.HasAttribute("number"))
	require.Same(t, field.Type.Attributes[1], field.GetAttribute("number"))
	field.RemoveAttribute("alias")
	require.False(t, field.HasAttribute("alias"))
	require.Nil(t, field.GetAttribute("alias"))
	require.Nil(t, (&ValueDecl{}).SetStringAttribute("name", "value"))
}
