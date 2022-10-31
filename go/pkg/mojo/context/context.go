package context

import (
	"context"
	"time"

	"github.com/mojo-lang/core/go/pkg/mojo/core"
)

type Context = context.Context

func Empty() Context {
	return context.Background()
}

type valuesContext struct {
	context.Context
	core.Options
}

func (*valuesContext) Deadline() (deadline time.Time, ok bool) {
	return
}

func (*valuesContext) Done() <-chan struct{} {
	return nil
}

func (*valuesContext) Err() error {
	return nil
}

func (c *valuesContext) Value(key interface{}) interface{} {
	if v, ok := c.Options[key.(string)]; ok {
		return v
	}
	return c.Context.Value(key)
}

func WithOptions(ctx context.Context, options core.Options) context.Context {
	return &valuesContext{
		Context: ctx,
		Options: options,
	}
}

func WithValues(ctx context.Context, kvs ...interface{}) context.Context {
	options := make(core.Options)
	for i := 0; i < len(kvs)-1; i += 2 {
		options[kvs[i].(string)] = kvs[i+1]
	}
	return WithOptions(ctx, options)
}

func Values(ctx context.Context, key string) []interface{} {
	var values []interface{}
	if c, ok := ctx.(*valuesContext); ok {
		if v, o := c.Options[key]; o {
			values = append(values, v)
		}

		values = append(values, Values(c.Context, key)...)
	}
	return values
}

func PreviousValue(ctx context.Context, key string, index int) interface{} {
	values := Values(ctx, key)
	if index < 0 || index >= len(values) {
		return nil
	}
	return values[index]
}

func GetBool(ctx context.Context, key string) bool {
	values := Values(ctx, key)
	if len(values) > 0 {
		if v, ok := values[0].(bool); ok {
			return v
		}
	}
	return false
}
