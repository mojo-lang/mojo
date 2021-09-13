package gokit

import (
	"fmt"
	"github.com/mojo-lang/core/go/pkg/logs"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/mojo-lang/mojo/go/pkg/ncraft/gokit"
	"github.com/mojo-lang/mojo/go/pkg/ncraft/gokit/generator/render"
	"github.com/mojo-lang/mojo/go/pkg/wand/builder"
	"github.com/pkg/errors"
	"io"
	"os"
	path2 "path"
	"path/filepath"
	"strings"
)

type Builder struct {
	builder.Builder
	Output string

	Type       string
	Repository string
}

func getPackageImport(pkg *lang.Package) string {
	repository := pkg.Repository
	if repository != nil {
		goPackageFullName := strings.ReplaceAll(pkg.FullName, ".", "/")
		return fmt.Sprintf("%s%s/go/pkg/%s", repository.Authority.Host, repository.Path, goPackageFullName)
	}
	return ""
}

func (b Builder) Build() error {
	logs.Infow("begin to compile mojo package.", "pwd", b.PWD, "path", b.Path)

	if b.Type == "client" {
		cb := &ClientBuilder{
			Builder:    b.Builder,
			Output:     b.Output,
			Repository: b.Repository,
		}
		return cb.Build()
	} else if b.Type == "sidecar" {
		sb := &SidecarBuilder{
			Builder:    b.Builder,
			Output:     b.Output,
			Repository: b.Repository,
		}
		return sb.Build()
	}

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
		Repository:    b.Repository,
		Output:        b.Output,
		PreviousFiles: make(map[string]io.Reader),
	}

	root := path2.Join(b.Output, path2.Base(b.Repository))
	prefixPath := b.Output
	if !strings.HasSuffix(prefixPath, "/") {
		prefixPath += "/"
	}
	err = filepath.Walk(root, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}

		if f.IsDir() {
			return nil
		}

		reader, err := os.Open(path)
		name := strings.TrimPrefix(path, prefixPath)
		conf.PreviousFiles[name] = reader
		return nil
	})
	if err != nil {
		return errors.Wrap(err, "filepath.Walk() failed")
	}

	for _, s := range services {
		err = gokit.GenerateService(s, conf)
		if err != nil {
			logs.Errorw("generate ncraft gokit failed", "pwd", b.PWD, "path", b.Path, "package", b.Package.FullName, "error", err.Error())
			return err
		}
	}

	return nil
}
