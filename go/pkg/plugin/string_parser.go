package plugin

import (
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/go/pkg/context"
)

type StringParser interface {
	ParseString(ctx context.Context, content string) (*lang.SourceFile, error)
}
