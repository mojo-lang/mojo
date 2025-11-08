package generator

import (
	"bytes"
	_ "embed"
	"text/template"

	"github.com/mojo-lang/mojo/go/pkg/util"
	"github.com/mojo-lang/mojo/packages/core/go/pkg/mojo/core/strcase"
	"github.com/pkg/errors"
	"strings"
)

//go:embed template/java/v1/SERVICE_http.java.tmpl
var javaServiceHttpFile string

//go:embed template/java/v1/constant/service_name_constants.java.tmpl
var javaServiceNameConstantsFile string

//go:embed template/java/v1/factory/SERVICE_http_fallback_factory.java.tmpl
var javaServiceHttpFallbackFactoryFile string

//go:embed template/pom.xml.tmpl
var pomFile string

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
// template and used within those template.
var FuncMap = template.FuncMap{
	"ToLower":          strings.ToLower,
	"ToUpper":          strings.ToUpper,
	"GoName":           strcase.ToCamel,
	"GoLocalName":      util.GoLocalName,
	"ToSnake":          strcase.ToSnake,
	"ToScreamingSnake": strcase.ToScreamingSnake,
	"ToKebab":          strcase.ToKebab,
	"ToCamel":          strcase.ToCamel,
	"ToLowerCamel":     strcase.ToLowerCamel,
	"CompactFullName":  CompactFullName,
	"IsMojoPackage":    func(pkg string) bool { return strings.HasPrefix(pkg, "mojo.") },
	"IsPackage":        func(srcPkg string, targetPkg string) bool { return strings.HasPrefix(srcPkg, targetPkg) },
}
