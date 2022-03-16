package plugin

import "github.com/mojo-lang/mojo/go/pkg/mojo/context"

const pluginsKey = "plugins"

func WithPlugins(ctx context.Context, plugins *Plugins) context.Context {
    return context.WithValues(ctx, pluginsKey, plugins)
}

func ContextPlugins(ctx context.Context) *Plugins {
    if plugins, ok := ctx.Value(pluginsKey).(*Plugins); ok {
        return plugins
    }
    return nil
}
