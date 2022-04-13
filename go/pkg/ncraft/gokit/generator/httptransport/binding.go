package httptransport

import (
    "fmt"
    "strconv"
    "strings"

    "github.com/mojo-lang/core/go/pkg/mojo/core/strcase"
    "github.com/mojo-lang/mojo/go/pkg/ncraft/gokit/compiler"
    "github.com/mojo-lang/mojo/go/pkg/ncraft/gokit/generator/httptransport/templates"
    "github.com/mojo-lang/mojo/go/pkg/ncraft/gokit/generator/types"
)

// Binding contains the distillation of information within an
// svcdef.HTTPBinding that's useful for templating http transport.
type Binding struct {
    // Label is the name of this method, plus the english word for the index of
    // this binding in this methods slice of bindings. So if this binding where
    // the first binding in the slice of bindings for the method "Sum", the
    // label for this binding would be "SumZero". If it where the third
    // binding, it would be named "SumTwo".
    Label string

    // PathTemplate is the full path template as it appeared in the http
    // annotation which this binding refers to.
    PathTemplate string

    // BasePath is the longest static portion of the full PathTemplate, and is
    // given to the net/http mux as the path for the route for this binding.
    BasePath string

    Verb string

    // a unwrapped map field will always be the last one if exist
    Fields []*Field

    BodyField *Field

    // A pointer back to the parent method of this binding. Used within some
    // binding methods
    Parent *Method
}

func httpCanCarryBody(method string) bool {
    switch method {
    case "post", "put", "patch":
        return true
    }
    return false
}

// NewBinding creates a Binding struct based on a svcdef.HTTPBinding. Because
// NewBinding requires access to some of it's parent method's fields, instead
// of passing a svcdef.HttpBinding directly, you instead pass a
// svcdef.InterfaceMethod and the index of the HTTPBinding within that methods
// "HTTPBinding" slice.
func NewBinding(i int, meth *types.InterfaceMethod) *Binding {
    binding := meth.Bindings[i]
    nBinding := Binding{
        Label:        strcase.ToCamel(meth.Name) + EnglishNumber(i),
        PathTemplate: binding.Path,
        BasePath:     basePath(binding.Path),
        Verb:         binding.Verb,
    }

    sortParams := func() []*types.HTTPParameter {
        var params []*types.HTTPParameter
        var unwrappedMap *types.HTTPParameter
        for _, param := range binding.Params {
            if param.Field.IsUnwrappedMap() {
                unwrappedMap = param
            } else {
                params = append(params, param)
            }
        }
        if unwrappedMap != nil {
            params = append(params, unwrappedMap)
        }
        return params
    }

    for _, param := range sortParams() {
        // The 'Field' attr of each HTTPParameter always point to it's bound
        // Methods RequestType

        newField := NewField(param, meth)
        if newField != nil {
            nBinding.Fields = append(nBinding.Fields, newField)
            if newField.Location == "body" {
                nBinding.BodyField = newField
            }
        }
    }

    if nBinding.BodyField == nil && httpCanCarryBody(nBinding.Verb) {
        for _, field := range nBinding.Fields {
            if field.Location == "query" && len(compiler.GetPackageName(field.GoType)) > 0 {
                nBinding.BodyField = field
                break
            }
        }
    }

    return &nBinding
}

func toCamel(fullName string) string {
    if len(fullName) == 0 {
        return ""
    }

    segments := strings.Split(fullName, ".")
    if len(segments) == 1 {
        return strcase.ToCamel(segments[0])
    }
    for i := 0; i < len(segments); i++ {
        segments[i] = strcase.ToCamel(segments[i])
    }
    return strings.Join(segments, ".")
}

func goTypeName(field *types.FieldType) string {
    if field.Map != nil {
        return goTypeName(field.Map.ValueType)
    } else {
        name := field.Name
        segments := strings.Split(name, ".")
        if len(segments) > 0 {
            return segments[len(segments)-1]
        }
    }
    return ""
}

