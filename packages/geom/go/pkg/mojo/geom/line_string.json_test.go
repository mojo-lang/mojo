package geom

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	emptyLineString = `{"type":"LineString","coordinates":[]}`
)

func TestLineStringCodec_Encode(t *testing.T) {
	l := &LineString{}

	str, _ := jsoniter.MarshalToString(l)
	assert.Equal(t, emptyLineString, str)
}

func TestLineStringCodec_Encode_2(t *testing.T) {
	ls := NewLineString(NewLngLat(116.898699, 36.678687), NewLngLat(116.910034, 36.671236))
	str, err := jsoniter.MarshalToString(ls)
	assert.NoError(t, err)
	assert.Equal(t, `{"type":"LineString","coordinates":[[116.898699,36.678687],[116.910034,36.671236]]}`, str)
}
