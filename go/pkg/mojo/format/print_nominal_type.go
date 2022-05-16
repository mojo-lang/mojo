package format

import (
    "github.com/mojo-lang/core/go/pkg/mojo/core"
    "github.com/mojo-lang/lang/go/pkg/mojo/lang"
    "github.com/mojo-lang/mojo/go/pkg/mojo/context"
)

const ignoreDocument = "ignore-document"

func (p *Printer) PrintNominalType(ctx context.Context, nominalType *lang.NominalType) *Printer {
    if nominalType == nil || p == nil || p.Error != nil {
        return p
    }

    switch nominalType.GetFullName() {
    case core.ArrayTypeFullName:
        p.PrintRaw("[")
        p.PrintNominalType(ctx, nominalType.GenericArguments[0])
        p.PrintRaw("]")
    case core.MapTypeFullName:
        p.PrintRaw("{")
        p.PrintNominalType(ctx, nominalType.GenericArguments[0])
        p.PrintRaw(": ")
        p.PrintNominalType(ctx, nominalType.GenericArguments[1])
        p.PrintRaw("}")
    case core.TupleTypeFullName:
        p.PrintRaw("(")
        for i, element := range nominalType.GenericArguments {
            if i > 0 {
                p.PrintRaw(", ")
            }
            p.PrintNominalType(ctx, element)
        }
        p.PrintRaw(")")
    case core.UnionTypeFullName:
        column := p.Cursor.Column
        if column > 3 {
            column = column - 3
        } else {
            column = 0
        }
        for i, element := range nominalType.GenericArguments {
            if i > 0 {
                if p.IsNewLine() {
                    p.PrintTo(Cursor{
                        Line:   p.Cursor.Line,
                        Column: column,
                    })
                }
                p.PrintRaw(" | ")
            }
            p.PrintNominalType(ctx, element)
        }
    case core.FunctionTypeFullName:
        result := nominalType.GenericArguments[0]
        var parameters []*lang.NominalType
        if len(nominalType.GenericArguments) > 0 {
            parameters = nominalType.GenericArguments[1:]
        }
        switch len(parameters) {
        case 0:
            p.PrintRaw("()")
        case 1:
            p.PrintNominalType(ctx, parameters[0])
        default:
            p.PrintRaw("(")
            for i, parameter := range parameters {
                if i > 0 {
                    p.PrintRaw(", ")
                }
                p.PrintNominalType(ctx, parameter)
            }
            p.PrintRaw(")")
        }
        p.PrintRaw(" -> ")
        p.PrintNominalType(ctx, result)
    default:
        if nominalType.PackageName == "mojo.core" {
            p.PrintRaw(nominalType.Name)
        } else {
            p.PrintRaw(nominalType.GetFullName())
        }
        p.PrintGenericArguments(ctx, nominalType.GenericArguments)
    }

    p.PrintAttributes(ctx, nominalType.Attributes)

    if nominalType.Document != nil && nominalType.Document.Following && !context.GetBool(ctx, ignoreDocument) {
        p.PrintDocument(ctx, nominalType.Document)
    }

    return p
}

func (p *Printer) PrintGenericArguments(ctx context.Context, genericArguments []*lang.NominalType) *Printer {
    if p == nil || p.Error != nil || len(genericArguments) == 0 {
        return p
    }

    p.PrintRaw("<")
    for i, argument := range genericArguments {
        if i > 0 {
            p.PrintRaw(", ")
        }
        p.PrintNominalType(ctx, argument)
    }
    p.PrintRaw(">")

    return p
}
