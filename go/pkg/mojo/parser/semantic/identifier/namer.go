package identifier

import (
	"errors"
	"strings"

	"github.com/mojo-lang/core/go/pkg/logs"
	"github.com/mojo-lang/core/go/pkg/mojo/core"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/go/pkg/mojo/context"
	"github.com/mojo-lang/mojo/go/pkg/mojo/plugin"
	"github.com/mojo-lang/mojo/go/pkg/util"
)

const namerName = "semantic.identifier-namer"

func init() {
	plugin.RegisterPlugin(NewNamer(nil))
}

type Namer struct {
	plugin.BasicPlugin
}

func NewNamer(options core.Options) *Namer {
	return &Namer{
		BasicPlugin: plugin.BasicPlugin{
			Name:          namerName,
			Group:         "semantic",
			GroupPriority: 3,
			Priority:      1,
			Creator: func(options core.Options) plugin.Plugin {
				return NewNamer(options)
			},
		},
	}
}

func (p *Namer) ParsePackage(ctx context.Context, pkg *lang.Package) error {
	if util.IsPackageProcessed(pkg, namerName) {
		logs.Infow("already processed, skip the plugin", "plugin", p.Name, "method", "ParsePackage", "pkg", pkg.FullName)
		return nil
	} else {
		logs.Infow("enter the plugin", "plugin", p.Name, "method", "ParsePackage", "pkg", pkg.FullName)
	}

	thisCtx := context.WithScopeType(ctx, pkg)
	thisScope := context.Scope(thisCtx)

	defer func() {
		scope := context.Scope(ctx)
		if scope != nil { // not global
			for key, value := range thisScope.Identifiers {
				// for core identifiers like String, ignore mojo.String to Global scope
				// only support mojo.core.String, String, core.String to Global scope
				if pkg.Name == "mojo" {
					if value.PackageName == "mojo.core" && !strings.HasPrefix(key, "core.") {
						scope.Identifiers[key] = value
						continue
					}
					// add core.String to Global scope
					if strings.HasPrefix(value.PackageName, "mojo.") {
						scope.Identifiers[key] = value
					}
				}
				scope.Identifiers[pkg.Name+"."+key] = value

				if value.PackageName == "mojo.core" {
					scope.Identifiers[key] = value
				}
			}
		}
	}()

	for _, file := range pkg.SourceFiles {
		if err := p.ParserSourceFile(thisCtx, file); err != nil {
			return err
		}
	}

	for _, child := range pkg.Children {
		if err := p.ParsePackage(thisCtx, child); err != nil {
			return err
		}
	}

	if !pkg.IsPadding() {
		util.SetPackageProcessed(pkg, namerName)
	}
	return nil
}

func (p *Namer) ParserSourceFile(ctx context.Context, file *lang.SourceFile) error {
	thisCtx := context.WithScopeType(ctx, file)
	thisScope := context.Scope(thisCtx)

	defer func() {
		scope := context.Scope(ctx)
		for key, value := range thisScope.Identifiers {
			scope.Identifiers[key] = value
		}
	}()

	pkg := context.Package(ctx)
	for _, statement := range file.Statements {
		switch statement.Statement.(type) {
		case *lang.Statement_Declaration:
			decl := statement.GetDeclaration()
			if decl == nil {
				return errors.New("declaration statement is nil")
			}

			if len(decl.GetPackageName()) == 0 {
				decl.SetPackageName(pkg.FullName)
			}

			identifier := thisScope.Declare(decl)
			if identifier != nil {
				identifier.SourceFileName = file.FullName
			}

			switch decl.Declaration.(type) {
			case *lang.Declaration_TypeAliasDecl:
				err := p.ParseTypeAlias(thisCtx, decl.GetTypeAliasDecl())
				if err != nil {
					return err
				}
			case *lang.Declaration_StructDecl:
				err := p.ParseStruct(thisCtx, decl.GetStructDecl())
				if err != nil {
					return err
				}
			case *lang.Declaration_EnumDecl:
				err := p.ParseEnum(thisCtx, decl.GetEnumDecl())
				if err != nil {
					return err
				}
			case *lang.Declaration_InterfaceDecl:
				err := p.ParseInterface(thisCtx, decl.GetInterfaceDecl())
				if err != nil {
					return err
				}
			}
		case *lang.Statement_Expression:
		default:
		}
	}
	return nil
}

