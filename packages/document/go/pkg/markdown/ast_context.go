package markdown

import "github.com/mojo-lang/mojo/packages/document/go/pkg/mojo/document"

type AstContext struct {
	Source   []byte
	Document *document.Document
	Stack    *Stack
}

type Stack []interface{}

func (b *Stack) Current() interface{} {
	if b != nil && len(*b) > 0 {
		return (*b)[len(*b)-1]
	}
	return nil
}

func (b *Stack) Previous() interface{} {
	if b != nil && len(*b) > 1 {
		return (*b)[len(*b)-2]
	}
	return nil
}

func (b *Stack) Push(block interface{}) {
	if b != nil {
		*b = append(*b, block)
	}
}

func (b *Stack) Popup() interface{} {
	if b != nil && len(*b) > 0 {
		block := b.Current()
		*b = (*b)[:len(*b)-1]
		return block
	}
	return nil
}
