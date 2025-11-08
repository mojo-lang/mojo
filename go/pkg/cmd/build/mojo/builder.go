package mojo

import (
	"path"
	"strings"

	"github.com/mojo-lang/mojo/packages/core/go/pkg/logs"
	"github.com/mojo-lang/mojo/packages/lang/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/go/pkg/cmd/build/builder"
	"github.com/mojo-lang/mojo/go/pkg/context"
	_ "github.com/mojo-lang/mojo/go/pkg/mojo/compiler"
	_ "github.com/mojo-lang/mojo/go/pkg/mojo/mpm"
	_ "github.com/mojo-lang/mojo/go/pkg/mojo/parser"
	"github.com/mojo-lang/mojo/go/pkg/plugin"
)

type Builder struct {
	builder.Builder
}

func (b Builder) Build() (*lang.Package, error) {
	logs.Infow("begin to parse mojo package.", "pwd", b.PWD, "path", b.Path)

	plugins := plugin.NewPlugins("mpm", "syntax", "semantic", "compiler")

	if strings.HasPrefix(b.Path, b.PWD) {
		b.Path = strings.TrimPrefix(b.Path, b.PWD)
	}
	pkg, err := plugins.ParsePath(context.Empty(), path.Join(b.PWD, b.Path))
	if err != nil {
		return nil, err
	}
	b.Package = pkg

	return pkg, err
}
