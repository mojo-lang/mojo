package plugin

import (
	"github.com/mojo-lang/mojo/packages/lang/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/go/pkg/context"
)

type FileParser interface {
	ParseFile(ctx context.Context, fileName string) (*lang.SourceFile, error)
}
