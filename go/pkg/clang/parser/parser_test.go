package parser

import (
	"testing"

	"github.com/alecthomas/assert"
)

func TestParser_ParseFile(t *testing.T) {
	cmdArgs := []string{"-I/Library/Developer/CommandLineTools/SDKs/MacOSX.sdk/usr/include"}
	file, err := New(nil).ParseFile("./testdata/func_call_with_var_args.h", cmdArgs)
	assert.NoError(t, err)
	assert.NotNil(t, file)
}
