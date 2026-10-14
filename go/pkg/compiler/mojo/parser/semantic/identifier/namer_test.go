package identifier

import (
	"testing"

	"github.com/mojo-lang/mojo/go/pkg/compiler/context"
	"github.com/mojo-lang/mojo/go/pkg/compiler/util"
	"github.com/mojo-lang/mojo/go/pkg/mojo/lang"
	"github.com/stretchr/testify/require"
)

func TestNamerDeclaresNestedTypesAndKeepsGenericParametersLocal(t *testing.T) {
	nested := &lang.StructDecl{Name: "Nested"}
	enum := &lang.EnumDecl{Name: "Status"}
	alias := &lang.TypeAliasDecl{Name: "Alias", GenericParameters: []*lang.GenericParameter{{Name: "V"}}}
	outer := &lang.StructDecl{Name: "Outer", StructDecls: []*lang.StructDecl{nested}, EnumDecls: []*lang.EnumDecl{enum}, TypeAliasDecls: []*lang.TypeAliasDecl{alias}, GenericParameters: []*lang.GenericParameter{{Name: "T"}}}
	method := &lang.FunctionDecl{Name: "Get"}
	service := &lang.InterfaceDecl{Name: "Service", Type: &lang.InterfaceType{Methods: []*lang.FunctionDecl{method}}}
	file := &lang.SourceFile{Name: "types.mojo", FullName: "sample/types.mojo", PackageName: "sample", Statements: []*lang.Statement{
		lang.NewStructDeclStatement(outer), lang.NewInterfaceDeclStatement(service),
		lang.NewEnumDeclStatement(&lang.EnumDecl{Name: "TopStatus"}),
		lang.NewTypeAliasDeclStatement(&lang.TypeAliasDecl{Name: "TopAlias"}),
	}}
	pkg := &lang.Package{Name: "sample", FullName: "sample", SourceFiles: []*lang.SourceFile{file}}
	ctx := context.WithScope(context.Empty())
	namer := NewNamer(nil)
	require.NoError(t, namer.ParsePackage(ctx, pkg))
	for _, name := range []string{"Outer", "Outer.Nested", "Outer.Status", "Outer.Alias", "Service", "TopStatus", "TopAlias"} {
		id := pkg.GetScope().Identifiers[name]
		require.NotNil(t, id, name)
		require.Equal(t, file.FullName, id.SourceFileName, name)
		require.Equal(t, "sample", id.PackageName, name)
		require.Same(t, id, context.Scope(ctx).Identifiers["sample."+name])
	}
	require.NotNil(t, outer.GetScope().Identifiers["T"])
	require.NotNil(t, alias.GetScope().Identifiers["V"])
	require.NotContains(t, pkg.GetScope().Identifiers, "Outer.T")
	require.NotContains(t, pkg.GetScope().Identifiers, "Outer.Alias.V")
	require.Equal(t, "sample", nested.PackageName)
	require.Equal(t, file.FullName, nested.SourceFileName)
	require.Equal(t, "sample", method.PackageName)
	require.Equal(t, file.FullName, method.SourceFileName)
	require.Equal(t, "Service", method.Enclosing.Name)
	require.True(t, util.IsPackageProcessed(pkg, namerName))
	count := len(pkg.GetScope().Identifiers)
	require.NoError(t, namer.ParsePackage(ctx, pkg))
	require.Len(t, pkg.GetScope().Identifiers, count)
}

func TestNamerExportsChildPackageNames(t *testing.T) {
	file := &lang.SourceFile{FullName: "root/child/item.mojo", PackageName: "root.child", Statements: []*lang.Statement{lang.NewStructDeclStatement(&lang.StructDecl{Name: "Item"})}}
	child := &lang.Package{Name: "child", FullName: "root.child", SourceFiles: []*lang.SourceFile{file}}
	root := &lang.Package{Name: "root", FullName: "root", Children: []*lang.Package{child}}
	ctx := context.WithScope(context.Empty())
	require.NoError(t, NewNamer(nil).ParsePackage(ctx, root))
	id := context.Scope(ctx).Identifiers["root.child.Item"]
	require.NotNil(t, id)
	require.Equal(t, "root.child", id.PackageName)
	require.Same(t, id, root.GetScope().Identifiers["child.Item"])
}

func TestNamerRejectsEmptyDeclarationWithoutMarkingPackage(t *testing.T) {
	file := &lang.SourceFile{Statements: []*lang.Statement{{Statement: &lang.Statement_Declaration{}}}}
	pkg := &lang.Package{Name: "invalid", FullName: "invalid", SourceFiles: []*lang.SourceFile{file}}
	err := NewNamer(nil).ParsePackage(context.WithScope(context.Empty()), pkg)
	require.ErrorContains(t, err, "declaration statement is nil")
	require.False(t, util.IsPackageProcessed(pkg, namerName))
}
