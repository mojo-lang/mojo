package plugin

import (
	"github.com/mojo-lang/mojo/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/go/pkg/compiler/context"
)

type FileParser interface {
	ParseFile(ctx context.Context, fileName string) (*lang.SourceFile, error)
}
