package gokit

import (
    "github.com/mojo-lang/core/go/pkg/logs"
    "github.com/mojo-lang/core/go/pkg/mojo/core"
    "github.com/mojo-lang/mojo/go/pkg/cmd/build/builder"
    "github.com/mojo-lang/mojo/go/pkg/mojo/context"
    "github.com/mojo-lang/mojo/go/pkg/ncraft/gokit"
    kitcompiler "github.com/mojo-lang/mojo/go/pkg/ncraft/gokit/compiler"
    "github.com/mojo-lang/mojo/go/pkg/ncraft/gokit/generator/render"
)

const HandlerInterface = `
type {{ToLowerCamel .Interface.Name}} struct{
	client  pb.{{.Interface.Name}}Server
}

// NewInterface returns a naive, stateless implementation of Interface.
func NewService() pb.{{.Interface.Name}}Server {
	var conf config.Config
	err := config.ScanKey("service", &conf)
	if err != nil {
		logs.Warnw("failed to get the server config", "error", err.Error())
	}
	logs.Debug("localhost http addr is ", conf.HttpAddr)
	return {{ToLowerCamel .Interface.Name}}{
		client: New{{ToLowerCamel .Interface.Name}}{
			client: pb.NewUserClient(conn),
		},
	}
}

`

const HandlerMethods = `
{{with $te := . }}
	{{range $i := $te.Interface.Methods}}
		// {{$i.Name}} implements Interface.
		func (s {{ToLowerCamel $te.Interface.Name}}) {{ToCamel $i.Name}}(ctx context.Context, in *{{PackageName $i.RequestType.Name}}.{{GoName $i.RequestType.Name}}) (*{{PackageName $i.ResponseType.Name}}.{{GoName $i.ResponseType.Name}}, error){
			logs.Debug("create {{ToCamel $i.Name}} request in sidecar")
			return s.client.{{ToCamel $i.Name}}(ctx, in)
		}
	{{end}}
{{- end}}
`

const HandlerProxyInterface = `
// NewInterface returns a naive, stateless implementation of Interface.
func NewService() pb.{{GoName .Interface.Name}}Server {
	return {{ToLowerCamel .Interface.Name}}{}
}

type {{ToLowerCamel .Interface.Name}} struct{
}

`

const HandlerProxyMethods = `
{{with $te := . }}
	{{range $i := $te.Interface.Methods}}
		// {{$i.Name}} implements Interface.
		func (s {{ToLowerCamel $te.Interface.Name}}) {{ToCamel $i.Name}}(ctx context.Context, in *{{PackageName $i.RequestType.Name}}.{{GoName $i.RequestType.Name}}) (*{{PackageName $i.ResponseType.Name}}.{{GoName $i.ResponseType.Name}}, error){
			var resp {{PackageName $i.ResponseType.Name}}.{{GoName $i.ResponseType.Name}}
			resp = {{PackageName $i.ResponseType.Name}}.{{GoName $i.ResponseType.Name}}{
				{{range $j := $i.ResponseType.Fields -}}
					// {{GoName $j.Name}}:
				{{end -}}
			}
			return &resp, nil
		}
	{{end}}
{{- end}}
`

type SidecarBuilder struct {
    builder.Builder
    Output     string
    Repository string
}

func (b SidecarBuilder) Build() error {
    logs.Infow("begin to compile mojo package.", "pwd", b.PWD, "path", b.Path)

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
    conf := render.Config{
        Repository: b.Repository,
        Output:     b.Output,
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
