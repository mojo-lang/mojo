package compiler

import (
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/mojo-lang/mojo/go/pkg/context"
)

type TypeDeclarationCompiler interface {
	CompileStruct(ctx *context.Context, decl *lang.StructDecl) error
	CompileTypeAlias(ctx *context.Context, decl *lang.TypeAliasDecl) error
	CompileEnum(ctx *context.Context, decl *lang.EnumDecl) error
	CompileInterface(ctx *context.Context, decl *lang.InterfaceDecl) error
}

type DummyCompiler struct{}

func (d *DummyCompiler) CompileStruct(ctx *context.Context, decl *lang.StructDecl) error {
	return nil
}
func (d *DummyCompiler) CompileTypeAlias(ctx *context.Context, decl *lang.TypeAliasDecl) error {
	return nil
}
func (d *DummyCompiler) CompileEnum(ctx *context.Context, decl *lang.EnumDecl) error {
	return nil
}
func (d *DummyCompiler) CompileInterface(ctx *context.Context, decl *lang.InterfaceDecl) error {
	return nil
}
