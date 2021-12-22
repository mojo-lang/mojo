package mpm

import (
	"github.com/mojo-lang/mojo/go/pkg/context"
	_ "github.com/mojo-lang/mojo/go/pkg/parser/semantic"
	_ "github.com/mojo-lang/mojo/go/pkg/parser/semantic/circle"
	_ "github.com/mojo-lang/mojo/go/pkg/parser/semantic/identifier"
	_ "github.com/mojo-lang/mojo/go/pkg/parser/syntax"
	"github.com/mojo-lang/mojo/go/pkg/plugin"
	"github.com/mojo-lang/mojo/go/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDependencyParser_ParsePackagePath(t *testing.T) {
	plugins := plugin.NewPlugins("mpm", "syntax", "semantic")
	pkg, err := plugins.ParsePackagePath(context.Empty(), "case-alias", test.AliasCaseFiles)
	assert.NoError(t, err)
	assert.NotNil(t, pkg)
}
