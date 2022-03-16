package render

import "io"

type Renderable interface {
    Render(string, *Data) (io.Reader, error)
}
