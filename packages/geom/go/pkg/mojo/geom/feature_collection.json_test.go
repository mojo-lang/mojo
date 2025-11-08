package geom

import (
	"github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	FeatureCollectionString = `{"type":"FeatureCollection","features":[{"type":"Feature","geometry":{"type":"Point","coordinates":[121.12,32.12]},"properties":{"key":"value"}},{"type":"Feature","geometry":{"type":"LineString","coordinates":[[121.12,32.12],[121.34,32.34]]},"properties":{"key":"value"}}]}`
	FeatureEmptyString      = `{"type":"FeatureCollection","features":[]}`
)

func TestUnmarshalFeatureCollection(t *testing.T) {
	fc := &FeatureCollection{}
	_ = jsoniter.ConfigDefault.UnmarshalFromString(FeatureCollectionString, fc)

	assert.Equal(t, 121.12, fc.Features[0].GetPoint().Coordinate.Longitude, "")
	assert.Equal(t, 32.12, fc.Features[0].GetPoint().Coordinate.Latitude, "")

	value, _ := fc.Features[0].GetStringProperty("key")
	assert.Equal(t, "value", value)

	assert.Equal(t, 32.12, fc.Features[1].GetLineString().Coordinates[0].Latitude, "")
	assert.Equal(t, 121.12, fc.Features[1].GetLineString().Coordinates[0].Longitude, "")
	assert.Equal(t, 32.34, fc.Features[1].GetLineString().Coordinates[1].Latitude, "")
	assert.Equal(t, 121.34, fc.Features[1].GetLineString().Coordinates[1].Longitude, "")

	value, _ = fc.Features[1].GetStringProperty("key")
	assert.Equal(t, "value", value)
}

func TestMarshalFeatureCollection(t *testing.T) {
	fc := NewFeatureCollection(
		NewPointFeatureFrom(121.12, 32.12).SetString("key", "value"),
		NewLineStringFeature(NewLngLat(121.12, 32.12), NewLngLat(121.34, 32.34)).SetString("key", "value"),
	)

	str, _ := jsoniter.ConfigDefault.MarshalToString(fc)
	assert.Equal(t, FeatureCollectionString, str, "")
}

func TestMarshalEmptyFeatureCollection(t *testing.T) {
	fc := NewFeatureCollection()
	str, _ := jsoniter.ConfigFastest.MarshalToString(fc)
	assert.Equal(t, FeatureEmptyString, str)
}
