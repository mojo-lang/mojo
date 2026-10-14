package lang

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTypeDeclarationsPreserveMetadataAndDependencies(t *testing.T) {
	scope, enclosing := NewScope(), &NominalType{Name: "Outer"}
	old := &Identifier{FullName: "sample.Existing"}
	added := &Identifier{FullName: "sample.Added"}
	for _, declaration := range []*Declaration{
		NewStructDeclaration(&StructDecl{Name: "Item", PackageName: "sample", SourceFileName: "sample.mojo", Implicit: true, Scope: scope, Enclosing: enclosing, ResolvedIdentifiers: []*Identifier{old}}),
		NewEnumDeclaration(&EnumDecl{Name: "Status", PackageName: "sample", SourceFileName: "sample.mojo", Implicit: true, Scope: scope, Enclosing: enclosing, ResolvedIdentifiers: []*Identifier{old}}),
		NewInterfaceDeclaration(&InterfaceDecl{Name: "Service", PackageName: "sample", SourceFileName: "sample.mojo", Implicit: true, Scope: scope, Enclosing: enclosing, ResolvedIdentifiers: []*Identifier{old}}),
		NewTypeAliasDeclaration(&TypeAliasDecl{Name: "ID", PackageName: "sample", SourceFileName: "sample.mojo", Implicit: true, Scope: scope, Enclosing: enclosing, ResolvedIdentifiers: []*Identifier{old}}),
	} {
		typ := NewTypeDeclarationFromDeclaration(declaration)
		require.Same(t, scope, typ.GetScope())
		require.Same(t, enclosing, typ.GetEnclosingType())
		require.Equal(t, "sample.mojo", typ.GetSourceFileName())
		require.True(t, typ.Implicit())
		require.NotNil(t, typ.GetDecl())
		require.False(t, typ.IsGeneric())
		typ.Merge([]*Identifier{added, old})
		getter := typ.GetDecl().(interface{ GetResolvedIdentifiers() []*Identifier })
		require.ElementsMatch(t, []*Identifier{old, added}, getter.GetResolvedIdentifiers())
		nominal := &NominalType{Name: "Item", PackageName: "sample", TypeDeclaration: typ, Enclosing: enclosing}
		identifier := nominal.NewIdentifier()
		require.NotEqual(t, Identifier_KIND_UNSPECIFIED, identifier.Kind)
		require.Equal(t, "sample.Outer.Item", identifier.FullName)
		require.Equal(t, "sample.mojo", identifier.SourceFileName)
		require.True(t, identifier.Implicit)
		require.Same(t, typ.GetDecl(), identifier.ToNominalType().TypeDeclaration.GetDecl())
	}
	generic := NewTypeDeclarationFromDeclaration(NewGenericParameterDeclaration(&GenericParameter{Name: "T", Enclosing: enclosing}))
	require.NotNil(t, generic.GetDecl())
	require.Same(t, enclosing, generic.GetEnclosingType())
	require.False(t, generic.Implicit())
	require.Nil(t, generic.GetScope())
	require.Nil(t, NewTypeDeclarationFromDeclaration(NewImportDeclaration(&ImportDecl{})))
}

