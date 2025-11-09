package plugin

import (
	"github.com/mojo-lang/mojo/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/go/pkg/compiler/context"
)

type ExpressionPreHook interface {
	PreExpression(ctx context.Context, expr *lang.Expression) error
}

type ExpressionPostHook interface {
	PostExpression(ctx context.Context, expr *lang.Expression)
}
