package format

import (
    "github.com/mojo-lang/core/go/pkg/mojo/core"
    "github.com/mojo-lang/lang/go/pkg/mojo/lang"
    "github.com/mojo-lang/mojo/go/pkg/mojo/context"
)

const ignoreDocument = "ignore-document"

func (p *Printer) PrintNominalType(ctx context.Context, nominalType *lang.NominalType) {
    if nominalType == nil || p == nil || p.Error != nil {
        return
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
        for i, element := range nominalType.GenericArguments {
            if i > 0 {
                p.PrintRaw("| ")
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

        if len(nominalType.GenericArguments) > 0 {
            p.PrintRaw("<")
            for i, argument := range nominalType.GenericArguments {
                if i > 0 {
                    p.PrintRaw(", ")
                    p.PrintNominalType(ctx, argument)
                }
            }
            p.PrintRaw(">")
        }
    }

    p.PrintAttributes(ctx, nominalType.Attributes)

    if nominalType.Document != nil && nominalType.Document.Following && !context.GetBool(ctx, ignoreDocument) {
        p.PrintDocument(ctx, nominalType.Document)
    }
}
