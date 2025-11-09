package core

import (
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
)

func TestFloatValueCodec_Decode(t *testing.T) {
	data := &Float32Value{}
	d := `{"code":"200","message":"OK","data":false}`
	err := jsoniter.Unmarshal([]byte(d), data)
	assert.Error(t, err)

	err = jsoniter.Unmarshal([]byte(`100`), data)
	assert.NoError(t, err)
	assert.Equal(t, float32(100), data.Val)

	err = jsoniter.Unmarshal([]byte(`0`), data)
	assert.NoError(t, err)
	assert.Equal(t, float32(0), data.Val)
}
