package lang

import (
	"testing"

	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/stretchr/testify/require"
)

func TestAttributesReplaceMarkersAndPreserveOtherValues(t *testing.T) {
	for _, tt := range []struct {
		name  string
		set   func([]*Attribute) []*Attribute
		check func(*testing.T, []*Attribute)
	}{
		{"boolean", func(a []*Attribute) []*Attribute { return SetBoolAttribute(a, "mojo.core.setting", false) }, func(t *testing.T, a []*Attribute) {
			v, err := GetBoolAttribute(a, "setting")
			require.NoError(t, err)
			require.False(t, v)
		}},
		{"integer", func(a []*Attribute) []*Attribute { return SetIntegerAttribute(a, "mojo.core.setting", -42) }, func(t *testing.T, a []*Attribute) {
			v, err := GetIntegerAttribute(a, "setting")
			require.NoError(t, err)
			require.Equal(t, int64(-42), v)
		}},
		{"string", func(a []*Attribute) []*Attribute { return SetStringAttribute(a, "mojo.core.setting", "new") }, func(t *testing.T, a []*Attribute) {
			v, err := GetStringAttribute(a, "setting")
			require.NoError(t, err)
			require.Equal(t, "new", v)
		}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			marker := &Attribute{PackageName: "mojo.core", Name: "setting"}
			other := NewStringAttribute("custom", "other", "keep")
			attributes := tt.set([]*Attribute{marker, other})
			require.Len(t, attributes, 2, "setting an existing marker must not create a duplicate")
			require.Same(t, marker, attributes[0])
			require.Same(t, other, attributes[1])
			tt.check(t, attributes)
			attributes = tt.set(attributes)
			require.Len(t, attributes, 2)
			tt.check(t, attributes)
			require.Equal(t, []*Attribute{other}, RemoveAttribute(attributes, "setting"))
			tt.check(t, tt.set(nil))
		})
	}
}

func TestAttributesLookupDefaultsAndStringLists(t *testing.T) {
	marker := &Attribute{Name: "required"}
	attrs := Attributes{nil, marker, NewIntegerAttribute("mojo.core", "number", 7), NewBoolAttribute("mojo.core", "optional")}
	require.Same(t, marker, attrs.GetRequiredAttribute())
	require.NotNil(t, attrs.GetNumberAttribute())
	require.NotNil(t, attrs.GetOptionalAttribute())
	require.True(t, attrs.HasAttribute("mojo.core.required"))
	b, err := GetBoolAttribute(attrs, "required")
	require.NoError(t, err)
	require.True(t, b)
	n, err := GetIntegerAttribute(attrs, "required")
	require.NoError(t, err)
	require.Zero(t, n)
	s, err := GetStringAttribute(attrs, "required")
	require.NoError(t, err)
	require.Empty(t, s)
	_, err = GetAttributeArguments(attrs, "required")
	require.Error(t, err)
	_, err = GetAttributeArgument(attrs, "missing")
	require.Error(t, err)
	_, err = GetBoolAttribute(attrs, "number")
	require.Error(t, err)
	_, err = GetIntegerAttribute(attrs, "missing")
	require.Error(t, err)
	_, err = GetStringAttribute(attrs, "missing")
	require.Error(t, err)
	for _, arguments := range [][]*Argument{
		{NewStringArgument("one"), NewStringArgument("two")},
		{NewArrayLiteralArgument(&ArrayLiteralExpr{Elements: []*Expression{NewStringLiteralExpressionFrom("one"), NewStringLiteralExpressionFrom("two")}})},
	} {
		values, err := GetStringValuesAttribute([]*Attribute{{Name: "tags", Arguments: arguments}}, "tags")
		require.NoError(t, err)
		require.Equal(t, []string{"one", "two"}, values)
	}
	_, err = GetStringValuesAttribute(nil, "tags")
	require.Error(t, err)
	_, err = GetStringValuesAttribute([]*Attribute{NewIntegerAttribute("", "tags", 1)}, "tags")
	require.Error(t, err)
	require.Nil(t, GetAttributes(42))
}

func TestAttributeMergeRespectsRepeatabilityAndDestination(t *testing.T) {
	first := NewStringAttribute("sample", "tag", "first")
	second := NewStringAttribute("sample", "tag", "second")
	require.Equal(t, []*Attribute{first}, MergeAttributes([]*Attribute{first}, []*Attribute{second}))
	declaration := &AttributeDecl{Attributes: []*Attribute{NewBoolAttribute("", "repeatable")}}
	first.Declaration, second.Declaration = declaration, declaration
	require.Equal(t, []*Attribute{first, second}, MergeAttributes([]*Attribute{first}, []*Attribute{second}))
	require.Empty(t, MergeAttributes(nil, nil))
	require.True(t, (&Attribute{Declaration: &AttributeDecl{DefaultValue: NewBoolLiteralExpressionFrom(true)}}).GetBool())
	require.False(t, (&Attribute{Declaration: &AttributeDecl{DefaultValue: NewBoolLiteralExpressionFrom(false)}}).GetBool())
	implicit := NewImplicitStringAttribute("sample", "tag", "value")
	require.True(t, implicit.Arguments[0].Value.GetStringLiteralExpr().Implicit)
}

func TestDisableGenerateOptions(t *testing.T) {
	_, err := GetDisableGenerateAttribute(nil)
	require.Error(t, err)
	all, err := GetDisableGenerateAttribute([]*Attribute{{Name: core.DisableGenerateAttributeName}})
	require.NoError(t, err)
	require.True(t, all.Including("any", "any"))
	attribute := &Attribute{Name: core.DisableGenerateAttributeName, Arguments: []*Argument{{Value: NewMapLiteralExpression(&MapLiteralExpr{Entries: []*MapLiteralExpr_Entry{
		{Key: "go", Value: NewArrayLiteralExpressionFrom(NewStringLiteralExpressionFrom("model"), NewStringLiteralExpressionFrom(""), NewIntegerLiteralExpressionFrom(3))},
		{Key: "protobuf", Value: NewNullLiteralExpressionFrom()},
		{Key: "openapi", Value: NewArrayLiteralExpressionFrom(NewStringLiteralExpressionFrom("all"))},
	}})}}}
	options, err := GetDisableGenerateAttribute([]*Attribute{attribute})
	require.NoError(t, err)
	require.Equal(t, []string{"model"}, options["go"])
	require.True(t, options.Including("go", "model"))
	require.False(t, options.Including("go", "client"))
	require.True(t, options.Including("protobuf", ""))
	require.False(t, options.Including("protobuf", "client"))
	require.True(t, options.Including("openapi", "schema"))
	require.False(t, options.Including("unknown", ""))
	var zero GenerateOption
	require.NotPanics(t, func() { require.NoError(t, zero.Parse(attribute)) })
	require.Equal(t, options, zero)
}
