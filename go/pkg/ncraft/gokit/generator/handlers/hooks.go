package handlers

import (
	"io"
	"strings"

	"github.com/mojo-lang/mojo/go/pkg/ncraft/data"
	"github.com/mojo-lang/mojo/go/pkg/ncraft/gokit/generator/handlers/templates"
	"github.com/mojo-lang/mojo/go/pkg/ncraft/render"
)

const HookPath = "pkg/NAME-service/handlers/hooks.go.tmpl"

// NewHook returns a new HookRender
func NewHook(prev io.Reader) render.Renderer {
	return &HookRender{
		prev: prev,
	}
}

type HookRender struct {
	prev io.Reader
}

// Render will return the existing file if it exists, otherwise it will return
// a brand new copy from the templates.
func (h *HookRender) Render(_ string, _ *data.Service) (io.Reader, error) {
	if h.prev == nil {
		return strings.NewReader(templates.Hook), nil
	} else {
		return h.prev, nil
	}
}
