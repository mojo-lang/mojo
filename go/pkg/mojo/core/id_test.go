package core

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIdConstructors(t *testing.T) {
	u, err := ParseUuid("00112233-4455-4677-8899-aabbccddeeff")
	require.NoError(t, err)
	for _, tt := range []struct {
		name  string
		input interface{}
		want  *Id
	}{
		{"integer", uint64(42), NewIntId(42)},
		{"string", "invoice-42", NewStringId("invoice-42")},
		{"uuid", u, NewUuidId(u)},
		{"unsupported", true, nil},
		{"nil", nil, nil},
	} {
		t.Run(tt.name, func(t *testing.T) { assert.Equal(t, tt.want, NewId(tt.input)) })
	}
}

func TestIdSettersAndString(t *testing.T) {
	id := &Id{}
	id.SetInt(math.MaxUint64)
	assert.Equal(t, uint64(math.MaxUint64), id.GetUint64Val())
	assert.Equal(t, "18446744073709551615", id.GetString())
	id.SetString("order-1")
	assert.Equal(t, "order-1", id.GetString())
	assert.Zero(t, id.GetUint64Val())
	id.SetUuidString("00112233-4455-4677-8899-aabbccddeeff")
	u := id.GetUuidVal()
	require.NotNil(t, u)
	assert.Len(t, u.Val, 16)
	id.SetUuidString("invalid")
	assert.Same(t, u, id.GetUuidVal(), "invalid input must preserve the previous ID")
	id.RegenerateUuid()
	require.Len(t, id.GetUuidVal().Val, 16)
	assert.NotEqual(t, u.Val, id.GetUuidVal().Val)
	require.Len(t, UuidId().GetUuidVal().Val, 16)
	assert.Empty(t, (&Id{}).GetString())
	assert.Empty(t, (*Id)(nil).GetString())
}
