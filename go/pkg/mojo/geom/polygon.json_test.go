package geom

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPolygonCodec_Codec(t *testing.T) {
	p := &Polygon{}

	err := jsoniter.UnmarshalFromString(polygonStr, p)
	assert.Nil(t, err)

	assert.Equal(t, 2, len(p.LineStrings))
	assert.Equal(t, 3, len(p.LineStrings[0].Coordinates))

	str, _ := jsoniter.MarshalToString(p)
	assert.Equal(t, polygonStr, str)
}

func TestPolygonCodec_Decode_Raw_Array(t *testing.T) {
	pstr := `[[]]`

	p := &Polygon{}
	err := jsoniter.UnmarshalFromString(pstr, p)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(p.LineStrings))
	assert.Equal(t, 0, len(p.LineStrings[0].Coordinates))

	pstr = `[[132,22],[133,21],[134,20],[132,22]]`
	p.LineStrings = nil
	err = jsoniter.UnmarshalFromString(pstr, p)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(p.LineStrings))

	pstr = `[[[132,22],[133,21],[134,20],[132,22]],[[132,22],[133,21],[134,20],[132,22]]]`
	p.LineStrings = nil
	err = jsoniter.UnmarshalFromString(pstr, p)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(p.LineStrings))
}
