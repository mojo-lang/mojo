package syntax

import (
	"github.com/mojo-lang/mojo/go/pkg/plugin"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParser_ParseString(t *testing.T) {
	file, pkg, err := plugin.NewPlugins("syntax").ParseString(`type A = Int`)
	assert.NoError(t, err)
	assert.NotNil(t, file)
	assert.NotNil(t, pkg)
}
