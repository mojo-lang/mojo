package mpm

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mojo-lang/mojo/go/pkg/context"
	_ "github.com/mojo-lang/mojo/go/pkg/mojo/compiler"
	_ "github.com/mojo-lang/mojo/go/pkg/mojo/parser"
	"github.com/mojo-lang/mojo/go/pkg/plugin"
)

func TestDependencyParser_ParsePath(t *testing.T) {
	plugins := plugin.NewPlugins("mpm", "syntax", "semantic", "compiler")
	pkg, err := plugins.ParsePath(context.Empty(), "../testdata/mojo-alias")
	assert.NoError(t, err)
	assert.NotNil(t, pkg)
}
