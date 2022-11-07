package printer

import (
	"github.com/mojo-lang/mojo/go/pkg/context"
)

const (
	ColumnsKey = "@printer/Columns"
)

func WithColumns(ctx context.Context, columns Columns) context.Context {
	return context.WithValues(ctx, ColumnsKey, columns)
}

func ContextColumns(ctx context.Context) Columns {
	if columns, ok := ctx.Value(ColumnsKey).(Columns); ok {
		return columns
	}
	return Columns{}
}
