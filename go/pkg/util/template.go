package util

import (
	"bytes"
	"io"
	"reflect"
	"strings"
	"text/template"

	"github.com/mojo-lang/core/go/pkg/mojo/core/strcase"
	"github.com/pkg/errors"
)

// ApplyTemplate is a helper methods that packages can call to render a
// templates with any data and func map
func ApplyTemplate(tmplName string, tmpl string, data interface{}, funcMap template.FuncMap) (io.Reader, error) {
	codeTemplate, err := template.New(tmplName).Funcs(funcMap).Parse(tmpl)
	if err != nil {
		return nil, errors.Wrap(err, "cannot create templates")
	}

	buffer := bytes.NewBuffer(nil)
	if err = codeTemplate.Execute(buffer, data); err != nil {
		return nil, errors.Wrapf(err, "attempting to execute templates %q", tmplName)
	}

	return buffer, nil
}

// FuncMap contains a series of utility functions to be passed into
// templates and used within those templates.
var FuncMap = template.FuncMap{
	"Last":         func(x int, a interface{}) bool { return x == reflect.ValueOf(a).Len()-1 },
	"ToLower":      strings.ToLower,
	"ToUpper":      strings.ToUpper,
	"GoName":       strcase.ToCamel,
	"GoLocalName":  GoLocalName,
	"ToSnake":      strcase.ToSnake,
	"ToKebab":      strcase.ToKebab,
	"ToCamel":      strcase.ToCamel,
	"ToLowerCamel": strcase.ToLowerCamel,
}

func GoLocalName(name string, blacklist ...string) string {
	n := strcase.ToLowerCamel(name)
	for _, b := range blacklist {
		if n == b {
			n += "Var"
		}
	}
	return n
}
