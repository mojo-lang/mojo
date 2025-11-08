package plugin

import (
	"github.com/mojo-lang/mojo/packages/lang/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/go/pkg/context"
)

type EnumCompiler interface {
	CompileEnum(ctx context.Context, decl *lang.EnumDecl) error
}

func IsEnumCompiler(c interface{}) bool {
	if _, ok := c.(EnumCompiler); ok {
		return true
	}
	return false
}
