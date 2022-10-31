package parser

import (
	"io/fs"
	"os"

	"github.com/mojo-lang/lang/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/go/pkg/mojo/context"
)

const fsCacheKey = "@fsCache"

type FsCache map[string]fs.FS

func WithFsCache(ctx context.Context, cache FsCache) context.Context {
	if ContextFsCache(ctx) == nil {
		return context.WithValues(ctx, fsCacheKey, cache)
	}
	return ctx
}

func ContextFsCache(ctx context.Context) FsCache {
	if cache, ok := ctx.Value(fsCacheKey).(FsCache); ok {
		return cache
	}
	return nil
}

const declaredPackageKey = "@declaredPackage"

func WithDeclaredPackage(ctx context.Context, pkg *lang.Package) context.Context {
	return context.WithValues(ctx, declaredPackageKey, pkg)
}

func ContextDeclaredPackage(ctx context.Context) *lang.Package {
	if pkg, ok := ctx.Value(declaredPackageKey).(*lang.Package); ok {
		return pkg
	}
	return nil
}

const workingDirKey = "workingDir"

func WithWorkingDir(ctx context.Context, workingDir string) context.Context {
	return context.WithValues(ctx, workingDirKey, workingDir)
}

func ContextWorkingDir(ctx context.Context) string {
	if dir, ok := ctx.Value(workingDirKey).(string); ok {
		return dir
	}

	if dir, err := os.Getwd(); err != nil {
		return ""
	} else {
		return dir
	}
}

const packageNameKey = "packageName"

func WithPackageName(ctx context.Context, pkgName string) context.Context {
	return context.WithValues(ctx, packageNameKey, pkgName)
}

func ContextPackageName(ctx context.Context) string {
	if name, ok := ctx.Value(packageNameKey).(string); ok {
		return name
	}
	return ""
}
