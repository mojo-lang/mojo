package printer

import (
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/go/pkg/context"
)

func (p *Printer) PrintTypeDeclaration(ctx context.Context, decl *lang.TypeDeclaration) *Printer {
	if decl == nil || p.GetError() != nil {
		return p
	}

	switch d := decl.TypeDeclaration.(type) {
	case *lang.TypeDeclaration_TypeAliasDecl:
		p.PrintTypeAliasDecl(ctx, d.TypeAliasDecl)
	case *lang.TypeDeclaration_EnumDecl:
		p.PrintEnumDecl(ctx, d.EnumDecl)
	case *lang.TypeDeclaration_StructDecl:
		p.PrintStructDecl(ctx, d.StructDecl)
	case *lang.TypeDeclaration_InterfaceDecl:
		p.PrintInterfaceDecl(ctx, d.InterfaceDecl)
	case *lang.TypeDeclaration_GenericParameter:
		p.PrintGenericParameter(ctx, d.GenericParameter)
	}

	return p
}
