package gokit

import (
    "fmt"
    "io"
    "os"
    path2 "path"
    "path/filepath"
    "strings"

    "github.com/mojo-lang/core/go/pkg/logs"
    "github.com/mojo-lang/core/go/pkg/mojo/core"
    "github.com/mojo-lang/lang/go/pkg/mojo/lang"
    "github.com/mojo-lang/mojo/go/pkg/cmd/build/builder"
    _go "github.com/mojo-lang/mojo/go/pkg/cmd/build/go"
    "github.com/mojo-lang/mojo/go/pkg/mojo/context"
    "github.com/mojo-lang/mojo/go/pkg/mojo/util"
    "github.com/mojo-lang/mojo/go/pkg/ncraft/gokit"
    kitcompiler "github.com/mojo-lang/mojo/go/pkg/ncraft/gokit/compiler"
    "github.com/mojo-lang/mojo/go/pkg/ncraft/gokit/generator/render"
    "github.com/pkg/errors"
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
        goPackageFullName := lang.PackageNameToPath(pkg.FullName)
        return fmt.Sprintf("%s%s/go/pkg/%s", repository.Authority.Host, repository.Path, goPackageFullName)
    }
    return ""
}

func (b Builder) Build() error {
    logs.Infow("gokit begin to compile mojo package.", "pwd", b.PWD, "path", b.Path)

    if len(b.Output) == 0 {
        if b.APIEnabled {
            b.Output = util.GetAbsolutePath(b.PWD, b.Path)
        } else {
            b.Output = util.GetAbsolutePath(b.PWD, path2.Join(b.Path, "../"))
        }
    }

    setDefaultRepository := func(ncraftType string) {
        if len(b.Repository) == 0 {
            if b.APIEnabled {
                b.Repository = path2.Join(b.Package.Repository.FormatWithoutSchema(), ncraftType)
            } else {
                b.Repository = b.Package.Repository.FormatWithoutSchema() + "-" + ncraftType
            }
        }
    }

    if b.Type == "client" {
        setDefaultRepository("client")
        b.Output = path2.Join(b.Output, path2.Base(b.Repository))
        cb := &ClientBuilder{
            Builder:    b.Builder,
            Output:     b.Output,
            Repository: b.Repository,
        }
        return cb.Build()
    } else if b.Type == "sidecar" {
        setDefaultRepository("sidecar")
        b.Output = path2.Join(b.Output, path2.Base(b.Repository))
        sb := &SidecarBuilder{
            Builder:    b.Builder,
            Output:     b.Output,
            Repository: b.Repository,
        }
        return sb.Build()
    }

    compiler := gokit.NewCompiler()
    pkgs := b.Package.GetAllPackages()
    options := make(core.Options)
    for _, pkg := range pkgs {
        options[pkg.FullName] = getPackageImport(pkg)
    }
    for _, pkg := range b.Package.ResolvedDependencies {
        options[pkg.FullName] = getPackageImport(pkg)
    }

    err := compiler.Compile(kitcompiler.WithGoPackageImports(context.Empty(), options), pkgs)
    if err != nil {
        logs.Errorw("failed to compile ncraft gokit", "package", b.Package.FullName, "error", err.Error())
        return err
    }

    services := compiler.Services
    setDefaultRepository("service-go")
    b.Output = path2.Join(b.Output, path2.Base(b.Repository))
    conf := render.Config{
        Repository:    b.Repository,
        ApiRepository: path2.Join(b.Package.Repository.FormatWithoutSchema(), "go"),
        Output:        b.Output,
        MixedInAPI:    b.APIEnabled,
        PreviousFiles: make(map[string]io.Reader),
    }

    prefixPath := b.Output
    if !strings.HasSuffix(prefixPath, "/") {
        prefixPath += "/"
    }
    if util.IsExist(b.Output) {
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

    _go.GoModTidy(b.Output)
    return nil
}
