package lang

import (
	"bytes"

	"github.com/mojo-lang/mojo/packages/core/go/pkg/logs"
	"github.com/mojo-lang/mojo/packages/document/go/pkg/markdown"
)

type DocumentGetter interface {
	GetDocument() *Document
}

func (x *Document) SetStartPosition(position *Position) {
	if x != nil {
		x.StartPosition = PatchPosition(x.StartPosition, position)
	}
}

func (x *Document) SetEndPosition(position *Position) {
	if x != nil {
		x.EndPosition = PatchPosition(x.EndPosition, position)
	}
}

func (x *Document) GetContent() string {
	if x != nil {
		content := bytes.Buffer{}
		for i, line := range x.Lines {
			if content.Len() == 0 && len(line.Content) == 0 {
				continue
			}
			content.WriteString(line.Content)

			if i < len(x.Lines)-1 {
				content.WriteString("\n")
			}
		}
		return content.String()
	}
	return ""
}

func (x *Document) Parse() *Document {
	if x != nil {
		mk := markdown.New()
		doc, err := mk.Parse(x.GetContent())
		if err != nil {
			logs.Warnw("failed to parse the document to markdown")
		} else {
			x.Structured = doc
		}
	}
	return x
}
