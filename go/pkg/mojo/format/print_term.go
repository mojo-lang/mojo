package format

import (
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/go/pkg/mojo/context"
)

func (p *Printer) PrintTerm(ctx context.Context, term *lang.Term) *Printer {
	if term == nil || p == nil || p.Error != nil {
		return p
	}

	comments := term.GetStartPosition().GetLeadingComments()
	if len(comments) > 0 {
		p.PrintComments(ctx, comments...)
	}

	if p.IsNewLine() {
		p.PrintIndent()
	}

	p.PrintRaw(term.Value)

	return p
}
