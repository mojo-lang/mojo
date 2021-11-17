package compiler

import (
	"github.com/mojo-lang/core/go/pkg/mojo/core"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/mojo-lang/mojo/go/pkg/context"
)

// used for union, tuple
type LabelFormatCompiler struct {
	DummyCompiler
}

func (l *LabelFormatCompiler) Compile(ctx *context.Context, pkg *lang.Package) error {
	for _, sourceFile := range pkg.SourceFiles {
		ctx.Open(sourceFile)
		for _, statement := range sourceFile.Statements {
			if decl := statement.GetDeclaration(); decl != nil {
				switch decl.Declaration.(type) {
				case *lang.Declaration_StructDecl:
					//structDecl := decl.GetStructDecl()
				case *lang.Declaration_TypeAliasDecl:
				case *lang.Declaration_EnumDecl:
				case *lang.Declaration_InterfaceDecl:
				}
			}
		}
		ctx.Close()
	}
	return nil
}

func (l *LabelFormatCompiler) CompileStruct(ctx *context.Context, decl *lang.StructDecl) error {
	//for _, s := range decl.StructDecls {
	//
	//}
	//
	//for _, field := range decl.Type.
	return nil
}

func (l *LabelFormatCompiler) CompileTypeAlias(ctx *context.Context, decl *lang.TypeAliasDecl) error {
	if decl.Type.GetFullName() == core.UnionTypeName {
		if lang.HasAttribute(decl.Attributes, core.LabelFormatAttributeName) {
		}
	}
	return nil
}
