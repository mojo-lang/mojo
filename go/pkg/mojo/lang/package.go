package lang

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/mojo-lang/mojo/go/pkg/mojo/core/strcase"
)

func NewGlobalPackage() *Package {
	return &Package{
		Summary:  "Global",
		Scope:    NewScope(),
		Implicit: true,
	}
}

func GetPackageName(fullPkgName string) string {
	if len(fullPkgName) == 0 {
		return ""
	}

	segments := strings.Split(fullPkgName, ".")
	return segments[len(segments)-1]
}

func GetPackageParentName(fullPkgName string) string {
	if len(fullPkgName) == 0 {
		return ""
	}

	segments := strings.Split(fullPkgName, ".")
	if len(segments) > 1 {
		return strings.Join(segments[:len(segments)-1], ".")
	}
	return ""
}

func GetPackageParentNames(fullPkgName string) []string {
	segments := strings.Split(fullPkgName, ".")
	var parents []string
	for i := 0; i < len(segments)-1; i++ {
		parents = append(parents, strings.Join(segments[:i+1], "."))
	}
	return parents
}

func GetGoPackageName(fullPkgName string) string {
	segments := strings.Split(fullPkgName, ".")
	if len(segments) > 0 {
		name := segments[len(segments)-1]
		if core.IsVersionTag(name) {
			if len(segments) > 1 {
				name = segments[len(segments)-2]
			}
		}
		return name
	}

	return ""
}

// PackageNameToPath the segment of the package name will to be a folder, using kebab-case style
func PackageNameToPath(fullName string) string {
	return strings.ReplaceAll(strcase.ToKebabWithIgnore(fullName, "."), ".", "/")
}

func (x *Package) ParentName() string {
	return GetPackageParentName(x.FullName)
}

func (x *Package) ParentNames() []string {
	return GetPackageParentNames(x.FullName)
}

func (x *Package) IsGlobal() bool {
	return len(x.Name) == 0 && x.Summary == "Global" && x.Scope != nil
}

// IsPadding auto generated package, used to padding package tree
// Global Package is also a padding package
func (x *Package) IsPadding() bool {
	return x.Implicit && len(x.SourceFiles) == 0
}

var versionTag = regexp.MustCompile(`^(v[0-9]+)|(alpha[0-9]*)|(beta[0-9]*)`)

func (x *Package) IsVersionTag() bool {
	if x != nil && versionTag.MatchString(x.Name) {
		return true
	}
	return false
}

func (x *Package) SetScope(scope *Scope) {
	if x != nil {
		x.Scope = scope
	}
}

func (x *Package) GetAllPackageArray() []*Package {
	packages := []*Package{x}

	for _, pkg := range x.Children {
		ps := pkg.GetAllPackageArray()
		packages = append(packages, ps...)
	}
	return packages
}

func (x *Package) GetAllPackages() map[string]*Package {
	if x != nil {
		packages := make(map[string]*Package)
		packages[x.FullName] = x

		for _, pkg := range x.Children {
			ps := pkg.GetAllPackages()
			for k, v := range ps {
				packages[k] = v
			}
		}
		return packages
	}
	return nil
}

func (x *Package) GetAllDependentPackages() map[string]*Package {
	if x != nil {
		packages := make(map[string]*Package)
		for _, d := range x.ResolvedDependencies {
			pkgs := d.GetAllPackages()
			for k, v := range pkgs {
				packages[k] = v
			}
		}
		return packages
	}

	return nil
}

func (x *Package) GetAllPackageCount() int {
	total := 1
	for _, pkg := range x.Children {
		total += pkg.GetAllPackageCount()
	}
	return total
}

func (x *Package) SetExtraBool(key string, value bool) {
	if x.ExtraInfo == nil {
		x.ExtraInfo = &core.Object{}
	}
	x.ExtraInfo.SetBool(key, value)
}

func (x *Package) SetExtraString(key string, value string) {
	if x.ExtraInfo == nil {
		x.ExtraInfo = &core.Object{}
	}
	x.ExtraInfo.SetString(key, value)
}

func (x *Package) GetExtraBool(key string) bool {
	if x.ExtraInfo == nil {
		return false
	}
	return x.ExtraInfo.GetBool(key)
}

func (x *Package) GetExtraString(key string) string {
	if x.ExtraInfo == nil {
		return ""
	}
	return x.ExtraInfo.GetString(key)
}

func (x *Package) GetSourceFile(name string) *SourceFile {
	if x != nil {
		for _, file := range x.SourceFiles {
			if file.FullName == name {
				return file
			}
		}
	}
	return nil
}

func (x *Package) GetOrganization() string {
	if x == nil {
		return ""
	}

	// first find only has organization
	for _, author := range x.Authors {
		if len(author.Author) == 0 && len(author.Organization) > 0 {
			return author.Organization
		}
	}

	for _, author := range x.Authors {
		if len(author.Organization) > 0 {
			return author.Organization
		}
	}

	return ""
}

