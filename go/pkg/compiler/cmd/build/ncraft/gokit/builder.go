package gokit

import (
	"fmt"
	"io"
	"os"
	path2 "path"
	"path/filepath"
	"strings"

	"github.com/mojo-lang/mojo/go/pkg/compiler/context"
	"github.com/mojo-lang/mojo/go/pkg/compiler/util"

	"github.com/mojo-lang/mojo/go/pkg/logs"
	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/mojo-lang/mojo/go/pkg/mojo/lang"
	"github.com/pkg/errors"

	"github.com/mojo-lang/mojo/go/pkg/compiler/cmd/build/builder"
	_go "github.com/mojo-lang/mojo/go/pkg/compiler/cmd/build/go"
	"github.com/mojo-lang/mojo/go/pkg/compiler/ncraft/compiler"
	"github.com/mojo-lang/mojo/go/pkg/compiler/ncraft/data"
	"github.com/mojo-lang/mojo/go/pkg/compiler/ncraft/gokit"
)

type Builder struct {
	builder.Builder
	Output string

	Type       string
	Repository string
}

func getPackageImport(pkg *lang.Package) string {
	return pkg.GetGoPackageImport()
}

func (b Builder) Build() error {
	switch b.Type {
	case "", "service", "client":
	default:
		return fmt.Errorf("unsupported ncraft build type %q; use service or client", b.Type)
	}

	logs.Infow("gokit begin to compile mojo package.", "pwd", b.PWD, "path", b.Path)

	if len(b.Output) == 0 {
		b.Output = b.GetAbsolutePath()
	} else {
		b.Output = util.GetAbsolutePath(b.PWD, b.Output)
	}

	setDefaultRepository := func(ncraftType string) {
		if len(b.Repository) == 0 {
			if b.APIEnabled || ncraftType == "service-go" || ncraftType == "client-go" {
				b.Repository = path2.Join(b.Package.Repository.FormatWithoutSchema(), ncraftType)
			} else {
				b.Repository = b.Package.Repository.FormatWithoutSchema() + "-" + ncraftType
			}
		}
	}

	if b.Type == "client" {
		setDefaultRepository("client-go")
		b.Output = path2.Join(b.Output, "client-go")
		cb := &ClientBuilder{
			Builder:    b.Builder,
			Output:     b.Output,
			Repository: b.Repository,
		}
		return cb.Build()
	}

	cmp := gokit.NewCompiler()
	options := make(core.Options)
	for _, pkg := range b.Package.GetAllPackages() {
		options[pkg.FullName] = getPackageImport(pkg)
	}
	for _, pkg := range b.Package.GetAllDependentPackages() {
		options[pkg.FullName] = getPackageImport(pkg)
	}

	err := cmp.CompilePackage(compiler.WithGoPackageImports(context.Empty(), options), b.Package)
	if err != nil {
		logs.Errorw("failed to compile ncraft gokit", "package", b.Package.FullName, "error", err.Error())
		return err
	}

	const ServiceRepository = "service-go"
	services := cmp.Services
	setDefaultRepository(ServiceRepository)
	b.Output = path2.Join(b.Output, path2.Base(b.Repository))
	conf := gokit.Options{
		Repository:    b.Repository,
		ApiRepository: path2.Join(b.Package.Repository.FormatWithoutSchema(), "go"),
		Output:        b.Output,
		MixedInAPI:    b.APIEnabled || core.IsExist(path2.Join(b.GetAbsolutePath(), "go/go.mod")),
		PreviousFiles: make(map[string]io.Reader),
	}
	if len(services) == 0 && len(cmp.Entities) == 0 {
		return nil
	}

	prefixPath := b.Output
	if !strings.HasSuffix(prefixPath, "/") {
		prefixPath += "/"
	}
	if core.IsExist(b.Output) {
		err = filepath.Walk(b.Output, func(path string, f os.FileInfo, err error) error {
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
	}

	for _, s := range services {
		err = gokit.GenerateService(s, conf)
		if err != nil {
			logs.Errorw("generate ncraft gokit failed", "pwd", b.PWD, "path", b.Path, "package", b.Package.FullName, "error", err.Error())
			return err
		}
	}
	if len(services) == 0 {
		if err := gokit.GenerateModelPackage(&data.Service{
			PackageName: b.Package.Name, Entities: cmp.Entities, Go: &data.GoService{},
		}, conf); err != nil {
			return err
		}
	}
	err = _go.GoModTidy(b.Output)

	//{
	//	b.Output = strings.TrimSuffix(b.Output, ServiceRepository)
	//	b.Repository = strings.TrimSuffix(b.Repository, ServiceRepository)
	//	b.Type = "client"
	//	cb := &ClientBuilder{
	//		Builder:    b.Builder,
	//		Output:     path2.Join(b.Output, "go/pkg"),
	//		Repository: b.Repository + "go",
	//	}
	//	return cb.Build()
	//}
	return err
}
