package gokit

import (
	"path/filepath"

	"github.com/mojo-lang/mojo/go/pkg/logs"
	"github.com/mojo-lang/mojo/go/pkg/mojo/core"

	"github.com/mojo-lang/mojo/go/pkg/compiler/cmd/build/builder"
	_go "github.com/mojo-lang/mojo/go/pkg/compiler/cmd/build/go"
	"github.com/mojo-lang/mojo/go/pkg/compiler/context"
	"github.com/mojo-lang/mojo/go/pkg/compiler/ncraft/compiler"
	"github.com/mojo-lang/mojo/go/pkg/compiler/ncraft/gokit"
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
	for _, pkg := range b.Package.GetAllDependentPackages() {
		options[pkg.FullName] = getPackageImport(pkg)
	}

	cmp := gokit.NewCompiler()
	err := cmp.CompilePackage(compiler.WithGoPackageImports(context.Empty(), options), b.Package)

	if err != nil {
		logs.Errorw("failed to compile ncraft gokit", "package", b.Package.FullName, "error", err.Error())
		return err
	}

	services := cmp.Services
	if len(services) == 0 {
		return nil
	}
	apiDir := filepath.Join(b.GetAbsolutePath(), "go")
	localAPI, err := filepath.Rel(b.Output, apiDir)
	if err != nil {
		return err
	}
	conf := gokit.Options{
		Repository:    b.Repository,
		Output:        b.Output,
		ApiRepository: b.Package.GoModName(),
		MixedInAPI:    b.APIEnabled || core.IsExist(filepath.Join(apiDir, "go.mod")),
		ApiLocalPath:  filepath.ToSlash(localAPI),
	}

	for _, s := range services {
		err = gokit.GenerateClient(s, conf)
		if err != nil {
			logs.Errorw("generate ncraft gokit failed", "pwd", b.PWD, "path", b.Path, "package", b.Package.FullName, "error", err.Error())
			return err
		}
	}

	return _go.GoModTidy(b.Output)
}
