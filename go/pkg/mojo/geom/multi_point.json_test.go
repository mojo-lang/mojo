package geom

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	multiPoints      = `{"type":"MultiPoint","coordinates":[[100,0],[101,1]]}`
	emptyMultiPoints = `{"type":"MultiPoint","coordinates":[]}`
)

func TestMultiPointCodec_Decode(t *testing.T) {
	mp := &MultiPoint{}

	err := jsoniter.UnmarshalFromString(multiPoints, mp)
	assert.Nil(t, err)

	assert.Equal(t, len(mp.Points), 2)
	assert.InDelta(t, mp.Points[0].Coordinate.Longitude, 100.0, 1e-5)
}

func TestMultiPointCodec_Encode(t *testing.T) {
	mp := &MultiPoint{}
	err := jsoniter.UnmarshalFromString(multiPoints, mp)
	assert.Nil(t, err)

	str, _ := jsoniter.MarshalToString(mp)
	assert.Equal(t, multiPoints, str)
}

func TestMultiPointCodec_EncodeEmpty(t *testing.T) {
	mp := &MultiPoint{}

	str, _ := jsoniter.MarshalToString(mp)
	assert.Equal(t, emptyMultiPoints, str)
}
