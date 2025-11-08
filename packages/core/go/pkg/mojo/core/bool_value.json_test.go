package core

import (
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
)

func TestBoolValueDecode(t *testing.T) {
	data := &BoolValue{}
	d := `{"code":"200","message":"OK","data":false}`
	err := jsoniter.Unmarshal([]byte(d), data)
	assert.Error(t, err)

	err = jsoniter.Unmarshal([]byte(`true`), data)
	assert.NoError(t, err)
	assert.True(t, data.Val)

	err = jsoniter.Unmarshal([]byte(`false`), data)
	assert.NoError(t, err)
	assert.False(t, data.Val)
}
