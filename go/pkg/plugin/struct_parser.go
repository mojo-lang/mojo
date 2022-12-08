package plugin

import (
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/go/pkg/context"
)

type StructParser interface {
	ParseStruct(ctx context.Context, decl *lang.StructDecl) error
}
