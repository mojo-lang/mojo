package printer

import (
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/go/pkg/context"
)

func (p *Printer) printPreDecl(ctx context.Context, decl interface{}, breaker *OnceLineBreaker) *OnceLineBreaker {
	hasComment := false
	if getter, ok := decl.(lang.StartPositionGetter); ok {
		if comments := getter.GetStartPosition().GetLeadingComments(); len(comments) > 0 {
			breaker.Break(p)
			p.PrintComments(ctx, comments...)
			hasComment = true
		}
	}

	if getter, ok := decl.(lang.DocumentGetter); ok {
		if document := getter.GetDocument(); document != nil && !document.Following {
			breaker.Break(p)
			if hasComment {
				p.BreakLine()
			}

			p.PrintDocument(ctx, document)
		}
	}

	if getter, ok := decl.(lang.AttributesGetter); ok {
		for _, attribute := range getter.GetAttributes() {
			breaker.Break(p)
			p.PrintAttribute(context.WithValues(ctx, needIndent, true), attribute)
			p.BreakLine()
		}
	}

	return breaker
}

func (p *Printer) printDeclGenericParameters(ctx context.Context, parameters []*lang.GenericParameter) *Printer {
	if len(parameters) > 0 {
		p.PrintRaw("<")
		for i, parameter := range parameters {
			if i > 0 {
				p.PrintRaw(" ")
			}
			p.PrintGenericParameter(ctx, parameter)

			if i < len(parameters)-1 {
				p.PrintRaw(",")
			}
		}
		p.PrintRaw(">")
	}
	return p
}
