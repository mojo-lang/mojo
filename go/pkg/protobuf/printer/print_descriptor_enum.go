package printer

import (
	"github.com/mojo-lang/mojo/packages/protobuf/go/pkg/mojo/protobuf/descriptor"

	"github.com/mojo-lang/mojo/go/pkg/context"
)

// PrintDescriptorEnum the enum definitions for this Enum.
func (p *Printer) PrintDescriptorEnum(ctx context.Context, enum *descriptor.Enum) *Printer {
	_ = ctx
	if enum.IsDeprecated() {
		p.PrintLine("// ", deprecationComment)
	}

	p.PrintLine("enum ", enum.GetName(), " {")
	p.Indent()
	for _, value := range enum.Values {
		if value.IsDeprecated() {
			p.BreakLine()
			p.PrintLine("// ", deprecationComment)
		}

		p.PrintLine(value.GetName(), "=", value.GetNumber(), ";")
	}
	p.Outdent()
	p.PrintLine("}")

	return p
}
