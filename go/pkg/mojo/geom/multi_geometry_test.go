package geom

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMultiGeometryBoundsAndEquality(t *testing.T) {
	a, b := NewLngLat(-10, -20), NewLngLat(30, 40)
	box := NewBoundingBox(a, b)
	points := NewMultiPointFrom(a, b)
	require.True(t, box.Equal(points.BoundingBox()))
	require.Equal(t, []*LngLat{a, b}, points.GetLngLats())
	require.True(t, points.Equal(NewMultiPoint(NewPoint(a), NewPoint(b))))
	require.False(t, points.Equal(NewMultiPointFrom(b, a)))
	require.False(t, points.Equal(NewMultiPointFrom(a)))
	require.Same(t, points, points.ToGeometry().GetMultiPoint())
	lines := NewMultiLineStringFrom([]*LngLat{a}, []*LngLat{b})
	require.True(t, box.Equal(lines.BoundingBox()))
	require.True(t, lines.Equal(NewMultiLineString(NewLineString(a), NewLineString(b))))
	require.False(t, lines.Equal(NewMultiLineString(NewLineString(b), NewLineString(a))))
	require.False(t, lines.Equal(NewMultiLineString()))
	require.Same(t, lines, lines.ToGeometry().GetMultiLineString())
	polygon := box.ToPolygon()
	polygons := NewMultiPolygon(polygon)
	require.True(t, box.Equal(polygons.BoundingBox()))
	require.True(t, polygons.Equal(NewMultiPolygon(polygon)))
	require.False(t, polygons.Equal(NewMultiPolygon()))
	require.InDelta(t, polygon.Area(), polygons.Area(), 1e-6)
	require.Same(t, polygons, polygons.ToGeometry().GetMultiPolygon())
}

func TestEmptyMultiGeometries(t *testing.T) {
	var points *MultiPoint
	var lines *MultiLineString
	var polygons *MultiPolygon
	require.True(t, points.Equal(nil))
	require.True(t, lines.Equal(nil))
	require.True(t, polygons.Equal(nil))
	require.Nil(t, points.ToGeometry())
	require.Nil(t, lines.ToGeometry())
	require.Nil(t, polygons.ToGeometry())
	require.Nil(t, points.BoundingBox())
	require.Nil(t, lines.BoundingBox())
	require.Nil(t, polygons.BoundingBox())
	require.Nil(t, NewMultiPoint().BoundingBox())
	require.Nil(t, NewMultiLineString(NewLineString()).BoundingBox())
	require.Nil(t, NewMultiPolygon(&Polygon{}).BoundingBox())
	require.Zero(t, polygons.Area())
}

func TestGeometryCollectionAggregatesBounds(t *testing.T) {
	first := NewPoint(NewLngLat(-10, 20)).ToGeometry()
	second := NewLineString(NewLngLat(30, -5), NewLngLat(50, 40)).ToGeometry()
	collection := NewGeometryCollection((&LineString{}).ToGeometry(), first)
	collection.AddGeometries(NewGeometryCollection(second).ToGeometry())
	box := collection.BoundingBox()
	require.NotNil(t, box)
	require.True(t, box.Equal(NewBoundingBox(NewLngLat(-10, -5), NewLngLat(50, 40))))
	require.True(t, box.Equal(collection.ToGeometry().BoundingBox()))
	require.Nil(t, NewGeometryCollection().BoundingBox())
}
