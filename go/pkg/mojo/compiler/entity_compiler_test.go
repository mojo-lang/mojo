package compiler

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mojo-lang/mojo/go/pkg/context"
	_ "github.com/mojo-lang/mojo/go/pkg/mojo/mpm"
	_ "github.com/mojo-lang/mojo/go/pkg/mojo/parser/semantic"
	_ "github.com/mojo-lang/mojo/go/pkg/mojo/parser/semantic/circle"
	_ "github.com/mojo-lang/mojo/go/pkg/mojo/parser/semantic/identifier"
	_ "github.com/mojo-lang/mojo/go/pkg/mojo/parser/syntax"
	"github.com/mojo-lang/mojo/go/pkg/plugin"
)

func TestEntityCompiler_CompilePackage(t *testing.T) {
	plugins := plugin.NewPlugins("mpm", "syntax", "semantic", "compiler")
	pkg, err := plugins.ParsePath(context.Empty(), "../testdata/mojo-entity")

	// Fixme
	// graph.Builder{Builder: builder.Builder{Package: pkg}}.Build()
	assert.NoError(t, err)
	assert.NotNil(t, pkg)
}
