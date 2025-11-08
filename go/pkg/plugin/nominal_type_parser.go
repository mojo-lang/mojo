package plugin

import (
	"github.com/mojo-lang/mojo/packages/lang/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/go/pkg/context"
)

type NominalTypeParser interface {
	ParseNominalType(ctx context.Context, typ *lang.NominalType) error
}
