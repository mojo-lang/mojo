package scaffolding

import (
	"bytes"
	_ "embed"
	"github.com/iancoleman/strcase"
	"github.com/pkg/errors"
	"strings"
	"text/template"
)

//go:embed template/.gitignore.tmpl
var gitignoreFile string

//go:embed template/readme.md.tmpl
var readmeFile string

//go:embed template/wand.mojo.tmpl
var wandFile string

// ApplyTemplate applies a template with a given name, executor context, and
// function map. Returns the output of the template on success, returns an
// error if template failed to execute.
func ApplyTemplate(name string, tmpl string, executor interface{}, funcMap template.FuncMap) (string, error) {
	codeTemplate := template.Must(template.New(name).Funcs(funcMap).Parse(tmpl))

	code := bytes.NewBuffer(nil)
	err := codeTemplate.Execute(code, executor)
	if err != nil {
		return "", errors.Wrapf(err, "attempting to execute template %q", name)
	}
	return code.String(), nil
}

// FuncMap contains a series of utility functions to be passed into
// templates and used within those templates.
var FuncMap = template.FuncMap{
	"ToLower":         strings.ToLower,
	"ToUpper":         strings.ToUpper,
	"GoName":          strcase.ToCamel,
	"ToSnake":         strcase.ToSnake,
	"ToKebab":         strcase.ToKebab,
	"ToCamel":         strcase.ToCamel,
	"ToLowerCamel":    strcase.ToLowerCamel,
}
