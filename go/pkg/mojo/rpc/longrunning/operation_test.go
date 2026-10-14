package longrunning

import (
	"testing"

	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/stretchr/testify/require"
)

func TestOperationValueAndResponse(t *testing.T) {
	operation := &Operation{Name: "job"}
	require.Same(t, operation, operation.SetValue("result"))
	require.Equal(t, "result", operation.GetValue().Get())
	require.Same(t, operation.Response, operation.GetValue())
	response := core.NewChecksum(core.Checksum_ALGORITHM_SHA256, []byte("payload"))
	require.Same(t, operation, operation.SetResponse(response))
	require.Same(t, response, operation.GetValue().Get())
	operation.SetValue(nil)
	require.Nil(t, operation.GetValue().Get())
	require.Equal(t, "job", operation.Name)
}

func TestOperationErrorAndNilReceiver(t *testing.T) {
	operation := &Operation{}
	failure := &core.Error{Message: "failed"}
	require.Same(t, operation, operation.SetError(failure))
	require.Same(t, failure, operation.GetError())
	operation.SetError(nil)
	require.Nil(t, operation.GetError())
	var absent *Operation
	require.Nil(t, absent.GetValue())
	require.Nil(t, absent.SetValue("ignored"))
	require.Nil(t, absent.SetResponse("ignored"))
	require.Nil(t, absent.SetError(failure))
}
