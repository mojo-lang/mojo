package printer

import (
	"github.com/mojo-lang/mojo/packages/lang/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/go/pkg/context"
)

func (p *Printer) PrintDeclaration(ctx context.Context, decl *lang.Declaration) *Printer {
	if decl == nil || p.GetError() != nil {
		return p
	}

	switch d := decl.Declaration.(type) {
	case *lang.Declaration_TypeAliasDecl:
		p.PrintTypeAliasDecl(ctx, d.TypeAliasDecl)
	case *lang.Declaration_EnumDecl:
		p.PrintEnumDecl(ctx, d.EnumDecl)
	case *lang.Declaration_StructDecl:
		p.PrintStructDecl(ctx, d.StructDecl)
	case *lang.Declaration_InterfaceDecl:
		p.PrintInterfaceDecl(ctx, d.InterfaceDecl)
	case *lang.Declaration_GenericParameter:
		p.PrintGenericParameter(ctx, d.GenericParameter)
	case *lang.Declaration_FunctionDecl:
		p.PrintFunctionDecl(ctx, d.FunctionDecl)
	case *lang.Declaration_AttributeDecl:
		p.PrintAttributeDecl(ctx, d.AttributeDecl)
	case *lang.Declaration_PackageDecl:
		p.PrintPackageDecl(ctx, d.PackageDecl)
	}

	return p
}
