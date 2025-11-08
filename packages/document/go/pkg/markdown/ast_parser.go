package markdown

import (
	"github.com/mojo-lang/mojo/packages/core/go/pkg/mojo/core"
	"github.com/yuin/goldmark/ast"
	extension "github.com/yuin/goldmark/extension/ast"
	"github.com/yuin/goldmark/text"

	"github.com/mojo-lang/mojo/packages/document/go/pkg/mojo/document"
)

// AstParser a parser for parsing the `ast.Node` to `document.Document`
type AstParser struct {
}

func (a *AstParser) Parse(source []byte, n ast.Node) (*document.Document, error) {
	ctx := &AstContext{
		Document: &document.Document{
			MetaData: nil,
		},
		Source: source,
		Stack:  &Stack{},
	}

	err := ast.Walk(n, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
		return a.parse(ctx, n, entering)
	})
	if err != nil {
		return nil, err
	}

	return ctx.Document, nil
}

func (a *AstParser) parse(ctx *AstContext, n ast.Node, entering bool) (ast.WalkStatus, error) {
	switch n.Kind() {
	case ast.KindDocument:
		return a.parseDocument(ctx, n, entering)
	case ast.KindHeading:
		return a.parseHeading(ctx, n, entering)
	case ast.KindBlockquote:
		return a.parseQuoteBlock(ctx, n, entering)
	case ast.KindCodeBlock:
		return a.parseCodeBlock(ctx, n, entering)
	case ast.KindFencedCodeBlock:
		return a.parseFencedCodeBlock(ctx, n, entering)
	case ast.KindHTMLBlock:
		return a.parseHTMLBlock(ctx, n, entering)
	case ast.KindList:
		return a.parseList(ctx, n, entering)
	case ast.KindListItem:
		return a.parseListItem(ctx, n, entering)
	case ast.KindParagraph:
		return a.parseParagraph(ctx, n, entering)
	case ast.KindTextBlock:
		return a.parseTextBlock(ctx, n, entering)
	case ast.KindThematicBreak:
		return a.parseThematicBreak(ctx, n, entering)
	case extension.KindTable:
		return a.parseTable(ctx, n, entering)
	case extension.KindTableHeader:
		return a.parseTableHeader(ctx, n, entering)
	case extension.KindTableRow:
		return a.parseTableRow(ctx, n, entering)
	case extension.KindTableCell:
		return a.parseTableCell(ctx, n, entering)
	case ast.KindAutoLink:
		return a.parseAutoLink(ctx, n, entering)
	case ast.KindCodeSpan:
		return a.parseCodeSpan(ctx, n, entering)
	case ast.KindEmphasis:
		return a.parseEmphasis(ctx, n, entering)
	case ast.KindImage:
		return a.parseImage(ctx, n, entering)
	case ast.KindLink:
		return a.parseLink(ctx, n, entering)
	case ast.KindRawHTML:
		return a.parseRawHTML(ctx, n, entering)
	case ast.KindText:
		return a.parseText(ctx, n, entering)
	case ast.KindString:
		return a.parseString(ctx, n, entering)
	}

	return ast.WalkContinue, nil
}

func (a *AstParser) parseDocument(ctx *AstContext, node ast.Node, entering bool) (ast.WalkStatus, error) {
	// nothing to do
	_ = node
	if entering {
		ctx.Stack.Push(ctx.Document)
	} else {
		ctx.Stack.Popup()
		// logs.Debug(ctx.Document.String())
	}
	return ast.WalkContinue, nil
}

func (a *AstParser) parseHeading(ctx *AstContext, node ast.Node, entering bool) (ast.WalkStatus, error) {
	n := node.(*ast.Heading)
	if entering {
		header := &document.Header{
			Attribute: nil,
			Level:     int64(n.Level),
		}

		switch v := ctx.Stack.Current().(type) {
		case *document.Document:
			v.Blocks = append(v.Blocks, document.NewHeaderBlock(header))
		}

		ctx.Stack.Push(header)
	} else {
		ctx.Stack.Popup()
	}

	return ast.WalkContinue, nil
}

