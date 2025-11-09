package shapefile

import (
	"errors"
	"fmt"
	"github.com/jonas-p/go-shp"
	geom2 "github.com/mojo-lang/mojo/go/pkg/mojo/geom"
)

func NewGeometry(shape shp.Shape) (*geom2.Geometry, *geom2.BoundingBox, error) {
	if shape == nil {
		return nil, nil, nil
	}

	box := shape.BBox()
	bbox := geom2.NewBoundingBox(geom2.NewLngLat(box.MinX, box.MinY), geom2.NewLngLat(box.MaxX, box.MaxY))

	switch s := shape.(type) {
	case *shp.Null:
		return nil, bbox, nil
	case *shp.Point:
		return geom2.NewPointGeometryFrom(s.X, s.Y), bbox, nil
	case *shp.PolyLine:
		ml := newLineStrings(s.Parts, s.Points)
		if len(ml) == 1 {
			return geom2.NewGeometry(ml[0]), bbox, nil
		} else if s.NumParts > 1 {
			return geom2.NewMultiLineStringGeometry(ml...), bbox, nil
		}
	case *shp.Polygon:
		ml := newLineStrings(s.Parts, s.Points)
		return geom2.NewPolygonGeometry(ml...), bbox, nil
	case *shp.MultiPoint:
		mp := &geom2.MultiPoint{}
		for _, p := range s.Points {
			mp.Points = append(mp.Points, geom2.NewPoint(geom2.NewLngLat(p.X, p.Y)))
		}
		return geom2.NewGeometry(mp), bbox, nil
	default:
		return nil, nil, errors.New(fmt.Sprintf("unsupported type for the Shape %T", s))
	}
	return nil, nil, nil
}

func GeometryType(geometry *geom2.Geometry) (shp.ShapeType, error) {
	var tpy shp.ShapeType
	switch geometry.GetGeometry().(type) {
	case *geom2.Geometry_Point:
		tpy = shp.POINT
	case *geom2.Geometry_LineString:
		tpy = shp.POLYLINE
	case *geom2.Geometry_Polygon:
		tpy = shp.POLYGON
	case *geom2.Geometry_MultiPoint:
		tpy = shp.MULTIPOINT
	default:
		return tpy, errors.New("not support geometry type except point, linestring, polygon or multi point")
	}
	return tpy, nil
}

func GeometryToShape(geometry *geom2.Geometry, bbox *geom2.BoundingBox) (shp.Shape, error) {
	if geometry == nil {
		return nil, nil
	}

	typ, err := GeometryType(geometry)
	if err != nil {
		return nil, err
	}

	if bbox == nil {
		bbox = geometry.BoundingBox()
	}

	switch g := geometry.Geometry.(type) {
	case *geom2.Geometry_Point:
		if typ == shp.POINT {
			point := &shp.Point{X: g.Point.GetLongitude(), Y: g.Point.GetLatitude()}
			return point, nil
		}
	case *geom2.Geometry_MultiPoint:
		if typ == shp.MULTIPOINT {
			points := &shp.MultiPoint{}
			for _, pt := range g.MultiPoint.GetPoints() {
				point := shp.Point{X: pt.GetLongitude(), Y: pt.GetLatitude()}
				points.Points = append(points.Points, point)
			}
			points.NumPoints = int32(len(points.Points))
			return points, nil
		}
	case *geom2.Geometry_LineString:
		if typ == shp.POLYLINE {
			var points []shp.Point
			for _, pt := range g.LineString.GetCoordinates() {
				points = append(points, shp.Point{X: pt.GetLongitude(), Y: pt.GetLatitude()})
			}
			line := shp.NewPolyLine([][]shp.Point{points})
			return line, nil
		}
	case *geom2.Geometry_Polygon:
		if typ == shp.POLYGON {
			var points [][]shp.Point
			for _, line := range g.Polygon.LineStrings {
				var pts []shp.Point
				for _, pt := range line.GetCoordinates() {
					pts = append(pts, shp.Point{
						X: pt.GetLongitude(),
						Y: pt.GetLatitude(),
					})
				}
				points = append(points, pts)
			}
			polygon := (*shp.Polygon)(shp.NewPolyLine(points))
			return polygon, nil
		}
	}

	return nil, nil
}

func newLineStrings(parts []int32, points []shp.Point) []*geom2.LineString {
	index := 0
	var ml []*geom2.LineString
	line := &geom2.LineString{}
	for i := 0; i < len(points); i++ {
		if index < len(parts)-1 && i == int(parts[index+1]) {
			ml = append(ml, line)
			line = &geom2.LineString{}
			index++
		}
		line.Coordinates = append(line.Coordinates, &geom2.LngLat{
			Longitude: points[i].X,
			Latitude:  points[i].Y,
		})
	}
	ml = append(ml, line)
	return ml
}
