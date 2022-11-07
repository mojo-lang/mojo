package gokit

import (
	"github.com/mojo-lang/core/go/pkg/logs"
	"github.com/mojo-lang/core/go/pkg/mojo/core"

	"github.com/mojo-lang/mojo/go/pkg/cmd/build/builder"
	"github.com/mojo-lang/mojo/go/pkg/context"
	"github.com/mojo-lang/mojo/go/pkg/ncraft/compiler"
	"github.com/mojo-lang/mojo/go/pkg/ncraft/gokit"
)

type ClientBuilder struct {
	builder.Builder
	Output string

	Repository string
}

func (b ClientBuilder) Build() error {
	logs.Infow("begin to compile mojo package.", "pwd", b.PWD, "path", b.Path)

	options := make(core.Options)
	for _, pkg := range b.Package.GetAllPackages() {
		options[pkg.FullName] = getPackageImport(pkg)
	}
	for _, pkg := range b.Package.ResolvedDependencies {
		options[pkg.FullName] = getPackageImport(pkg)
	}

	cmp := gokit.NewCompiler()
	err := cmp.CompilePackage(compiler.WithGoPackageImports(context.Empty(), options), b.Package)

	if err != nil {
		logs.Errorw("failed to compile ncraft gokit", "package", b.Package.FullName, "error", err.Error())
		return err
	}

	services := cmp.Services
	conf := gokit.Options{
		Repository: b.Repository,
		Output:     b.Output,
	}

	for _, s := range services {
		err = gokit.GenerateClient(s, conf)
		if err != nil {
			logs.Errorw("generate ncraft gokit failed", "pwd", b.PWD, "path", b.Path, "package", b.Package.FullName, "error", err.Error())
			return err
		}
	}

	return nil
}
