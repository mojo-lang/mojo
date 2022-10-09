package util

import (
	"bytes"
	"github.com/mojo-lang/core/go/pkg/mojo/core/strcase"
	"github.com/pkg/errors"
	"io"
	"text/template"
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
		return nil, errors.Wrap(err, "templates error")
	}

	return buffer, nil
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
