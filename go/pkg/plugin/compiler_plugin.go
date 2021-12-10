package plugin

import (
	"github.com/mojo-lang/mojo/go/pkg/plugin/compiler"
)

type compilerPlugin struct {
	basicPlugin
	compiler.PackageCompiler
}
