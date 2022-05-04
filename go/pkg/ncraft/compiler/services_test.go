package compiler

import (
    "github.com/mojo-lang/core/go/pkg/mojo/core"
    _ "github.com/mojo-lang/mojo/go/pkg/mojo/compiler"
    "github.com/mojo-lang/mojo/go/pkg/mojo/context"
    _ "github.com/mojo-lang/mojo/go/pkg/mojo/mpm"
    _ "github.com/mojo-lang/mojo/go/pkg/mojo/parser"
    "github.com/mojo-lang/mojo/go/pkg/mojo/plugin"
    "github.com/mojo-lang/mojo/go/test"
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestServices_CompileInterface(t *testing.T) {
    plugins := plugin.NewPlugins("mpm", "syntax", "semantic", "compiler")
    pkg, err := plugins.ParsePackagePath(context.Empty(), "mojo-ncraft", test.NCraftCaseFiles)
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
