package formatter

import (
	"github.com/mojo-lang/mojo/go/pkg/context"
	"github.com/mojo-lang/mojo/go/pkg/mojo/parser/syntax"
	"github.com/mojo-lang/mojo/go/pkg/mojo/printer"
)

type Formatter struct {
}

func New() *Formatter {
	return &Formatter{}
}

func (f *Formatter) FormatFile(ctx context.Context, filename string) error {
	return nil
}

func (f *Formatter) FormatPath(ctx context.Context, path string) error {
	return nil
}

func (f *Formatter) FormatString(ctx context.Context, content string) (string, error) {
	file, err := syntax.New(nil).ParseString(ctx, content)
	if err != nil {
		return content, err
	}

	p := printer.New(nil).PrintSourceFile(ctx, file)
	if p.GetError() != nil {
		return content, p.GetError()
	}

	return p.Buffer.String(), nil
}
