package geom

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	emptyPoint = `{"type":"Point","coordinates":[]}`
)

func TestPointCodec_EncodeEmpty(t *testing.T) {
	p := &Point{}

	str, _ := jsoniter.ConfigFastest.MarshalToString(p)
	assert.Equal(t, emptyPoint, str)
}
