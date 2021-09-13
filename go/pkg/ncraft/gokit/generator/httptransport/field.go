package httptransport

// Field contains the distillation of information within an svcdef.Field that's
// useful for templating http transport.
type Field struct {
	Name      string
	ParamName string

	// The name of this field, but passed through the CamelCase function.
	// Removes underscores, adds camelcase; "client_id" becomes "ClientId".
	//CamelName string

	FullName      string
	FullParamName string

	Unwrapped bool
	Enclosing    *Field

	// The name of this field, but run through the camelcase function and with
	// the first letter lowercased. "example_name" becomes "exampleName".
	// LowCamelName is how the names of fields should appear when marshaled to
	// JSON, according to the gRPC language guide.
	//LowCamelName string

	// The go-compatible name for this variable, for use in auto generated go
	// code.
	LocalName string

	// The location within the the http request that this field is to be found.
	Location string

	// The type within the Go language that's used to represent the original
	// field that this field refers to.
	GoType string

	// In order to support common ways of serializing simple parameters,
	// a set of style values are defined.
	Style string

	// The string form of the function to be used to convert the incoming
	// string msg from a string into it's intended type.
	ConvertFunc string

	// Used in determining if a convert func will need error checking logic
	ConvertFuncNeedsErrorCheck bool

	// The string form of a type cast from 64 to 32bit if the GoType is 32bit
	// as the ConvertFunc will always use return a 64bit type
	TypeConversion string

	// Indicates whether this field represents a basic protobuf type such as
	// one of the ints, floats, strings, bools, etc. Since we can only create
	// automatic marshaling of base types, if this is false a warning is given
	// to the user.
	IsBaseType bool

	// Protobuf Enums need to be handled uniquely when parsing queryparameters
	IsEnum bool

	IsMap bool

	// Repeated is true if this arg corresponds to a protobuf field which is
	// given an identifier of "repeated", meaning it will represented in Go as
	// a slice of it's type.
	Repeated bool
}

// GenQueryUnmarshaler returns the generated code for server-side unmarshaling
// of a query parameter into it's correct field on the request struct.
func (f *Field) GenQueryUnmarshaler() (string, error) {
	queryParamLogic := `
{{if eq .ParamName .FullParamName}}
{{.LocalName}}Strs, {{.LocalName}}Ok := {{.Location}}Params["{{.ParamName}}"]
if {{.LocalName}}Ok {
	parsedQueryParams["{{.ParamName}}"] = true
{{else}}
{{.LocalName}}Strs, {{.LocalName}}Ok := {{.Location}}Params["{{.ParamName}}"]
if !{{.LocalName}}Ok {
	if {{.LocalName}}Strs, {{.LocalName}}Ok = {{.Location}}Params["{{.FullParamName}}"]; {{.LocalName}}Ok {
		parsedQueryParams["{{.FullParamName}}"] = true
	}
} else {
	parsedQueryParams["{{.ParamName}}"] = true
}
if {{.LocalName}}Ok {
{{end}}{{if and .IsBaseType (not .Repeated) }} 
{{.LocalName}}Str := ""
if len({{.LocalName}}Strs) > 0 {
	{{.LocalName}}Str = {{.LocalName}}Strs[0]
}
{{end}}
`

	pathParamLogic := `
{{.LocalName}}Str := {{.Location}}Params["{{.ParamName}}"]
{{if not (eq .ParamName .FullParamName)}}
if len({{.LocalName}}Str) == 0 {
	{{.LocalName}}Str = {{.Location}}Params["{{.FullParamName}}"]
}{{end}}`

	genericLogic := `{{.ConvertFunc}}{{if .ConvertFuncNeedsErrorCheck}}
if err != nil {
	return nil, errors.Wrap(err, fmt.Sprintf("Error while extracting {{.LocalName}} from {{.Location}}, {{.Location}}Params: %v", {{.Location}}Params))
}{{end -}}
{{define "init_parent"}}{{if .Enclosing}}
	{{template "init_parent" .Enclosing}}
	if req.{{.Enclosing.FullName}} == nil {
		req.{{.Enclosing.FullName}} = &{{.Enclosing.GoType}}{}
	}
{{- end}}
{{- end}}
{{- template "init_parent" .}}
req.{{.FullName}} = {{.TypeConversion}}
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

	req.{{.FullName}}[key] = v
}
`

	mergedLogic := queryParamLogic + genericLogic + "}"
	if f.Location == "path" {
		mergedLogic = pathParamLogic + genericLogic
	} else if f.Location == "query" && f.IsMap && f.Unwrapped {
		mergedLogic = unwrappedMapLogic
	}

	code, err := ApplyTemplate("FieldEncodeLogic", mergedLogic, f, TemplateFunctions)
	if err != nil {
		return "", err
	}
	code = FormatCode(code)
	return code, nil
}
