// Package httptransport provides functions and template helpers for templating
// the http-transport of a go-kit based service.
package httptransport

import (
	"bytes"
	"github.com/pkg/errors"
	"go/format"
	"text/template"

	"github.com/mojo-lang/mojo/go/pkg/ncraft/gokit/generator/httptransport/templates"
	"github.com/mojo-lang/mojo/go/pkg/ncraft/gokit/generator/types"
)

// Helper is the base struct for the data structure containing all the
// information necessary to correctly template the HTTP transport functionality
// of a service. Helper must be built from a Service.
type Helper struct {
	Methods        []*Method
	ServerTemplate func(interface{}) (string, error)
	ClientTemplate func(interface{}) (string, error)
}

// NewHelper builds a helper struct from a service declaration. The other
// "New*" functions in this file are there to make this function smaller and
// more testable.
func NewHelper(svc *types.Interface) *Helper {
	// The HTTPAssistFuncs global is a group of function literals defined
	// within templates.go
	rv := Helper{
		ServerTemplate: GenServerTemplate,
		ClientTemplate: GenClientTemplate,
	}
	for _, meth := range svc.Methods {
		if len(meth.Bindings) > 0 {
			nMeth := NewMethod(meth)
			rv.Methods = append(rv.Methods, nMeth)
		}
	}
	return &rv
}

func GenServerTemplate(exec interface{}) (string, error) {
	code, err := ApplyTemplate("ServerTemplate", templates.ServerTemplate, exec, TemplateFunctions)
	if err != nil {
		return "", err
	}
	code = FormatCode(code)
	return code, nil
}

func GenClientTemplate(exec interface{}) (string, error) {
	code, err := ApplyTemplate("ClientTemplate", templates.ClientTemplate, exec, TemplateFunctions)
	if err != nil {
		return "", err
	}
	code = FormatCode(code)
	return code, nil
}

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
