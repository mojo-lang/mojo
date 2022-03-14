package format

import (
    "github.com/mojo-lang/lang/go/pkg/mojo/lang"
    "github.com/mojo-lang/mojo/go/pkg/mojo/context"
)

func (p *Printer) PrintStructDecl(ctx context.Context, decl *lang.StructDecl) {
    if decl == nil || p == nil || p.Error != nil {
        return
    }

    breaker := &OnceLineBreaker{}
    p.printPreDecl(ctx, decl, breaker).
        Break(p).
        PrintLine("type", " ", decl.Name).
        printDeclGenericParameters(ctx, decl.GenericParameters)

    if decl.Type != nil {
        var lastInheritDocument *lang.Document
        if len(decl.Type.Inherits) > 0 {
            p.PrintTerm(ctx, lang.NewSymbolTerm(decl.Type.InheritePosition, lang.TermTypeSymbol, ": "))

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
            }
            p.BreakLine()

            p.Indent()

            for _, d := range decl.TypeAliasDecls {
                p.PrintTypeAliasDecl(ctx, d)
                p.BreakLine()
            }

            for _, d := range decl.EnumDecls {
                p.PrintEnumDecl(ctx, d)
                p.BreakLine()
            }

            for _, d := range decl.StructDecls {
                p.PrintStructDecl(ctx, d)
                p.BreakLine()
            }

            for _, field := range decl.Type.Fields {
                p.PrintStructField(ctx, field)
                p.BreakLine()
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

func (p *Printer) PrintStructField(ctx context.Context, decl *lang.ValueDecl) *Printer {
    if decl == nil || p == nil || p.Error != nil {
        return p
    }

    breaker := &OnceLineBreaker{}
    p.printPreDecl(ctx, decl, breaker).
        Break(p).
        PrintLine(decl.Name, " ").
        PrintNominalType(ctx, decl.Type)

    if decl.Document != nil && decl.Document.Following {
        p.PrintDocument(ctx, decl.Document)
    }

    if comments := decl.GetEndPosition().GetTailingComments(); len(comments) > 0 {
        p.PrintComments(ctx, comments...)
    }

    return p
}
