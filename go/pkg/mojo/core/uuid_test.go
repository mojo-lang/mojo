package core

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewUuidAndConversion(t *testing.T) {
	u := NewUuid()
	require.Len(t, u.Val, 16)
	want, err := uuid.FromBytes(u.Val)
	require.NoError(t, err)
	assert.Equal(t, uuid.Version(4), want.Version())
	assert.Equal(t, want, u.ToUUID())
	assert.Equal(t, want, u.ToUUID(), "conversion must not generate a new UUID")
	assert.NotEqual(t, u.Val, NewUuid().Val)
}

func TestUuidBinaryRoundTrip(t *testing.T) {
	input := []byte{0, 17, 34, 51, 68, 85, 70, 119, 136, 153, 170, 187, 204, 221, 238, 255}
	var u Uuid
	require.NoError(t, u.UnmarshalBinary(input))
	assert.Equal(t, input, u.Val)
	encoded, err := u.MarshalBinary()
	require.NoError(t, err)
	assert.Equal(t, input, encoded)
	input[0] = 255
	assert.Equal(t, byte(0), u.Val[0], "unmarshal must own its bytes")
	for _, length := range []int{0, 15, 17} {
		before := append([]byte(nil), u.Val...)
		assert.Error(t, u.UnmarshalBinary(make([]byte, length)))
		assert.Equal(t, before, u.Val)
	}
}
