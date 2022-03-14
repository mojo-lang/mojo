package format

import (
    "github.com/mojo-lang/lang/go/pkg/mojo/lang"
    "github.com/mojo-lang/mojo/go/pkg/mojo/context"
)

func (p *Printer) PrintArgument(ctx context.Context, argument *lang.Argument) {
    if argument == nil || p == nil || p.Error != nil {
        return
    }

    if len(argument.Label) > 0 {
        p.PrintRaw(argument.Label, ": ")
    }
    p.PrintExpression(ctx, argument.Value)
}
