// Package httptransport provides functions and templates helpers for templating
// the http-transport of a go-kit based service.
package httptransport

import (
    "bytes"
    "fmt"
    "github.com/mojo-lang/mojo/go/pkg/mojo/util"
    "github.com/mojo-lang/mojo/go/pkg/ncraft/data"
    _go "github.com/mojo-lang/mojo/go/pkg/ncraft/go"
    "github.com/mojo-lang/mojo/go/pkg/ncraft/gokit/generator/httptransport/templates"
    "io"
    "io/ioutil"
    "text/template"
)

const ServerHttpTransportPath = "pkg/NAME-service/svc/transport_http.go.tmpl"

type ServerHttpTransport struct {
}

func NewServerHttpTransport(ds *data.Service) (*ServerHttpTransport, error) {
    for _, method := range ds.Interface.Methods {
        for _, binding := range method.Bindings {
            for _, param := range binding.Parameters {
                decoder, check := createDecodeConvertFunc(param, ds.FuncMap)
                param.Go.ConvertFunc = decoder
                param.Go.ConvertFuncNeedsErrorCheck = check

                param.Go.TypeConversion = createDecodeTypeConversion(param)

                str, err := createQueryUnmarshaler(param, ds.FuncMap)
                if err != nil {
                    return nil, err
                }
                param.Go.QueryUnmarshaler = str
            }

            if serverDecode, err := createServerDecode(binding, ds.FuncMap); err != nil {
                return nil, err
            } else {
                binding.Extensions["ServerDecode"] = string(serverDecode)
            }
        }
    }

    return &ServerHttpTransport{}, nil
}

func (h *ServerHttpTransport) Render(tmpl string, ds *data.Service) (io.Reader, error) {
    code, err := util.ApplyTemplate("ServerTemplate", templates.ServerTemplate, ds, ds.FuncMap)
    if err != nil {
        return nil, err
    }
    if bs, err := _go.FormatCode(code); err != nil {
        return nil, err
    } else {
        return bytes.NewBuffer(bs), nil
    }
}

// createServerDecode returns the generated code for the server-side decoding of
// a http request into its request struct.
func createServerDecode(binding *data.HTTPBinding, funcMap template.FuncMap) ([]byte, error) {
    code, err := util.ApplyTemplate("ServerDecodeTemplate", templates.ServerDecodeTemplate, binding, funcMap)
    if err != nil {
        return nil, err
    }
    return _go.FormatCode(code)
}

