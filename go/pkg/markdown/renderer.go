package markdown

import (
	"bytes"
	"fmt"
	"go/format"
	"io"
	"strings"

	"github.com/mojo-lang/mojo/go/pkg/mojo/document"
)

func FormatDocument(doc *document.Document) (string, error) {
	out := &bytes.Buffer{}
	r := &Renderer{}
	r.blocks(doc.Blocks, out, "\n")
	return out.String(), nil
}

func FormatBlock(block *document.Block) (string, error) {
	return FormatBlocks(block)
}

func FormatBlocks(blocks ...*document.Block) (string, error) {
	out := &bytes.Buffer{}
	r := &Renderer{}
	r.blocks(blocks, out, "\n")
	return out.String(), nil
}

func FormatInlines(inlines ...*document.Inline) (string, error) {
	out := &bytes.Buffer{}
	r := &Renderer{}
	r.inlines(inlines, out)
	return out.String(), nil
}

func FormatInline(inline *document.Inline) (string, error) {
	return FormatInlines(inline)
}

type Renderer struct {
	inTableCell bool
}

func (r *Renderer) RenderDocument(doc *document.Document, writer io.Writer) error {
	out := &bytes.Buffer{}
	r.blocks(doc.Blocks, out, "\n")
	_, err := writer.Write(out.Bytes())
	return err
}

func (r *Renderer) RenderBlock(block *document.Block, writer io.Writer) error {
	out := &bytes.Buffer{}
	r.block(block, out)

	_, err := writer.Write(out.Bytes())
	return err
}

func (r *Renderer) blocks(blocks []*document.Block, out Writer, lineBreak string) {
	for i, block := range blocks {
		if i > 0 {
			out.WriteString(lineBreak)
		}
		r.block(block, out)
	}
}

func (r *Renderer) block(block *document.Block, out Writer) {
	switch block.Block.(type) {
	case *document.Block_Header:
		r.header(block.GetHeader(), out)
	case *document.Block_Plain:
		r.plain(block.GetPlain(), out)
	case *document.Block_Paragraph:
		r.paragraph(block.GetParagraph(), out)
	case *document.Block_OrderedList:
		r.orderedList(block.GetOrderedList(), out)
	case *document.Block_BulletList:
		r.bulletList(block.GetBulletList(), out)
	case *document.Block_DefinitionList:
		r.definitionList(block.GetDefinitionList(), out)
	case *document.Block_Table:
		r.table(block.GetTable(), out)
	case *document.Block_CodeBlock:
		r.codeBlock(block.GetCodeBlock(), out)
	case *document.Block_QuoteBlock:
		r.quoteBlock(block.GetQuoteBlock(), out)
	case *document.Block_LineBlock:
		r.lineBlock(block.GetLineBlock(), out)
	case *document.Block_Division:
		r.division(block.GetDivision(), out)
	}
}

func (r *Renderer) header(header *document.Header, out Writer) {
	fmt.Fprint(out, "\n")
	fmt.Fprint(out, strings.Repeat("#", int(header.Level)), " ")

	r.inlines(header.Text, out)

	// out.WriteString("\n")
}

func (r *Renderer) plain(pain *document.Plain, out Writer) {
	r.inlines(pain.Inlines, out)
}

func (r *Renderer) paragraph(paragraph *document.Paragraph, out Writer) {
	r.inlines(paragraph.Inlines, out)

	// if !text() {
	//	out.Truncate(marker)
	//	return
	// }
	// out.WriteString("\n")
}

func (r *Renderer) orderedList(list *document.OrderedList, out Writer) {
	for i, item := range list.Items {
		if i > 0 {
			out.WriteString("\n")
		}

		out.WriteString("1.")
		r.blocks(item.Vals, NewIndentWriter(out, 1), "\n")
	}
}

func (r *Renderer) bulletList(list *document.BulletList, out Writer) {
	for i, item := range list.Items {
		if i > 0 {
			out.WriteString("\n")
		}
		out.WriteString("-")
		r.blocks(item.Vals, NewIndentWriter(out, 1), "\n")
	}
}

func (r *Renderer) definitionList(list *document.DefinitionList, out Writer) {
}

func (r *Renderer) table(table *document.Table, out Writer) {
	for _, cell := range table.Header.Vals {
		out.WriteByte('|')
		out.WriteByte(' ')

		r.tableCell(cell, out)

		// for i := r.stringWidth(cell); i < r.columnWidths[column]; i++ {
		//	out.WriteByte(' ')
		// }

		out.WriteByte(' ')
	}
	out.WriteString("|\n")

	for range table.Header.Vals {
		out.WriteByte('|')
		// out.WriteByte(':')
		out.WriteByte('-')
		out.WriteByte('-')
		// out.WriteByte(':')
		out.WriteByte('-')
	}
	out.WriteString("|\n")

	for _, row := range table.Rows {
		for _, cell := range row.Vals {
			out.WriteByte('|')
			out.WriteByte(' ')
			r.tableCell(cell, out)
			out.WriteByte(' ')
		}
		out.WriteString("|\n")
	}
}

