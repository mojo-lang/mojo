package syntax

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mojo-lang/mojo/go/pkg/context"
)

func TestServiceDefVisitor_VisitRpc_Stream(t *testing.T) {
	stream := `syntax = "proto2";
service T {
rpc test ( in ) returns (out);
}`
	file, err := New(nil).ParseString(context.Empty(), stream)
	assert.NoError(t, err)
	assert.NotNil(t, file)
}
