package templates

//const HandlerMethods = `
//{{ with $te := .}}
//		{{range $i := .Methods}}
//		// {{.Name}} implements Interface.
//		func (s *{{ToLowerCamel $te.InterfaceName}}Interface) {{ToCamel .Name}}(ctx context.Context, in *{{PackageName $i.RequestType.Name}}.{{GoName .RequestType.Name}}) (*{{PackageName $i.ResponseType.Name}}.{{GoName .ResponseType.Name}}, error){
//			var resp {{PackageName $i.ResponseType.Name}}.{{GoName .ResponseType.Name}}
//			resp = {{PackageName $i.ResponseType.Name}}.{{GoName .ResponseType.Name}}{
//				{{range $j := $i.ResponseType.Message.Fields -}}
//					// {{GoName $j.Name}}:
//				{{end -}}
//			}
//			return &resp, nil
//		}
//		{{end}}
//{{- end}}
//`

const Handlers = `
package handlers

import (
	"context"
	
	{{range $i := .ExternalMessageImports}}
	"{{$i}}"
	{{- end}}
	pb "{{.ApiImportPath -}}"
)

var ({{range $name := .ExternalStructs}}
	_ = {{PackageName $name}}.{{$name}}{}
{{- end}}{{range $name := .ExternalEnums}}
	_ = {{PackageName $name}}.{{$name}}(0)
{{- end}})

`

const HandlerInterface = `
type {{ToLowerCamel .Interface.Name}} struct{
}

// NewInterface returns a naive, stateless implementation of Interface.
func NewService() pb.{{GoName .Interface.Name}}Server {
	return {{ToLowerCamel .Interface.Name}}{}
}

`

const HandlerMethods = `
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