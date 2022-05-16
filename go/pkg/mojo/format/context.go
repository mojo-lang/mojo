package format

import (
    "github.com/mojo-lang/mojo/go/pkg/mojo/context"
)

const (
    VerticalLinesKey = "@format/VerticalLines"
)

func WithVerticalLines(ctx context.Context, lines VerticalLines) context.Context {
    return context.WithValues(ctx, VerticalLinesKey, lines)
}

func ContextVerticalLines(ctx context.Context) VerticalLines {
    if lines, ok := ctx.Value(VerticalLinesKey).(VerticalLines); ok {
        return lines
    }
    return VerticalLines{}
}
