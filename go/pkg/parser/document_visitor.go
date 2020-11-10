package parser

import "github.com/mojo-lang/lang/go/pkg/lang"

type DocumentVisitor struct {
	*BaseMojoParserVisitor
}

func NewDocumentVisitor() *DocumentVisitor {
	return &DocumentVisitor{}
}

func GetDocument(ctx IDocumentContext) *lang.Document {
	if ctx != nil {
		visitor := NewDocumentVisitor()
		return ctx.Accept(visitor).(*lang.Document)
	}
	return nil
}

func GetFollowingDocument(ctx IFollowingDocumentContext) *lang.Document {
	if ctx != nil {
		visitor := NewDocumentVisitor()
		return ctx.Accept(visitor).(*lang.Document)
	}
	return nil
}

func (d *DocumentVisitor) VisitDocument(ctx *DocumentContext) interface{} {
	lines := ctx.AllLineDocument()
	document := &lang.Document{}
	for _, line := range lines {
		lineDocument := &lang.LineDocument{}
		lineDocument.Line = line.GetText()
		//lineDocument.StartPosition = GetPosition(line.GetStart())
		//lineDocument.EndPosition = GetPosition(line.GetStop())

		document.Lines = append(document.Lines, lineDocument)
	}

	return document
}

func (d *DocumentVisitor) VisitFollowingDocument (ctx *FollowingDocumentContext) interface{} {
	document := &lang.Document{}
	return document
}