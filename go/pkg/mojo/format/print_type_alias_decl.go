package format

import (
    "github.com/mojo-lang/lang/go/pkg/mojo/lang"
    "github.com/mojo-lang/mojo/go/pkg/mojo/context"
)

func (p *Printer) PrintTypeAliasDecl(ctx context.Context, decl *lang.TypeAliasDecl) {
    if decl == nil || p == nil || p.Error != nil {
        return
    }

    breaker := &OnceLineBreaker{}
    p.printPreDecl(ctx, decl, breaker).
        Break(p).
        PrintLine("type", " ", decl.Name).
        printDeclGenericParameters(ctx, decl.GenericParameters).
        PrintRaw(" = ").
        PrintNominalType(ctx, decl.Type)
}