func (x *Package) HasChild(name string) bool {
	for _, p := range x.Children {
		if p.FullName == name {
			return true
		}
	}
	return false
}

func (x *Package) GoModName() string {
	if x.Repository != nil {
		return fmt.Sprintf("%s%s/go", x.Repository.GetAuthority().GetHost(), x.Repository.GetPath())
	}
	return ""
}

func (x *Package) GoFullPackageName() string {
	if len(x.FullName) > 0 {
		segments := strings.Split(x.FullName, ".")
		for i := 0; i < len(segments); i++ {
			segments[i] = strcase.ToKebab(segments[i])
		}
		fullName := strings.Join(segments, "/")
		return fmt.Sprintf("%s/pkg/%s", x.GoModName(), fullName)
	}
	return ""
}

func (x *Package) GoPackageName() string {
	if x.Repository != nil {
		name := x.Name
		if core.IsVersionTag(name) {
			segments := strings.Split(x.FullName, ".")
			if len(segments) > 1 {
				name = segments[len(segments)-2]
			}
		}
		return name
	}
	return ""
}

func (x *Package) GetIdentifier(name string) *Identifier {
	if x != nil {
		id := x.Scope.GetIdentifier(name)
		if id != nil {
			return id
		}

		for _, pkg := range x.Children {
			id = pkg.GetIdentifier(name)
			if id != nil {
				return id
			}
		}

		for _, pkg := range x.ResolvedDependencies {
			id = pkg.GetIdentifier(name)
			if id != nil {
				return id
			}
		}
	}
	return nil
}

// GetResolvedIdentifier resolved identifier may be defined in self package or defined in the dependent package
func (x *Package) GetResolvedIdentifier(fullName string) *Identifier {
	if x != nil {
		for _, file := range x.SourceFiles {
			if id := FindIdentifier(file.ResolvedIdentifiers, fullName); id != nil {
				return id
			}
		}

		for _, pkg := range x.ResolvedDependencies {
			if id := pkg.GetResolvedIdentifier(fullName); id != nil {
				return id
			}
		}

		return x.GetIdentifier(fullName)
	}
	return nil
}

func (x *Package) DirectlyContainsService() bool {
	if x != nil {
		if len(x.Children) > 0 {
			for _, pkg := range x.Children {
				if pkg.IsVersionTag() && pkg.directlyContainsService() {
					return true
				}
			}
		}

		return x.directlyContainsService()
	}
	return false
}

func (x *Package) directlyContainsService() bool {
	if x != nil {
		for _, file := range x.SourceFiles {
			for _, statement := range file.Statements {
				if i := statement.GetDeclaration().GetInterfaceDecl(); i != nil {
					return true
				}
			}
		}
	}
	return false
}

func (x *Package) GetEntityNode(name string) *EntityNode {
	if x != nil && x.EntityRelationSet != nil {
		typeName := GetTypeTypeName(name)
		if node := x.EntityRelationSet.GetNode(typeName); node != nil {
			return node
		}
		if node := x.EntityRelationSet.GetNode(name); node != nil {
			return node
		}

		for _, pkg := range x.Children {
			if node := pkg.GetEntityNode(name); node != nil {
				return node
			}
		}

		for _, pkg := range x.ResolvedDependencies {
			if node := pkg.GetEntityNode(name); node != nil {
				return node
			}
		}
	}
	return nil
}

func (x *Package) SetEntityNode(name string, node *EntityNode) *Package {
	if x != nil {
		if x.EntityRelationSet == nil {
			x.EntityRelationSet = NewEntityRelationSet()
		}
		x.EntityRelationSet.Nodes[name] = node
	}
	return x
}

func (x *Package) GetEntityEdge(name string) *EntityEdge {
	if x != nil && x.EntityRelationSet != nil {
		return x.EntityRelationSet.Edges[name]
	}
	return nil
}

func (x *Package) SetEntityEdge(name string, edge *EntityEdge) *Package {
	if x != nil {
		if x.EntityRelationSet == nil {
			x.EntityRelationSet = NewEntityRelationSet()
		}
		x.EntityRelationSet.Edges[name] = edge
	}
	return x
}

func (x *Package) DeleteEntityEdge(name string) *Package {
	if x != nil && x.EntityRelationSet != nil {
		delete(x.EntityRelationSet.Edges, name)
	}
	return x
}

func (x *Package) GetGoPackageImport() string {
	if x != nil {
		repository := x.Repository
		if repository != nil {
			goPackageFullName := PackageNameToPath(x.FullName)
			return fmt.Sprintf("%s%s/go/pkg/%s", repository.Authority.Host, repository.Path, goPackageFullName)
		}
		return ""
	}
	return ""
}
