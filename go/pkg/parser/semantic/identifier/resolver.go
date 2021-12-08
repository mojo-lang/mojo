package identifier

import (
	"errors"
	"fmt"
	"github.com/mojo-lang/core/go/pkg/logs"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/mojo-lang/mojo/go/pkg/parser/semantic/plugin"
)

func init() {
	resolver := &Resolver{}
	plugin.AddPlugin("identifier-resolver", 3, resolver)
}

type Resolver struct {
}

func (p *Resolver) Parse(ctx *plugin.Context, pkg *lang.Package, options map[string]interface{}) error {
	ctx.Open(pkg)
	if len(options) > 0 {
		ctx.SetOptions(options)
	}

	defer func() {
		ctx.Close()
	}()

	if disabled, ok := ctx.GetOption(pkg.FullName + ".disable").(bool); !ok || !disabled {
		for _, file := range pkg.SourceFiles {
			if err := p.ParserSourceFile(ctx, file); err != nil {
				return err
			}

			if len(file.UnresolvedIdentifiers) > 0 {
				for _, identifier := range file.UnresolvedIdentifiers {
					logs.Errorw("unresolved identifier", "name", identifier.Name, "package", pkg.FullName, "file", file.Name)
				}

				return fmt.Errorf("there are unresolved identifiers in the file (%s)", file.Name)
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

func (p *Resolver) ParserSourceFile(ctx *plugin.Context, file *lang.SourceFile) error {
	ctx.Open(file)

	defer func() {
		resolveds := lang.MergeDependencies(file.ResolvedIdentifiers)
		file.ResolvedIdentifiers = []*lang.Identifier{}
		for _, resolved := range resolveds {
			if resolved.SourceFileName != file.FullName {
				file.ResolvedIdentifiers = append(file.ResolvedIdentifiers, resolved)
			}
		}

		unresolveds := lang.MergeUnresolvedIdentifiers(file.UnresolvedIdentifiers)
		file.UnresolvedIdentifiers = unresolveds

		// push to parent
		//switch value := ctx.CurrentValue().(type) {
		//case *lang.Package:
		//	value.ResolvedIdentifiers = append(value.ResolvedIdentifiers, resolveds...)
		//	value.UnresolvedIdentifiers = append(value.UnresolvedIdentifiers, unresolveds...)
		//}
	}()

	for _, statement := range file.Statements {
		switch statement.Statement.(type) {
		case *lang.Statement_Declaration:
			decl := statement.GetDeclaration()
			if decl == nil {
				return errors.New("declaration statement is nil")
			}

			switch decl.Declaration.(type) {
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
			case *lang.Declaration_TypeAliasDecl:
				err := p.ParseTypeAlias(ctx, decl.GetTypeAliasDecl())
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

func (p *Resolver) ParseStruct(ctx *plugin.Context, decl *lang.StructDecl) error {
	var resolveds []*lang.Identifier
	var unresolveds []*lang.Identifier

	ctx.Open(decl)
	defer func() {
		resolveds = append(resolveds, decl.ResolvedIdentifiers...)
		resolveds = lang.MergeDependencies(resolveds)
		decl.ResolvedIdentifiers = resolveds

		unresolveds = append(unresolveds, decl.UnresolvedIdentifiers...)
		unresolveds = lang.MergeUnresolvedIdentifiers(unresolveds)
		decl.UnresolvedIdentifiers = unresolveds

		ctx.Close()

		// push to parent
		switch value := ctx.CurrentValue().(type) {
		case *lang.StructDecl:
			value.ResolvedIdentifiers = append(value.ResolvedIdentifiers, resolveds...)
			value.UnresolvedIdentifiers = append(value.UnresolvedIdentifiers, unresolveds...)
		case *lang.SourceFile:
			value.ResolvedIdentifiers = append(value.ResolvedIdentifiers, resolveds...)
			value.UnresolvedIdentifiers = append(value.UnresolvedIdentifiers, unresolveds...)
		}
	}()

	// 分析依赖的symbol，并补全decl指针
	t := decl.Type
	if t != nil {
		for _, field := range t.Fields {
			rs, urs := resolveNominalType(ctx, field.Type)
			resolveds = append(resolveds, rs...)
			unresolveds = append(unresolveds, urs...)
		}

		for _, inherited := range t.Inherits {
			rs, urs := resolveNominalType(ctx, inherited)
			resolveds = append(resolveds, rs...)
			unresolveds = append(unresolveds, urs...)
		}
	}

	for _, structDecl := range decl.StructDecls {
		if err := p.ParseStruct(ctx, structDecl); err != nil {
			return err
		}
	}

	for _, enumDecl := range decl.EnumDecls {
		if err := p.ParseEnum(ctx, enumDecl); err != nil {
			return err
		}
	}

	for _, aliasDecl := range decl.TypeAliasDecls {
		if err := p.ParseTypeAlias(ctx, aliasDecl); err != nil {
			return err
		}
	}

	return nil
}

func (p *Resolver) ParseEnum(ctx *plugin.Context, decl *lang.EnumDecl) error {
	var resolveds []*lang.Identifier
	var unresolveds []*lang.Identifier

	ctx.Open(decl)
	defer func() {
		resolveds = append(resolveds, decl.ResolvedIdentifiers...)
		resolveds = lang.MergeDependencies(resolveds)
		decl.ResolvedIdentifiers = resolveds

		unresolveds = append(unresolveds, decl.UnresolvedIdentifiers...)
		unresolveds = lang.MergeUnresolvedIdentifiers(unresolveds)
		decl.UnresolvedIdentifiers = unresolveds

		ctx.Close()

		// push to parent
		switch value := ctx.CurrentValue().(type) {
		case *lang.StructDecl:
			value.ResolvedIdentifiers = append(value.ResolvedIdentifiers, resolveds...)
			value.UnresolvedIdentifiers = append(value.UnresolvedIdentifiers, unresolveds...)
		case *lang.SourceFile:
			value.ResolvedIdentifiers = append(value.ResolvedIdentifiers, resolveds...)
			value.UnresolvedIdentifiers = append(value.UnresolvedIdentifiers, unresolveds...)
		}
	}()

	if decl.Type.UnderlyingType != nil {
		resolveds, unresolveds = resolveNominalType(ctx, decl.Type.UnderlyingType)
	}

	return nil
}

func (p *Resolver) ParseInterface(ctx *plugin.Context, decl *lang.InterfaceDecl) error {
	var resolveds []*lang.Identifier
	var unresolveds []*lang.Identifier

	ctx.Open(decl)
	defer func() {
		resolveds = append(resolveds, decl.ResolvedIdentifiers...)
		resolveds = lang.MergeDependencies(resolveds)
		decl.ResolvedIdentifiers = resolveds

		unresolveds = append(unresolveds, decl.UnresolvedIdentifiers...)
		unresolveds = lang.MergeUnresolvedIdentifiers(unresolveds)
		decl.UnresolvedIdentifiers = unresolveds

		ctx.Close()

		// push to parent
		switch value := ctx.CurrentValue().(type) {
		case *lang.SourceFile:
			value.ResolvedIdentifiers = append(value.ResolvedIdentifiers, resolveds...)
			value.UnresolvedIdentifiers = append(value.UnresolvedIdentifiers, unresolveds...)
		}
	}()

	for _, method := range decl.Type.Methods {
		signature := method.Signature
		for _, parameter := range signature.Parameters {
			rs, urs := resolveNominalType(ctx, parameter.Type)
			resolveds = append(resolveds, rs...)
			unresolveds = append(unresolveds, urs...)
		}

		if signature.Result != nil {
			rs, urs := resolveNominalType(ctx, signature.Result)
			resolveds = append(resolveds, rs...)
			unresolveds = append(unresolveds, urs...)
		}
	}

	for _, inherited := range decl.Type.Inherits {
		rs, urs := resolveNominalType(ctx, inherited)
		resolveds = append(resolveds, rs...)
		unresolveds = append(unresolveds, urs...)
	}

	//for _, aliasDecl := range decl.TypeAliasDecls {
	//	if err := resolveNominalType(ctx, aliasDecl.Type); err != nil {
	//		return err
	//	}
	//
	//	for _, argument := range aliasDecl.Type.GenericArguments {
	//		if err := resolveNominalType(ctx, argument); err != nil {
	//			return err
	//		}
	//	}
	//}

	return nil
}

func (p *Resolver) ParseTypeAlias(ctx *plugin.Context, decl *lang.TypeAliasDecl) error {
	var resolveds []*lang.Identifier
	var unresolveds []*lang.Identifier

	ctx.Open(decl)
	defer func() {
		resolveds = lang.MergeDependencies(resolveds)
		decl.ResolvedIdentifiers = resolveds

		unresolveds = lang.MergeUnresolvedIdentifiers(unresolveds)
		decl.UnresolvedIdentifiers = unresolveds

		ctx.Close()

		// push to parent
		switch value := ctx.CurrentValue().(type) {
		case *lang.StructDecl:
			value.ResolvedIdentifiers = append(value.ResolvedIdentifiers, resolveds...)
			value.UnresolvedIdentifiers = append(value.UnresolvedIdentifiers, unresolveds...)
		case *lang.InterfaceDecl:
			value.ResolvedIdentifiers = append(value.ResolvedIdentifiers, resolveds...)
			value.UnresolvedIdentifiers = append(value.UnresolvedIdentifiers, unresolveds...)
		case *lang.SourceFile:
			value.ResolvedIdentifiers = append(value.ResolvedIdentifiers, resolveds...)
			value.UnresolvedIdentifiers = append(value.UnresolvedIdentifiers, unresolveds...)
		}
	}()

	resolveds, unresolveds = resolveNominalType(ctx, decl.Type)
	return nil
}

func (p *Resolver) ParseImport(ctx *plugin.Context, decl *lang.ImportDecl) error {
	return nil
}

func resolveNominalType(ctx *plugin.Context, t *lang.NominalType) ([]*lang.Identifier, []*lang.Identifier) {
	fullName := t.Name
	enclosingNames := lang.GetEnclosingNames(t.EnclosingType)
	var identifier *lang.Identifier
	if len(t.PackageName) > 0 {
		fullName = lang.GetFullName(t.PackageName, enclosingNames, t.Name)
		identifier = ctx.Lookup(fullName)
	} else {
		identifier = ctx.Lookup(t.Name)
		if identifier == nil && len(enclosingNames) > 0 {
			identifier = ctx.Lookup(lang.GetFullName("", enclosingNames, t.Name))
		}
	}

	var resolveds []*lang.Identifier
	var unresolveds []*lang.Identifier

	if identifier != nil {
		t.PackageName = identifier.PackageName
		t.TypeDeclaration = lang.NewTypeDeclarationFromDeclaration(identifier.Declaration)
		t.EnclosingType = identifier.Declaration.GetEnclosingType()
		resolveds = append(resolveds, identifier)
	} else {
		ident := &lang.Identifier{
			PackageName:   t.PackageName,
			StartPosition: t.StartPosition,
			Name:          t.Name,
			FullName:      fullName,
		}
		unresolveds = append(unresolveds, ident)
	}

	for _, argument := range t.GenericArguments {
		rs, urs := resolveNominalType(ctx, argument)
		resolveds = append(resolveds, rs...)
		unresolveds = append(unresolveds, urs...)
	}

	for _, attribute := range t.Attributes {
		for _, argument := range attribute.GenericArguments {
			rs, urs := resolveNominalType(ctx, argument)
			resolveds = append(resolveds, rs...)
			unresolveds = append(unresolveds, urs...)
		}
	}

	return resolveds, unresolveds
}
