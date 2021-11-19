package httptransport

import (
	"strings"
	"text/template"

	"github.com/mojo-lang/core/go/pkg/mojo/core/strcase"
	"github.com/mojo-lang/mojo/go/pkg/ncraft/gokit/compiler"
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
	"GoName":                   strcase.ToCamel,
	"Contains":                 strings.Contains,
	"PackageName":              compiler.GetPackageName,
	"FieldArrayElementType":    FieldArrayElementType,
	"IsArrayElementStringType": IsArrayElementStringType,
}
