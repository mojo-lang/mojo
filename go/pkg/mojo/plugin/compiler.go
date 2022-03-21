package plugin

import (
    "errors"
    "github.com/mojo-lang/lang/go/pkg/mojo/lang"
    "github.com/mojo-lang/mojo/go/pkg/mojo/context"
    "github.com/mojo-lang/mojo/go/pkg/mojo/plugin/compiler"
    "github.com/mojo-lang/mojo/go/pkg/mojo/plugin/hook"
)

const (
    compilePackageMethod    = "CompilePackage"
    compileSourceFileMethod = "CompileSourceFile"
    compileStructMethod     = "CompileStruct"
    compileInterfaceMethod  = "CompileInterface"
    compileEnumMethod       = "CompileEnum"
    compileTypeAliasMethod  = "CompileTypeAlias"
)

func CompilePackage(p interface{}, ctx context.Context, pkg *lang.Package) error {
    if p == nil {
        return nil
    }

    if !ContainsAnyMethod(p, compilePackageMethod, compileSourceFileMethod, compileStructMethod,
        compileInterfaceMethod, compileEnumMethod, compileTypeAliasMethod) {
        return SkipError{}
    }

    if hk, ok := p.(hook.PackagePreHook); ok {
        err := hk.PrePackage(ctx, pkg)
        if errors.Is(err, SkipError{}) {
            return nil
        } else {
            return err
        }
    }

    defer func() {
        if hook, ok := p.(hook.PackagePostHook); ok {
            hook.PostPackage(ctx, pkg)
        }
    }()

    if pkgCompiler, ok := p.(compiler.PackageCompiler); ok {
        return pkgCompiler.CompilePackage(ctx, pkg)
    }

    thisCtx := context.WithType(ctx, pkg)
    for _, file := range pkg.SourceFiles {
        if err := CompileSourceFile(p, thisCtx, file); err != nil && !errors.Is(err, SkipError{}) {
            return err
        }
    }

    for _, child := range pkg.Children {
        if err := CompilePackage(p, thisCtx, child); err != nil && !errors.Is(err, SkipError{}) {
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
        return SkipError{}
    }

    if hk, ok := p.(hook.SourceFilePreHook); ok {
        err := hk.PreSourceFile(ctx, file)
        if errors.Is(err, SkipError{}) {
            return nil
        } else {
            return err
        }
    }
    defer func() {
        if hk, ok := p.(hook.SourceFilePostHook); ok {
            hk.PostSourceFile(ctx, file)
        }
    }()

    if fileCompiler, ok := p.(compiler.SourceFileCompiler); ok {
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
        return SkipError{}
    }

    if hk, ok := p.(hook.InterfacePreHook); ok {
        err := hk.PreInterface(ctx, decl)
        if errors.Is(err, SkipError{}) {
            return nil
        } else {
            return err
        }
    }
    defer func() {
        if hk, ok := p.(hook.InterfacePostHook); ok {
            hk.PostInterface(ctx, decl)
        }
    }()

    if compiler, ok := p.(compiler.InterfaceCompiler); ok {
        return compiler.CompileInterface(ctx, decl)
    }

    return nil
}

func CompileStruct(p interface{}, ctx context.Context, decl *lang.StructDecl) error {
    if p == nil {
        return nil
    }

    if !ContainsAnyMethod(p, compileStructMethod) {
        return SkipError{}
    }

    if hk, ok := p.(hook.StructPreHook); ok {
        err := hk.PreStruct(ctx, decl)
        if errors.Is(err, SkipError{}) {
            return nil
        } else {
            return err
        }
    }
    defer func() {
        if hk, ok := p.(hook.StructPostHook); ok {
            hk.PostStruct(ctx, decl)
        }
    }()

    if compiler, ok := p.(compiler.StructCompiler); ok {
        return compiler.CompileStruct(ctx, decl)
    }

    return nil
}

func CompileTypeAlias(p interface{}, ctx context.Context, decl *lang.TypeAliasDecl) error {
    if p == nil {
        return nil
    }

    if !ContainsAnyMethod(p, compileTypeAliasMethod) {
        return SkipError{}
    }

    if hk, ok := p.(hook.TypeAliasPreHook); ok {
        err := hk.PreTypeAlias(ctx, decl)
        if errors.Is(err, SkipError{}) {
            return nil
        } else {
            return err
        }
    }
    defer func() {
        if hk, ok := p.(hook.TypeAliasPostHook); ok {
            hk.PostTypeAlias(ctx, decl)
        }
    }()

    if compiler, ok := p.(compiler.TypeAliasCompiler); ok {
        return compiler.CompileTypeAlias(ctx, decl)
    }

    return nil
}

func CompileEnum(p interface{}, ctx context.Context, decl *lang.EnumDecl) error {
    if p == nil {
        return nil
    }

    if !ContainsAnyMethod(p, compileEnumMethod) {
        return SkipError{}
    }

    if hk, ok := p.(hook.EnumPreHook); ok {
        err := hk.PreEnum(ctx, decl)
        if errors.Is(err, SkipError{}) {
            return nil
        } else {
            return err
        }
    }
    defer func() {
        if hk, ok := p.(hook.EnumPostHook); ok {
            hk.PostEnum(ctx, decl)
        }
    }()

    if compiler, ok := p.(compiler.EnumCompiler); ok {
        return compiler.CompileEnum(ctx, decl)
    }

    return nil
}
