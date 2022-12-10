package generator

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mojo-lang/mojo/go/pkg/context"
	_ "github.com/mojo-lang/mojo/go/pkg/mojo/compiler"
	_ "github.com/mojo-lang/mojo/go/pkg/mojo/mpm"
	_ "github.com/mojo-lang/mojo/go/pkg/mojo/parser"
	"github.com/mojo-lang/mojo/go/pkg/plugin"
	"github.com/mojo-lang/mojo/go/pkg/protobuf/converter"
)

func TestGenerator_GenerateDescriptors(t *testing.T) {
	plugins := plugin.NewPlugins("mpm", "syntax", "semantic", "compiler")
	pkg, err := plugins.ParsePath(context.Empty(), "../../mojo/testdata/mojo-enum")

	if assert.NoError(t, err) {
		compiler := converter.New()
		err = compiler.CompilePackage(context.Empty(), pkg)
		if assert.NoError(t, err) {
			files := compiler.Descriptors.Filter(pkg.FullName, false)
			out, err := New().GenerateDescriptors(files)
			_ = out
			assert.NoError(t, err)
		}
	}
}
