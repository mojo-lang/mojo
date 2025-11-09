// Package httptransport provides functions and templates helpers for templating
// the http-transport of a go-kit based service.
package httptransport

import (
	"bytes"
	"io"
	"text/template"

	"github.com/mojo-lang/mojo/go/pkg/compiler/ncraft/data"
	_go "github.com/mojo-lang/mojo/go/pkg/compiler/ncraft/go"
	"github.com/mojo-lang/mojo/go/pkg/compiler/ncraft/gokit/generator/httptransport/templates"
	"github.com/mojo-lang/mojo/go/pkg/compiler/util"
)

const ServerHttpTransportPath = "pkg/NAME-service/svc/transport_http.go.tmpl"

type ServerHttpTransport struct {
}

func NewServerHttpTransport(ds *data.Service) (*ServerHttpTransport, error) {
	for _, method := range ds.Interface.Methods {
		for _, binding := range method.Bindings {
			for _, param := range binding.Parameters {
				str, err := createParamUnmarshaler(param, ds.FuncMap)
				if err != nil {
					return nil, err
				}
				param.Go.ParamUnmarshaler = str
			}

			if serverDecode, err := createServerDecode(binding, ds.FuncMap); err != nil {
				return nil, err
			} else {
				binding.Extensions["ServerDecode"] = string(serverDecode)
			}
		}
	}

	return &ServerHttpTransport{}, nil
}

func (h *ServerHttpTransport) Render(tmpl string, ds *data.Service) (io.Reader, error) {
	code, err := util.ApplyTemplate("ServerTemplate", templates.ServerTemplate, ds, ds.FuncMap)
	if err != nil {
		return nil, err
	}
	if bs, err := _go.FormatCode(code); err != nil {
		return nil, err
	} else {
		return bytes.NewBuffer(bs), nil
	}
}

// createServerDecode returns the generated code for the server-side decoding of
// a http request into its request struct.
func createServerDecode(binding *data.HTTPBinding, funcMap template.FuncMap) ([]byte, error) {
	code, err := util.ApplyTemplate("ServerDecodeTemplate", templates.ServerDecodeTemplate, binding, funcMap)
	if err != nil {
		return nil, err
	}
	return _go.FormatCode(code)
}
