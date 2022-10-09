package syntax

import (
	_ "embed"
	"github.com/alecthomas/assert"
	"testing"
)

//go:embed test_data/func_call_with_var_args.c
var funcCallWithVarArgsFile string

func TestParser_ParseString(t *testing.T) {
	file, err := New(nil).ParseString(funcCallWithVarArgsFile)
	assert.NoError(t, err)
	assert.NotNil(t, file)
}
