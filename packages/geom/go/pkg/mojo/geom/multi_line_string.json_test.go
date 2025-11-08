package geom

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	multiLineString      = `{"type":"MultiLineString","coordinates":[[[170,45],[180,45]],[[-180,45],[-170,45]]]}`
	emptyMultiLineString = `{"type":"MultiLineString","coordinates":[]}`
)

func TestMultiLineStringCodec_Decode(t *testing.T) {
	ml := &MultiLineString{}

	err := jsoniter.UnmarshalFromString(multiLineString, ml)
	assert.Nil(t, err)

	assert.Equal(t, len(ml.LineStrings), 2)
	assert.Equal(t, len(ml.LineStrings[0].Coordinates), 2)

	assert.InDelta(t, ml.LineStrings[0].Coordinates[0].Longitude, 170.0, 1e-5)
}

func TestMultiLineStringCodec_Encode(t *testing.T) {
	ml := &MultiLineString{}
	err := jsoniter.UnmarshalFromString(multiLineString, ml)
	assert.Nil(t, err)

	str, _ := jsoniter.MarshalToString(ml)
	assert.Equal(t, multiLineString, str)
}

func TestMultiLineStringCodec_EncodeEmpty(t *testing.T) {
	ml := &MultiLineString{}

	str, _ := jsoniter.MarshalToString(ml)
	assert.Equal(t, emptyMultiLineString, str)
}
