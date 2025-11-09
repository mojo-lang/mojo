package syntax

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed testdata/func_call_with_var_args.c
var funcCallWithVarArgsFile string

func TestParser_ParseString(t *testing.T) {
	file, err := New(nil).ParseString(funcCallWithVarArgsFile)
	assert.NoError(t, err)
	assert.NotNil(t, file)
}
