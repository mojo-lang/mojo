package printer

import (
	"github.com/mojo-lang/mojo/packages/lang/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/go/pkg/context"
)

func (p *Printer) PrintAttributeDecl(ctx context.Context, decl *lang.AttributeDecl) *Printer {
	if decl == nil || p.GetError() != nil {
		return p
	}

	breaker := &OnceLineBreaker{}
	p.printPreDecl(ctx, decl, breaker).
		Break(p).
		PrintLine("attribute", " ", decl.Name).
		printDeclGenericParameters(ctx, decl.GenericParameters).
		PrintRaw(": ")

	if nominalType := decl.GetNominalType(); nominalType != nil {
		p.PrintNominalType(ctx, decl.GetNominalType())
	} else if structType := decl.GetStructType(); structType != nil {
		p.PrintTerm(ctx, lang.NewSymbolTerm(structType.StartPosition, lang.TermTypeStart, " {"))
		p.BreakLine()

		p.Indent()

		for _, field := range structType.Fields {
			p.PrintStructField(ctx, field)
			p.BreakLine()
		}

		p.Outdent()
		p.PrintTerm(ctx, lang.NewSymbolTerm(structType.EndPosition, lang.TermTypeEnd, "}"))
	}

	return p
}
