package compiler

import (
    "github.com/mojo-lang/core/go/pkg/mojo/core"
    "github.com/mojo-lang/mojo/go/pkg/mojo/context"
    "github.com/mojo-lang/mojo/go/pkg/ncraft/data"
)

const GoPackageImportKey = "@go-package-import/"
const DataMethodKey = "@data.method"

func WithGoPackageImports(ctx context.Context, options core.Options) context.Context {
    newOptions := make(core.Options)
    for key, value := range options {
        newOptions[GoPackageImportKey+key] = value
    }
    return context.WithOptions(ctx, newOptions)
}

func GoPackageImport(ctx context.Context, key string) interface{} {
    return ctx.Value(GoPackageImportKey + key)
}

func WithDataMethod(ctx context.Context, method *data.Method) context.Context {
    return context.WithValues(ctx, DataMethodKey, method)
}

func DataMethod(ctx context.Context) *data.Method {
    if method, ok := ctx.Value(DataMethodKey).(*data.Method); ok {
        return method
    }
    return nil
}
