package boot

import (
	"github.com/mojo-lang/mojo/go/pkg/ncraft/boot"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/mojo-lang/mojo/go/pkg/context"
	"github.com/mojo-lang/mojo/go/pkg/util"

	"github.com/mojo-lang/core/go/pkg/logs"
	"github.com/mojo-lang/core/go/pkg/mojo/core"
	"github.com/pkg/errors"

	"github.com/mojo-lang/mojo/go/pkg/cmd/build/builder"
)

type Builder struct {
	builder.Builder
	Output string

	Type       string
	Repository string
}

//func getPackageImport(pkg *lang.Package) string {
//	repository := pkg.Repository
//	if repository != nil {
//		goPackageFullName := lang.PackageNameToPath(pkg.FullName)
//		return fmt.Sprintf("%s%s/go/pkg/%s", repository.Authority.Host, repository.Path, goPackageFullName)
//	}
//	return ""
//}

func (b Builder) Build() error {
	logs.Infow("boot begin to compile mojo package.", "pwd", b.PWD, "path", b.Path)

	if len(b.Output) == 0 {
		if b.APIEnabled {
			b.Output = util.GetAbsolutePath(b.PWD, b.Path)
		} else {
			b.Output = util.GetAbsolutePath(b.PWD, path.Join(b.Path, "../"))
		}
	}

	setDefaultRepository := func(typ string) {
		if len(b.Repository) == 0 {
			if b.APIEnabled {
				b.Repository = path.Join(b.Package.Repository.FormatWithoutSchema(), typ)
			} else {
				b.Repository = b.Package.Repository.FormatWithoutSchema() + "-" + typ
			}
		}
	}

	cmp := boot.NewCompiler()
	options := make(core.Options)
	//for _, pkg := range b.Package.GetAllPackages() {
	//	options[pkg.FullName] = getPackageImport(pkg)
	//}
	//for _, pkg := range b.Package.GetAllDependentPackages() {
	//	options[pkg.FullName] = getPackageImport(pkg)
	//}

	err := cmp.CompilePackage(context.WithOptions(context.Empty(), options), b.Package)
	if err != nil {
		logs.Errorw("failed to compile ncraft boot", "package", b.Package.FullName, "error", err.Error())
		return err
	}

	services := cmp.Services
	setDefaultRepository("service-java")
	b.Output = path.Join(b.Output, path.Base(b.Repository))
	conf := boot.Options{
		Repository:    b.Repository,
		Output:        b.Output,
		MixedInAPI:    b.APIEnabled,
		PreviousFiles: make(map[string]io.Reader),
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
		err = boot.GenerateService(s, conf)
		if err != nil {
			logs.Errorw("generate ncraft boot failed", "pwd", b.PWD, "path", b.Path, "package", b.Package.FullName, "error", err.Error())
			return err
		}
	}

	return nil
}
