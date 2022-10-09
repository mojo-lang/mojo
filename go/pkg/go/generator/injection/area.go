package injection

import "go/token"

type Area interface {
	GetStart() token.Pos
	GetEnd() token.Pos
	GetContent() string
}

type TextArea struct {
	Start   token.Pos
	End     token.Pos
	Content string
}

func (a TextArea) GetStart() token.Pos {
	return a.Start
}

func (a TextArea) GetEnd() token.Pos {
	return a.End
}

func (a TextArea) GetContent() string {
	return a.Content
}
