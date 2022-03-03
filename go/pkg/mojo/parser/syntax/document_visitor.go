package syntax

import (
	"github.com/mojo-lang/core/go/pkg/logs"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"strings"
)

type DocumentVisitor struct {
	*BaseMojoParserVisitor
}

func NewDocumentVisitor() *DocumentVisitor {
	return &DocumentVisitor{}
}

func GetFreeFloatingDocument(ctx IFreeFloatingDocumentContext) *lang.Document {
	if ctx != nil {
		if document, ok := ctx.Accept(NewDocumentVisitor()).(*lang.Document); ok {
			return document
		}
	}
	return nil
}

func GetDocument(ctx IDocumentContext) *lang.Document {
	if ctx != nil {
		if document, ok := ctx.Accept(NewDocumentVisitor()).(*lang.Document); ok {
			return document
		}
	}
	return nil
}

func GetFollowingDocument(ctx IFollowingDocumentContext) *lang.Document {
	if ctx != nil {
		visitor := NewDocumentVisitor()
		if document, ok := ctx.Accept(visitor).(*lang.Document); ok {
			return document
		}
	}
	return nil
}

func GetEovDocument(ctx IEovWithDocumentContext) *lang.Document {
	if ctx != nil {
		visitor := NewDocumentVisitor()
		if document, ok := ctx.Accept(visitor).(*lang.Document); ok {
			return document
		} else {
			logs.Errorw("failed to parse the following document in EovWithDocument context")
		}
	}
	return nil
}

func GetEosDocument(ctx IEosWithDocumentContext) *lang.Document {
	if ctx != nil {
		visitor := NewDocumentVisitor()
		if document, ok := ctx.Accept(visitor).(*lang.Document); ok {
			return document
		} else {
			logs.Errorw("failed to parse the following document in EosWithDocument context")
		}
	}
	return nil
}

func (d *DocumentVisitor) VisitDocument(ctx *DocumentContext) interface{} {
	lines := ctx.AllLINE_DOCUMENT()
	document := &lang.Document{}
	for _, line := range lines {
		lineDocument := &lang.Document_Line{}

		lineDocument.StartPosition = GetPosition(line.GetSymbol())
		lineDocument.EndPosition = &lang.Position{
			Line:   lineDocument.StartPosition.Line,
			Column: lineDocument.StartPosition.Column + int64(len(line.GetText())),
		}

		if strings.HasPrefix(line.GetText(), "/// ") {
			lineDocument.Text = strings.TrimPrefix(line.GetText(), "/// ")
		} else {
			lineDocument.Text = strings.TrimPrefix(line.GetText(), "///")
		}

		document.Lines = append(document.Lines, lineDocument)
	}

	if len(lines) > 0 {
		document.StartPosition = document.Lines[0].StartPosition
		document.EndPosition = document.Lines[len(document.Lines)-1].EndPosition
		return document
	}

	return nil
}

func (d *DocumentVisitor) VisitFollowingDocument(ctx *FollowingDocumentContext) interface{} {
	document := &lang.Document{}

	line := &lang.Document_Line{}
	line.StartPosition = GetPosition(ctx.GetStart())
	line.EndPosition = GetPosition(ctx.GetStop())
	line.Following = true

	if strings.HasPrefix(ctx.GetText(), "//< ") {
		line.Text = strings.TrimPrefix(ctx.GetText(), "//< ")
	} else {
		line.Text = strings.TrimPrefix(ctx.GetText(), "//<")
	}

	document.Lines = append(document.Lines, line)
	return document
}

func (d *DocumentVisitor) VisitEovWithDocument(ctx *EovWithDocumentContext) interface{} {
	if ctx != nil {
		return GetFollowingDocument(ctx.FollowingDocument())
	}
	return nil
}

func (d *DocumentVisitor) VisitEosWithDocument(ctx *EosWithDocumentContext) interface{} {
	if ctx != nil {
		return GetFollowingDocument(ctx.FollowingDocument())
	}
	return nil
}

func (d *DocumentVisitor) VisitFreeFloatingDocument(ctx *FreeFloatingDocumentContext) interface{} {
	if ctx != nil {
		return GetDocument(ctx.Document())
	}
	return nil
}
