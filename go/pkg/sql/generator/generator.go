package generator

import (
	"context"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
)

type Generator struct {
}

func (g *Generator) GenerateExpr(ctx context.Context, expr *lang.Expression) (string, []interface{}, error) {
	return "", nil, nil
}
