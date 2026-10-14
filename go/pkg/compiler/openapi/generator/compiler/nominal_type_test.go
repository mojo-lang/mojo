package compiler

import (
	"testing"

	"github.com/mojo-lang/mojo/go/pkg/compiler/context"
	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/mojo-lang/mojo/go/pkg/mojo/lang"
	"github.com/mojo-lang/mojo/go/pkg/mojo/openapi"
	"github.com/stretchr/testify/require"
)

func coreType(name string, arguments ...*lang.NominalType) *lang.NominalType {
	return &lang.NominalType{PackageName: "mojo.core", Name: name, GenericArguments: arguments}
}

func TestNominalTypeCompilerScalars(t *testing.T) {
	for _, tt := range []struct {
		name   string
		typ    openapi.Schema_Type
		format string
	}{
		{"String", openapi.Schema_TYPE_STRING, ""}, {"Bytes", openapi.Schema_TYPE_STRING, "Bytes"},
		{"Bool", openapi.Schema_TYPE_BOOLEAN, ""}, {"Null", openapi.Schema_TYPE_NULL, ""},
		{"Int", openapi.Schema_TYPE_INTEGER, ""}, {"UInt", openapi.Schema_TYPE_INTEGER, ""},
		{"Int32", openapi.Schema_TYPE_INTEGER, "Int32"}, {"UInt64", openapi.Schema_TYPE_INTEGER, "UInt64"},
		{"Positive", openapi.Schema_TYPE_INTEGER, "Positive"}, {"Negative", openapi.Schema_TYPE_INTEGER, "Negative"},
		{"Float32", openapi.Schema_TYPE_NUMBER, "Float32"}, {"Float64", openapi.Schema_TYPE_NUMBER, "Float64"},
		{"Float", openapi.Schema_TYPE_NUMBER, "Float32"}, {"Double", openapi.Schema_TYPE_NUMBER, "Float64"},
	} {
		t.Run(tt.name, func(t *testing.T) {
			compiler := &NominalTypeCompiler{}
			got, err := compiler.Compile(coreType(tt.name))
			require.NoError(t, err)
			require.Equal(t, tt.typ, got.Type)
			require.Equal(t, tt.format, got.Format)
			got.Format = "changed"
			again, err := compiler.Compile(coreType(tt.name))
			require.NoError(t, err)
			require.Equal(t, tt.format, again.Format)
		})
	}
}

func TestNominalTypeCompilerContainersAndConstants(t *testing.T) {
	compiler := &NominalTypeCompiler{}
	array := coreType("Array", coreType("String"))
	array.Attributes = lang.SetIntegerAttribute(nil, core.MinLengthAttributeName, 2)
	array.Attributes = lang.SetIntegerAttribute(array.Attributes, core.MaxLengthAttributeName, 5)
	got, err := compiler.Compile(array)
	require.NoError(t, err)
	require.Equal(t, openapi.Schema_TYPE_ARRAY, got.Type)
	require.Equal(t, uint64(2), got.MinItems)
	require.Equal(t, uint64(5), got.MaxItems)
	require.Equal(t, openapi.Schema_TYPE_STRING, got.Items.GetSchema().Type)
	got, err = compiler.Compile(coreType("Map", coreType("String"), coreType("Int32")))
	require.NoError(t, err)
	require.Equal(t, openapi.Schema_TYPE_OBJECT, got.Type)
	require.Equal(t, openapi.Schema_TYPE_INTEGER, got.AdditionalProperties.GetSchema().Type)
	got, err = compiler.Compile(coreType("Union", coreType("String"), coreType("Bool")))
	require.NoError(t, err)
	require.Len(t, got.OneOf, 2)
	require.Equal(t, openapi.Schema_TYPE_STRING, got.OneOf[0].GetSchema().Type)
	require.Equal(t, openapi.Schema_TYPE_BOOLEAN, got.OneOf[1].GetSchema().Type)
	constant := coreType("String")
	constant.Attributes = lang.SetStringAttribute(nil, core.ConstAttributeName, "ready")
	got, err = compiler.Compile(constant)
	require.NoError(t, err)
	require.Len(t, got.Enum, 1)
	require.Equal(t, "ready", got.Enum[0].GetString())
}

func TestNominalTypeCompilerEnumAndFormatOverride(t *testing.T) {
	enum := &lang.EnumDecl{Name: "State", PackageName: "sample", Type: &lang.EnumType{Enumerators: []*lang.ValueDecl{
		{Name: "UNSPECIFIED"}, {Name: "Ready"}, {Name: "Running", Attributes: lang.SetStringAttribute(nil, core.AliasAttributeName, "active")},
	}}}
	typ := &lang.NominalType{Name: "State", PackageName: "sample", TypeDeclaration: lang.NewTypeDeclarationFromDeclaration(lang.NewEnumDeclaration(enum))}
	compiler := &NominalTypeCompiler{}
	got, err := compiler.Compile(typ)
	require.NoError(t, err)
	require.Equal(t, "sample.State", got.Title)
	require.Len(t, got.Enum, 2)
	require.Equal(t, "Ready", got.Enum[0].GetString())
	require.Equal(t, "active", got.Enum[1].GetString())
	typ.Attributes = []*lang.Attribute{{Name: core.TypeFormatAttributeName, GenericArguments: []*lang.NominalType{coreType("String")}}}
	got, err = compiler.Compile(typ)
	require.NoError(t, err)
	require.Equal(t, openapi.Schema_TYPE_STRING, got.Type)
	require.Empty(t, got.Enum)
}

func TestCompileNominalTypePropagatesUnknownTypeErrors(t *testing.T) {
	for _, typ := range []*lang.NominalType{
		{Name: "Unknown", PackageName: "sample"},
		coreType("Array", &lang.NominalType{Name: "Unknown", PackageName: "sample"}),
		coreType("Map", coreType("String"), &lang.NominalType{Name: "Unknown", PackageName: "sample"}),
		coreType("Union", coreType("String"), &lang.NominalType{Name: "Unknown", PackageName: "sample"}),
	} {
		schema, err := compileNominalType(context.Empty(), typ)
		require.ErrorContains(t, err, "sample.Unknown")
		require.Nil(t, schema)
	}
}
