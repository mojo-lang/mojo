package geom

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCircleCodec_Decode(t *testing.T) {
	c := &Circle{}
	cstr := `"131.122,23.11,100"`
	err := jsoniter.UnmarshalFromString(cstr, c)
	assert.Nil(t, err)
	assert.NotNil(t, c.Center)
	assert.Equal(t, float32(100), c.Radius)

	cstr = `{"type":"Feature","geometry":{"type":"Point","coordinates":[131.122,23.11]},"properties":{"subType":"Circle","radius":100}}`
	c.Reset()
	err = jsoniter.UnmarshalFromString(cstr, c)
	assert.Nil(t, err)
	assert.NotNil(t, c.Center)
	assert.Equal(t, float32(100), c.Radius)
}
