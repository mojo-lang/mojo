package plugin

import (
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/go/pkg/context"
)

type ExpressionParser interface {
	ParseExpression(ctx context.Context, expr *lang.Expression) error
}
