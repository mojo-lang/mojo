package format

import (
    "github.com/mojo-lang/mojo/go/pkg/mojo/context"
    "github.com/mojo-lang/lang/go/pkg/mojo/lang"
)

func (p *Printer) PrintTerm(ctx context.Context, term *lang.Term) *Printer {
    if term == nil || p == nil || p.Error != nil {
        return p
    }

    comments := term.GetStartPosition().GetLeadingComments()
    if len(comments) > 0 {
        p.PrintComments(ctx, comments...)
    }

    p.PrintRaw(term.Value)

    return p
}