func NewField(param *types.HTTPParameter, meth *types.InterfaceMethod) *Field {
    field := param.Field
    newField := Field{
        Name:          strcase.ToCamel(field.Name),
        ParamName:     strcase.ToSnake(field.Name),
        FullName:      toCamel(field.FullName),
        FullParamName: strcase.ToSnake(strings.ReplaceAll(field.FullName, ".", "_")),
        Location:      param.Location,
        Repeated:      field.Type.ArrayType,
        IsMap:         field.Type.Map != nil,
        Unwrapped:     field.Unwrap,
        GoType:        goTypeName(field.Type),
        LocalName:     fmt.Sprintf("%s%s", strcase.ToLowerCamel(field.Name), strcase.ToCamel(meth.Name)),
    }

    if field.Enclosing != nil {
        newField.Enclosing = NewField(&types.HTTPParameter{Field: field.Enclosing}, meth)
    }

    if field.Type.Message == nil && field.Type.Enum == nil && field.Type.Map == nil {
        newField.IsBaseType = true
    } else {
        if pkgName := compiler.GetPackageName(newField.GoType); len(pkgName) > 0 {
            newField.GoType = pkgName + "." + newField.GoType
        }
    }

    // Modify GoType to reflect pointer or repeated status
    if field.Type.StarExpr && field.Type.ArrayType {
        newField.GoType = "[]*" + newField.GoType
    } else if field.Type.ArrayType {
        newField.GoType = "[]" + newField.GoType
    } else if field.Type.Map != nil {
        newField.GoType = "map[" + goTypeName(field.Type.Map.KeyType) + "]" + newField.GoType
    }

    // IsEnum needed for ConvertFunc and TypeConversion logic just below
    newField.IsEnum = field.Type.Enum != nil
    newField.ConvertFunc, newField.ConvertFuncNeedsErrorCheck = createDecodeConvertFunc(newField)
    newField.TypeConversion = createDecodeTypeConversion(newField)

    // Enums are allowed in query/path parameters, skip warning
    if newField.IsEnum {
        return nil
    }

    // Emit warnings for certain cases
    if !newField.IsBaseType && newField.Location != "body" {
        //log.Warnf(
        //	"%s.%s is a non-base type specified to be located outside of "+
        //		"the body. Non-base types outside the body may result in "+
        //		"generated code which fails to compile.",
        //	meth.Name,
        //	newField.Name)
    }
    if newField.Repeated && newField.Location == "path" {
        //log.Warnf(
        //	"%s.%s is a repeated field specified to be in the path. "+
        //		"Repeated fields are not supported in the path and may"+
        //		"result in generated code which fails to compile.",
        //	meth.Name,
        //	newField.Name)
    }
    return &newField
}

// GenServerDecode returns the generated code for the server-side decoding of
// an http request into its request struct.
func (b *Binding) GenServerDecode() (string, error) {
    code, err := ApplyTemplate("ServerDecodeTemplate", templates.ServerDecodeTemplate, b, TemplateFunctions)
    if err != nil {
        return "", err
    }
    code = FormatCode(code)
    return code, nil
}

// GenClientEncode returns the generated code for the client-side encoding of
// that clients request struct into the correctly formatted http request.
func (b *Binding) GenClientEncode() (string, error) {
    code, err := ApplyTemplate("ClientEncodeTemplate", templates.ClientEncodeTemplate, b, TemplateFunctions)
    if err != nil {
        return "", err
    }
    code = FormatCode(code)
    return code, nil
}

