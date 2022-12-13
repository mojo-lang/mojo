package parser

import (
	"testing"

	"github.com/alecthomas/assert"
)

func TestParser_ParseFile1(t *testing.T) {
	file, err := New(nil).ParseFile("./testdata/func_call_with_var_args.h", nil)
	assert.NoError(t, err)
	assert.NotNil(t, file)
}

func TestParser_ParseFile2(t *testing.T) {
	file, err := New(nil).ParseFile("./testdata/bt.c", nil)
	assert.NoError(t, err)
	assert.NotNil(t, file)
}
