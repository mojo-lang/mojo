package plugin

import (
	"github.com/mojo-lang/mojo/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/go/pkg/compiler/context"
)

type SourceFileCompiler interface {
	CompileSourceFile(ctx context.Context, file *lang.SourceFile) error
}

func IsSourceFileCompiler(c interface{}) bool {
	if _, ok := c.(SourceFileCompiler); ok {
		return true
	}
	return false
}
