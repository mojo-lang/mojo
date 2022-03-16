package parser

import (
    "github.com/mojo-lang/lang/go/pkg/mojo/lang"
    "github.com/mojo-lang/mojo/go/pkg/mojo/context"
    "io/fs"
)

type PackageParser interface {
    ParsePackage(ctx context.Context, pkg *lang.Package) error
}

type SourceFileParser interface {
    ParseSourceFile(ctx context.Context, file *lang.SourceFile) error
}

type StructParser interface {
    ParseStruct(ctx context.Context, decl *lang.StructDecl) error
}

type InterfaceParser interface {
    ParseInterface(ctx context.Context, decl *lang.InterfaceDecl) error
}

type EnumParser interface {
    ParseEnum(ctx context.Context, decl *lang.EnumDecl) error
}

type TypeAliasParser interface {
    ParseTypeAlias(ctx context.Context, decl *lang.TypeAliasDecl) error
}

type ExpressionParser interface {
    ParseExpression(ctx context.Context, expr *lang.Expression) error
}

type FileParser interface {
    ParseFile(ctx context.Context, fileName string, fileSys fs.FS) (*lang.SourceFile, error)
}

type PackagePathParser interface {
    ParsePackagePath(ctx context.Context, pkgPath string, fileSys fs.FS) (*lang.Package, error)
}
