package converter

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mojo-lang/mojo/go/pkg/compiler/context"
	_ "github.com/mojo-lang/mojo/go/pkg/compiler/mojo/mpm"
	_ "github.com/mojo-lang/mojo/go/pkg/compiler/mojo/parser/semantic"
	_ "github.com/mojo-lang/mojo/go/pkg/compiler/mojo/parser/semantic/circle"
	_ "github.com/mojo-lang/mojo/go/pkg/compiler/mojo/parser/semantic/identifier"
	_ "github.com/mojo-lang/mojo/go/pkg/compiler/mojo/parser/syntax"
	"github.com/mojo-lang/mojo/go/pkg/compiler/plugin"
	"github.com/mojo-lang/mojo/go/pkg/compiler/protobuf/generator"
)

func TestConvert_CompilePackages(t *testing.T) {
	plugins := plugin.NewPlugins("mpm", "syntax", "semantic", "compiler")
	pkg, err := plugins.ParsePath(context.Empty(), "../testdata/mojo-test")

	if assert.NoError(t, err) {
		compiler := New()
		err = compiler.CompilePackage(context.Empty(), pkg)
		if assert.NoError(t, err) {
			files := compiler.Descriptors.Filter(pkg.FullName, false)
			out, err := generator.New().GenerateDescriptors(files)
			assert.NoError(t, err)
			assert.Equal(t, 1, len(out))
		}
	}
}