func (a *AstParser) parseQuoteBlock(ctx *AstContext, node ast.Node, entering bool) (ast.WalkStatus, error) {
	_ = node.(*ast.Blockquote)
	if entering {
		block := document.NewQuoteBlockBlock()
		switch v := ctx.Stack.Current().(type) {
		case *document.Document:
			v.Blocks = append(v.Blocks, block)
		case *document.OrderedList_Blocks:
			v.Vals = append(v.Vals, block)
		case *document.BulletList_Blocks:
			v.Vals = append(v.Vals, block)
		}

		ctx.Stack.Push(block.GetQuoteBlock())
	} else {
		ctx.Stack.Popup()
	}

	return ast.WalkContinue, nil
}

func (a *AstParser) parseCodeBlock(ctx *AstContext, node ast.Node, entering bool) (ast.WalkStatus, error) {
	_ = node.(*ast.CodeBlock)
	if entering {
		codeBlock := &document.CodeBlock{}
		block := document.NewCodeBlockBlock(codeBlock)
		switch v := ctx.Stack.Current().(type) {
		case *document.Document:
			v.Blocks = append(v.Blocks, block)
		case *document.OrderedList_Blocks:
			v.Vals = append(v.Vals, block)
		case *document.BulletList_Blocks:
			v.Vals = append(v.Vals, block)
		}

		ctx.Stack.Push(codeBlock)
	} else {
		ctx.Stack.Popup()
	}
	return ast.WalkContinue, nil
}

func (a *AstParser) parseFencedCodeBlock(ctx *AstContext, node ast.Node, entering bool) (ast.WalkStatus, error) {
	n := node.(*ast.FencedCodeBlock)
	if entering {
		codeBlock := &document.CodeBlock{
			Attribute: nil,
			Language:  string(n.Language(ctx.Source)),
			Lines:     getLines(ctx, n.Lines()),
		}
		block := document.NewCodeBlockBlock(codeBlock)
		switch v := ctx.Stack.Current().(type) {
		case *document.Document:
			v.Blocks = append(v.Blocks, block)
		case *document.OrderedList_Blocks:
			v.Vals = append(v.Vals, block)
		case *document.BulletList_Blocks:
			v.Vals = append(v.Vals, block)
		}

		ctx.Stack.Push(codeBlock)
	} else {
		ctx.Stack.Popup()
	}
	return ast.WalkContinue, nil
}

func (a *AstParser) parseHTMLBlock(ctx *AstContext, node ast.Node, entering bool) (ast.WalkStatus, error) {
	// nothing to do
	_ = ctx
	_ = node
	_ = entering
	return ast.WalkContinue, nil
}

func (a *AstParser) parseList(ctx *AstContext, node ast.Node, entering bool) (ast.WalkStatus, error) {
	n := node.(*ast.List)
	if entering {
		_ = n
		if n.Start == 0 {
			list := &document.BulletList{}
			block := document.NewBulletListBlock(list)
			switch v := ctx.Stack.Current().(type) {
			case *document.Document:
				v.Blocks = append(v.Blocks, block)
			case *document.OrderedList_Blocks:
				v.Vals = append(v.Vals, block)
			case *document.BulletList_Blocks:
				v.Vals = append(v.Vals, block)
			}
			ctx.Stack.Push(list)
		} else {
			list := &document.OrderedList{
				Attribute: &document.ListAttribute{
					BeginNumber: int64(n.Start),
					NumberStyle: document.ListAttribute_NUMBER_STYLE_DECIMAL,
				},
			}

			switch n.Marker {
			case '.':
				list.Attribute.NumberDelimiter = document.ListAttribute_NUMBER_DELIMITER_PERIOD
			case ')':
				list.Attribute.NumberDelimiter = document.ListAttribute_NUMBER_DELIMITER_ONE_PARENT
			}

			block := document.NewOrderedListBlock(list)
			switch v := ctx.Stack.Current().(type) {
			case *document.Document:
				v.Blocks = append(v.Blocks, block)
			case *document.OrderedList_Blocks:
				v.Vals = append(v.Vals, block)
			case *document.BulletList_Blocks:
				v.Vals = append(v.Vals, block)
			}
			ctx.Stack.Push(list)
		}
	} else {
		ctx.Stack.Popup()
	}

	return ast.WalkContinue, nil
}

