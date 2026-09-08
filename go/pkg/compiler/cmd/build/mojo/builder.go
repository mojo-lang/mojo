package mojo

import (
	"github.com/mojo-lang/mojo/go/pkg/logs"
	"github.com/mojo-lang/mojo/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/go/pkg/compiler/cmd/build/builder"
	"github.com/mojo-lang/mojo/go/pkg/compiler/context"
	_ "github.com/mojo-lang/mojo/go/pkg/compiler/mojo/compiler"
	_ "github.com/mojo-lang/mojo/go/pkg/compiler/mojo/mpm"
	_ "github.com/mojo-lang/mojo/go/pkg/compiler/mojo/parser"
	"github.com/mojo-lang/mojo/go/pkg/compiler/plugin"
)

type Builder struct {
	builder.Builder
	PackageName string
}

func (b Builder) Build() (*lang.Package, error) {
	logs.Infow("begin to parse mojo package.", "pwd", b.PWD, "path", b.Path)

	plugins := plugin.NewPlugins("mpm", "syntax", "semantic", "compiler")

	pkg, err := plugins.ParsePath(plugin.WithPackageName(context.Empty(), b.PackageName), b.GetAbsolutePath())
	if err != nil {
		return nil, err
	}
	b.Package = pkg

	return pkg, err
}
