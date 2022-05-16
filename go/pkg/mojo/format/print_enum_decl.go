package format

import (
    "github.com/mojo-lang/lang/go/pkg/mojo/lang"
    "github.com/mojo-lang/mojo/go/pkg/mojo/context"
    "strings"
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

        newCtx := WithVerticalLines(ctx, p.calcEnumVerticalLines(ctx, decl.Type))
        for _, enumerator := range decl.Type.Enumerators {
            p.PrintEnumEnumerator(newCtx, enumerator)
            if !p.IsNewLine() {
                p.BreakLine()
            }
        }

        p.Outdent()

        p.PrintTerm(ctx, lang.NewSymbolTerm(decl.Type.EndPosition, lang.TermTypeEnd, "}"))
    }
}

func (p *Printer) calcEnumVerticalLines(ctx context.Context, enum *lang.EnumType) VerticalLines {
    nameMax := 0
    attributesMax := 0
    for _, enumerator := range enum.Enumerators {
        if len(enumerator.Name) > nameMax {
            nameMax = len(enumerator.Name)
        }

        sb := &strings.Builder{}
        printer := NewPrinter(Config{}, sb)
        printer.PrintAttributes(context.Empty(), enumerator.Attributes)
        s := sb.String()
        if len(s) > attributesMax {
            attributesMax = len(s)
        }
    }

    lines := VerticalLines{}
    lines = append(lines, nameMax)
    if attributesMax > 0 {
        lines = append(lines, attributesMax)
    }
    return lines
}

func (p *Printer) PrintEnumEnumerator(ctx context.Context, decl *lang.ValueDecl) {
    if decl == nil || p == nil || p.Error != nil {
        return
    }
    vLines := ContextVerticalLines(ctx)
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

    vLines.PrintTo(0, p)
    p.PrintAttributes(ctx, decl.Attributes)

    if decl.Document != nil && decl.Document.Following {
        vLines.PrintTo(1, p)
        p.PrintDocument(ctx, decl.Document)
    }

    if comments := decl.GetEndPosition().GetTailingComments(); len(comments) > 0 {
        p.PrintComments(ctx, comments...)
    }
}
