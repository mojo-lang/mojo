package compiler

import (
	"github.com/mojo-lang/mojo/go/pkg/mojo/context"
	_ "github.com/mojo-lang/mojo/go/pkg/mojo/mpm"
	_ "github.com/mojo-lang/mojo/go/pkg/mojo/parser/semantic"
	_ "github.com/mojo-lang/mojo/go/pkg/mojo/parser/semantic/circle"
	_ "github.com/mojo-lang/mojo/go/pkg/mojo/parser/semantic/identifier"
	_ "github.com/mojo-lang/mojo/go/pkg/mojo/parser/syntax"
	"github.com/mojo-lang/mojo/go/pkg/mojo/plugin"
	"github.com/mojo-lang/mojo/go/pkg/mojo/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEntityCompiler_CompilePackage(t *testing.T) {
	plugins := plugin.NewPlugins("mpm", "syntax", "semantic", "compiler")
	pkg, err := plugins.ParsePackagePath(context.Empty(), "mojo-entity", testdata.EntityCaseFiles)

	//Fixme
	//graph.Builder{Builder: builder.Builder{Package: pkg}}.Build()
	assert.NoError(t, err)
	assert.NotNil(t, pkg)
}
