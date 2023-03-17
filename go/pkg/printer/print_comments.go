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
			text := c.BlockComment.Content
			if c.BlockComment.IsInLine() {
				if c.BlockComment.HeadEmbedded {
					p.PrintRaw(c.BlockComment.Content)
				} else {
					p.PrintLine(c.BlockComment.Content)
				}
				if !c.BlockComment.TailEmbedded {
					p.BreakLine()
				}
			} else {
				lines := strings.Split(text, "\n")
				for i, line := range lines {
					if i > 0 {
						p.BreakLine()
					}
					if i == 0 {
						if c.BlockComment.HeadEmbedded {
							p.PrintRaw(line)
						} else {
							p.PrintLine(line)
						}
					} else {
						p.PrintRaw(line)
					}
				}

				if !c.BlockComment.TailEmbedded {
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
					p.PrintRaw(" // ", line.Content)
					p.BreakLine()
				}
			} else {
				for _, line := range multiLine.Lines {
					if len(line.Content) == 0 || strings.HasPrefix(line.Content, " ") {
						p.PrintLine("//", line.Content)
					} else {
						p.PrintLine("// ", line.Content)
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
