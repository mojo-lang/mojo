package printer

import (
	"strings"

	"github.com/mojo-lang/lang/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/go/pkg/context"
)

func (p *Printer) PrintComments(ctx context.Context, comments ...*lang.Comment) *Printer {
	if len(comments) == 0 || p == nil || p.Error != nil {
		return p
	}

	for i, comment := range comments {
		if i > 0 {
			p.BreakLine()
		}

		switch c := comment.Comment.(type) {
		case *lang.Comment_BlockComment:
			text := c.BlockComment.Text
			if c.BlockComment.IsInLine() {
				if c.BlockComment.HeadEmbeded {
					p.PrintRaw(c.BlockComment.Text)
				} else {
					p.PrintLine(c.BlockComment.Text)
				}
				if !c.BlockComment.TailEmbeded {
					p.BreakLine()
				}
			} else {
				lines := strings.Split(text, "\n")
				for i, line := range lines {
					if i > 0 {
						p.BreakLine()
					}
					if i == 0 {
						if c.BlockComment.HeadEmbeded {
							p.PrintRaw(line)
						} else {
							p.PrintLine(line)
						}
					} else {
						p.PrintRaw(line)
					}
				}

				if !c.BlockComment.TailEmbeded {
					p.BreakLine()
				}
			}
		case *lang.Comment_MultiLineComment:
			multiLine := c.MultiLineComment
			if multiLine.IsFollowing() {
				var cursor Cursor
				for i, line := range multiLine.Lines {
					if i > 0 {
						p.PrintTo(Cursor{
							Line:   cursor.Line + 1,
							Column: cursor.Column,
						})
					}

					cursor = p.Cursor
					p.PrintRaw(" // ", line.Text)
					p.BreakLine()
				}
			} else {
				for _, line := range multiLine.Lines {
					if len(line.Text) == 0 || strings.HasPrefix(line.Text, " ") {
						p.PrintLine("//", line.Text)
					} else {
						p.PrintLine("// ", line.Text)
					}
				}
				p.BreakLine()
			}
		case *lang.Comment_Document:
			p.PrintDocument(ctx, c.Document)
		}
	}

	return p
}
