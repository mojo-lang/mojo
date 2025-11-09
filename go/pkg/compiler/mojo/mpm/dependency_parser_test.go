package mpm

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mojo-lang/mojo/go/pkg/compiler/context"
	_ "github.com/mojo-lang/mojo/go/pkg/compiler/mojo/compiler"
	_ "github.com/mojo-lang/mojo/go/pkg/compiler/mojo/parser"
	"github.com/mojo-lang/mojo/go/pkg/compiler/plugin"
)

func TestDependencyParser_ParsePath(t *testing.T) {
	plugins := plugin.NewPlugins("mpm", "syntax", "semantic", "compiler")
	pkg, err := plugins.ParsePath(context.Empty(), "../testdata/mojo-alias")
	assert.NoError(t, err)
	assert.NotNil(t, pkg)
}
