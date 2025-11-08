package geom

import (
	"github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
	"testing"
)

const PointString = `{"type":"Point","coordinates":[121.12,32.12]}`
const LineStringString = `{"type":"LineString","coordinates":[[121.12,32.12],[121.34,32.34]]}`
const PolygonString = `{"type":"Polygon","coordinates":[[[121.12,32.12],[121.34,32.34],[121.12,32.12]]]}`
const MultiPolygonString = `{"type":"MultiPolygon","coordinates":[[[[120,30],[120.1,30.1],[120.2,30.2],[120,30]]],[[[120.3,30.3],[120.4,30.4],[120.5,30.5],[120.3,30.3]]]]}`
const GeometryCollectionString = `{"type":"GeometryCollection","geometries":[{"type":"Point","coordinates":[121.12,32.12]}]}`

func TestUnmarshalPoint(t *testing.T) {
	geometry := &Geometry{}
	_ = jsoniter.ConfigDefault.UnmarshalFromString(PointString, geometry)
	assert.Equal(t, 121.12, geometry.GetPoint().Coordinate.Longitude, "")
	assert.Equal(t, 32.12, geometry.GetPoint().Coordinate.Latitude, "")
}

func TestMarshalPoint(t *testing.T) {
	geometry := NewGeometry(&Point{Coordinate: &LngLat{Longitude: 121.12, Latitude: 32.12}})

	str, _ := jsoniter.ConfigDefault.MarshalToString(geometry)
	assert.Equal(t, PointString, str)
}

func TestUnmarshalLineString(t *testing.T) {
	geometry := &Geometry{}
	_ = jsoniter.ConfigDefault.UnmarshalFromString(LineStringString, geometry)
	assert.Equal(t, 121.12, geometry.GetLineString().Coordinates[0].Longitude, "")
	assert.Equal(t, 32.12, geometry.GetLineString().Coordinates[0].Latitude, "")
}

func TestMarshalLineString(t *testing.T) {
	geometry := NewGeometry(&LineString{Coordinates: []*LngLat{{Longitude: 121.12, Latitude: 32.12}, {Longitude: 121.34, Latitude: 32.34}}})

	str, _ := jsoniter.ConfigDefault.MarshalToString(geometry)
	assert.Equal(t, LineStringString, str)
}

func TestUnmarshalPolygon(t *testing.T) {
	geometry := &Geometry{}
	_ = jsoniter.ConfigDefault.UnmarshalFromString(PolygonString, geometry)
	assert.Equal(t, 121.12, geometry.GetPolygon().LineStrings[0].Coordinates[0].Longitude, "")
	assert.Equal(t, 32.12, geometry.GetPolygon().LineStrings[0].Coordinates[0].Latitude, "")
}

func TestMarshalPolygon(t *testing.T) {
	geometry := NewGeometry(&Polygon{LineStrings: []*LineString{{Coordinates: []*LngLat{{Longitude: 121.12, Latitude: 32.12}, {Longitude: 121.34, Latitude: 32.34}, {Longitude: 121.12, Latitude: 32.12}}}}})

	str, _ := jsoniter.ConfigDefault.MarshalToString(geometry)
	assert.Equal(t, PolygonString, str)
}

func TestUnmarshalGeometryCollection(t *testing.T) {
	geometry := &Geometry{}
	_ = jsoniter.ConfigDefault.UnmarshalFromString(GeometryCollectionString, geometry)
	assert.Equal(t, 121.12, geometry.GetGeometryCollection().Geometries[0].GetPoint().Coordinate.Longitude, "")
	assert.Equal(t, 32.12, geometry.GetGeometryCollection().Geometries[0].GetPoint().Coordinate.Latitude, "")
}

func TestMarshalGeometryCollection(t *testing.T) {
	geometry := NewGeometry(&GeometryCollection{Geometries: []*Geometry{NewPointGeometryFrom(121.12, 32.12)}})
	str, _ := jsoniter.ConfigDefault.MarshalToString(geometry)
	assert.Equal(t, GeometryCollectionString, str)
}

func TestNewGeometryWithMultiPolygon(t *testing.T) {
	coordinates := GenerateGeometry()
	s, _ := jsoniter.ConfigDefault.MarshalToString(coordinates)
	assert.Equal(t, MultiPolygonString, s)
}

func TestGeometry_Type(t *testing.T) {
	{
		geometry := NewGeometry(&Point{Coordinate: &LngLat{Longitude: 121.12, Latitude: 32.12}})
		assert.Equal(t, geometry.Type(), PointType)
	}

	{
		geometry := NewGeometry(&LineString{Coordinates: []*LngLat{{Longitude: 121.12, Latitude: 32.12}, {Longitude: 121.34, Latitude: 32.34}}})
		assert.Equal(t, geometry.Type(), LineStringType)
	}

	{
		geometry := NewGeometry(&Polygon{LineStrings: []*LineString{{Coordinates: []*LngLat{{Longitude: 121.12, Latitude: 32.12}, {Longitude: 121.34, Latitude: 32.34}, {Longitude: 121.12, Latitude: 32.12}}}}})
		assert.Equal(t, geometry.Type(), PolygonType)
	}
}

func GenerateGeometry() *Geometry {
	lls1 := make([]*LngLat, 0, 4)
	lls1 = append(lls1, &LngLat{
		Longitude: 120,
		Latitude:  30,
	})
	lls1 = append(lls1, &LngLat{
		Longitude: 120.1,
		Latitude:  30.1,
	})
	lls1 = append(lls1, &LngLat{
		Longitude: 120.2,
		Latitude:  30.2,
	})
	lls2 := make([]*LngLat, 0, 4)
	lls2 = append(lls2, &LngLat{
		Longitude: 120.3,
		Latitude:  30.3,
	})
	lls2 = append(lls2, &LngLat{
		Longitude: 120.4,
		Latitude:  30.4,
	})
	lls2 = append(lls2, &LngLat{
		Longitude: 120.5,
		Latitude:  30.5,
	})
	coordinates := NewMultiPolygonGeometry(NewPolygonFrom(lls1...), NewPolygonFrom(lls2...))
	return coordinates
}
