package httptransport

import (
	gogen "github.com/golang/protobuf/protoc-gen-go/generator"
	"github.com/iancoleman/strcase"
	"github.com/mojo-lang/mojo/go/pkg/ncraft/gokit/compiler"
	"strings"
	"text/template"
)

// TemplateFunctions contains a series of utility functions to be passed into
// templates and used within those templates.
var TemplateFunctions = template.FuncMap{
	"ToLower":                  strings.ToLower,
	"ToUpper":                  strings.ToUpper,
	"ToSnake":                  strcase.ToSnake,
	"ToKebab":                  strcase.ToKebab,
	"ToCamel":                  strcase.ToCamel,
	"ToLowerCamel":             strcase.ToLowerCamel,
	"Title":                    strings.Title,
	"GoName":                   gogen.CamelCase,
	"Contains":                 strings.Contains,
	"PackageName":              compiler.GetPackageName,
	"FieldArrayElementType":    FieldArrayElementType,
	"IsArrayElementStringType": IsArrayElementStringType,
}
