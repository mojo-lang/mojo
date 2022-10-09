package syntax

import (
	"github.com/alecthomas/assert"
	"testing"
)

func TestServiceDefVisitor_VisitRpc_Stream(t *testing.T) {
	stream := `syntax = "proto3";
service T {
rpc test ( in ) returns (stream out);
}`
	file, err := New(nil).ParseString(stream)
	assert.NoError(t, err)
	assert.NotNil(t, file)
}
