package printer

import (
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/go/pkg/context"
)

func (p *Printer) PrintTypeAliasDecl(ctx context.Context, decl *lang.TypeAliasDecl) *Printer {
	if decl == nil || p.GetError() != nil {
		return p
	}

	breaker := &OnceLineBreaker{}
	p.printPreDecl(ctx, decl, breaker).
		Break(p).
		PrintLine("type", " ", decl.Name).
		printDeclGenericParameters(ctx, decl.GenericParameters).
		PrintRaw(" = ").
		PrintNominalType(ctx, decl.Type)

	return p
}
