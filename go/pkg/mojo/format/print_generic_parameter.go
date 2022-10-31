package format

import (
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/go/pkg/mojo/context"
)

func (p *Printer) PrintGenericParameter(ctx context.Context, parameter *lang.GenericParameter) *Printer {
	if parameter == nil || p == nil || p.Error != nil {
		return p
	}

	p.PrintRaw(parameter.Name)

	if parameter.Constraint != nil {
		p.PrintRaw(": ").PrintNominalType(ctx, parameter.Constraint)
	}

	return p
}
