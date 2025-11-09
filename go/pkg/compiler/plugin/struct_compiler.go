package plugin

import (
	"github.com/mojo-lang/mojo/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/go/pkg/compiler/context"
)

type StructCompiler interface {
	CompileStruct(ctx context.Context, decl *lang.StructDecl) error
}

func IsStructCompiler(c interface{}) bool {
	if _, ok := c.(StructCompiler); ok {
		return true
	}
	return false
}
