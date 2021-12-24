package syntax

import (
	"github.com/mojo-lang/mojo/go/pkg/mojo/plugin"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParser_ParseString(t *testing.T) {
	file, err := plugin.NewPlugins("syntax").ParseString(`type A = Int`)
	assert.NoError(t, err)
	assert.NotNil(t, file)
}
