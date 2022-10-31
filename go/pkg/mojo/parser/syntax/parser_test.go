package syntax

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mojo-lang/mojo/go/pkg/mojo/plugin"
)

func TestParser_ParseString(t *testing.T) {
	file, err := plugin.NewPlugins("syntax").ParseString(`type A = Int`)
	assert.NoError(t, err)
	assert.NotNil(t, file)
}
