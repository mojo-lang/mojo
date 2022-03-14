package format

import (
    "github.com/mojo-lang/lang/go/pkg/mojo/lang"
    "github.com/mojo-lang/mojo/go/pkg/mojo/context"
)

func (p *Printer) PrintAttribute(ctx context.Context, attribute *lang.Attribute) {
    if attribute == nil || p == nil || p.Error != nil {
        return
    }

    if comments := attribute.GetStartPosition().GetLeadingComments(); len(comments) > 0 {
        p.PrintComments(ctx, comments...)
    }

    if context.GetBool(ctx, needIndent) {
        p.PrintLine("@", attribute.GetFullName())
    } else {
        p.PrintRaw("@", attribute.GetFullName())
    }

    if len(attribute.GenericArguments) > 0 {
        p.PrintRaw("<")
        for i, argument := range attribute.GenericArguments {
            if i > 0 {
                p.PrintRaw(", ")
            }
            p.PrintNominalType(ctx, argument)
        }
        p.PrintRaw(">")
    }

    if len(attribute.Arguments) > 0 {
        p.PrintRaw("(")
        for i, argument := range attribute.Arguments {
            if i > 0 {
                p.PrintRaw(", ")
            }
            p.PrintArgument(ctx, argument)
        }
        p.PrintRaw(")")
    }
}

func (p *Printer) PrintAttributes(ctx context.Context, attributes lang.Attributes) {
    if p == nil || p.Error != nil {
        return
    }

    if attribute := attributes.GetRequiredAttribute(); attribute != nil {
        if attribute.GetBool() {
            p.PrintRaw("!")
        }
    }
    if attribute := attributes.GetOptionalAttribute(); attribute != nil {
        if attribute.GetBool() {
            p.PrintRaw("?")
        }
    }
    if attribute := attributes.GetNumberAttribute(); attribute != nil {
        if val, err := attribute.GetInteger(); err != nil {
            p.Error = err
            return
        } else {
            p.PrintRaw(" @", val)
        }
    }

    for _, attribute := range attributes {
        if attribute.IsNumber() || attribute.IsRequired() && attribute.IsOptional() {
            continue
        }

        p.PrintRaw(" ")
        p.PrintAttribute(ctx, attribute)
    }
}