// PathSections returns a slice of strings for templating the creation of a
// fully assembled URL with the correct fields in the correct locations.
//
// For example, let's say there's a method "Sum" which accepts a "SumRequest",
// and SumRequest has two fields, 'a' and 'b'. Additionally, lets say that this
// binding for "Sum" has a path of "/sum/{a}". If we call the PathSection()
// method on this binding, it will return a slice that looks like the
// following slice literal:
//
//     []string{
//         "\"\"",
//         "\"sum\"",
//         "fmt.Sprint(req.A)",
//     }
func (b *Binding) PathSections() []string {
    isEnum := make(map[string]struct{})
    for _, v := range b.Fields {
        if v.IsEnum {
            isEnum[v.Name] = struct{}{}
        }
    }

    rv := []string{}
    parts := strings.Split(b.PathTemplate, "/")
    for _, part := range parts {
        if len(part) > 2 && part[0] == '{' && part[len(part)-1] == '}' {
            name := RemoveBraces(part)
            if _, ok := isEnum[strcase.ToCamel(name)]; ok {
                convert := fmt.Sprintf("fmt.Sprintf(\"%%d\", req.%v)", strcase.ToCamel(name))
                rv = append(rv, convert)
                continue
            }
            convert := fmt.Sprintf("fmt.Sprint(req.%v)", strcase.ToCamel(name))
            rv = append(rv, convert)
        } else {
            // Add quotes around things which'll be embeded as string literals,
            // so that the 'fmt.Sprint' lines will be unquoted and thus
            // evaluated as code.
            rv = append(rv, `"`+part+`"`)
        }
    }
    return rv
}

func FieldArrayElementType(t string) string {
    if t == "[]byte" {
        return t
    }

    if strings.HasPrefix(t, "[]*") {
        return strings.TrimPrefix(t, "[]*")
    }
    return strings.TrimPrefix(t, "[]")
}

func IsArrayElementStringType(t string) bool {
    return FieldArrayElementType(t) == "string"
}

