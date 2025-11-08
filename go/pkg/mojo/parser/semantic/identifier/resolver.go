package identifier

import (
	"errors"
	"fmt"

	"github.com/mojo-lang/mojo/packages/core/go/pkg/logs"
	"github.com/mojo-lang/mojo/packages/core/go/pkg/mojo/core"
	"github.com/mojo-lang/mojo/packages/lang/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/go/pkg/context"
	"github.com/mojo-lang/mojo/go/pkg/plugin"
	"github.com/mojo-lang/mojo/go/pkg/util"
)

const resolverName = "semantic.identifier-resolver"

type Resolver struct {
	plugin.BasicPlugin
}

func init() {
	plugin.RegisterPlugin(NewResolver(nil))
}

func NewResolver(options core.Options) *Resolver {
	_ = options
	return &Resolver{
		BasicPlugin: plugin.BasicPlugin{
			Name:          resolverName,
			Group:         "semantic",
			GroupPriority: 3,
			Priority:      3,
			Creator: func(options core.Options) plugin.Plugin {
				return NewResolver(options)
			},
		},
	}
}

func (p *Resolver) ParsePackage(ctx context.Context, pkg *lang.Package) error {
	if util.IsPackageProcessed(pkg, resolverName) {
		logs.Infow("already processed, skip the plugin", "plugin", p.Name, "method", "ParsePackage", "pkg", pkg.FullName)
		return nil
	} else {
		logs.Infow("enter the plugin", "plugin", p.Name, "method", "ParsePackage", "pkg", pkg.FullName)
	}

	thisCtx := context.WithScopeType(ctx, pkg)

	for _, file := range pkg.SourceFiles {
		if err := p.ParserSourceFile(thisCtx, file); err != nil {
			return err
		}

		if len(file.UnresolvedIdentifiers) > 0 {
			for _, identifier := range file.UnresolvedIdentifiers {
				logs.Errorw("unresolved identifier", "name", identifier.Name, "package", pkg.FullName, "file", file.Name)
			}

			return fmt.Errorf("there are unresolved identifiers in the file (%s)", file.Name)
		}
	}

	for _, child := range pkg.Children {
		if err := p.ParsePackage(thisCtx, child); err != nil {
			return err
		}
	}

	if !pkg.IsPadding() {
		util.SetPackageProcessed(pkg, resolverName)
	}
	return nil
}

