package templates

// const HandlerMethods = `
// {{ with $te := .}}
//		{{range $i := .Methods}}
//		// {{.Name}} implements Interface.
//		func (s *{{ToLowerCamel $te.InterfaceName}}Interface) {{ToCamel .Name}}(ctx context.Context, in *{{GoPackageName $i.Request.Name}}.{{GoName .Request.Name}}) (*{{GoPackageName $i.Response.Name}}.{{GoName .Response.Name}}, error){
//			var resp {{GoPackageName $i.Response.Name}}.{{GoName .Response.Name}}
//			resp = {{GoPackageName $i.Response.Name}}.{{GoName .Response.Name}}{
//				{{range $j := $i.Response.Message.Fields -}}
//					// {{GoName $j.Name}}:
//				{{end -}}
//			}
//			return &resp, nil
//		}
//		{{end}}
// {{- end}}
// `

const Handlers = `
package handlers

import (
	"context"
	
	{{range $i := .Go.ImportedTypePaths}}
	"{{$i}}"
	{{- end}}

	// this service api
	pb "{{.Go.ApiImportPath -}}"
)
{{if .HasImported}}
var ({{range $msg := .ImportedMessages}}
	_ = {{$msg.Go.PackageName}}.{{$msg.Name}}{}
{{- end}}{{range $enum := .ImportedEnums}}
	_ = {{$enum.Go.PackageName}}.{{$enum.Name}}(0)
{{- end}}){{end}}

`

const HandlerInterface = `
type {{ToLowerCamel .Interface.ServerName}} struct{
	pb.Unimplemented{{GoName .Interface.ServerName}}
}

// NewService returns a naive, stateless implementation of Interface.
func NewService() pb.{{GoName .Interface.ServerName}} {
	return {{ToLowerCamel .Interface.ServerName}}{}
}

`

const HandlerMethods = `
{{with $te := . }}
	{{range $i := $te.Interface.Methods}}
		// {{GoName $i.Name}} implements Interface.
		func (s {{ToLowerCamel $te.Interface.ServerName}}) {{GoName $i.Name}}(ctx context.Context, in *{{GoPackageName $i.Request.Name}}.{{GoName $i.Request.Name}}) (*{{GoPackageName $i.Response.Name}}.{{GoName $i.Response.Name}}, error){
			resp := &{{GoPackageName $i.Response.Name}}.{{GoName $i.Response.Name}}{
				{{range $j := $i.Response.Fields -}}
					// {{GoName $j.Name}}:
				{{end -}}
			}
			return resp, nil
		}
	{{end}}
{{- end}}
`