// createDecodeConvertFunc creates a go string representing the function to
// convert the string form of the field to it's correct go type.
func createDecodeConvertFunc(f Field) (string, bool) {
    needsErrorCheck := true

    fType := ""
    switch f.GoType {
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

    if f.Repeated {
    }

    if f.IsEnum && !f.Repeated {
        fType = "%s, err := strconv.ParseInt(%s, 10, 32)"
        return fmt.Sprintf(fType, f.LocalName, f.LocalName+"Str"), true
    }

    // Use json unmarshalling for any custom/repeated messages
    if !f.IsBaseType || f.Repeated {
        // Args representing single custom message types are represented as
        // pointers. To do a bare assignment to a pointer, our rvalue must be a
        // pointer as well. So we special case args of a single custom message
        // type so that the variable LocalName is declared as a pointer.
        singleCustomTypeUnmarshalTmpl := `
var {{.LocalName}} *{{.GoType}}
{{.LocalName}}Str := ""
if len({{.LocalName}}Strs) > 0 {
	{{.LocalName}}Str = {{.LocalName}}Strs[0]
}
{{.LocalName}} = &{{.GoType}}{}
err = jsoniter.ConfigFastest.Unmarshal([]byte({{.LocalName}}Str), {{.LocalName}})
if err != nil {
	return nil, errors.Wrapf(err, "couldn't decode {{.LocalName}} from %v", {{.LocalName}}Str)
}`

        mapTypeUmarshalTmpl := `
{{.LocalName}}Str := ""
if len({{.LocalName}}Strs) > 0 {
	{{.LocalName}}Str = {{.LocalName}}Strs[0]
}
{{.LocalName}} := make({{.GoType}})
err = jsoniter.ConfigFastest.Unmarshal([]byte({{.LocalName}}Str), {{.LocalName}})
if err != nil {
	return nil, errors.Wrapf(err, "couldn't decode {{.LocalName}} from %v", {{.LocalName}}Str)
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
        //var {{.LocalName}} {{.GoType}}
        //{{- if and (and .IsBaseType .Repeated) (not (Contains .GoType "[]byte"))}}
        //err = json.Unmarshal([]byte({{.LocalName}}Str), &{{.LocalName}})
        //if err != nil {
        //	{{.LocalName}}Str = "[" + {{.LocalName}}Str + "]"
        //}
        //{{- end}}
        //err = json.Unmarshal([]byte({{.LocalName}}Str), &{{.LocalName}})`

        //		errorCheckingTmpl := `
        //if err != nil {
        //	return nil, errors.Wrapf(err, "couldn't decode {{.LocalName}} from %v", {{.LocalName}}Str)
        //}`
        repeatedTmpl := `{{- if eq .Location "query"}}
if len({{.LocalName}}Strs) == 1 {
	{{.LocalName}}Strs = ParseArrayStr({{.LocalName}}Strs[0], ",")
}{{- else -}}
{{.LocalName}}Strs := ParseArrayStr({{.LocalName}}Str, ",")
{{- end}}
var {{.LocalName}} {{.GoType}}
for _, str := range {{.LocalName}}Strs {
	{{if .IsBaseType}}var v {{FieldArrayElementType .GoType}}{{else}}v := &{{FieldArrayElementType .GoType}}{}{{end}}
	err = jsoniter.ConfigFastest.UnmarshalFromString({{if IsArrayElementStringType .GoType}}"\"" + str + "\""{{else}}str{{end}}, {{if .IsBaseType}}&v{{else}}v{{end}})
	if err != nil {
		return nil, errors.Wrapf(err, "couldn't decode {{.Name}} field from %v", str)
	}
	{{.LocalName}} = append({{.LocalName}}, v)
}`

        var jsonConvTmpl string
        if f.IsMap {
            jsonConvTmpl = mapTypeUmarshalTmpl
        } else if f.Repeated {
            jsonConvTmpl = repeatedTmpl
        } else {
            jsonConvTmpl = singleCustomTypeUnmarshalTmpl
        }
        code, err := ApplyTemplate("UnmarshalNonBaseType", jsonConvTmpl, f, TemplateFunctions)
        if err != nil {
            panic(fmt.Sprintf("Couldn't apply template: %v", err))
        }
        return code, false
    }
    return fmt.Sprintf(fType, f.LocalName, f.LocalName+"Str"), needsErrorCheck
}

// createDecodeTypeConversion creates a go string that converts a 64 bit type
// to a 32 bit type as strconv.ParseInt, ParseUInt, and ParseFloat always
// return the 64 bit type. If the type is not a 64 bit integer type or is
// repeated, then returns the LocalName of that Field.
func createDecodeTypeConversion(f Field) string {
    if f.Repeated {
        // Equivalent of the 'default' case below, but taken early for repeated
        // types.
        return f.LocalName
    }
    fType := ""
    switch f.GoType {
    case "uint32", "int32", "float32":
        fType = f.GoType + "(%s)"
    default:
        fType = "%s"
    }
    if f.IsEnum {
        fType = f.GoType + "(%s)"
    }
    return fmt.Sprintf(fType, f.LocalName)
}

// The 'basePath' of a path is the section from the start of the string till
// the first '{' character.
func basePath(path string) string {
    parts := strings.Split(path, "{")
    return parts[0]
}

// DigitEnglish is a map of runes of digits zero to nine to their lowercase
// english language spellings.
var DigitEnglish = map[rune]string{
    '0': "zero",
    '1': "one",
    '2': "two",
    '3': "three",
    '4': "four",
    '5': "five",
    '6': "six",
    '7': "seven",
    '8': "eight",
    '9': "nine",
}

// EnglishNumber takes an integer and returns the english words that represents
// that number, in base ten. Examples:
//     1  -> "One"
//     5  -> "Five"
//     10 -> "OneZero"
//     48 -> "FourEight"
func EnglishNumber(i int) string {
    n := strconv.Itoa(i)
    rv := ""
    for _, c := range n {
        if engl, ok := DigitEnglish[rune(c)]; ok {
            rv += strings.Title(engl)
        }
    }
    return rv
}
