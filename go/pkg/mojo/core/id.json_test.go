package core

import (
	"testing"
	"unsafe"

	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIdJSON(t *testing.T) {
	for _, tt := range []struct{ input, output string }{
		{`42`, `"42"`}, {`18446744073709551615`, `"18446744073709551615"`},
		{`"invoice-42"`, `"invoice-42"`}, {`""`, `""`},
		{`"00112233-4455-4677-8899-aabbccddeeff"`, `"00112233-4455-4677-8899-aabbccddeeff"`},
	} {
		t.Run(tt.input, func(t *testing.T) {
			var id Id
			require.NoError(t, jsoniter.UnmarshalFromString(tt.input, &id))
			encoded, err := jsoniter.MarshalToString(&id)
			require.NoError(t, err)
			assert.Equal(t, tt.output, encoded)
			decoded, err := NewIdCodec().DecodeAny(jsoniter.Get([]byte(tt.input)))
			require.NoError(t, err)
			assert.Equal(t, id.Format(), decoded.Format())
		})
	}
	for _, input := range []string{`true`, `null`, `{}`, `[]`, `-1`, `1.5`, `18446744073709551616`} {
		t.Run("invalid/"+input, func(t *testing.T) {
			var id Id
			assert.Error(t, jsoniter.UnmarshalFromString(input, &id))
			_, err := NewIdCodec().DecodeAny(jsoniter.Get([]byte(input)))
			assert.Error(t, err)
		})
	}
}

func TestIdCodecIsEmpty(t *testing.T) {
	codec := NewIdCodec()
	assert.True(t, codec.IsEmpty(nil))
	assert.True(t, codec.IsEmpty(unsafe.Pointer(&Id{})))
	assert.False(t, codec.IsEmpty(unsafe.Pointer(NewIntId(0))))
	assert.False(t, codec.IsEmpty(unsafe.Pointer(NewStringId(""))))
}
