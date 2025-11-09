package geom

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var (
	wkt1 = `GEOMETRYCOLLECTION (POINT (120.123456 30.654321),` +
		`MULTIPOINT (120.123456 30.654321,60.123456 15.654321),` +
		`LINESTRING (120.123456 30.654321,60.123456 15.654321),` +
		`MULTILINESTRING ((120.123456 30.654321,60.123456 15.654321),(120.123456 30.654321,60.123456 15.654321)),` +
		`POLYGON ((120.123456 30.123456,120.234567 30.123456,120.234567 30.234567,120.123456 30.234567,120.123456 30.123456)),` +
		`POLYGON ((120.123456 30.123456,120.234567 30.123456,120.234567 30.234567,120.123456 30.234567,120.123456 30.123456),` +
		`(120.123456 30.123556,120.123456 30.234467,120.234467 30.234467,120.234467 30.123556,120.123456 30.123556)),` +
		`MULTIPOLYGON (((120.123456 30.123456,120.234567 30.123456,120.234567 30.234567,120.123456 30.234567,120.123456 30.123456)),` +
		`((120.123456 30.123456,120.234567 30.123456,120.234567 30.234567,120.123456 30.234567,120.123456 30.123456),` +
		`(120.123456 30.123556,120.123456 30.234467,120.234467 30.234467,120.234467 30.123556,120.123456 30.123556))))`

	wkt2 = `GEOMETRYCOLLECTION (
	POINT (120.123456 30.654321),
	MULTIPOINT (120.123456 30.654321,60.123456 15.654321),
	LINESTRING (120.123456 30.654321,60.123456 15.654321),
	MULTILINESTRING (
		(120.123456 30.654321,60.123456 15.654321),
		(120.123456 30.654321,60.123456 15.654321)
	),
	POLYGON (
		(120.123456 30.123456,120.234567 30.123456,120.234567 30.234567,120.123456 30.234567,120.123456 30.123456)
	),
	POLYGON (
		(120.123456 30.123456,120.234567 30.123456,120.234567 30.234567,120.123456 30.234567,120.123456 30.123456),
		(120.123456 30.123556,120.123456 30.234467,120.234467 30.234467,120.234467 30.123556,120.123456 30.123556)
	),
	MULTIPOLYGON (
		(
			(120.123456 30.123456,120.234567 30.123456,120.234567 30.234567,120.123456 30.234567,120.123456 30.123456)
		),
		(
			(120.123456 30.123456,120.234567 30.123456,120.234567 30.234567,120.123456 30.234567,120.123456 30.123456),
			(120.123456 30.123556,120.123456 30.234467,120.234467 30.234467,120.234467 30.123556,120.123456 30.123556)
		)
	)
)`
)

func TestWKTEncode(t *testing.T) {
	p := NewPointGeometryFrom(120.123456, 30.654321)
	multiPoint := NewMultiPointGeometryFrom(
		&LngLat{
			Longitude: 120.123456,
			Latitude:  30.654321,
		},
		&LngLat{
			Longitude: 60.123456,
			Latitude:  15.654321,
		},
	)
	line := NewLineStringGeometry(
		&LngLat{
			Longitude: 120.123456,
			Latitude:  30.654321,
		},
		&LngLat{
			Longitude: 60.123456,
			Latitude:  15.654321,
		},
	)
	multiLine := NewMultiLineStringGeometry(
		&LineString{
			Coordinates: []*LngLat{
				{
					Longitude: 120.123456,
					Latitude:  30.654321,
				},
				{
					Longitude: 60.123456,
					Latitude:  15.654321,
				},
			},
		},
		&LineString{
			Coordinates: []*LngLat{
				{
					Longitude: 120.123456,
					Latitude:  30.654321,
				},
				{
					Longitude: 60.123456,
					Latitude:  15.654321,
				},
			},
		},
	)
	polygon1 := &Polygon{
		LineStrings: []*LineString{
			{
				Coordinates: []*LngLat{
					{
						Longitude: 120.123456,
						Latitude:  30.123456,
					},
					{
						Longitude: 120.234567,
						Latitude:  30.123456,
					},
					{
						Longitude: 120.234567,
						Latitude:  30.234567,
					},
					{
						Longitude: 120.123456,
						Latitude:  30.234567,
					},
					{
						Longitude: 120.123456,
						Latitude:  30.123456,
					},
				},
			},
		},
	}
	poly1 := &Geometry{}
	_ = poly1.SetValue(polygon1)

	polygon2 := &Polygon{
		LineStrings: []*LineString{
			{
				Coordinates: []*LngLat{
					{
						Longitude: 120.123456,
						Latitude:  30.123456,
					},
					{
						Longitude: 120.234567,
						Latitude:  30.123456,
					},
					{
						Longitude: 120.234567,
						Latitude:  30.234567,
					},
					{
						Longitude: 120.123456,
						Latitude:  30.234567,
					},
					{
						Longitude: 120.123456,
						Latitude:  30.123456,
					},
				},
			},
			{
				Coordinates: []*LngLat{
					{
						Longitude: 120.123456,
						Latitude:  30.123556,
					},
					{
						Longitude: 120.123456,
						Latitude:  30.234467,
					},
					{
						Longitude: 120.234467,
						Latitude:  30.234467,
					},
					{
						Longitude: 120.234467,
						Latitude:  30.123556,
					},
					{
						Longitude: 120.123456,
						Latitude:  30.123556,
					},
				},
			},
		},
	}
	poly2 := &Geometry{}
	_ = poly2.SetValue(polygon2)

	mPolygon := NewMultiPolygonGeometry(polygon1, polygon2)

	collection := NewGeometryCollectionGeometry(p, multiPoint, line, multiLine, poly1, poly2, mPolygon)

	w := collection.ToWKT()
	assert.Equal(t, len(wkt1), len(w))
	assert.Equal(t, wkt1, w)
}

func TestWKTDecode(t *testing.T) {
	g, err := NewGeometryFromWKT(wkt2)
	if err != nil {
		t.Error(err)
		return
	}

	w := g.ToWKT()
	wkt2 = strings.ReplaceAll(wkt2, "\n", "")
	wkt2 = strings.ReplaceAll(wkt2, "\t", "")
	assert.Equal(t, wkt2, w)
}
