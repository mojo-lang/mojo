package format

import (
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/go/pkg/mojo/context"
)

func (p *Printer) PrintInterfaceDecl(ctx context.Context, decl *lang.InterfaceDecl) {
	if decl == nil || p == nil || p.Error != nil {
		return
	}

	breaker := &OnceLineBreaker{}
	p.printPreDecl(ctx, decl, breaker).
		Break(p).
		PrintLine("interface", " ", decl.Name).
		printDeclGenericParameters(ctx, decl.GenericParameters)

	if decl.Type != nil {
		var lastInheritDocument *lang.Document
		if len(decl.Type.Inherits) > 0 {
			p.PrintTerm(ctx, lang.NewSymbolTerm(decl.Type.InheritePosition, lang.TermTypeSymbol, ": "))

			for i, inherit := range decl.Type.Inherits {
				if i > 0 {
					p.PrintRaw(", ").PrintNominalType(ctx, inherit)
				}
			}
		}

		if !decl.IsBodyEmpty() || len(decl.Type.EndPosition.LeadingComments) > 0 {
			p.PrintTerm(ctx, lang.NewSymbolTerm(decl.Type.StartPosition, lang.TermTypeStart, " {"))
			if lastInheritDocument != nil {
				p.PrintDocument(ctx, lastInheritDocument)
			}
			p.BreakLine()
			p.Indent()

			newCtx := context.WithType(ctx, decl)
			for _, d := range decl.TypeAliasDecls {
				p.PrintTypeAliasDecl(newCtx, d)
			}

			for _, method := range decl.Type.Methods {
				p.PrintFunctionDecl(newCtx, method)
			}

			p.Outdent()
			p.PrintTerm(ctx, lang.NewSymbolTerm(decl.Type.EndPosition, lang.TermTypeEnd, "}"))
		} else {
			if lastInheritDocument != nil {
				p.PrintDocument(ctx, lastInheritDocument)
			}

			if len(decl.Type.StartPosition.LeadingComments) > 0 {
				p.PrintComments(ctx, decl.Type.StartPosition.LeadingComments...)
			}
		}
	}
}
