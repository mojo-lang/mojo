package plugin

import (
	"github.com/mojo-lang/mojo/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/go/pkg/compiler/context"
)

type StringParser interface {
	ParseString(ctx context.Context, content string) (*lang.SourceFile, error)
}
