package planar

import (
	"context"
	geom2 "github.com/mojo-lang/mojo/go/pkg/mojo/geom"
)

func simplifyPolygon(ctx context.Context, simplifier Simplifier, plg []*geom2.LineString, isClosed bool) ([]*geom2.LineString, error) {
	var lines []*geom2.LineString
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
			var ps []*geom2.LngLat
			for _, p := range ls {
				ps = append(ps, geom2.NewLngLat(p[0], p[1]))
			}
			lines = append(lines, geom2.NewLineString(ps...))
		}
	}
	return lines, nil
}

func SimplifyLineString(ctx context.Context, simplifier Simplifier, linestring *geom2.LineString) (*geom2.LineString, error) {
	ls, err := simplifyPolygon(ctx, simplifier, []*geom2.LineString{linestring}, false)
	if err != nil {
		return nil, err
	}
	return ls[0], nil
}

func SimplifyMultiLineString(ctx context.Context, simplifier Simplifier, multiLinestring *geom2.MultiLineString) (*geom2.MultiLineString, error) {
	ls, err := simplifyPolygon(ctx, simplifier, multiLinestring.LineStrings, false)
	if err != nil {
		return nil, err
	}
	return geom2.NewMultiLineString(ls...), nil
}

func SimplifyPolygon(ctx context.Context, simplifier Simplifier, polygon *geom2.Polygon) (*geom2.Polygon, error) {
	ls, err := simplifyPolygon(ctx, simplifier, polygon.LineStrings, true)
	if err != nil {
		return nil, err
	}
	return geom2.NewPolygon(ls...), nil
}

func SimplifyMultiPolygon(ctx context.Context, simplifier Simplifier, multiPolygon *geom2.MultiPolygon) (*geom2.MultiPolygon, error) {
	ps := make([]*geom2.Polygon, len(multiPolygon.Polygons))
	for i, p := range multiPolygon.Polygons {
		polygon, err := SimplifyPolygon(ctx, simplifier, p)
		if err != nil {
			return nil, err
		}
		ps[i] = polygon
	}
	return geom2.NewMultiPolygon(ps...), nil
}

// Simplify will simplify the provided geometry using the provided simplifier.
// If the simplifier is nil, no simplification will be attempted.
func Simplify(ctx context.Context, simplifier Simplifier, geometry *geom2.Geometry) (*geom2.Geometry, error) {
	if simplifier == nil {
		return geometry, nil
	}

	switch g := geometry.Geometry.(type) {
	case *geom2.Geometry_GeometryCollection:
		geos := g.GeometryCollection.Geometries
		coll := make([]*geom2.Geometry, len(geos))
		for i := range geos {
			geo, err := Simplify(ctx, simplifier, geos[i])
			if err != nil {
				return nil, err
			}
			coll[i] = geo
		}
		return geom2.NewGeometryCollectionGeometry(coll...), nil
	case *geom2.Geometry_MultiPolygon:
		mp, err := SimplifyMultiPolygon(ctx, simplifier, g.MultiPolygon)
		if err != nil {
			return nil, err
		}
		return geom2.NewMultiPolygonGeometry(mp.Polygons...), nil
	case *geom2.Geometry_Polygon:
		p, err := SimplifyPolygon(ctx, simplifier, g.Polygon)
		if err != nil {
			return nil, err
		}
		return geom2.NewPolygonGeometry(p.LineStrings...), nil
	case *geom2.Geometry_MultiLineString:
		ml, err := SimplifyMultiLineString(ctx, simplifier, g.MultiLineString)
		if err != nil {
			return nil, err
		}
		return geom2.NewMultiLineStringGeometry(ml.LineStrings...), nil
	case *geom2.Geometry_LineString:
		ls, err := SimplifyLineString(ctx, simplifier, g.LineString)
		if err != nil {
			return nil, err
		}
		return geom2.NewLineStringGeometry(ls.Coordinates...), nil
	default: // Points, MultiPoints or anything else.
		return geometry, nil
	}
}
