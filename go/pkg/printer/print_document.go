package printer

import (
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/go/pkg/context"
)

func (p *Printer) PrintDocument(ctx context.Context, document *lang.Document) {
	if document == nil || p == nil || p.Error != nil {
		return
	}

	if document.Following {
		var cursor Cursor
		for i, line := range document.Lines {
			if i > 0 {
				p.PrintTo(Cursor{
					Line:   cursor.Line + 1,
					Column: cursor.Column,
				})
			}

			cursor = p.Cursor
			p.PrintRaw(" //< ", line.Text)
			p.BreakLine()
		}
	} else {
		for _, line := range document.Lines {
			p.PrintLine("/// ", line.Text)
			p.BreakLine()
		}
	}
}
