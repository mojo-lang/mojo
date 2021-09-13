package generator

import (
	"bytes"
	_ "embed"
	"github.com/iancoleman/strcase"
	"github.com/pkg/errors"
	"strings"
	"text/template"
)

//go:embed template/go/BOXED_ARRAY.ext.go.tmpl
var goArrayExtFile string

//go:embed template/go/BOXED_ARRAY.json.go.tmpl
var goArrayJsonFile string

//go:embed template/go/BOXED_UNION.ext.go.tmpl
var goUnionExtFile string

//go:embed template/go/BOXED_UNION.json.go.tmpl
var goUnionJsonFile string

//go:embed template/go/ENUM.fmt.go.tmpl
var goEnumExtFile string

//go:embed template/go/ENUM.json.go.tmpl
var goEnumJsonFile string

//go:embed template/go.mod.tmpl
var goModFile string

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

func CompactFullName(name string) string {
	return strings.ReplaceAll(name, "_", "")
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
	"CompactFullName": CompactFullName,
	"IsMojoPackage":   func(pkg string) bool { return strings.HasPrefix(pkg, "mojo.") },
	"IsPackage":       func(srcPkg string, targetPkg string) bool { return strings.HasPrefix(srcPkg, targetPkg) },
}
