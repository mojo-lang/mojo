package lang

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPackageTreeAndIdentifierResolution(t *testing.T) {
	root := &Package{Name: "sample", FullName: "sample", Scope: NewScope()}
	child := &Package{Name: "child", FullName: "sample.child", Scope: NewScope()}
	grandchild := &Package{Name: "v1", FullName: "sample.child.v1", Scope: NewScope()}
	dependency := &Package{Name: "shared", FullName: "shared", Scope: NewScope()}
	root.Children, child.Children = []*Package{child}, []*Package{grandchild}
	root.ResolvedDependencies = map[string]*Package{"shared": dependency}
	rootID := root.Scope.Declare(NewStructDeclaration(&StructDecl{Name: "Item", PackageName: "sample"}))
	childID := child.Scope.Declare(NewEnumDeclaration(&EnumDecl{Name: "Status", PackageName: "sample.child"}))
	dependencyID := dependency.Scope.Declare(NewStructDeclaration(&StructDecl{Name: "Common", PackageName: "shared"}))
	fileID := &Identifier{Name: "Remote", FullName: "external.Remote"}
	file := &SourceFile{FullName: "sample/item.mojo", ResolvedIdentifiers: []*Identifier{fileID}}
	root.SourceFiles = []*SourceFile{file}
	dependency.SourceFiles = []*SourceFile{{ResolvedIdentifiers: []*Identifier{{FullName: "external.Other"}}}}
	require.Equal(t, []*Package{root, child, grandchild}, root.GetAllPackageArray())
	require.Equal(t, 3, root.GetAllPackageCount())
	require.Equal(t, map[string]*Package{"sample": root, "sample.child": child, "sample.child.v1": grandchild}, root.GetAllPackages())
	require.Equal(t, map[string]*Package{"shared": dependency}, root.GetAllDependentPackages())
	require.True(t, root.HasChild("sample.child"))
	require.False(t, root.HasChild("shared"))
	require.Same(t, rootID, root.GetIdentifier("Item"))
	require.Same(t, childID, root.GetIdentifier("Status"))
	require.Same(t, dependencyID, root.GetIdentifier("Common"))
	require.Nil(t, root.GetIdentifier("missing"))
	require.Same(t, fileID, root.GetResolvedIdentifier("external.Remote"))
	require.NotNil(t, root.GetResolvedIdentifier("external.Other"))
	require.Same(t, rootID, root.GetResolvedIdentifier("Item"))
	require.Nil(t, root.GetResolvedIdentifier("missing"))
	require.Same(t, file, root.GetSourceFile(file.FullName))
	require.Nil(t, root.GetSourceFile("missing"))
	require.Equal(t, "sample.child", grandchild.ParentName())
	require.Equal(t, []string{"sample", "sample.child"}, grandchild.ParentNames())
}

func TestPackageMetadataAndServiceDetection(t *testing.T) {
	global := NewGlobalPackage()
	require.True(t, global.IsGlobal())
	require.True(t, global.IsPadding())
	global.SourceFiles = []*SourceFile{{}}
	require.False(t, global.IsPadding())
	pkg := &Package{}
	require.False(t, pkg.GetExtraBool("enabled"))
	require.Empty(t, pkg.GetExtraString("name"))
	pkg.SetExtraBool("enabled", true)
	pkg.SetExtraString("name", "sample")
	require.True(t, pkg.GetExtraBool("enabled"))
	require.Equal(t, "sample", pkg.GetExtraString("name"))
	pkg.Authors = []*Package_Author{{Author: "Ada", Organization: "Team"}, {Organization: "Company"}}
	require.Equal(t, "Company", pkg.GetOrganization())
	pkg.Authors = pkg.Authors[:1]
	require.Equal(t, "Team", pkg.GetOrganization())
	pkg.Authors = nil
	require.Empty(t, pkg.GetOrganization())
	service := &SourceFile{Statements: []*Statement{NewStructDeclStatement(&StructDecl{Name: "Item"}), NewInterfaceDeclStatement(&InterfaceDecl{Name: "Store"})}}
	version := &Package{Name: "v1", SourceFiles: []*SourceFile{service}}
	pkg.Children = []*Package{version}
	require.True(t, pkg.DirectlyContainsService())
	version.Name = "internal"
	require.False(t, pkg.DirectlyContainsService())
	pkg.SourceFiles = []*SourceFile{service}
	require.True(t, pkg.DirectlyContainsService())
	for _, name := range []string{"v1", "v12", "alpha", "alpha2", "beta3"} {
		require.True(t, (&Package{Name: name}).IsVersionTag(), name)
	}
	for _, name := range []string{"", "preview", "v1extra", "mybeta", "beta2more"} {
		require.False(t, (&Package{Name: name}).IsVersionTag(), name)
	}
}

func TestPackageEntityLookupAcrossChildrenAndDependencies(t *testing.T) {
	root, child, dependency := &Package{}, &Package{}, &Package{}
	root.Children = []*Package{child}
	root.ResolvedDependencies = map[string]*Package{"dep": dependency}
	local, nested, remote := &EntityNode{}, &EntityNode{}, &EntityNode{}
	child.SetEntityNode("Nested", nested)
	dependency.SetEntityNode("Remote", remote)
	require.Same(t, nested, root.GetEntityNode("sample.Nested"), "parents without a local relation set must still search children")
	require.Same(t, remote, root.GetEntityNode("sample.Remote"))
	require.Same(t, root, root.SetEntityNode("sample.Local", local))
	require.Same(t, local, root.GetEntityNode("sample.Local"))
	require.Nil(t, root.GetEntityNode("Missing"))
	edge := &EntityEdge{}
	require.Same(t, root, root.SetEntityEdge("owns", edge))
	require.Same(t, edge, root.GetEntityEdge("owns"))
	require.Same(t, root, root.DeleteEntityEdge("owns"))
	require.Nil(t, root.GetEntityEdge("owns"))
	var absent *Package
	require.Nil(t, absent.GetAllPackages())
	require.Nil(t, absent.GetAllDependentPackages())
	require.Nil(t, absent.GetIdentifier("x"))
	require.Nil(t, absent.GetResolvedIdentifier("x"))
	require.Nil(t, absent.GetSourceFile("x"))
	require.Empty(t, absent.GetOrganization())
	require.False(t, absent.DirectlyContainsService())
	require.Nil(t, absent.SetEntityNode("x", local))
	require.Nil(t, absent.SetEntityEdge("x", edge))
	require.Nil(t, absent.DeleteEntityEdge("x"))
	require.Nil(t, absent.GetEntityNode("x"))
	require.Nil(t, absent.GetEntityEdge("x"))
}
