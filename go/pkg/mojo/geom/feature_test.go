package geom

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
	"testing"
)

const PointFeatureString = `{"type":"Feature","geometry":{"type":"Point","coordinates":[121.12,32.12]},"properties":{"key":"value"}}`
const LineStringFeatureString = `{"type":"Feature","geometry":{"type":"LineString","coordinates":[[121.12,32.12],[121.34,32.34]]},"properties":{"key":"value"}}`
const PolygonFeatureString = `{"type":"Feature","geometry":{"type":"Polygon","coordinates":[[[121.12,32.12],[121.34,32.34],[121.12,32.12]]]},"properties":{"key":"value"}}`
const NilValuePropertyFeatureString = `{"type":"Feature","geometry":{"type":"LineString","coordinates":[[27762.55375,-30728.908429]]},"properties":{"EntityHand":"32","ExtendedEn":null,"Layer":"0","Linetype":null,"SubClasses":"AcDbEntity:AcDbPolyline","Text":null}}`
const NilGeometryFeatureString = `{"type":"Feature","geometry":null,"properties":null}`
const NilGeometryFeatureStringV2 = `{"type":"Feature","geometry":null,"properties":{}}`

func TestUnmarshalPointFeature(t *testing.T) {
	feature := &Feature{}
	_ = jsoniter.ConfigDefault.UnmarshalFromString(PointFeatureString, feature)
	assert.Equal(t, 121.12, feature.GetPoint().Coordinate.Longitude, "")
	assert.Equal(t, 32.12, feature.GetPoint().Coordinate.Latitude, "")

	value, _ := feature.GetStringProperty("key")
	assert.Equal(t, "value",value)
}

func TestUnmarshalPropertyFeature(t *testing.T) {
	feature := &Feature{}
	_ = jsoniter.ConfigDefault.UnmarshalFromString(NilValuePropertyFeatureString, feature)
	assert.Equal(t, 27762.55375, feature.GetLineString().Coordinates[0].Longitude, "")
	assert.Equal(t, -30728.908429, feature.GetLineString().Coordinates[0].Latitude, "")

	value, _ := feature.GetStringProperty("Layer")
	assert.Equal(t, "0", value)

	out, err := jsoniter.ConfigFastest.Marshal(feature)
	assert.NoError(t, err)
	assert.Equal(t, NilValuePropertyFeatureString, string(out))

	_ = jsoniter.ConfigDefault.UnmarshalFromString(NilGeometryFeatureString, feature)
	out, err = jsoniter.ConfigFastest.Marshal(feature)
	assert.NoError(t, err)
	assert.Equal(t, NilGeometryFeatureString, string(out))

	_ = jsoniter.ConfigDefault.UnmarshalFromString(NilGeometryFeatureStringV2, feature)
	out, err = jsoniter.ConfigFastest.Marshal(feature)
	assert.NoError(t, err)
	assert.Equal(t, NilGeometryFeatureString, string(out))
}

func TestMarshalPointFeature(t *testing.T) {
	feature := NewPointFeatureFrom(121.12, 32.12)
	feature.SetString("key", "value")

	str, _ := jsoniter.ConfigDefault.MarshalToString(feature)
	assert.Equal(t, PointFeatureString, str, "")
}

func TestUnmarshalLineStringFeature(t *testing.T) {
	feature := &Feature{}
	_ = jsoniter.ConfigDefault.UnmarshalFromString(LineStringFeatureString, feature)
	assert.Equal(t, 32.12, feature.GetLineString().Coordinates[0].Latitude, "")
	assert.Equal(t, 121.12, feature.GetLineString().Coordinates[0].Longitude, "")
	assert.Equal(t, 32.34, feature.GetLineString().Coordinates[1].Latitude, "")
	assert.Equal(t, 121.34, feature.GetLineString().Coordinates[1].Longitude, "")

	value, _ := feature.GetStringProperty("key")
	assert.Equal(t, "value", value)
}

func TestMarshalLineStringFeature(t *testing.T) {
	feature := NewLineStringFeature(NewLngLat(121.12, 32.12), NewLngLat(121.34, 32.34))
	feature.SetString("key", "value")

	str, _ := jsoniter.ConfigDefault.MarshalToString(feature)
	assert.Equal(t, LineStringFeatureString, str, "")
}
