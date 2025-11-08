package printer

import (
	"github.com/mojo-lang/mojo/packages/lang/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/go/pkg/context"
)

func (p *Printer) PrintArgument(ctx context.Context, argument *lang.Argument) *Printer {
	if argument == nil || p.GetError() != nil {
		return p
	}

	if len(argument.Label) > 0 {
		p.PrintRaw(argument.Label, ": ")
	}
	p.PrintExpression(ctx, argument.Value)

	return p
}
