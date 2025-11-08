package markdown

import (
	"errors"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/text"

	"github.com/mojo-lang/mojo/packages/document/go/pkg/mojo/document"
)

type Parser struct {
	MdEngine goldmark.Markdown
	Node     ast.Node
}

func ParseMarkdown(src []byte) (*document.Document, error) {
	return NewParser().ParseBytes(src)
}

func ParseBlocks(src []byte) ([]*document.Block, error) {
	doc, err := NewParser().ParseBytes(src)
	if err != nil {
		return nil, err
	}
	return doc.Blocks, nil
}

func NewParser() *Parser {
	return &Parser{
		MdEngine: goldmark.New(
			goldmark.WithExtensions(
				extension.GFM,
				extension.Table,
				extension.DefinitionList,
			),
			goldmark.WithParserOptions(
				parser.WithAutoHeadingID(),
			),
			goldmark.WithRendererOptions(
				html.WithHardWraps(),
				html.WithXHTML(),
			),
		),
	}
}

func (p *Parser) ParseBytes(src []byte) (*document.Document, error) {
	node := p.MdEngine.Parser().Parse(text.NewReader(src))
	if node != nil {
		p.Node = node
		return p.parse(src, node)
	}

	return nil, errors.New("failed to parse the markdown file")
}

func (p *Parser) parse(src []byte, node ast.Node) (*document.Document, error) {
	astParser := &AstParser{}
	return astParser.Parse(src, node)
}
