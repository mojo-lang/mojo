package syntax

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mojo-lang/mojo/go/pkg/context"
)

func TestServiceDefVisitor_VisitRpc_Stream(t *testing.T) {
	str := `syntax = "proto3";
service T {
rpc test ( in ) returns (stream out);
}`
	file, err := New(nil).ParseString(context.Empty(), str)
	assert.NoError(t, err)
	assert.NotNil(t, file)
}
