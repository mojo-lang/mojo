package httptransport

import (
	"bytes"
	"io"
	"text/template"

	"github.com/mojo-lang/mojo/go/pkg/ncraft/data"
	_go "github.com/mojo-lang/mojo/go/pkg/ncraft/go"
	"github.com/mojo-lang/mojo/go/pkg/ncraft/gokit/generator/httptransport/templates"
	"github.com/mojo-lang/mojo/go/pkg/util"
)

type ClientHttpTransport struct {
}

func NewClientHttpTransport(ds data.Service) (*ClientHttpTransport, error) {
	for _, method := range ds.Interface.Methods {
		for _, binding := range method.Bindings {
			if encoder, err := createClientEncode(binding, ds.FuncMap); err != nil {
				return nil, err
			} else {
				binding.Extensions["ClientEncoder"] = encoder
			}
		}
	}

	return &ClientHttpTransport{}, nil
}

func (h *ClientHttpTransport) Render(tmpl string, ds *data.Service) (io.Reader, error) {
	code, err := util.ApplyTemplate(tmpl, templates.ClientTemplate, ds, ds.FuncMap)
	if err != nil {
		return nil, err
	}
	if bs, err := _go.FormatCode(code); err != nil {
		return nil, err
	} else {
		return bytes.NewReader(bs), nil
	}
}

// createClientEncode returns the generated code for the client-side encoding of
// that clients request struct into the correctly formatted http request.
func createClientEncode(binding *data.HTTPBinding, funcMap template.FuncMap) (string, error) {
	code, err := util.ApplyTemplate("ClientEncodeTemplate", templates.ClientEncodeTemplate, binding, funcMap)
	if err != nil {
		return "", err
	}
	bs, err := _go.FormatCode(code)
	if err != nil {
		return "", err
	}
	return string(bs), nil
}
