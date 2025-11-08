package printer

import (
	"github.com/mojo-lang/mojo/packages/lang/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/go/pkg/context"
)

func (p *Printer) PrintStructDecl(ctx context.Context, decl *lang.StructDecl) *Printer {
	if decl == nil || p.GetError() != nil {
		return p
	}

	breaker := &OnceLineBreaker{}
	p.printPreDecl(ctx, decl, breaker).
		Break(p).
		PrintLine("type", " ", decl.Name).
		printDeclGenericParameters(ctx, decl.GenericParameters)

	if decl.Type != nil {
		var lastInheritDocument *lang.Document
		if len(decl.Type.Inherits) > 0 {
			p.PrintTerm(ctx, lang.NewSymbolTerm(decl.Type.InheritPosition, lang.TermTypeSymbol, ": "))

			for i, inherit := range decl.Type.Inherits {
				if i > 0 {
					p.PrintRaw(", ")
				}
				p.PrintNominalType(ctx, inherit)
			}
		}

		if !decl.IsBodyEmpty() || len(decl.Type.EndPosition.LeadingComments) > 0 {
			p.PrintTerm(ctx, lang.NewSymbolTerm(decl.Type.StartPosition, lang.TermTypeStart, " {"))
			if lastInheritDocument != nil {
				p.PrintDocument(ctx, lastInheritDocument)
				p.BreakLine()
			}

			p.Indent()

			hasDecl := false
			for _, d := range decl.TypeAliasDecls {
				if hasDecl {
					p.BreakLine()
				} else {
					hasDecl = true
				}

				p.PrintTypeAliasDecl(ctx, d)
			}

			for _, d := range decl.EnumDecls {
				if hasDecl {
					p.BreakLine()
				} else {
					hasDecl = true
				}

				p.PrintEnumDecl(ctx, d)
			}

			for _, d := range decl.StructDecls {
				if hasDecl {
					p.BreakLine()
				} else {
					hasDecl = true
				}

				p.PrintStructDecl(ctx, d)
			}

			for _, field := range decl.Type.Fields {
				if hasDecl {
					p.BreakLine()
				} else {
					hasDecl = true
				}

				p.PrintStructField(ctx, field)
			}

			p.Outdent()
			if !p.IsNewLine() {
				p.BreakLine()
			}
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

	return p
}

func (p *Printer) PrintStructField(ctx context.Context, decl *lang.ValueDecl) *Printer {
	if decl == nil || p.GetError() != nil {
		return p
	}

	breaker := &OnceLineBreaker{}
	p.printPreDecl(ctx, decl, breaker).
		Break(p).
		PrintLine(decl.Name, ": ").
		PrintNominalType(ctx, decl.Type)

	if decl.Document != nil && decl.Document.Following {
		p.PrintDocument(ctx, decl.Document)
	}

	if comments := decl.GetEndPosition().GetTailingComments(); len(comments) > 0 {
		p.PrintComments(ctx, comments...)
	}

	return p
}
