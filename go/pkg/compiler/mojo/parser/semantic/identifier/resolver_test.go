package identifier

import (
	"testing"

	"github.com/mojo-lang/mojo/go/pkg/compiler/context"
	"github.com/mojo-lang/mojo/go/pkg/compiler/util"
	"github.com/mojo-lang/mojo/go/pkg/mojo/lang"
	"github.com/stretchr/testify/require"
)

func TestResolverLinksDeclarationsAndDeduplicatesFileDependencies(t *testing.T) {
	external := &lang.StructDecl{Name: "Value", PackageName: "external", SourceFileName: "external/value.mojo"}
	ctx := context.WithScope(context.Empty())
	id := context.Scope(ctx).Declare(lang.NewStructDeclaration(external))
	id.SourceFileName = external.SourceFileName
	context.Scope(ctx).Identifiers["external.Value"] = id
	reference := func() *lang.NominalType { return &lang.NominalType{Name: "Value", PackageName: "external"} }
	local := &lang.StructDecl{Name: "Local"}
	field := &lang.ValueDecl{Name: "value", Type: reference()}
	nested := &lang.StructDecl{Name: "Nested", Enclosing: &lang.NominalType{Name: "Item"}, Type: &lang.StructType{Fields: []*lang.ValueDecl{{Name: "nested_value", Type: reference()}}}}
	enum := &lang.EnumDecl{Name: "Status", Enclosing: &lang.NominalType{Name: "Item"}, Type: &lang.EnumType{UnderlyingType: reference()}}
	alias := &lang.TypeAliasDecl{Name: "Alias", Enclosing: &lang.NominalType{Name: "Item"}, Type: reference()}
	item := &lang.StructDecl{Name: "Item", Type: &lang.StructType{Fields: []*lang.ValueDecl{field, {Name: "local", Type: &lang.NominalType{Name: "Local"}}}, Inherits: []*lang.NominalType{reference()}},
		StructDecls: []*lang.StructDecl{nested}, EnumDecls: []*lang.EnumDecl{enum}, TypeAliasDecls: []*lang.TypeAliasDecl{alias}}
	source := &lang.SourceFile{Name: "item.mojo", FullName: "sample/item.mojo", PackageName: "sample", Statements: []*lang.Statement{
		lang.NewStructDeclStatement(local), lang.NewStructDeclStatement(item),
		lang.NewEnumDeclStatement(&lang.EnumDecl{Name: "TopEnum", Type: &lang.EnumType{UnderlyingType: reference()}}),
		lang.NewTypeAliasDeclStatement(&lang.TypeAliasDecl{Name: "TopAlias", Type: reference()}),
		lang.NewInterfaceDeclStatement(&lang.InterfaceDecl{Name: "Service", Type: &lang.InterfaceType{Inherits: []*lang.NominalType{reference()}, Methods: []*lang.FunctionDecl{
			{Name: "Get", Signature: lang.NewFunctionSignature(reference(), &lang.ValueDecl{Name: "input", Type: reference()})},
		}}}),
	}}
	pkg := &lang.Package{Name: "sample", FullName: "sample", SourceFiles: []*lang.SourceFile{source}}
	require.NoError(t, NewNamer(nil).ParsePackage(ctx, pkg))
	resolver := NewResolver(nil)
	require.NoError(t, resolver.ParsePackage(ctx, pkg))
	require.Same(t, external, field.Type.TypeDeclaration.GetStructDecl())
	require.Same(t, local, item.Type.Fields[1].Type.TypeDeclaration.GetStructDecl())
	require.Same(t, item, nested.Enclosing.TypeDeclaration.GetStructDecl())
	require.Same(t, item, enum.Enclosing.TypeDeclaration.GetStructDecl())
	require.Same(t, item, alias.Enclosing.TypeDeclaration.GetStructDecl())
	require.Equal(t, []*lang.Identifier{id}, source.ResolvedIdentifiers)
	require.Empty(t, source.UnresolvedIdentifiers)
	require.True(t, util.IsPackageProcessed(pkg, resolverName))
	require.NoError(t, resolver.ParsePackage(ctx, pkg))
	require.Len(t, source.ResolvedIdentifiers, 1)
}

func TestResolverReportsMissingTypesInGenericAndAttributeArguments(t *testing.T) {
	typ := &lang.NominalType{Name: "Container", GenericArguments: []*lang.NominalType{{Name: "Missing"}}, Attributes: []*lang.Attribute{{Name: "format", GenericArguments: []*lang.NominalType{{Name: "Unknown"}}}}}
	decl := &lang.StructDecl{Name: "Item", Type: &lang.StructType{Fields: []*lang.ValueDecl{{Name: "value", Type: typ}, {Name: "again", Type: &lang.NominalType{Name: "Missing"}}}}}
	file := &lang.SourceFile{Name: "bad.mojo", FullName: "sample/bad.mojo", PackageName: "sample", Statements: []*lang.Statement{lang.NewStructDeclStatement(decl)}}
	pkg := &lang.Package{Name: "sample", FullName: "sample", SourceFiles: []*lang.SourceFile{file}}
	ctx := context.WithScope(context.Empty())
	require.NoError(t, NewNamer(nil).ParsePackage(ctx, pkg))
	err := NewResolver(nil).ParsePackage(ctx, pkg)
	require.ErrorContains(t, err, "bad.mojo")
	var names []string
	for _, id := range file.UnresolvedIdentifiers {
		names = append(names, id.Name)
	}
	require.ElementsMatch(t, []string{"Container", "Missing", "Unknown"}, names)
	require.False(t, util.IsPackageProcessed(pkg, resolverName))
}

func TestResolverResolvesEnclosedAndGenericParameterNames(t *testing.T) {
	ctx := context.WithScope(context.Empty())
	nested := &lang.StructDecl{Name: "Inner", PackageName: "sample", Enclosing: &lang.NominalType{Name: "Outer"}}
	id := context.Scope(ctx).Declare(lang.NewStructDeclaration(nested))
	context.Scope(ctx).Identifiers["Outer.Inner"] = id
	parameter := &lang.GenericParameter{Name: "T"}
	generic := context.Scope(ctx).Declare(lang.NewGenericParameterDeclaration(parameter))
	typ := &lang.NominalType{Name: "Inner", Enclosing: &lang.NominalType{Name: "Outer"}, GenericArguments: []*lang.NominalType{{Name: "T"}}}
	resolved, unresolved := resolveNominalType(ctx, typ)
	require.Empty(t, unresolved)
	require.ElementsMatch(t, []*lang.Identifier{id, generic}, resolved)
	require.Same(t, nested, typ.TypeDeclaration.GetStructDecl())
	require.Same(t, parameter, typ.GenericArguments[0].TypeDeclaration.GetGenericParameter())
	require.Equal(t, "sample", typ.PackageName)
}

func TestResolverPropagatesChildErrors(t *testing.T) {
	child := &lang.Package{Name: "child", FullName: "root.child", SourceFiles: []*lang.SourceFile{{Statements: []*lang.Statement{{Statement: &lang.Statement_Declaration{}}}}}}
	parent := &lang.Package{Name: "root", FullName: "root", Children: []*lang.Package{child}}
	err := NewResolver(nil).ParsePackage(context.WithScope(context.Empty()), parent)
	require.ErrorContains(t, err, "declaration statement is nil")
	require.False(t, util.IsPackageProcessed(parent, resolverName))
	require.False(t, util.IsPackageProcessed(child, resolverName))
}
