package template

import (
	"bytes"
	"github.com/pkg/errors"
	"go/format"
	"text/template"
)

// ApplyTemplate applies a template with a given name, executor context, and
// function map. Returns the output of the template on success, returns an
// error if template failed to execute.
func ApplyTemplate(name string, tmpl string, executor interface{}, fncs template.FuncMap) (string, error) {
	codeTemplate := template.Must(template.New(name).Funcs(fncs).Parse(tmpl))

	code := bytes.NewBuffer(nil)
	err := codeTemplate.Execute(code, executor)
	if err != nil {
		return "", errors.Wrapf(err, "attempting to execute template %q", name)
	}
	return code.String(), nil
}

// FormatCode takes a string representing some go code and attempts to format
// that code. If formating fails, the original source code is returned.
func FormatCode(code string) string {
	formatted, err := format.Source([]byte(code))

	if err != nil {
		// Set formatted to code so at least we get something to examine
		formatted = []byte(code)
	}

	return string(formatted)
}