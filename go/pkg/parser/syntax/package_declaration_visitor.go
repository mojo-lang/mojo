package syntax

import (
	"github.com/mojo-lang/core/go/pkg/mojo/core"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
)

type PackageDeclarationVisitor struct {
	*BaseMojoParserVisitor
}

func NewPackageDeclarationVisitor() *PackageDeclarationVisitor {
	visitor := &PackageDeclarationVisitor{}
	return visitor
}

func (p *PackageDeclarationVisitor) VisitPackageDeclaration(ctx *PackageDeclarationContext) interface{} {
	packageDecl := &lang.PackageDecl{}
	packageDecl.Name = GetPackageIdentifier(ctx.PackageIdentifier())
	packageDecl.PackageLiteralExpr = GetObjectLiteral(ctx.ObjectLiteral())

	return parsePackage(packageDecl)
}

func parsePackage(decl *lang.PackageDecl) *lang.PackageDecl {
	decl.Package = &lang.Package{}
	decl.Package.FullName = decl.Name
	decl.Package.Name = lang.GetPackageName(decl.Name)

	literalExpr := decl.PackageLiteralExpr
	for _, field := range literalExpr.Fields {
		if name := field.Name.GetIdentifierExpr(); name != nil && name.Identifier != nil {
			switch name.Identifier.Name {
			case "authors":
				if array := field.Value.GetArrayLiteralExpr(); array != nil {
					for _, element := range array.Elements {
						if obj := element.GetObjectLiteralExpr(); obj != nil {
							author := parseAuthor(obj)
							if author != nil {
								decl.Package.Authors = append(decl.Package.Authors, author)
							}
						}
					}
				}
			case "license":
				if value := field.Value.GetStringLiteralExpr(); value != nil {
					decl.Package.License = value.Value
				}
			case "dependencies":
				if value := field.Value.GetObjectLiteralExpr(); value != nil {
					decl.Package.Dependencies = parseDependencies(value)
				}
			case "repository":
				if value := field.Value.GetStringLiteralExpr(); value != nil {
					url, _ := core.ParseUrl(value.Value)
					decl.Package.Repository = url
				}
			case "version":
				if value := field.Value.GetStringLiteralExpr(); value != nil {
					v, _ := core.ParseVersion(value.Value)
					decl.Package.Version = v
				}
			}
		}
	}
	return decl
}

func parseDependencies(obj *lang.ObjectLiteralExpr) map[string]*lang.Package_Requirement {
	dependencies := make(map[string]*lang.Package_Requirement)
	for _, f := range obj.Fields {
		if k := f.Name.GetStringLiteralExpr(); k != nil {
			pkgName := k.Value
			if v := f.Value.GetStringLiteralExpr(); v != nil {
				dependencies[pkgName] = &lang.Package_Requirement{
					Version: lang.NewPackageRequirementVersion(v.Value),
				}
			} else if v := f.Value.GetObjectLiteralExpr(); v != nil {
				dependencies[pkgName] = parseRequirement(v)
			}
		}
	}
	return dependencies
}

func parseRequirement(obj *lang.ObjectLiteralExpr) *lang.Package_Requirement {
	requirement := &lang.Package_Requirement{}
	for _, field := range obj.Fields {
		if name := field.Name.GetIdentifierExpr(); name != nil && name.Identifier != nil {
			switch name.Identifier.Name {
			case "path":
				if value := field.Value.GetStringLiteralExpr(); value != nil {
					requirement.Path = value.Value
				}
			case "version":
				if value := field.Value.GetStringLiteralExpr(); value != nil {
					requirement.Version = lang.NewPackageRequirementVersion(value.Value)
				}
			}
		}
	}
	return requirement
}

func parseAuthor(obj *lang.ObjectLiteralExpr) *lang.Package_Author {
	author := &lang.Package_Author{}
	for _, field := range obj.Fields {
		if name := field.Name.GetIdentifierExpr(); name != nil && name.Identifier != nil {
			switch name.Identifier.Name {
			case "author":
				if value := field.Value.GetStringLiteralExpr(); value != nil {
					author.Author = value.Value
				}
			case "organization":
				if value := field.Value.GetStringLiteralExpr(); value != nil {
					author.Organization = value.Value
				}
			}
		}
	}
	return author
}
