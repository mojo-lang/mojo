package plugin

import (
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/go/pkg/context"
)

type ExpressionPreHook interface {
	PreExpression(ctx context.Context, expr *lang.Expression) error
}

type ExpressionPostHook interface {
	PostExpression(ctx context.Context, expr *lang.Expression)
}
