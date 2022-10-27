package syntax

import (
	"github.com/alecthomas/assert"
	"github.com/mojo-lang/mojo/go/pkg/protobuf2/parser/syntax/testdata"
	"testing"
)

func TestParser_ParseString_Demo(t *testing.T) {
	file, err := New(nil).ParseString(testdata.DemoProto)
	assert.NoError(t, err)
	assert.NotNil(t, file)
}

func TestParser_ParseString_Extend(t *testing.T) {
	file, err := New(nil).ParseString(testdata.ExtendProto)
	assert.NoError(t, err)
	assert.NotNil(t, file)
}
