package compiler

import (
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/mojo-lang/mojo/go/pkg/context"
)

type PackageCompiler interface {
	CompilePackage(ctx context.Context, pkg *lang.Package) error
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
