package core

import (
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
)

func TestDurationCodec(t *testing.T) {
	duration := &Duration{Seconds: 10, Nanoseconds: 1000}
	str, err := jsoniter.ConfigFastest.Marshal(duration)
	assert.NoError(t, err)
	assert.Equal(t, "\"10.000001s\"", string(str))

	d := &Duration{}
	e := jsoniter.ConfigFastest.Unmarshal(str, d)
	assert.NoError(t, e)
	assert.Equal(t, int64(10), d.Seconds)
	assert.Equal(t, int32(1000), d.Nanoseconds)
}
