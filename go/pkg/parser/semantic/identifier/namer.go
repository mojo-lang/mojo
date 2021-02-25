package identifier

import (
	"errors"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/mojo-lang/mojo/go/pkg/parser/semantic/plugin"
)

func init() {
	parser := &Namer{}
	plugin.AddPlugin("identifier-namer", 1, parser)
}

type Namer struct {
}

func (p *Namer) Parse(ctx *plugin.Context, pkg *lang.Package, options map[string]interface{}) error {
	ctx.Open(pkg)
	if len(options) > 0 {
		ctx.SetOptions(options)
	}

	defer func() {
		scope := ctx.CurrentScope()
		ctx.Close()

		if ctx.CurrentScope() != nil { // not global
			for key, value := range scope.Identifiers {
				ctx.CurrentScope().Identifiers[pkg.Name+"."+key] = value

				if value.Package == "mojo.core" {
					ctx.CurrentScope().Identifiers[key] = value
				}
			}
		}
	}()

	if disabled, ok := ctx.GetOption(pkg.FullName + ".disable").(bool); !ok || !disabled {
		for _, file := range pkg.SourceFiles {
			if err := p.ParserSourceFile(ctx, file); err != nil {
				return err
			}
		}
	}

	for _, child := range pkg.Children {
		if err := p.Parse(ctx, child, nil); err != nil {
			return err
		}
	}

	return nil
}

func (p *Namer) ParserSourceFile(ctx *plugin.Context, file *lang.SourceFile) error {
	ctx.Open(file)
	defer func() {
		scope := ctx.CurrentScope()
		ctx.Close()

		for key, value := range scope.Identifiers {
			ctx.CurrentScope().Identifiers[key] = value
		}
	}()

	pkg := ctx.GetPackage()

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

			identifier := ctx.Declare(decl)
			if identifier != nil {
				identifier.SourceFile = file.FullName
			}

			switch decl.Declaration.(type) {
			case *lang.Declaration_TypeAliasDecl:
				err := p.ParseTypeAlias(ctx, decl.GetTypeAliasDecl())
				if err != nil {
					return err
				}
			case *lang.Declaration_StructDecl:
				err := p.ParseStruct(ctx, decl.GetStructDecl())
				if err != nil {
					return err
				}
			case *lang.Declaration_EnumDecl:
				err := p.ParseEnum(ctx, decl.GetEnumDecl())
				if err != nil {
					return err
				}
			case *lang.Declaration_InterfaceDecl:
				err := p.ParseInterface(ctx, decl.GetInterfaceDecl())
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

func (p *Namer) ParseStruct(ctx *plugin.Context, decl *lang.StructDecl) error {
	ctx.Open(decl)
	file := ctx.GetSourceFile()
	decl.SourceFileName = file.FullName
	decl.PackageName = file.PackageName

	for _, declaration := range decl.StructDecls {
		declaration.PackageName = file.PackageName
		declaration.SourceFileName = file.FullName

		identifier := ctx.Declare(lang.NewStructDecl(declaration))
		identifier.SourceFile = file.FullName
		p.ParseStruct(ctx, declaration)
	}

	for _, declaration := range decl.EnumDecls {
		declaration.PackageName = file.PackageName
		declaration.SourceFileName = file.FullName

		identifier := ctx.Declare(lang.NewEnumDecl(declaration))
		identifier.SourceFile = file.FullName
		p.ParseEnum(ctx, declaration)
	}

	for _, declaration := range decl.TypeAliasDecls {
		declaration.PackageName = file.PackageName
		declaration.SourceFileName = file.FullName

		identifier := ctx.Declare(lang.NewTypeAliasDecl(declaration))
		identifier.SourceFile = file.FullName
		p.ParseTypeAlias(ctx, declaration)
	}

	for _, parameter := range decl.GenericParameters {
		identifier := ctx.Declare(lang.NewGenericParameterDeclaration(parameter))
		identifier.SourceFile = file.FullName
	}

	scope := ctx.CurrentScope()
	ctx.Close()

	for key, value := range scope.Identifiers {
		if value.Kind != lang.Identifier_KIND_GENERIC_PARAMETER {
			ctx.CurrentScope().Identifiers[decl.Name+"."+key] = value
		}
	}

	return nil
}

func (p *Namer) ParseEnum(ctx *plugin.Context, decl *lang.EnumDecl) error {
	file := ctx.GetSourceFile()
	decl.SourceFileName = file.FullName
	decl.PackageName = file.PackageName
	return nil
}

func (p *Namer) ParseTypeAlias(ctx *plugin.Context, decl *lang.TypeAliasDecl) error {
	ctx.Open(decl)

	file := ctx.GetSourceFile()
	decl.SourceFileName = file.FullName
	decl.PackageName = file.PackageName

	for _, parameter := range decl.GenericParameters {
		identifier := ctx.Declare(lang.NewGenericParameterDeclaration(parameter))
		identifier.SourceFile = file.FullName
	}

	ctx.Close()
	return nil
}

func (p *Namer) ParseInterface(ctx *plugin.Context, decl *lang.InterfaceDecl) error {
	ctx.Open(decl)
	file := ctx.GetSourceFile()
	decl.SourceFileName = file.FullName
	decl.PackageName = file.PackageName

	//for _, declaration := range decl.TypeAliasDecls {
	//	identifier := ctx.Declare(lang.NewTypeAliasDecl(declaration))
	//	ctx.Current().Identifiers[identifier.Name] = identifier
	//
	//	p.ParseTypeAlias(ctx, declaration)
	//}

	innerScope := ctx.CurrentScope()
	ctx.Close()

	for key, value := range innerScope.Identifiers {
		ctx.CurrentScope().Identifiers[decl.Name+"."+key] = value
	}

	return nil
}

func (p *Namer) ParseImport(ctx *plugin.Context, decl *lang.ImportDecl) error {
	return nil
}
