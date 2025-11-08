package planar

import (
	"context"

	"github.com/mojo-lang/mojo/packages/geom/go/pkg/mojo/geom"
)

func simplifyPolygon(ctx context.Context, simplifier Simplifier, plg []*geom.LineString, isClosed bool) ([]*geom.LineString, error) {
	var lines []*geom.LineString
	for _, l := range plg {
		points := make([][2]float64, len(l.Coordinates))
		for i, coord := range l.Coordinates {
			points[i] = [2]float64{coord.Longitude, coord.Latitude}
		}

		ls, err := simplifier.Simplify(ctx, points, isClosed)
		if err != nil {
			return nil, err
		}
		if len(ls) > 2 || !isClosed {
			var ps []*geom.LngLat
			for _, p := range ls {
				ps = append(ps, geom.NewLngLat(p[0], p[1]))
			}
			lines = append(lines, geom.NewLineString(ps...))
		}
	}
	return lines, nil
}

func SimplifyLineString(ctx context.Context, simplifier Simplifier, linestring *geom.LineString) (*geom.LineString, error) {
	ls, err := simplifyPolygon(ctx, simplifier, []*geom.LineString{linestring}, false)
	if err != nil {
		return nil, err
	}
	return ls[0], nil
}

func SimplifyMultiLineString(ctx context.Context, simplifier Simplifier, multiLinestring *geom.MultiLineString) (*geom.MultiLineString, error) {
	ls, err := simplifyPolygon(ctx, simplifier, multiLinestring.LineStrings, false)
	if err != nil {
		return nil, err
	}
	return geom.NewMultiLineString(ls...), nil
}

func SimplifyPolygon(ctx context.Context, simplifier Simplifier, polygon *geom.Polygon) (*geom.Polygon, error) {
	ls, err := simplifyPolygon(ctx, simplifier, polygon.LineStrings, true)
	if err != nil {
		return nil, err
	}
	return geom.NewPolygon(ls...), nil
}

func SimplifyMultiPolygon(ctx context.Context, simplifier Simplifier, multiPolygon *geom.MultiPolygon) (*geom.MultiPolygon, error) {
	ps := make([]*geom.Polygon, len(multiPolygon.Polygons))
	for i, p := range multiPolygon.Polygons {
		polygon, err := SimplifyPolygon(ctx, simplifier, p)
		if err != nil {
			return nil, err
		}
		ps[i] = polygon
	}
	return geom.NewMultiPolygon(ps...), nil
}

// Simplify will simplify the provided geometry using the provided simplifier.
// If the simplifier is nil, no simplification will be attempted.
func Simplify(ctx context.Context, simplifier Simplifier, geometry *geom.Geometry) (*geom.Geometry, error) {
	if simplifier == nil {
		return geometry, nil
	}

	switch g := geometry.Geometry.(type) {
	case *geom.Geometry_GeometryCollection:
		geos := g.GeometryCollection.Geometries
		coll := make([]*geom.Geometry, len(geos))
		for i := range geos {
			geo, err := Simplify(ctx, simplifier, geos[i])
			if err != nil {
				return nil, err
			}
			coll[i] = geo
		}
		return geom.NewGeometryCollectionGeometry(coll...), nil
	case *geom.Geometry_MultiPolygon:
		mp, err := SimplifyMultiPolygon(ctx, simplifier, g.MultiPolygon)
		if err != nil {
			return nil, err
		}
		return geom.NewMultiPolygonGeometry(mp.Polygons...), nil
	case *geom.Geometry_Polygon:
		p, err := SimplifyPolygon(ctx, simplifier, g.Polygon)
		if err != nil {
			return nil, err
		}
		return geom.NewPolygonGeometry(p.LineStrings...), nil
	case *geom.Geometry_MultiLineString:
		ml, err := SimplifyMultiLineString(ctx, simplifier, g.MultiLineString)
		if err != nil {
			return nil, err
		}
		return geom.NewMultiLineStringGeometry(ml.LineStrings...), nil
	case *geom.Geometry_LineString:
		ls, err := SimplifyLineString(ctx, simplifier, g.LineString)
		if err != nil {
			return nil, err
		}
		return geom.NewLineStringGeometry(ls.Coordinates...), nil
	default: // Points, MultiPoints or anything else.
		return geometry, nil
	}
}
