package scaffolding

import (
	"bytes"
	_ "embed"
	"strings"
	"text/template"

	"github.com/mojo-lang/core/go/pkg/mojo/core/strcase"
	"github.com/pkg/errors"
)

//go:embed template/.gitignore.tmpl
var gitignoreFile string

//go:embed template/readme.md.tmpl
var readmeFile string

//go:embed template/package.mojo.tmpl
var packageFile string

//go:embed template/hello-world/v1/echo.mojo
var helloWorldEcho string

//go:embed template/hello-world/v1/hello_world.mojo
var helloWorldService string

// ApplyTemplate applies a templates with a given name, executor context, and
// function map. Returns the output of the templates on success, returns an
// error if templates failed to execute.
func ApplyTemplate(name string, tmpl string, executor interface{}, funcMap template.FuncMap) (string, error) {
	codeTemplate := template.Must(template.New(name).Funcs(funcMap).Parse(tmpl))

	code := bytes.NewBuffer(nil)
	err := codeTemplate.Execute(code, executor)
	if err != nil {
		return "", errors.Wrapf(err, "attempting to execute templates %q", name)
	}
	return code.String(), nil
}

// FuncMap contains a series of utility functions to be passed into
// templates and used within those templates.
var FuncMap = template.FuncMap{
	"ToLower":      strings.ToLower,
	"ToUpper":      strings.ToUpper,
	"GoName":       strcase.ToCamel,
	"ToSnake":      strcase.ToSnake,
	"ToKebab":      strcase.ToKebab,
	"ToCamel":      strcase.ToCamel,
	"ToLowerCamel": strcase.ToLowerCamel,
}
