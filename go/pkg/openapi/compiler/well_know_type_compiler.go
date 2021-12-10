package compiler

import (
	"github.com/mojo-lang/core/go/pkg/mojo/core"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/mojo-lang/mojo/go/pkg/context"
	"github.com/mojo-lang/openapi/go/pkg/mojo/openapi"
)

type WellKnowTypeCompiler struct {
}

func (w *WellKnowTypeCompiler) Compile(ctx context.Context, nominalType *lang.NominalType) (*openapi.ReferenceableSchema, error) {
	switch nominalType.GetFullName() {
	case core.UrlTypeName:
		return openapi.NewReferenceableSchema(&openapi.Schema{
			Type:   openapi.Schema_TYPE_STRING,
			Format: "url",
		}), nil
	default:
		return nil, nil
	}
}
