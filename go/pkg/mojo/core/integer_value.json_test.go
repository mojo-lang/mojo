package core

import (
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
)

func TestInt32ValueCodec_Decode(t *testing.T) {
	data := &Int32Value{}
	d := `{"code":"200","message":"OK","data":false}`
	err := jsoniter.Unmarshal([]byte(d), data)
	assert.Error(t, err)

	err = jsoniter.Unmarshal([]byte(`100`), data)
	assert.NoError(t, err)
	assert.Equal(t, int32(100), data.Val)

	err = jsoniter.Unmarshal([]byte(`0`), data)
	assert.NoError(t, err)
	assert.Equal(t, int32(0), data.Val)
}
