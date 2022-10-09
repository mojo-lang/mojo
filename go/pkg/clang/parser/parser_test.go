package parser

import (
	"github.com/alecthomas/assert"
	"testing"
)

func TestParser_ParseFile(t *testing.T) {
	cmdArgs := []string{"-I/Library/Developer/CommandLineTools/SDKs/MacOSX.sdk/usr/include"}
	file, err := New(nil).ParseFile("./test-data/func_call_with_var_args.h", cmdArgs)
	assert.NoError(t, err)
	assert.NotNil(t, file)
}
