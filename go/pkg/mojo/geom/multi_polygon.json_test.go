package geom

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	multiPolygon      = `{"type":"MultiPolygon","coordinates":[[[[113.167297,23.522705],[113.166503,23.515686]],[[113.167297,23.522705],[113.167297,23.522705]]],[[[113.167297,23.522705],[113.166503,23.515686]],[[113.167297,23.522705],[113.167297,23.522705]]]]}`
	emptyMultiPolygon = `{"type":"MultiPolygon","coordinates":[]}`
)

func TestMultiPolygonCodec_Decode(t *testing.T) {
	mp := &MultiPolygon{}

	err := jsoniter.UnmarshalFromString(multiPolygon, mp)
	assert.Nil(t, err)

	assert.Equal(t, len(mp.Polygons), 2)
	assert.Equal(t, len(mp.Polygons[0].LineStrings), 2)
	assert.Equal(t, len(mp.Polygons[0].LineStrings[0].Coordinates), 2)

	str, _ := jsoniter.MarshalToString(mp)
	assert.Equal(t, multiPolygon, str)
}

func TestMultiPolygonCodec_EncodeEmpty(t *testing.T) {
	mp := &MultiPolygon{}

	str, _ := jsoniter.ConfigFastest.MarshalToString(mp)
	assert.Equal(t, emptyMultiPolygon, str)
}

func TestMultiPolygonCodec_Decode_Raw_Array(t *testing.T) {
	pstr := `[[[]]]`

	p := &MultiPolygon{}
	err := jsoniter.UnmarshalFromString(pstr, p)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(p.Polygons))
	assert.Equal(t, 1, len(p.Polygons[0].LineStrings))
	assert.Equal(t, 0, len(p.Polygons[0].LineStrings[0].Coordinates))

	pstr = `[[[132,22],[133,21],[134,20],[132,22]]]`
	p.Polygons = nil
	err = jsoniter.UnmarshalFromString(pstr, p)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(p.Polygons))

	pstr = `[[[[132,22],[133,21],[134,20],[132,22]]],[[[132,22],[133,21],[134,20],[132,22]]]]`
	p.Polygons = nil
	err = jsoniter.UnmarshalFromString(pstr, p)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(p.Polygons))
}
