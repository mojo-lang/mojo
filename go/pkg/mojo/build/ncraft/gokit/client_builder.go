package gokit

import (
	"github.com/mojo-lang/core/go/pkg/logs"
	"github.com/mojo-lang/mojo/go/pkg/mojo/build/builder"
	"github.com/mojo-lang/mojo/go/pkg/ncraft/gokit"
	"github.com/mojo-lang/mojo/go/pkg/ncraft/gokit/generator/render"
)

type ClientBuilder struct {
	builder.Builder
	Output string

	Repository string
}

func (b ClientBuilder) Build() error {
	logs.Infow("begin to compile mojo package.", "pwd", b.PWD, "path", b.Path)

	compiler := gokit.NewCompiler()
	pkgs := b.Package.GetAllPackage()
	for _, pkg := range pkgs {
		compiler.Context.GoPackageImports[pkg.FullName] = getPackageImport(pkg)
	}
	for _, pkg := range b.Package.ResolvedDependencies {
		compiler.Context.GoPackageImports[pkg.FullName] = getPackageImport(pkg)
	}

	err := compiler.Compile(pkgs)
	if err != nil {
		logs.Errorw("failed to compile ncraft gokit", "package", b.Package.FullName, "error", err.Error())
		return err
	}

	services := compiler.Services
	conf := render.Config{
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
