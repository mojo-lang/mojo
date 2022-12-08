package plugin

import (
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/go/pkg/context"
)

type TypeAliasCompiler interface {
	CompileTypeAlias(ctx context.Context, decl *lang.TypeAliasDecl) error
}

func IsTypeAliasCompiler(c interface{}) bool {
	if _, ok := c.(TypeAliasCompiler); ok {
		return true
	}
	return false
}
