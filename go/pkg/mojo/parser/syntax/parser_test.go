package syntax

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mojo-lang/mojo/go/pkg/plugin"
)

func TestParser_ParseString(t *testing.T) {
	file, err := plugin.NewPlugins("syntax").ParseString(context.Background(), `type A = Int`)
	assert.NoError(t, err)
	assert.NotNil(t, file)
}

func TestParser_ParseExpression(t *testing.T) {
	expr, err := ParseExpression(`id in ['foo']`)
	assert.NoError(t, err)
	assert.NotNil(t, expr)
}

func TestParser_ParseExpression2(t *testing.T) {
	expr, err := ParseExpression(`id in ['foo'] and id >= 'bar' and id < 'baz'`)
	assert.NoError(t, err)
	assert.NotNil(t, expr)
}
