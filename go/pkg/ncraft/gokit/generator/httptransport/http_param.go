package httptransport

import (
    "github.com/mojo-lang/mojo/go/pkg/mojo/util"
    "github.com/mojo-lang/mojo/go/pkg/ncraft/data"
    _go "github.com/mojo-lang/mojo/go/pkg/ncraft/go"
    "text/template"
)

func createParamUnmarshaler(parameter *data.HTTPParameter, funcMap template.FuncMap) (string, error) {
    initParent := `{{define "init_parent"}}{{if .GetEnclosingField}}
	{{template "init_parent" .GetEnclosingField}}
	if req.{{GoName .GetEnclosingField.FullName}} == nil {
		req.{{GoName .GetEnclosingField.FullName}} = &{{.GetEnclosingField.GetGoTypeName}}{}
	}{{end}}
{{end}}
{{if .GetGoType.IsPointer}}if req.{{GoName .GetField.Name}} == nil {
    req.{{GoName .GetField.Name}} = &{{.GetGoType.Name}}{}
}{{end}}
{{template "init_parent" .}}`

    query := `if err = mjhttp.UnmarshalQueryParam(queryParams, {{if not .GetGoType.IsPointer}}&{{end}}req.{{GoName .Field.FullName}}, "{{.Name}}"{{if ne .Name .FullName}}, "{{.FullName}}"{{end}}); err != nil {
    return nil, nhttp.WrapError(err, 400, "cannot unmarshal the {{.Name}} {{if ne .Name .FullName}}or {{.FullName}}{{end}} query parameter")
}
`
    path := `if err = mjhttp.UnmarshalPathParam(pathParams, {{if not .GetGoType.IsPointer}}&{{end}}req.{{GoName .Field.FullName}}, "{{.Name}}"{{if ne .Name .FullName}}, "{{.FullName}}"{{end}}); err != nil {
    return nil, nhttp.WrapError(err, 400, "cannot unmarshal the {{.Name}} {{if ne .Name .FullName}}or {{.FullName}}{{end}} path parameter")
}
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