func (a *AstParser) parseListItem(ctx *AstContext, node ast.Node, entering bool) (ast.WalkStatus, error) {
	_ = node.(*ast.ListItem)
	if entering {
		switch v := ctx.Stack.Current().(type) {
		case *document.OrderedList:
			blocks := &document.OrderedList_Blocks{}
			v.Items = append(v.Items, blocks)
			ctx.Stack.Push(blocks)
		case *document.BulletList:
			blocks := &document.BulletList_Blocks{}
			v.Items = append(v.Items, blocks)
			ctx.Stack.Push(blocks)
		}
	} else {
		ctx.Stack.Popup()
	}
	return ast.WalkContinue, nil
}

func (a *AstParser) parseParagraph(ctx *AstContext, node ast.Node, entering bool) (ast.WalkStatus, error) {
	_ = node.(*ast.Paragraph)
	if entering {
		block := document.NewParagraphBlock()
		switch v := ctx.Stack.Current().(type) {
		case *document.Document:
			v.Blocks = append(v.Blocks, block)
		case *document.OrderedList_Blocks:
			v.Vals = append(v.Vals, block)
		case *document.BulletList_Blocks:
			v.Vals = append(v.Vals, block)
		}

		ctx.Stack.Push(block.GetParagraph())
	} else {
		ctx.Stack.Popup()
	}
	return ast.WalkContinue, nil
}

func (a *AstParser) parseTextBlock(ctx *AstContext, node ast.Node, entering bool) (ast.WalkStatus, error) {
	_ = node.(*ast.TextBlock)
	if entering {
		block := document.NewLineBlockBlock()
		switch v := ctx.Stack.Current().(type) {
		case *document.Document:
			v.Blocks = append(v.Blocks, block)
		case *document.OrderedList_Blocks:
			v.Vals = append(v.Vals, block)
		case *document.BulletList_Blocks:
			v.Vals = append(v.Vals, block)
		}

		ctx.Stack.Push(block.GetLineBlock())
	} else {
		ctx.Stack.Popup()
	}

	return ast.WalkContinue, nil
}

func (a *AstParser) parseThematicBreak(ctx *AstContext, node ast.Node, entering bool) (ast.WalkStatus, error) {
	_ = node.(*ast.TextBlock)
	if entering {
		block := &document.Division{}

		switch v := ctx.Stack.Current().(type) {
		case *document.Document:
			v.Blocks = append(v.Blocks, document.NewDivisionBlock(block))
		}

		ctx.Stack.Push(block)
	} else {
		ctx.Stack.Popup()
	}

	return ast.WalkContinue, nil
}

func (a *AstParser) parseTable(ctx *AstContext, node ast.Node, entering bool) (ast.WalkStatus, error) {
	n := node.(*extension.Table)
	if entering {
		_ = n
		table := &document.Table{}

		switch v := ctx.Stack.Current().(type) {
		case *document.Document:
			v.Blocks = append(v.Blocks, document.NewTableBlock(table))
		}

		ctx.Stack.Push(table)
	} else {
		ctx.Stack.Popup()
	}

	return ast.WalkContinue, nil
}

func (a *AstParser) parseTableHeader(ctx *AstContext, node ast.Node, entering bool) (ast.WalkStatus, error) {
	n := node.(*extension.TableHeader)
	if entering {
		_ = n
		header := &document.Table_Header{
			Vals: nil,
		}
		if table, ok := ctx.Stack.Current().(*document.Table); ok {
			table.Header = header
		}
		ctx.Stack.Push(header)
	} else {
		ctx.Stack.Popup()
	}

	return ast.WalkContinue, nil
}

