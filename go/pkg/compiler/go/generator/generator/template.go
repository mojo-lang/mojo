package generator

import (
	"bytes"
	_ "embed"
	"strings"
	"text/template"

	"github.com/mojo-lang/mojo/go/pkg/compiler/util"

	"github.com/mojo-lang/mojo/go/pkg/mojo/core/strcase"
	"github.com/pkg/errors"
)

//go:embed template/go/BOXED_ARRAY.ext.go.tmpl
var goArrayExtFile string

//go:embed template/go/BOXED_ARRAY.json.go.tmpl
var goArrayJsonFile string

//go:embed template/go/BOXED_MAP.ext.go.tmpl
var goMapExtFile string

//go:embed template/go/BOXED_MAP.json.go.tmpl
var goMapJsonFile string

//go:embed template/go/BOXED_UNION.ext.go.tmpl
var goUnionExtFile string

//go:embed template/go/BOXED_UNION.json.go.tmpl
var goUnionJsonFile string

//go:embed template/go/ENUM.fmt.go.tmpl
var goEnumFmtFile string

//go:embed template/go/ENUM.json.go.tmpl
var goEnumJsonFile string

//go:embed template/go/DB_JSON.sql.go.tmpl
var goDbJSONSqlFile string

//go:embed template/go/DB_JSON.json.go.tmpl
var goDbJSONJSONFile string

//go:embed template/go/FORMAT_JSON.json.go.tmpl
var goFormatJSONFile string

//go:embed template/go/ARRAY_RESPONSE.json.go.tmpl
var goArrayResponseJsonFile string

//go:embed template/go.mod.tmpl
var goModFile string

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

func CompactFullName(name string) string {
	return strings.ReplaceAll(name, "_", "")
}

// FuncMap contains a series of utility functions to be passed into
// templates and used within those templates.
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
