package render

import (
	"io"

	"github.com/mojo-lang/mojo/go/pkg/ncraft/data"
)

type Renderer interface {
	Render(string, *data.Service) (io.Reader, error)
}
