package printer

import (
	"github.com/mojo-lang/mojo/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/go/pkg/compiler/context"
)

func (p *Printer) PrintFunctionDecl(ctx context.Context, decl *lang.FunctionDecl) *Printer {
	if decl == nil || p.GetError() != nil {
		return p
	}

	breaker := &OnceLineBreaker{}
	newCtx := context.WithType(ctx, decl)
	p.printPreDecl(newCtx, decl, breaker).
		Break(p).
		PrintFunctionName(newCtx, decl.Name).
		printDeclGenericParameters(newCtx, decl.GenericParameters)

	ignoredBlockStartSymbol := p.PrintFunctionSignature(newCtx, decl.Signature)
	if ignoredBlockStartSymbol {
		p.PrintBlockStatement(context.WithValues(newCtx, ignoreBlockStartSymbol, ignoredBlockStartSymbol), decl.Body)
	} else {
		p.PrintBlockStatement(newCtx, decl.Body)
	}

	return p
}

func (p *Printer) PrintFunctionName(ctx context.Context, name string) *Printer {
	if context.InterfaceDecl(ctx) != nil {
		p.PrintLine(name)
	} else {
		p.PrintLine("func", " ", name)
	}
	return p
}

func (p *Printer) PrintFunctionSignature(ctx context.Context, signature *lang.FunctionSignature) (ignoredBlockStartSymbol bool) {
	funcDecl := context.FunctionDecl(ctx)
	paramHasDoc := signature.Parameter.HasFollowingDocument()
	ignoredBlockStartSymbol = false

	p.PrintRaw("(")
	column := p.P.Cursor.Column
	var lastParamDocument *lang.Document
	for i, param := range signature.Parameter.Decls {
		if i > 0 {
			if paramHasDoc {
				if !p.IsNewLine() { // there is no document after the parameter, break line first
					p.BreakLine()
				}

				p.PrintTo(Cursor{
					Line:   p.P.Cursor.Line,
					Column: column,
				})
			} else {
				p.PrintRaw(", ")
			}
		}

		if len(param.Name) > 0 {
			p.PrintRaw(param.Name, " ")
		}

		p.PrintNominalType(ctx, param.Type)

		if i == len(signature.Parameter.Decls)-1 && param.Document != nil {
			lastParamDocument = param.Document
		} else {
			p.PrintDocument(ctx, param.Document)
		}
	}
	p.PrintRaw(")")

	if signature.Result == nil && funcDecl.Body != nil {
		p.PrintRaw(" {")
		ignoredBlockStartSymbol = true
	}

	if lastParamDocument != nil {
		p.PrintDocument(ctx, lastParamDocument)
	}

	if result := signature.GetResult(); result != nil {
		if paramHasDoc {
			if !p.IsNewLine() {
				p.BreakLine()
			}

			p.PrintTo(Cursor{
				Line:   p.P.Cursor.Line,
				Column: column,
			})
		} else {
			p.PrintRaw(" ")
		}

		p.PrintRaw("-> ")
		p.PrintNominalType(context.WithValues(ctx, ignoreDocument, true), result.Type)

		if funcDecl.Body != nil {
			p.PrintRaw(" {")
			ignoredBlockStartSymbol = true
		}

		p.PrintDocument(ctx, result.Type.Document)
	}

	return
}
