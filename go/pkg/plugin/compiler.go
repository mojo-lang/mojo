package plugin

import (
	"github.com/mojo-lang/core/go/pkg/mojo/core"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/go/pkg/context"
)

const (
	compilePackageMethod    = "CompilePackage"
	compileSourceFileMethod = "CompileSourceFile"
	compileStructMethod     = "CompileStruct"
	compileInterfaceMethod  = "CompileInterface"
	compileEnumMethod       = "CompileEnum"
	compileTypeAliasMethod  = "CompileTypeAlias"
)

func IsCompiler(c interface{}) bool {
	return IsPackageCompiler(c) ||
		IsSourceFileCompiler(c) ||
		IsStructCompiler(c) ||
		IsTypeAliasCompiler(c) ||
		IsEnumCompiler(c) ||
		IsInterfaceCompiler(c)
}

func CompilePackage(p interface{}, ctx context.Context, pkg *lang.Package) error {
	if p == nil {
		return nil
	}

	if !ContainsAnyMethod(p, compilePackageMethod, compileSourceFileMethod, compileStructMethod,
		compileInterfaceMethod, compileEnumMethod, compileTypeAliasMethod) {
		return core.NewSkipError()
	}

	if hk, ok := p.(PackagePreHook); ok {
		err := hk.PrePackage(ctx, pkg)
		if core.IsSkipError(err) {
			return nil
		} else {
			return err
		}
	}

	defer func() {
		if hook, ok := p.(PackagePostHook); ok {
			hook.PostPackage(ctx, pkg)
		}
	}()

	if pkgCompiler, ok := p.(PackageCompiler); ok {
		return pkgCompiler.CompilePackage(ctx, pkg)
	}

	thisCtx := context.WithType(ctx, pkg)
	for _, file := range pkg.SourceFiles {
		if err := CompileSourceFile(p, thisCtx, file); err != nil && !core.IsSkipError(err) {
			return err
		}
	}

	for _, child := range pkg.Children {
		if err := CompilePackage(p, thisCtx, child); err != nil && !core.IsSkipError(err) {
			return err
		}
	}
	return nil
}

func CompileSourceFile(p interface{}, ctx context.Context, file *lang.SourceFile) error {
	if p == nil {
		return nil
	}

	if !ContainsAnyMethod(p, compileSourceFileMethod, compileStructMethod,
		compileInterfaceMethod, compileEnumMethod, compileTypeAliasMethod) {
		return core.NewSkipError()
	}

	if hk, ok := p.(SourceFilePreHook); ok {
		err := hk.PreSourceFile(ctx, file)
		if core.IsSkipError(err) {
			return nil
		} else {
			return err
		}
	}
	defer func() {
		if hk, ok := p.(SourceFilePostHook); ok {
			hk.PostSourceFile(ctx, file)
		}
	}()

	if fileCompiler, ok := p.(SourceFileCompiler); ok {
		return fileCompiler.CompileSourceFile(ctx, file)
	}

	thisCtx := context.WithType(ctx, file)
	for _, statement := range file.Statements {
		switch statement.Statement.(type) {
		case *lang.Statement_Declaration:
			switch decl := statement.GetDeclaration().Declaration.(type) {
			case *lang.Declaration_StructDecl:
				if err := CompileStruct(p, thisCtx, decl.StructDecl); err != nil {
					return err
				}
			case *lang.Declaration_InterfaceDecl:
				if err := CompileInterface(p, thisCtx, decl.InterfaceDecl); err != nil {
					return err
				}
			case *lang.Declaration_TypeAliasDecl:
				if err := CompileTypeAlias(p, thisCtx, decl.TypeAliasDecl); err != nil {
					return err
				}
			case *lang.Declaration_EnumDecl:
				if err := CompileEnum(p, thisCtx, decl.EnumDecl); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func CompileInterface(p interface{}, ctx context.Context, decl *lang.InterfaceDecl) error {
	if p == nil {
		return nil
	}

	if !ContainsAnyMethod(p, compileInterfaceMethod) {
		return core.NewSkipError()
	}

	if hk, ok := p.(InterfacePreHook); ok {
		err := hk.PreInterface(ctx, decl)
		if core.IsSkipError(err) {
			return nil
		} else {
			return err
		}
	}
	defer func() {
		if hk, ok := p.(InterfacePostHook); ok {
			hk.PostInterface(ctx, decl)
		}
	}()

	if compiler, ok := p.(InterfaceCompiler); ok {
		return compiler.CompileInterface(ctx, decl)
	}

	return nil
}

func CompileStruct(p interface{}, ctx context.Context, decl *lang.StructDecl) error {
	if p == nil {
		return nil
	}

	if !ContainsAnyMethod(p, compileStructMethod) {
		return core.NewSkipError()
	}

	if hk, ok := p.(StructPreHook); ok {
		if err := hk.PreStruct(ctx, decl); err != nil {
			if core.IsSkipError(err) {
				return nil
			} else {
				return err
			}
		}
	}
	defer func() {
		if hk, ok := p.(StructPostHook); ok {
			hk.PostStruct(ctx, decl)
		}
	}()

	if compiler, ok := p.(StructCompiler); ok {
		return compiler.CompileStruct(ctx, decl)
	}

	return nil
}

func CompileTypeAlias(p interface{}, ctx context.Context, decl *lang.TypeAliasDecl) error {
	if p == nil {
		return nil
	}

	if !ContainsAnyMethod(p, compileTypeAliasMethod) {
		return core.NewSkipError()
	}

	if hk, ok := p.(TypeAliasPreHook); ok {
		err := hk.PreTypeAlias(ctx, decl)
		if core.IsSkipError(err) {
			return nil
		} else {
			return err
		}
	}
	defer func() {
		if hk, ok := p.(TypeAliasPostHook); ok {
			hk.PostTypeAlias(ctx, decl)
		}
	}()

	if compiler, ok := p.(TypeAliasCompiler); ok {
		return compiler.CompileTypeAlias(ctx, decl)
	}

	return nil
}

func CompileEnum(p interface{}, ctx context.Context, decl *lang.EnumDecl) error {
	if p == nil {
		return nil
	}

	if !ContainsAnyMethod(p, compileEnumMethod) {
		return core.NewSkipError()
	}

	if hk, ok := p.(EnumPreHook); ok {
		err := hk.PreEnum(ctx, decl)
		if core.IsSkipError(err) {
			return nil
		} else {
			return err
		}
	}
	defer func() {
		if hk, ok := p.(EnumPostHook); ok {
			hk.PostEnum(ctx, decl)
		}
	}()

	if compiler, ok := p.(EnumCompiler); ok {
		return compiler.CompileEnum(ctx, decl)
	}

	return nil
}
