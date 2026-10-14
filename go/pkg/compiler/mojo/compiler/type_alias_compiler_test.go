package compiler

import (
	"testing"

	"github.com/mojo-lang/mojo/go/pkg/compiler/context"
	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/mojo-lang/mojo/go/pkg/mojo/lang"
	"github.com/stretchr/testify/require"
)

func aliasReference(decl *lang.TypeAliasDecl) *lang.NominalType {
	return &lang.NominalType{Name: decl.Name, PackageName: decl.PackageName, TypeDeclaration: lang.NewTypeDeclarationFromDeclaration(lang.NewTypeAliasDeclaration(decl))}
}

func TestTypeAliasCompilerExpandsChainAndPreservesAttributes(t *testing.T) {
	dependency := &lang.Identifier{FullName: "mojo.core.String", SourceFileName: "mojo/core/string.mojo"}
	base := &lang.NominalType{Name: "String", PackageName: "mojo.core", Attributes: lang.SetStringAttribute(nil, core.LabelAttributeName, "value")}
	inner := &lang.TypeAliasDecl{Name: "Text", Type: base, ResolvedIdentifiers: []*lang.Identifier{dependency}}
	outer := &lang.TypeAliasDecl{Name: "ID", Type: aliasReference(inner)}
	use := aliasReference(outer)
	use.Attributes = lang.SetIntegerAttribute(nil, core.NumberAttributeName, 7)
	use.Attributes = lang.SetStringAttribute(use.Attributes, lang.OriginalTypeAliasName, "ID<String>")
	compiler := NewTypeAliasCompiler(nil, "test", 1)
	ids, got, err := compiler.compileNominalType(context.Empty(), use)
	require.NoError(t, err)
	require.Equal(t, "mojo.core.String", got.GetFullName())
	require.NotSame(t, base, got)
	require.Equal(t, []*lang.Identifier{dependency}, ids)
	number, err := got.GetIntegerAttribute(core.NumberAttributeName)
	require.NoError(t, err)
	require.Equal(t, int64(7), number)
	label, err := got.GetStringAttribute(core.LabelAttributeName)
	require.NoError(t, err)
	require.Equal(t, "value", label)
	original, err := got.GetStringAttribute(lang.OriginalTypeAliasName)
	require.NoError(t, err)
	require.Equal(t, "ID<String>", original)
	require.Len(t, base.Attributes, 1)
	require.Equal(t, "ID", use.Name)
}

func TestTypeAliasCompilerDefersUnresolvedAndGenericTypes(t *testing.T) {
	genericAlias := &lang.TypeAliasDecl{Name: "Box", Type: &lang.NominalType{Name: "T"}, GenericParameters: []*lang.GenericParameter{{Name: "T"}}}
	instantiated := aliasReference(&lang.TypeAliasDecl{Name: "IDs", Type: &lang.NominalType{
		Name: "Array", GenericArguments: []*lang.NominalType{{Name: "String"}},
		TypeDeclaration: lang.NewStructTypeDeclaration(&lang.StructDecl{Name: "Array"}),
	}})
	genericUse := aliasReference(&lang.TypeAliasDecl{Name: "Alias", Type: &lang.NominalType{Name: "String"}})
	genericUse.GenericArguments = []*lang.NominalType{{Name: "String"}}
	for _, typ := range []*lang.NominalType{nil, {Name: "Unresolved"}, {Name: "Item", TypeDeclaration: lang.NewStructTypeDeclaration(&lang.StructDecl{Name: "Item"})}, aliasReference(genericAlias), genericUse, instantiated} {
		ids, got, err := (&TypeAliasCompiler{}).compileNominalType(context.Empty(), typ)
		require.NoError(t, err)
		require.Nil(t, got)
		require.Empty(t, ids)
	}
}

func TestTypeAliasCompilerRebuildsDependencies(t *testing.T) {
	dependency := &lang.Identifier{FullName: "external.Value", SourceFileName: "external/value.mojo"}
	local := &lang.Identifier{FullName: "sample.Local", SourceFileName: "sample/item.mojo"}
	alias := &lang.TypeAliasDecl{Name: "Alias", Type: &lang.NominalType{Name: "Value", PackageName: "external"}, ResolvedIdentifiers: []*lang.Identifier{dependency}}
	field := &lang.ValueDecl{Name: "value", Type: aliasReference(alias)}
	argument := aliasReference(alias)
	nested := &lang.StructDecl{Name: "Nested", Type: &lang.StructType{Fields: []*lang.ValueDecl{field}}}
	decl := &lang.StructDecl{Name: "Item", StructDecls: []*lang.StructDecl{nested}, Type: &lang.StructType{
		Inherits: []*lang.NominalType{aliasReference(alias)},
		Fields:   []*lang.ValueDecl{{Name: "values", Type: &lang.NominalType{Name: "Array", GenericArguments: []*lang.NominalType{argument}}}},
	}}
	file := &lang.SourceFile{FullName: "sample/item.mojo", Statements: []*lang.Statement{
		lang.NewStructDeclStatement(decl),
		lang.NewEnumDeclStatement(&lang.EnumDecl{Name: "Status", ResolvedIdentifiers: []*lang.Identifier{local, dependency}}),
		lang.NewTypeAliasDeclStatement(alias),
		lang.NewInterfaceDeclStatement(&lang.InterfaceDecl{Name: "Service", ResolvedIdentifiers: []*lang.Identifier{dependency}}),
	}}
	childFile := &lang.SourceFile{FullName: "sample/child.mojo", Statements: []*lang.Statement{lang.NewTypeAliasDeclStatement(alias)}}
	pkg := &lang.Package{FullName: "sample", SourceFiles: []*lang.SourceFile{file}, Children: []*lang.Package{{FullName: "sample.child", SourceFiles: []*lang.SourceFile{childFile}}}}
	require.NoError(t, (&TypeAliasCompiler{}).CompilePackage(context.Empty(), pkg))
	require.Equal(t, "external.Value", field.Type.GetFullName())
	require.Equal(t, "external.Value", decl.Type.Inherits[0].GetFullName())
	require.Equal(t, "external.Value", decl.Type.Fields[0].Type.GenericArguments[0].GetFullName())
	require.Equal(t, []*lang.Identifier{dependency}, decl.ResolvedIdentifiers)
	require.Equal(t, []*lang.Identifier{dependency}, file.ResolvedIdentifiers)
	require.Equal(t, []*lang.Identifier{dependency}, childFile.ResolvedIdentifiers)
}
