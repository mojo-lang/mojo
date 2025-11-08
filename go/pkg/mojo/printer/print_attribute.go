package printer

import (
	"github.com/mojo-lang/mojo/packages/lang/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/go/pkg/context"
)

func (p *Printer) PrintAttribute(ctx context.Context, attribute *lang.Attribute) *Printer {
	if attribute == nil || p.GetError() != nil {
		return p
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

	return p
}

func (p *Printer) PrintAttributes(ctx context.Context, attributes lang.Attributes) *Printer {
	if p.GetError() != nil {
		return p
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
			return p.SetError(err)
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
	return p
}
