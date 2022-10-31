package httptransport

import (
	"text/template"

	"github.com/mojo-lang/mojo/go/pkg/ncraft/data"
	_go "github.com/mojo-lang/mojo/go/pkg/ncraft/go"
	"github.com/mojo-lang/mojo/go/pkg/util"
)

func createParamUnmarshaler(parameter *data.HTTPParameter, funcMap template.FuncMap) (string, error) {
	initParent := `{{define "init_parent"}}{{if .GetEnclosingField}}
	{{template "init_parent" .GetEnclosingField}}
	if req.{{GoName .GetEnclosingField.FullName}} == nil {
		req.{{GoName .GetEnclosingField.FullName}} = &{{.GetEnclosingField.GetGoTypeName}}{}
	}{{end}}
{{end}}
{{template "init_parent" .}}
{{if .GetGoType.IsPointer}}{{ToLowerCamel .Field.FullName}}Initialized := false
if req.{{GoName .Field.FullName}} == nil {
    {{ToLowerCamel .Field.FullName}}Initialized = true
    req.{{GoName .Field.FullName}} = &{{.GetGoType.Name}}{}
}
{{end}}`

	query := `err = mjhttp.UnmarshalQueryParam(queryParams, {{if not .GetGoType.IsPointer}}&{{end}}req.{{GoName .Field.FullName}}, "{{.Name}}"{{if ne .Name .FullName}}, "{{.FullName}}"{{end}})
    {{if .GetGoType.IsPointer}}if err != nil {
        if core.IsNotFoundError(err) {
            if {{ToLowerCamel .Field.FullName}}Initialized {
                req.{{GoName .GetField.Name}} = nil
            }
        } else {
            return nil, nhttp.WrapError(err, 400, "cannot unmarshal the {{.Name}} {{if ne .Name .FullName}}or {{.FullName}}{{end}} query parameter")
        }
    }
    {{- else -}}
    if err != nil && !core.IsNotFoundError(err) {
        return nil, nhttp.WrapError(err, 400, "cannot unmarshal the {{.Name}} {{if ne .Name .FullName}}or {{.FullName}}{{end}} query parameter")
    }
    {{end}}
`
	path := `err = mjhttp.UnmarshalPathParam(pathParams, {{if not .GetGoType.IsPointer}}&{{end}}req.{{GoName .Field.FullName}}, "{{.Name}}"{{if ne .Name .FullName}}, "{{.FullName}}"{{end}})
    {{if .GetGoType.IsPointer}}
     if err != nil {
        if core.IsNotFoundError(err) {
            {{ToLowerCamel .Field.FullName}}Initialized {
                req.{{GoName .GetField.Name}} = nil
            }
        } else {
            return nil, nhttp.WrapError(err, 400, "cannot unmarshal the {{.Name}} {{if ne .Name .FullName}}or {{.FullName}}{{end}} query parameter")
        }
    }
    {{- else -}}
    if err != nil && !core.IsNotFoundError(err) {
        return nil, nhttp.WrapError(err, 400, "cannot unmarshal the {{.Name}} {{if ne .Name .FullName}}or {{.FullName}}{{end}} query parameter")
    }
    {{end}}
`
	tmpl := ""
	if parameter.Location == "path" {
		tmpl = initParent + path
	} else {
		tmpl = initParent + query
	}

	code, err := util.ApplyTemplate("Field Unmarshal", tmpl, parameter, funcMap)
	if err != nil {
		return "", err
	}

	bytes, err := _go.FormatCode(code)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
