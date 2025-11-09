package core

import (
	"encoding/json"
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
)

const anyStr = `{"@type":"mojo.core.Error","code":"404","message":"something wrong"}`

var anyError = &Error{Code: NotFound, Message: "something wrong"}

const checksumStr = `{"@type":"mojo.core.Checksum","value":"SHA256:cd42404d52ad55ccfa9aca4adc828aa5800ad9d385a0671fbcbf724118320619"}`

var anyChecksum = NewChecksum(Checksum_ALGORITHM_SHA256, []byte("value"))

func TestAnyCodec_Decode(t *testing.T) {
	any := &Any{}
	err := jsoniter.ConfigFastest.UnmarshalFromString(anyStr, any)

	assert.NoError(t, err)
	assert.Equal(t, anyError.Code.Code, any.Get().(*Error).Code.Code)
	assert.Equal(t, anyError.Message, any.Get().(*Error).Message)
}

func TestAnyCodec_Decode2(t *testing.T) {
	any := &Any{}
	err := json.Unmarshal([]byte(anyStr), any)

	assert.NoError(t, err)
	assert.Equal(t, anyError.Code.Code, any.Get().(*Error).Code.Code)
	assert.Equal(t, anyError.Message, any.Get().(*Error).Message)
}

func TestAnyCodec_Decode3(t *testing.T) {
	any := &Any{}
	err := json.Unmarshal([]byte(checksumStr), any)

	assert.NoError(t, err)
	assert.Equal(t, anyChecksum.Algorithm, any.Get().(*Checksum).Algorithm)
	assert.Equal(t, anyChecksum.Val, any.Get().(*Checksum).Val)
}

func TestAnyCodec_Encode(t *testing.T) {
	any := NewAny(anyError)
	out, err := jsoniter.ConfigFastest.MarshalToString(any)

	assert.NoError(t, err)
	assert.Equal(t, anyStr, out)
}

func TestAnyCodec_Encode2(t *testing.T) {
	any := NewAny(anyError)
	out, err := json.Marshal(any)

	assert.NoError(t, err)
	assert.Equal(t, anyStr, string(out))
}

func TestAnyCodec_Encode3(t *testing.T) {
	any := NewAny(anyChecksum)
	out, err := jsoniter.ConfigFastest.MarshalToString(any)

	assert.NoError(t, err)
	assert.Equal(t, checksumStr, out)
}
