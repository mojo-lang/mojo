package compiler

import (
	"testing"

	"github.com/mojo-lang/mojo/packages/core/go/pkg/mojo/core"
	"github.com/stretchr/testify/assert"

	"github.com/mojo-lang/mojo/go/pkg/context"
	_ "github.com/mojo-lang/mojo/go/pkg/mojo/compiler"
	_ "github.com/mojo-lang/mojo/go/pkg/mojo/mpm"
	_ "github.com/mojo-lang/mojo/go/pkg/mojo/parser"
	"github.com/mojo-lang/mojo/go/pkg/plugin"
)

func TestServices_CompileInterface(t *testing.T) {
	plugins := plugin.NewPlugins("mpm", "syntax", "semantic", "compiler")
	pkg, err := plugins.ParsePath(context.Empty(), "../testdata/mojo-ncraft")
	assert.NoError(t, err)
	assert.NotNil(t, pkg)
	if pkg == nil {
		return
	}

	options := make(core.Options)
	for _, p := range pkg.GetAllPackages() {
		options[p.FullName] = p.GetGoPackageImport()
	}
	for _, p := range pkg.GetAllDependentPackages() {
		options[p.FullName] = p.GetGoPackageImport()
	}
	services, err := CompilePackage(WithGoPackageImports(context.Empty(), options), pkg)
	assert.NoError(t, err)
	assert.NotEmpty(t, services)
}