func (a *AstParser) parseTableRow(ctx *AstContext, node ast.Node, entering bool) (ast.WalkStatus, error) {
	n := node.(*extension.TableRow)
	if entering {
		_ = n
		row := &document.Table_Row{
			Vals: nil,
		}

		if table, ok := ctx.Stack.Current().(*document.Table); ok {
			table.Rows = append(table.Rows, row)
		}

		ctx.Stack.Push(row)
	} else {
		ctx.Stack.Popup()
	}

	return ast.WalkContinue, nil
}

func (a *AstParser) parseTableCell(ctx *AstContext, node ast.Node, entering bool) (ast.WalkStatus, error) {
	n := node.(*extension.TableCell)
	if entering {
		_ = n
		cell := &document.Table_Cell{
			Vals: nil,
		}

		switch v := ctx.Stack.Current().(type) {
		case *document.Table_Header:
			v.Vals = append(v.Vals, cell)
		case *document.Table_Row:
			v.Vals = append(v.Vals, cell)
		}
		ctx.Stack.Push(cell)
	} else {
		ctx.Stack.Popup()
	}

	return ast.WalkContinue, nil
}

func (a *AstParser) addInline(ctx *AstContext, inline *document.Inline) {
	switch value := ctx.Stack.Current().(type) {
	case *document.Header:
		value.Text = append(value.Text, inline)
	case *document.Paragraph:
		value.Inlines = append(value.Inlines, inline)
	case *document.Table_Cell:
		value.Vals = append(value.Vals, document.NewPainBlock(inline))
	case *document.LineBlock:
		if len(value.Lines) == 0 {
			value.Lines = append(value.Lines, &document.Line{Vals: []*document.Inline{inline}})
		} else {
			last := value.Lines[len(value.Lines)-1]
			last.Vals = append(last.Vals, inline)
		}
	}
}

func (a *AstParser) parseAutoLink(ctx *AstContext, node ast.Node, entering bool) (ast.WalkStatus, error) {
	n := node.(*ast.AutoLink)
	if entering {
		_ = n
		link := &document.Link{}
		a.addInline(ctx, document.NewLinkInline(link))
		ctx.Stack.Push(link)
	} else {
		ctx.Stack.Popup()
	}
	return ast.WalkContinue, nil
}

func (a *AstParser) parseCodeSpan(ctx *AstContext, node ast.Node, entering bool) (ast.WalkStatus, error) {
	_ = node.(*ast.CodeSpan)
	if entering {
		code := &document.Code{}
		a.addInline(ctx, document.NewCodeInline(code))
		ctx.Stack.Push(code)
	} else {
		ctx.Stack.Popup()
	}
	return ast.WalkContinue, nil
}

func (a *AstParser) parseEmphasis(ctx *AstContext, node ast.Node, entering bool) (ast.WalkStatus, error) {
	n := node.(*ast.Emphasis)
	if entering {
		if n.Level <= 1 {
			emphasized := &document.Emphasized{}
			a.addInline(ctx, document.NewEmphasizedInline())
			ctx.Stack.Push(emphasized)
		} else {
			strong := &document.Strong{}
			a.addInline(ctx, document.NewStrongInline())
			ctx.Stack.Push(strong)
		}
	} else {
		ctx.Stack.Popup()
	}
	return ast.WalkContinue, nil
}

func (a *AstParser) parseImage(ctx *AstContext, node ast.Node, entering bool) (ast.WalkStatus, error) {
	n := node.(*ast.Image)
	if entering {
		url, err := core.NewUrl(string(n.Destination))
		if err != nil {
			return ast.WalkStop, err
		}
		image := &document.Image{
			Target: &document.Target{
				Title: string(n.Title),
				Url:   url,
			},
		}
		a.addInline(ctx, document.NewImageInline(image))
		ctx.Stack.Push(image)
	} else {
		ctx.Stack.Popup()
	}
	return ast.WalkContinue, nil
}

