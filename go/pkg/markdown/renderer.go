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
	for _, item := range list.Items {
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

		r.blocks(cell.Vals, out, "<br>")

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
			r.blocks(cell.Vals, out, "<br>")
			out.WriteByte(' ')
		}
		out.WriteString("|\n")
	}
}

func (r *Renderer) codeBlock(block *document.CodeBlock, out Writer) {
	out.WriteString("```")
	out.WriteString(block.Language)
	out.WriteString("\n")

	code := block.GetCode()
	if formattedCode, ok := formatCode(block.Language, code); ok {
		out.Write(formattedCode)
	} else {
		out.Write(code)
	}

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
	out.WriteByte('`')
	out.WriteString(code.Content)
	out.WriteByte('`')
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
