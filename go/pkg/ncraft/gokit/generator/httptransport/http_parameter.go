package httptransport

import (
    "github.com/mojo-lang/mojo/go/pkg/mojo/util"
    "github.com/mojo-lang/mojo/go/pkg/ncraft/data"
    _go "github.com/mojo-lang/mojo/go/pkg/ncraft/go"
    "text/template"
)

// createQueryUnmarshaler returns the generated code for server-side unmarshaling
// of a query parameter into it's correct field on the request struct.
func createQueryUnmarshaler(parameter *data.HTTPParameter, funcMap template.FuncMap) (string, error) {
    queryParamLogic := `
{{if eq .Name .FullName}}
{{.Go.LocalName}}Strs, {{.Go.LocalName}}Ok := {{.Location}}Params["{{.Name}}"]
if {{.Go.LocalName}}Ok {
	parsedQueryParams["{{.Name}}"] = true
{{else}}
{{.Go.LocalName}}Strs, {{.Go.LocalName}}Ok := {{.Go.Location}}Params["{{.Name}}"]
if !{{.Go.LocalName}}Ok {
	if {{.Go.LocalName}}Strs, {{.Go.LocalName}}Ok = {{.Location}}Params["{{.FullName}}"]; {{.Go.LocalName}}Ok {
		parsedQueryParams["{{.FullName}}"] = true
	}
} else {
	parsedQueryParams["{{.Name}}"] = true
}
if {{.Go.LocalName}}Ok {
{{end}}{{if and .IsScalar (not .IsArray) }} 
{{.Go.LocalName}}Str := ""
if len({{.Go.LocalName}}Strs) > 0 {
	{{.Go.LocalName}}Str = {{.Go.LocalName}}Strs[0]
}
{{end}}
`

    pathParamLogic := `
{{.Go.LocalName}}Str := {{.Location}}Params["{{.Name}}"]
{{if not (eq .Name .FullName)}}
if len({{.Go.LocalName}}Str) == 0 {
	{{.Go.LocalName}}Str = {{.Location}}Params["{{.FullName}}"]
}
{{end}}
`

    genericLogic := `{{.Go.ConvertFunc}}{{if .Go.ConvertFuncNeedsErrorCheck}}
if err != nil {
	return nil, errors.Wrap(err, fmt.Sprintf("Error while extracting {{.Go.LocalName}} from {{.Location}}, {{.Location}}Params: %v", {{.Location}}Params))
}{{end}}
{{define "init_parent"}}{{if .GetEnclosingField}}
	{{template "init_parent" .GetEnclosingField}}
	if req.{{GoName .GetEnclosingField.FullName}} == nil {
		req.{{GoName .GetEnclosingField.FullName}} = &{{.GetEnclosingField.GetGoTypeName}}{}
	}
{{end}}
{{end}}
{{template "init_parent" .}}
req.{{GoName .Field.FullName}} = {{.Go.TypeConversion}}
`

    unwrappedMapLogic := `
for key, value := range {{.Location}}Params {
	if _, ok := parsedQueryParams[key]; ok {
		continue
	}
	
	decodeStr := ""
	if len(value) == 0 {
		decodeStr = value[0]
	} else {
		decodeStr = "[" + strings.Join(value, ",") + "]"
	}

	v := &core.Value{}
	err = jsoniter.ConfigFastest.UnmarshalFromString(decodeStr, v)
	if err != nil {
		return nil, errors.Wrapf(err, "couldn't decode {{.Name}} field from %v", decodeStr)
	}

	req.{{.Field.FullName}}[key] = v
}
`

    mergedLogic := queryParamLogic + genericLogic + "}"
    if parameter.Location == "path" {
        mergedLogic = pathParamLogic + genericLogic
    } else if parameter.Location == "query" && parameter.Field.Type.IsMap && parameter.Field.Exploded {
        mergedLogic = unwrappedMapLogic
    }

    code, err := util.ApplyTemplate("FieldEncodeLogic", mergedLogic, parameter, funcMap)
    if err != nil {
        return "", err
    }

    bytes, err := _go.FormatCode(code)
    if err != nil {
        return "", err
    }

    return string(bytes), nil
}