func (a *AstParser) parseLink(ctx *AstContext, node ast.Node, entering bool) (ast.WalkStatus, error) {
	n := node.(*ast.Link)
	if entering {
		_ = n
		link := &document.Link{}
		a.addInline(ctx, document.NewLinkInline(link))
		ctx.Stack.Push(link)
	} else {
		ctx.Stack.Popup()
	}
	return ast.WalkContinue, nil
}

func (a *AstParser) parseText(ctx *AstContext, node ast.Node, entering bool) (ast.WalkStatus, error) {
	n := node.(*ast.Text)
	if entering {
		isLineBreak := func() bool { return n.SoftLineBreak() || n.HardLineBreak() }
		content := string(n.Text(ctx.Source))
		inline := document.NewTextInline(content)
		if v := ctx.Stack.Current(); v != nil {
			switch value := v.(type) {
			case *document.Header:
				value.Text = append(value.Text, inline)
			case *document.Table_Cell:
				value.Vals = append(value.Vals, document.NewPainBlock(inline))
			case *document.Paragraph:
				value.Inlines = append(value.Inlines, inline)
			case *document.LineBlock:
				if len(value.Lines) == 0 {
					value.Lines = append(value.Lines, &document.Line{Vals: []*document.Inline{inline}})
				} else if isLineBreak() {
					last := value.Lines[len(value.Lines)-1]
					last.Vals = append(last.Vals, inline)

					value.Lines = append(value.Lines, &document.Line{})
				} else {
					last := value.Lines[len(value.Lines)-1]
					last.Vals = append(last.Vals, inline)
				}
			case *document.Emphasized:
				value.Vals = append(value.Vals, inline)
			case *document.Strong:
				value.Vals = append(value.Vals, inline)
			case *document.Image:
				value.Description = append(value.Description, inline)
			case *document.Link:
				value.Description = append(value.Description, inline)
			case *document.Code:
				value.Content = content
			}
		}
	} else {
	}
	return ast.WalkContinue, nil
}

func (a *AstParser) parseString(ctx *AstContext, node ast.Node, entering bool) (ast.WalkStatus, error) {
	n := node.(*ast.String)
	if entering {
		_ = n
		content := string(n.Text(ctx.Source))
		inline := document.NewTextInline(content)
		if v := ctx.Stack.Current(); v != nil {
			switch value := v.(type) {
			case *document.Header:
				value.Text = append(value.Text, inline)
			case *document.Table_Cell:
				value.Vals = append(value.Vals, document.NewPainBlock(inline))
			case *document.Paragraph:
				value.Inlines = append(value.Inlines, inline)
			case *document.LineBlock:
				if len(value.Lines) == 0 {
					value.Lines = append(value.Lines, &document.Line{Vals: []*document.Inline{inline}})
				} else {
					last := value.Lines[len(value.Lines)-1]
					last.Vals = append(last.Vals, inline)
				}
			case *document.Emphasized:
				value.Vals = append(value.Vals, inline)
			case *document.Code:
				value.Content = content
			}
		}
	} else {
	}
	return ast.WalkContinue, nil
}

func (a *AstParser) parseRawHTML(ctx *AstContext, node ast.Node, entering bool) (ast.WalkStatus, error) {
	// nothing to do
	_ = ctx
	_ = node
	_ = entering
	return ast.WalkContinue, nil
}

func getLines(ctx *AstContext, lines *text.Segments) []*document.Line {
	var ls []*document.Line
	for i := 0; i < lines.Len(); i++ {
		segment := lines.At(i)
		content := string(segment.Value(ctx.Source))
		line := &document.Line{
			Vals: []*document.Inline{document.NewTextInline(content)},
		}
		ls = append(ls, line)
	}
	return ls
}