func (p *Namer) ParseStruct(ctx context.Context, decl *lang.StructDecl) error {
	thisCtx := context.WithScopeType(ctx, decl)
	thisScope := context.Scope(thisCtx)

	file := context.SourceFile(ctx)
	decl.SourceFileName = file.FullName
	decl.PackageName = file.PackageName

	for _, declaration := range decl.StructDecls {
		declaration.PackageName = file.PackageName
		declaration.SourceFileName = file.FullName

		identifier := thisScope.Declare(lang.NewStructDeclaration(declaration))
		identifier.SourceFileName = file.FullName
		p.ParseStruct(thisCtx, declaration)
	}

	for _, declaration := range decl.EnumDecls {
		declaration.PackageName = file.PackageName
		declaration.SourceFileName = file.FullName

		identifier := thisScope.Declare(lang.NewEnumDeclaration(declaration))
		identifier.SourceFileName = file.FullName
		p.ParseEnum(thisCtx, declaration)
	}

	for _, declaration := range decl.TypeAliasDecls {
		declaration.PackageName = file.PackageName
		declaration.SourceFileName = file.FullName

		identifier := thisScope.Declare(lang.NewTypeAliasDeclaration(declaration))
		identifier.SourceFileName = file.FullName
		p.ParseTypeAlias(thisCtx, declaration)
	}

	for _, parameter := range decl.GenericParameters {
		identifier := thisScope.Declare(lang.NewGenericParameterDeclaration(parameter))
		identifier.SourceFileName = file.FullName
	}

	for key, value := range thisScope.Identifiers {
		if value.Kind != lang.Identifier_KIND_GENERIC_PARAMETER {
			context.Scope(ctx).Identifiers[decl.Name+"."+key] = value
		}
	}

	return nil
}

func (p *Namer) ParseEnum(ctx context.Context, decl *lang.EnumDecl) error {
	file := context.SourceFile(ctx)
	decl.SourceFileName = file.FullName
	decl.PackageName = file.PackageName
	return nil
}

func (p *Namer) ParseTypeAlias(ctx context.Context, decl *lang.TypeAliasDecl) error {
	thisCtx := context.WithScopeType(ctx, decl)
	thisScope := context.Scope(thisCtx)

	file := context.SourceFile(ctx)
	decl.SourceFileName = file.FullName
	decl.PackageName = file.PackageName

	for _, parameter := range decl.GenericParameters {
		identifier := thisScope.Declare(lang.NewGenericParameterDeclaration(parameter))
		identifier.SourceFileName = file.FullName
	}

	return nil
}

func (p *Namer) ParseInterface(ctx context.Context, decl *lang.InterfaceDecl) error {
	thisCtx := context.WithScopeType(ctx, decl)
	thisScope := context.Scope(thisCtx)

	file := context.SourceFile(ctx)
	decl.SourceFileName = file.FullName
	decl.PackageName = file.PackageName

	for _, method := range decl.GetType().GetMethods() {
		method.PackageName = decl.PackageName
		method.SourceFileName = decl.SourceFileName
		method.Enclosing = &lang.NominalType{
			Name: decl.Name,
		}
	}

	// for _, declaration := range decl.TypeAliasDecls {
	//	identifier := ctx.Declare(lang.NewTypeAliasDecl(declaration))
	//	ctx.Current().Identifiers[identifier.Name] = identifier
	//
	//	p.ParseTypeAlias(ctx, declaration)
	// }

	for key, value := range thisScope.Identifiers {
		context.Scope(ctx).Identifiers[decl.Name+"."+key] = value
	}

	return nil
}

func (p *Namer) ParseImport(ctx context.Context, decl *lang.ImportDecl) error {
	return nil
}
