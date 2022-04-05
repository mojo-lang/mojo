package mpm

import (
    "github.com/mojo-lang/mojo/go/pkg/mojo/context"
    _ "github.com/mojo-lang/mojo/go/pkg/mojo/parser"
    "github.com/mojo-lang/mojo/go/pkg/mojo/plugin"
    "github.com/mojo-lang/mojo/go/test"
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestDependencyParser_ParsePackagePath(t *testing.T) {
    plugins := plugin.NewPlugins("mpm", "syntax", "semantic")
    pkg, err := plugins.ParsePackagePath(context.Empty(), "mojo-alias", test.AliasCaseFiles)
    assert.NoError(t, err)
    assert.NotNil(t, pkg)
}
