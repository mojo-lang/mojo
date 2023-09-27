package syntax

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mojo-lang/mojo/go/pkg/context"
	"github.com/mojo-lang/mojo/go/pkg/protobuf2/parser/syntax/testdata"
)

func TestParser_ParseString_Demo(t *testing.T) {
	file, err := New(nil).ParseString(context.Empty(), testdata.DemoProto)
	assert.NoError(t, err)
	assert.NotNil(t, file)
}

func TestParser_ParseString_Extend(t *testing.T) {
	file, err := New(nil).ParseString(context.Empty(), testdata.ExtendProto)
	assert.NoError(t, err)
	assert.NotNil(t, file)
}
