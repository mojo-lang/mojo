package format

import (
    "github.com/mojo-lang/mojo/go/pkg/mojo/context"
    "github.com/mojo-lang/lang/go/pkg/mojo/lang"
)

func (p *Printer) PrintTypeDeclaration(ctx context.Context, decl *lang.TypeDeclaration) {
    if decl == nil || p == nil || p.Error != nil {
        return
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
}
