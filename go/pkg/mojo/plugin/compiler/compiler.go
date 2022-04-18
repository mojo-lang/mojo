package compiler

import (
    "github.com/mojo-lang/lang/go/pkg/mojo/lang"
    "github.com/mojo-lang/mojo/go/pkg/mojo/context"
)

type PackageCompiler interface {
    CompilePackage(ctx context.Context, pkg *lang.Package) error
}

type SourceFileCompiler interface {
    CompileSourceFile(ctx context.Context, file *lang.SourceFile) error
}

type StructCompiler interface {
    CompileStruct(ctx context.Context, decl *lang.StructDecl) error
}

type TypeAliasCompiler interface {
    CompileTypeAlias(ctx context.Context, decl *lang.TypeAliasDecl) error
}

type EnumCompiler interface {
    CompileEnum(ctx context.Context, decl *lang.EnumDecl) error
}

type InterfaceCompiler interface {
    CompileInterface(ctx context.Context, decl *lang.InterfaceDecl) error
}

func IsPackageCompiler(c interface{}) bool {
    if _, ok := c.(PackageCompiler); ok {
        return true
    }
    return false
}

func IsSourceFileCompiler(c interface{}) bool {
    if _, ok := c.(SourceFileCompiler); ok {
        return true
    }
    return false
}

func IsStructCompiler(c interface{}) bool {
    if _, ok := c.(StructCompiler); ok {
        return true
    }
    return false
}

func IsTypeAliasCompiler(c interface{}) bool {
    if _, ok := c.(TypeAliasCompiler); ok {
        return true
    }
    return false
}

func IsEnumCompiler(c interface{}) bool {
    if _, ok := c.(EnumCompiler); ok {
        return true
    }
    return false
}

func IsInterfaceCompiler(c interface{}) bool {
    if _, ok := c.(InterfaceCompiler); ok {
        return true
    }
    return false
}

func IsCompiler(c interface{}) bool {
    return IsPackageCompiler(c) ||
        IsSourceFileCompiler(c) ||
        IsStructCompiler(c) ||
        IsTypeAliasCompiler(c) ||
        IsEnumCompiler(c) ||
        IsInterfaceCompiler(c)
}
