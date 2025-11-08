package plugin

import (
	"github.com/mojo-lang/mojo/packages/lang/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/go/pkg/context"
)

type InterfaceCompiler interface {
	CompileInterface(ctx context.Context, decl *lang.InterfaceDecl) error
}

func IsInterfaceCompiler(c interface{}) bool {
	if _, ok := c.(InterfaceCompiler); ok {
		return true
	}
	return false
}