// Render cells separately so nested blocks cannot introduce physical rows or columns.
func (r *Renderer) tableCell(cell *document.Table_Cell, out Writer) {
	var content bytes.Buffer
	cellRenderer := *r
	cellRenderer.inTableCell = true
	cellRenderer.blocks(cell.Vals, &content, "<br>")
	text := strings.ReplaceAll(content.String(), "\r\n", "\n")
	text = strings.ReplaceAll(text, "\r", "\n")
	text = strings.ReplaceAll(text, "\n", "<br>")
	slashes := 0
	for _, char := range text {
		if char == '|' && slashes == 0 {
			out.WriteByte('\\')
		}
		out.WriteString(string(char))
		if char == '\\' {
			slashes++
		} else {
			slashes = 0
		}
	}
}

func (r *Renderer) codeBlock(block *document.CodeBlock, out Writer) {
	code := block.GetCode()
	if formattedCode, ok := formatCode(block.Language, code); ok {
		code = formattedCode
	}
	if r.inTableCell {
		// Fenced blocks cannot live inside pipe tables. Keep each code line as an
		// inline code span, with explicit breaks between lines.
		text := strings.ReplaceAll(string(code), "\r\n", "\n")
		text = strings.ReplaceAll(text, "\r", "\n")
		for i, line := range strings.Split(strings.TrimSuffix(text, "\n"), "\n") {
			if i > 0 {
				out.WriteString("<br>")
			}
			if line != "" {
				r.code(&document.Code{Content: line}, out)
			}
		}
		return
	}
	out.WriteString("```")
	out.WriteString(block.Language)
	out.WriteString("\n")
	out.Write(code)
	out.WriteString("\n```\n")
}

func (r *Renderer) quoteBlock(blockQuote *document.QuoteBlock, out Writer) {
	// lines := bytes.Split(text, []byte("\n"))
	// for i, line := range lines {
	//	if i == len(lines)-1 {
	//		continue
	//	}
	//	out.WriteString(">")
	//	if len(line) != 0 {
	//		out.WriteString(" ")
	//		out.Write(line)
	//	}
	//	out.WriteString("\n")
	// }
}

func (r *Renderer) lineBlock(lineBlock *document.LineBlock, out Writer) {
	for i, line := range lineBlock.Lines {
		if i > 0 {
			out.WriteString("\n")
		}
		r.inlines(line.Vals, out)
	}
}

func (r *Renderer) division(division *document.Division, out Writer) {
	if division.Attribute != nil {
		if division.Attribute.Identifier == "toc" && len(division.Content) > 0 {
			r.block(division.Content[0], out)
		}
	}
}

func (r *Renderer) inlines(inlines []*document.Inline, out Writer) {
	for _, inline := range inlines {
		r.inline(inline, out)
	}
}

func (r *Renderer) inline(inline *document.Inline, out Writer) {
	switch inline.Inline.(type) {
	case *document.Inline_Text:
		out.WriteString(inline.GetText().Val)
	case *document.Inline_Emphasized:
		r.emphasized(inline.GetEmphasized(), out)
	case *document.Inline_Strong:
		r.strong(inline.GetStrong(), out)
	case *document.Inline_Link:
		r.link(inline.GetLink(), out)
	case *document.Inline_Image:
		r.image(inline.GetImage(), out)
	case *document.Inline_Code:
		r.code(inline.GetCode(), out)
	}
}

func (r *Renderer) link(link *document.Link, out Writer) {
	if link != nil && link.Target != nil {
		out.WriteString("[")
		r.inlines(link.Description, out)
		out.WriteString("](")
		out.WriteString(link.Target.Url.Format())

		if len(link.Target.Title) != 0 {
			out.WriteString(` "`)
			out.WriteString(link.Target.Title)
			out.WriteString(`"`)
		}
		out.WriteString(")")
	}
}

func (r *Renderer) image(image *document.Image, out Writer) {
	out.WriteString("![")
	r.inlines(image.Description, out)
	out.WriteString("](")
	out.WriteString(image.Target.Url.Format())

	if len(image.Target.Title) != 0 {
		out.WriteString(` "`)
		out.WriteString(image.Target.Title)
		out.WriteString(`"`)
	}
	out.WriteString(")")
}

func (r *Renderer) code(code *document.Code, out Writer) {
	delimiter := "`"
	for strings.Contains(code.Content, delimiter) {
		delimiter += "`"
	}
	padding := strings.HasPrefix(code.Content, "`") || strings.HasSuffix(code.Content, "`") ||
		(strings.HasPrefix(code.Content, " ") && strings.HasSuffix(code.Content, " ") && strings.TrimSpace(code.Content) != "")
	out.WriteString(delimiter)
	if padding {
		out.WriteByte(' ')
	}
	content := code.Content
	if r.inTableCell {
		content = strings.ReplaceAll(content, "|", "\\|")
	}
	out.WriteString(content)
	if padding {
		out.WriteByte(' ')
	}
	out.WriteString(delimiter)
}

func (r *Renderer) emphasized(emphasized *document.Emphasized, out Writer) {
	out.WriteByte('*')
	r.inlines(emphasized.Vals, out)
	out.WriteByte('*')
}

func (r *Renderer) strong(strong *document.Strong, out Writer) {
	out.WriteString("***")
	r.inlines(strong.Vals, out)
	out.WriteString("***")
}

func formatCode(lang string, text []byte) (formattedCode []byte, ok bool) {
	switch lang {
	case "Go", "go":
		gofmt, err := format.Source(text)
		if err != nil {
			return nil, false
		}
		return gofmt, true
	default:
		return nil, false
	}
}
