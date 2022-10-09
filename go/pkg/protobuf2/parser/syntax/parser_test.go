package syntax

import (
	"github.com/alecthomas/assert"
	"github.com/mojo-lang/mojo/go/pkg/protobuf2/parser/syntax/test_data"
	"testing"
)

func TestParser_ParseString_Demo(t *testing.T) {
	file, err := New(nil).ParseString(test_data.DemoProto)
	assert.NoError(t, err)
	assert.NotNil(t, file)
}

func TestParser_ParseString_Extend(t *testing.T) {
	file, err := New(nil).ParseString(test_data.ExtendProto)
	assert.NoError(t, err)
	assert.NotNil(t, file)
}
