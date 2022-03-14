package format

import (
    "github.com/mojo-lang/mojo/go/pkg/mojo/context"
    "github.com/mojo-lang/lang/go/pkg/mojo/lang"
)

func (p *Printer) PrintEnumDecl(ctx context.Context, decl *lang.EnumDecl) {
    if decl == nil || p == nil || p.Error != nil {
        return
    }

    breaker := &OnceLineBreaker{}
    p.printPreDecl(ctx, decl, breaker).
        Break(p).
        PrintLine("enum", " ", decl.Name).
        printDeclGenericParameters(ctx, decl.GenericParameters)

    if decl.Type != nil {
        var lastInheritDocument *lang.Document
        if decl.Type.UnderlyingType != nil {
            p.PrintTerm(ctx, lang.NewSymbolTerm(decl.Type.UnderlyingTypePosition, lang.TermTypeSymbol, ": "))
            p.PrintNominalType(ctx, decl.Type.UnderlyingType)
        }

        p.PrintTerm(ctx, lang.NewSymbolTerm(decl.Type.StartPosition, lang.TermTypeStart, " {"))
        if lastInheritDocument != nil {
            p.PrintDocument(ctx, lastInheritDocument)
        }
        p.BreakLine()

        p.Indent()
        for _, enumerator := range decl.Type.Enumerators {
            p.PrintEnumEnumerator(ctx, enumerator)
            if !p.IsNewLine() {
                p.BreakLine()
            }
        }
        p.Outdent()

        p.PrintTerm(ctx, lang.NewSymbolTerm(decl.Type.EndPosition, lang.TermTypeEnd, "}"))
    }
}

func (p *Printer) PrintEnumEnumerator(ctx context.Context, decl *lang.ValueDecl) {
    if decl == nil || p == nil || p.Error != nil {
        return
    }

    breaker := OnceLineBreaker{}
    if comments := decl.GetStartPosition().GetLeadingComments(); len(comments) > 0 {
        breaker.Break(p)
        p.PrintComments(ctx, comments...)
    }

    if decl.Document != nil && !decl.Document.Following {
        breaker.Break(p)
        p.PrintDocument(ctx, decl.Document)
    }

    p.PrintLine(decl.Name)
    p.PrintAttributes(ctx, decl.Attributes)

    if decl.Document != nil && decl.Document.Following {
        p.PrintDocument(ctx, decl.Document)
    }

    if comments := decl.GetEndPosition().GetTailingComments(); len(comments) > 0 {
        p.PrintComments(ctx, comments...)
    }
}
