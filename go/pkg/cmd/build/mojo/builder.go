package mojo

import (
	"github.com/mojo-lang/core/go/pkg/logs"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/mojo-lang/mojo/go/pkg/cmd/build/builder"
	_ "github.com/mojo-lang/mojo/go/pkg/mojo/compiler"
	"github.com/mojo-lang/mojo/go/pkg/mojo/context"
	_ "github.com/mojo-lang/mojo/go/pkg/mojo/mpm"
	_ "github.com/mojo-lang/mojo/go/pkg/mojo/parser"
	"github.com/mojo-lang/mojo/go/pkg/mojo/plugin"
	"github.com/mojo-lang/mojo/go/pkg/mojo/plugin/parser"
	"os"
	"strings"
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
	pkg, err := plugins.ParsePackagePath(parser.WithWorkingDir(context.Empty(), b.PWD), b.Path, os.DirFS(b.PWD))
	if err != nil {
		return nil, err
	}
	b.Package = pkg

	return pkg, err
}
