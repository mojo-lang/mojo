package plugin

import (
	"github.com/mojo-lang/mojo/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/go/pkg/compiler/context"
)

type NominalTypeCompiler interface {
	CompileNominalType(ctx context.Context, typ *lang.NominalType) error
}

func IsNominalTypeCompiler(c interface{}) bool {
	if _, ok := c.(NominalTypeCompiler); ok {
		return true
	}
	return false
}