func TestNominalTypeClassificationAndGenericNames(t *testing.T) {
	for _, name := range []string{"Bool", "Int8", "Int16", "Int32", "Int64", "UInt8", "UInt16", "UInt32", "UInt64", "Float32", "Float64", "String", "Bytes", "Positive", "Negative"} {
		typ := &NominalType{TypeDeclaration: NewStructTypeDeclaration(&StructDecl{Name: name, PackageName: "mojo.core"})}
		require.True(t, typ.IsStruct())
		require.True(t, typ.IsScalar(), name)
	}
	for _, name := range []string{"Byte", "Int", "UInt", "Size", "Float", "Double"} {
		typ := &NominalType{TypeDeclaration: NewTypeDeclarationFromDeclaration(NewTypeAliasDeclaration(&TypeAliasDecl{Name: name, PackageName: "mojo.core"}))}
		require.True(t, typ.IsTypeAlias())
		require.True(t, typ.IsScalar(), name)
	}
	for _, name := range []string{"Array", "Map", "Tuple", "Intersection", "Union"} {
		typ := &NominalType{TypeDeclaration: NewStructTypeDeclaration(&StructDecl{Name: name, PackageName: "mojo.core"})}
		require.False(t, typ.IsScalar())
		require.Equal(t, name == "Array", typ.IsArrayType())
		require.Equal(t, name == "Map", typ.IsMapType())
		require.Equal(t, name == "Tuple", typ.IsTupleType())
		require.Equal(t, name == "Intersection", typ.IsIntersectionType())
		require.Equal(t, name == "Union", typ.IsUnionType())
	}
	parameter := &NominalType{Name: "T", TypeDeclaration: NewTypeDeclarationFromDeclaration(NewGenericParameterDeclaration(&GenericParameter{Name: "T"}))}
	array := NewArrayNominalType(parameter)
	require.True(t, parameter.IsGenericParameter())
	require.True(t, array.IsGeneric())
	require.False(t, array.IsInstantiatedGeneric())
	array.GenericArguments[0] = &NominalType{PackageName: "sample", Name: "Item"}
	require.True(t, array.IsInstantiatedGeneric())
	require.Equal(t, "mojo.core.Array<sample.Item>", array.GetGenericFullName())
	mapping := NewStringMapNominalType(array)
	require.Equal(t, "mojo.core.Map<mojo.core.String,mojo.core.Array<sample.Item>>", mapping.GetGenericFullName())
	require.Equal(t, mapping.GetGenericFullName(), NewMapNominalType(StringNominalType, array).GetGenericFullName())
	require.Equal(t, "mojo.core.Union<mojo.core.String,sample.Item>", NewUnionNominalType(StringNominalType, array.GenericArguments[0]).GetGenericFullName())
	var absent *NominalType
	require.False(t, absent.IsScalar())
	require.False(t, absent.IsGeneric())
	require.False(t, absent.IsStruct())
	require.False(t, absent.IsEnum())
	require.False(t, absent.IsInterface())
	require.False(t, absent.IsTypeAlias())
	require.False(t, absent.IsArrayType())
	require.Nil(t, absent.NewIdentifier())
	require.Empty(t, absent.GetGenericName())
	require.Empty(t, absent.GetGenericFullName())
}

func TestIdentifierEnclosingOrderAndSourceFileTypes(t *testing.T) {
	id := NewIdentifier("Alpha", "Zulu", "Middle", "Item")
	require.Equal(t, []string{"Alpha", "Zulu", "Middle"}, id.EnclosingNames(), "nested identifiers follow lexical nesting, not alphabetical order")
	require.Equal(t, []string{"Alpha", "Zulu", "Middle", "Item"}, id.FullNames())
	require.Nil(t, NewIdentifier())
	require.True(t, (&Identifier{Name: "Array<T>"}).IsGenericInstantiated())
	structure, alias := &StructDecl{Name: "Item"}, &TypeAliasDecl{Name: "ID"}
	file := &SourceFile{Name: "Array<T>", Statements: []*Statement{
		NewStructDeclStatement(structure), NewTypeAliasDeclStatement(alias), NewEnumDeclStatement(&EnumDecl{Name: "Status"}), NewInterfaceDeclStatement(&InterfaceDecl{Name: "Store"}), NewExpressionStatement(NewStringLiteralExpressionFrom("skip")),
	}}
	require.True(t, file.IsGenericInstantiated())
	require.Len(t, file.GetTypeDeclarations(), 4)
	require.Equal(t, []*StructDecl{structure}, file.GetStructDecls())
	require.Equal(t, []*TypeAliasDecl{alias}, file.GetTypeAliasDecls())
	scope := NewScope()
	file.SetScope(scope)
	require.Same(t, scope, file.Scope)
}
