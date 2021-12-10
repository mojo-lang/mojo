package compiler

import (
	"github.com/mojo-lang/core/go/pkg/mojo/core"
	"github.com/mojo-lang/mojo/go/pkg/context"
)

const GoPackageImportKey = "@go-package-import/"

func WithGoPackageImports(ctx context.Context, options core.Options) context.Context {
	newOptions := make(core.Options)
	for key, value := range options {
		newOptions[GoPackageImportKey + key] = value
	}
	return context.WithOptions(ctx, newOptions)
}

func GoPackageImport(ctx context.Context, key string) interface{} {
	return ctx.Value(GoPackageImportKey + key)
}