// createDecodeConvertFunc creates a go string representing the function to
// convert the string form of the field to it's correct go type.
func createDecodeConvertFunc(f *data.HTTPParameter, funcMap template.FuncMap) (string, bool) {
    needsErrorCheck := true

    fType := ""
    switch f.GetGoTypeName() {
    case "uint32":
        fType = "%s, err := strconv.ParseUint(%s, 10, 32)"
    case "uint64":
        fType = "%s, err := strconv.ParseUint(%s, 10, 64)"
    case "int32":
        fType = "%s, err := strconv.ParseInt(%s, 10, 32)"
    case "int64":
        fType = "%s, err := strconv.ParseInt(%s, 10, 64)"
    case "bool":
        fType = "%s, err := strconv.ParseBool(%s)"
    case "float32":
        fType = "%s, err := strconv.ParseFloat(%s, 32)"
    case "float64":
        fType = "%s, err := strconv.ParseFloat(%s, 64)"
    case "string":
        fType = "%s := %s"
        needsErrorCheck = false
    }

    if f.IsArray() {
    }

    if f.IsEnum() && !f.IsArray() {
        fType = "%s, err := strconv.ParseInt(%s, 10, 32)"
        return fmt.Sprintf(fType, f.Go.LocalName, f.Go.LocalName+"Str"), true
    }

    // Use json unmarshalling for any custom/repeated messages
    if !f.GetGoType().IsBaseType || f.IsArray() {
        // Args representing single custom message types are represented as
        // pointers. To do a bare assignment to a pointer, our rvalue must be a
        // pointer as well. So we special case args of a single custom message
        // type so that the variable LocalName is declared as a pointer.
        singleCustomTypeUnmarshalTmpl := `
var {{.Go.LocalName}} *{{.GetGoTypeName}}
{{.Go.LocalName}}Str := ""
if len({{.Go.LocalName}}Strs) > 0 {
	{{.Go.LocalName}}Str = {{.Go.LocalName}}Strs[0]
}
{{.Go.LocalName}} = &{{.GetGoTypeName}}{}
err = jsoniter.ConfigFastest.Unmarshal([]byte({{.Go.LocalName}}Str), {{.Go.LocalName}})
if err != nil {
	return nil, errors.Wrapf(err, "couldn't decode {{.Go.LocalName}} from %v", {{.Go.LocalName}}Str)
}`

        mapTypeUmarshalTmpl := `
{{.Go.LocalName}}Str := ""
if len({{.Go.LocalName}}Strs) > 0 {
	{{.Go.LocalName}}Str = {{.Go.LocalName}}Strs[0]
}
{{.Go.LocalName}} := make({{.GetGoTypeName}})
err = jsoniter.ConfigFastest.Unmarshal([]byte({{.Go.LocalName}}Str), {{.Go.LocalName}})
if err != nil {
	return nil, errors.Wrapf(err, "couldn't decode {{.Go.LocalName}} from %v", {{.Go.LocalName}}Str)
}
`
        // All repeated args of any type are represented as slices, and bare
        // assignments to a slice accept a slice as the rvalue. As a result,
        // LocalName will be declared as a slice, and json.Unmarshal handles
        // everything else for us. Addititionally, if a type is a Base type and
        // is repeated, we first attempt to unmarshal the string we're
        // provided, and if that fails, we try to unmarshal the string
        // surrounded by square brackets. If THAT fails, then the string does
        // not represent a valid JSON string and an error is returned.
        //		repeatedUnmarshalTmpl := `
        //var {{.Go.LocalName}} {{.GetGoTypeName}}
        //{{- if and (and .IsBaseType .Repeated) (not (Contains .GoType "[]byte"))}}
        //err = json.Unmarshal([]byte({{.Go.LocalName}}Str), &{{.Go.LocalName}})
        //if err != nil {
        //	{{.Go.LocalName}}Str = "[" + {{.Go.LocalName}}Str + "]"
        //}
        //{{- end}}
        //err = json.Unmarshal([]byte({{.Go.LocalName}}Str), &{{.Go.LocalName}})`

        //		errorCheckingTmpl := `
        //if err != nil {
        //	return nil, errors.Wrapf(err, "couldn't decode {{.Go.LocalName}} from %v", {{.Go.LocalName}}Str)
        //}`
        repeatedTmpl := `{{- if eq .Location "query"}}
if len({{.Go.LocalName}}Strs) == 1 {
	{{.Go.LocalName}}Strs = ParseArrayStr({{.Go.LocalName}}Strs[0], ",")
}{{- else -}}
{{.Go.LocalName}}Strs := ParseArrayStr({{.Go.LocalName}}Str, ",")
{{- end}}
var {{.Go.LocalName}} {{.GetGoTypeName}}
for _, str := range {{.Go.LocalName}}Strs {
	{{if .GetElementGoType.IsBaseType}}var v {{GoFieldArrayElementType .GetElementGoTypeName}}{{else}}v := &{{.GetElementGoTypeName}}{}{{end}}
	err = jsoniter.ConfigFastest.UnmarshalFromString({{if GoIsArrayElementStringType .GetGoTypeName}}"\"" + str + "\""{{else}}str{{end}}, {{if .GetElementGoType.IsBaseType}}&v{{else}}v{{end}})
	if err != nil {
		return nil, errors.Wrapf(err, "couldn't decode {{.Name}} field from %v", str)
	}
	{{.Go.LocalName}} = append({{.Go.LocalName}}, v)
}`

        var jsonConvTmpl string
        if f.IsMap() {
            jsonConvTmpl = mapTypeUmarshalTmpl
        } else if f.Field.Type.IsArray {
            jsonConvTmpl = repeatedTmpl
        } else {
            jsonConvTmpl = singleCustomTypeUnmarshalTmpl
        }
        code, err := util.ApplyTemplate("UnmarshalNonBaseType", jsonConvTmpl, f, funcMap)
        if err != nil {
            panic(fmt.Sprintf("Couldn't apply templates: %v", err))
        }
        bs, err := ioutil.ReadAll(code)

        return string(bs), false
    }
    return fmt.Sprintf(fType, f.Go.LocalName, f.Go.LocalName+"Str"), needsErrorCheck
}

// createDecodeTypeConversion creates a go string that converts a 64 bit type
// to a 32 bit type as strconv.ParseInt, ParseUInt, and ParseFloat always
// return the 64 bit type. If the type is not a 64 bit integer type or is
// repeated, then returns the LocalName of that Field.
func createDecodeTypeConversion(f *data.HTTPParameter) string {
    if f.Field.Type.IsArray {
        // Equivalent of the 'default' case below, but taken early for repeated
        // types.
        return f.Go.LocalName
    }
    fType := ""
    switch f.GetGoTypeName() {
    case "uint32", "int32", "float32":
        fType = f.Field.Type.Go.Name + "(%s)"
    default:
        fType = "%s"
    }
    if f.IsEnum() {
        fType = f.GetGoTypeName() + "(%s)"
    }
    return fmt.Sprintf(fType, f.Go.LocalName)
}