func (p *Resolver) ParserSourceFile(ctx context.Context, file *lang.SourceFile) error {
	thisCtx := context.WithScopeType(ctx, file)

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
		// switch value := ctx.CurrentValue().(type) {
		// case *lang.Package:
		//	value.ResolvedIdentifiers = append(value.ResolvedIdentifiers, resolveds...)
		//	value.UnresolvedIdentifiers = append(value.UnresolvedIdentifiers, unresolveds...)
		// }
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
			case *lang.Declaration_TypeAliasDecl:
				err := p.ParseTypeAlias(thisCtx, decl.GetTypeAliasDecl())
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

func (p *Resolver) ParseStruct(ctx context.Context, decl *lang.StructDecl) error {
	thisCtx := context.WithScopeType(ctx, decl)
	var resolveds []*lang.Identifier
	var unresolveds []*lang.Identifier

	defer func() {
		resolveds = append(resolveds, decl.ResolvedIdentifiers...)
		resolveds = lang.MergeDependencies(resolveds)
		decl.ResolvedIdentifiers = resolveds

		unresolveds = append(unresolveds, decl.UnresolvedIdentifiers...)
		unresolveds = lang.MergeUnresolvedIdentifiers(unresolveds)
		decl.UnresolvedIdentifiers = unresolveds

		// push to parent
		switch value := context.TypeValue(ctx).(type) {
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
			rs, urs := resolveNominalType(thisCtx, field.Type)
			resolveds = append(resolveds, rs...)
			unresolveds = append(unresolveds, urs...)
		}

		for _, inherited := range t.Inherits {
			rs, urs := resolveNominalType(thisCtx, inherited)
			resolveds = append(resolveds, rs...)
			unresolveds = append(unresolveds, urs...)
		}
	}

	for _, structDecl := range decl.StructDecls {
		if structDecl.Enclosing != nil {
			structDecl.Enclosing.TypeDeclaration = lang.NewStructTypeDeclaration(decl)
		}
		if err := p.ParseStruct(thisCtx, structDecl); err != nil {
			return err
		}
	}

	for _, enumDecl := range decl.EnumDecls {
		if enumDecl.Enclosing != nil {
			enumDecl.Enclosing.TypeDeclaration = lang.NewStructTypeDeclaration(decl)
		}
		if err := p.ParseEnum(thisCtx, enumDecl); err != nil {
			return err
		}
	}

	for _, aliasDecl := range decl.TypeAliasDecls {
		if aliasDecl.Enclosing != nil {
			aliasDecl.Enclosing.TypeDeclaration = lang.NewStructTypeDeclaration(decl)
		}
		if err := p.ParseTypeAlias(thisCtx, aliasDecl); err != nil {
			return err
		}
	}

	return nil
}

func (p *Resolver) ParseEnum(ctx context.Context, decl *lang.EnumDecl) error {
	var resolveds []*lang.Identifier
	var unresolveds []*lang.Identifier
	thisCtx := context.WithScopeType(ctx, decl)

	defer func() {
		resolveds = append(resolveds, decl.ResolvedIdentifiers...)
		resolveds = lang.MergeDependencies(resolveds)
		decl.ResolvedIdentifiers = resolveds

		unresolveds = append(unresolveds, decl.UnresolvedIdentifiers...)
		unresolveds = lang.MergeUnresolvedIdentifiers(unresolveds)
		decl.UnresolvedIdentifiers = unresolveds

		// push to parent
		switch value := context.TypeValue(ctx).(type) {
		case *lang.StructDecl:
			value.ResolvedIdentifiers = append(value.ResolvedIdentifiers, resolveds...)
			value.UnresolvedIdentifiers = append(value.UnresolvedIdentifiers, unresolveds...)
		case *lang.SourceFile:
			value.ResolvedIdentifiers = append(value.ResolvedIdentifiers, resolveds...)
			value.UnresolvedIdentifiers = append(value.UnresolvedIdentifiers, unresolveds...)
		}
	}()

	if decl.Type.UnderlyingType != nil {
		resolveds, unresolveds = resolveNominalType(thisCtx, decl.Type.UnderlyingType)
	}

	return nil
}

func (p *Resolver) ParseInterface(ctx context.Context, decl *lang.InterfaceDecl) error {
	var resolveds []*lang.Identifier
	var unresolveds []*lang.Identifier
	thisCtx := context.WithScopeType(ctx, decl)

	defer func() {
		resolveds = append(resolveds, decl.ResolvedIdentifiers...)
		resolveds = lang.MergeDependencies(resolveds)
		decl.ResolvedIdentifiers = resolveds

		unresolveds = append(unresolveds, decl.UnresolvedIdentifiers...)
		unresolveds = lang.MergeUnresolvedIdentifiers(unresolveds)
		decl.UnresolvedIdentifiers = unresolveds

		// push to parent
		switch value := context.TypeValue(ctx).(type) {
		case *lang.SourceFile:
			value.ResolvedIdentifiers = append(value.ResolvedIdentifiers, resolveds...)
			value.UnresolvedIdentifiers = append(value.UnresolvedIdentifiers, unresolveds...)
		}
	}()

	for _, method := range decl.Type.Methods {
		signature := method.Signature
		for _, parameter := range signature.GetParameters() {
			rs, urs := resolveNominalType(thisCtx, parameter.Type)
			resolveds = append(resolveds, rs...)
			unresolveds = append(unresolveds, urs...)
		}

		if signature.GetResultType() != nil {
			rs, urs := resolveNominalType(thisCtx, signature.GetResultType())
			resolveds = append(resolveds, rs...)
			unresolveds = append(unresolveds, urs...)
		}
	}

	for _, inherited := range decl.Type.Inherits {
		rs, urs := resolveNominalType(thisCtx, inherited)
		resolveds = append(resolveds, rs...)
		unresolveds = append(unresolveds, urs...)
	}

	// for _, aliasDecl := range decl.TypeAliasDecls {
	//	if err := resolveNominalType(thisCtx, aliasDecl.Type); err != nil {
	//		return err
	//	}
	//
	//	for _, argument := range aliasDecl.Type.GenericArguments {
	//		if err := resolveNominalType(thisCtx, argument); err != nil {
	//			return err
	//		}
	//	}
	// }

	return nil
}

func (p *Resolver) ParseTypeAlias(ctx context.Context, decl *lang.TypeAliasDecl) error {
	var resolveds []*lang.Identifier
	var unresolveds []*lang.Identifier
	thisCtx := context.WithScopeType(ctx, decl)
	defer func() {
		resolveds = lang.MergeDependencies(resolveds)
		decl.ResolvedIdentifiers = resolveds

		unresolveds = lang.MergeUnresolvedIdentifiers(unresolveds)
		decl.UnresolvedIdentifiers = unresolveds

		// push to parent
		switch value := context.TypeValue(ctx).(type) {
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

	resolveds, unresolveds = resolveNominalType(thisCtx, decl.Type)
	return nil
}

func (p *Resolver) ParseImport(ctx context.Context, decl *lang.ImportDecl) error {
	return nil
}

func resolveNominalType(ctx context.Context, t *lang.NominalType) ([]*lang.Identifier, []*lang.Identifier) {
	fullName := t.Name
	enclosingNames := lang.GetEnclosingNames(t.Enclosing)
	var identifier *lang.Identifier
	if len(t.PackageName) > 0 {
		fullName = lang.GetFullName(t.PackageName, enclosingNames, t.Name)
		identifier = context.LookupIdentifier(ctx, fullName)
	} else if len(enclosingNames) > 0 {
		identifier = context.LookupIdentifier(ctx, lang.GetFullName("", enclosingNames, t.Name))
	} else {
		identifier = context.LookupIdentifier(ctx, t.Name)
	}

	var resolveds []*lang.Identifier
	var unresolveds []*lang.Identifier

	if identifier != nil {
		t.PackageName = identifier.PackageName
		t.TypeDeclaration = lang.NewTypeDeclarationFromDeclaration(identifier.Declaration)
		t.Enclosing = identifier.Declaration.GetEnclosingType()
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
