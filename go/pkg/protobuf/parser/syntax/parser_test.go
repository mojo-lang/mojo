package syntax

import (
	"testing"

	"github.com/alecthomas/assert"

	proto2 "github.com/mojo-lang/mojo/go/pkg/protobuf2/parser/syntax/testdata"
	proto3 "github.com/mojo-lang/mojo/go/pkg/protobuf3/parser/syntax/testdata"
)

func TestParser_ParseString(t *testing.T) {
	file, err := New(nil).ParseString(proto3.AddressBookProto)
	assert.NoError(t, err)
	assert.NotNil(t, file)
}

func TestParser_ParseString2(t *testing.T) {
	file, err := New(nil).ParseString(proto2.DemoProto)
	assert.NoError(t, err)
	assert.NotNil(t, file)
}
