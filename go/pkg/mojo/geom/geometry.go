package geom

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
)

type GeometryType int

const (
	UnknownType GeometryType = iota
	PointType
	MultiPointType
	LineStringType
	MultiLineStringType
	PolygonType
	MultiPolygonType
	GeometryCollectionType
)

func NewGeometry(val interface{}) *Geometry {
	geometry := &Geometry{}
	_ = geometry.SetValue(val)
	return geometry
}

func NewPointGeometryFrom(coordinates ...float64) *Geometry {
	return NewGeometry(NewPointFrom(coordinates...))
}

func NewPointGeometry(ll *LngLat) *Geometry {
	return NewGeometry(NewPoint(ll))
}

func NewMultiPointGeometry(points ...*Point) *Geometry {
	return NewGeometry(NewMultiPoint(points...))
}

func NewMultiPointGeometryFrom(points ...*LngLat) *Geometry {
	return NewGeometry(NewMultiPointFrom(points...))
}

func NewLineStringGeometry(coordinates ...*LngLat) *Geometry {
	return NewGeometry(NewLineString(coordinates...))
}

func NewMultiLineStringGeometry(lineStrings ...*LineString) *Geometry {
	return NewGeometry(NewMultiLineString(lineStrings...))
}

func NewMultiLineStringGeometryFrom(coordinates ...[]*LngLat) *Geometry {
	return NewGeometry(NewMultiLineStringFrom(coordinates...))
}

func NewPolygonGeometry(lineStrings ...*LineString) *Geometry {
	return NewGeometry(NewPolygon(lineStrings...))
}

func NewPolygonGeometryFrom(coordinates ...*LngLat) *Geometry {
	return NewGeometry(NewPolygonFrom(coordinates...))
}

func NewMultiPolygonGeometry(p ...*Polygon) *Geometry {
	return NewGeometry(NewMultiPolygon(p...))
}

func NewGeometryCollectionGeometry(geometries ...*Geometry) *Geometry {
	return NewGeometry(NewGeometryCollection(geometries...))
}

func (x *Geometry) Type() GeometryType {
	switch x.Geometry.(type) {
	case *Geometry_Point:
		return PointType
	case *Geometry_MultiPoint:
		return MultiPointType
	case *Geometry_LineString:
		return LineStringType
	case *Geometry_MultiLineString:
		return MultiLineStringType
	case *Geometry_Polygon:
		return PolygonType
	case *Geometry_MultiPolygon:
		return MultiPolygonType
	case *Geometry_GeometryCollection:
		return GeometryCollectionType
	default:
		return UnknownType
	}
}

func (x *Geometry) TypeString() string {
	switch x.Geometry.(type) {
	case *Geometry_Point:
		return "Point"
	case *Geometry_MultiPoint:
		return "MultiPoint"
	case *Geometry_LineString:
		return "LineString"
	case *Geometry_MultiLineString:
		return "MultiLineString"
	case *Geometry_Polygon:
		return "Polygon"
	case *Geometry_MultiPolygon:
		return "MultiPolygon"
	case *Geometry_GeometryCollection:
		return "GeometryCollection"
	default:
		return ""
	}
}

func (x *Geometry) SetValue(val interface{}) error {
	switch m := val.(type) {
	case *Point:
		x.Geometry = &Geometry_Point{Point: m}
	case *MultiPoint:
		x.Geometry = &Geometry_MultiPoint{MultiPoint: m}
	case *LineString:
		x.Geometry = &Geometry_LineString{LineString: m}
	case *MultiLineString:
		x.Geometry = &Geometry_MultiLineString{MultiLineString: m}
	case *Polygon:
		x.Geometry = &Geometry_Polygon{Polygon: m}
	case *MultiPolygon:
		x.Geometry = &Geometry_MultiPolygon{m}
	case *GeometryCollection:
		x.Geometry = &Geometry_GeometryCollection{GeometryCollection: m}
	case *core.Object:
		if bytes, err := jsoniter.Marshal(m); err != nil {
			return err
		} else {
			geometry := &Geometry{}
			if err = jsoniter.Unmarshal(bytes, geometry); err != nil {
				return err
			}
			x.Geometry = geometry.Geometry
		}
	case *core.Value:
		if obj := m.GetObject(); obj != nil {
			return x.SetValue(obj)
		} else {
			return fmt.Errorf("Place.Value has unexpected core.Value type %s", m.GetKind().String())
		}
	default:
		return fmt.Errorf("Place.Value has unexpected type %T", m)
	}
	return nil
}

func (x *Geometry) ToObject() *core.Object {
	if x != nil {
		if bytes, err := jsoniter.Marshal(x); err != nil {
			return nil
		} else {
			obj := &core.Object{}
			if err = jsoniter.Unmarshal(bytes, obj); err != nil {
				return nil
			}
			return obj
		}
	}
	return nil
}

func NewGeometryFrom(obj *core.Object) (*Geometry, error) {
	if obj == nil {
		return nil, nil
	}

	if bytes, err := jsoniter.Marshal(obj); err != nil {
		return nil, err
	} else {
		g := &Geometry{}
		if err = jsoniter.Unmarshal(bytes, g); err != nil {
			return nil, err
		}
		return g, nil
	}
}

// Area
// Calculate the geometry area if the geometry is polygon or polygon-convertable.
func (x *Geometry) Area() float64 {
	if x == nil {
		return 0
	}

	switch g := x.Geometry.(type) {
	case *Geometry_Point, *Geometry_MultiPoint:
		return 0
	case *Geometry_LineString:
		if g.LineString.IsClosed() {
			return NewPolygon(g.LineString).Area()
		}
	case *Geometry_MultiLineString:
		area := 0.0
		for _, l := range g.MultiLineString.LineStrings {
			if l.IsClosed() {
				area += NewPolygon(l).Area()
			}
		}
		return area
	case *Geometry_Polygon:
		return g.Polygon.Area()
	case *Geometry_MultiPolygon:
		return g.MultiPolygon.Area()
	case *Geometry_GeometryCollection:
		area := 0.0
		for _, geo := range g.GeometryCollection.Geometries {
			area += geo.Area()
		}
		return area
	}

	return 0
}

func (x *Geometry) BoundingBox() *BoundingBox {
	if x == nil {
		return nil
	}

	switch g := x.Geometry.(type) {
	case *Geometry_Point:
		return g.Point.BoundingBox()
	case *Geometry_MultiPoint:
		return g.MultiPoint.BoundingBox()
	case *Geometry_LineString:
		return g.LineString.BoundingBox()
	case *Geometry_MultiLineString:
		return g.MultiLineString.BoundingBox()
	case *Geometry_Polygon:
		return g.Polygon.BoundingBox()
	case *Geometry_MultiPolygon:
		return g.MultiPolygon.BoundingBox()
	case *Geometry_GeometryCollection:
		return g.GeometryCollection.BoundingBox()
	}
	return nil
}

func (x *Geometry) Contains(point *LngLat) bool {
	if x == nil {
		return false
	}

	switch g := x.Geometry.(type) {
	case *Geometry_Polygon:
		return g.Polygon.Contains(point)
	case *Geometry_MultiPolygon:
		return g.MultiPolygon.Contains(point)
	case *Geometry_GeometryCollection:
		for _, geo := range g.GeometryCollection.Geometries {
			if geo.Contains(point) {
				return true
			}
		}
	}
	return false
}

func (x *Geometry) Equal(geometry *Geometry) bool {
	if x != nil && geometry != nil {
		switch g := x.Geometry.(type) {
		case *Geometry_Point:
			if p, ok := geometry.Geometry.(*Geometry_Point); ok {
				return g.Point.Equal(p.Point)
			}
		case *Geometry_MultiPoint:
			if p, ok := geometry.Geometry.(*Geometry_MultiPoint); ok {
				return g.MultiPoint.Equal(p.MultiPoint)
			}
		case *Geometry_LineString:
			if p, ok := geometry.Geometry.(*Geometry_LineString); ok {
				return g.LineString.Equal(p.LineString)
			}
		case *Geometry_MultiLineString:
			if p, ok := geometry.Geometry.(*Geometry_MultiLineString); ok {
				return g.MultiLineString.Equal(p.MultiLineString)
			}
		case *Geometry_Polygon:
			if p, ok := geometry.Geometry.(*Geometry_Polygon); ok {
				return g.Polygon.Equal(p.Polygon)
			}
		case *Geometry_MultiPolygon:
			if p, ok := geometry.Geometry.(*Geometry_MultiPolygon); ok {
				return g.MultiPolygon.Equal(p.MultiPolygon)
			}
		case *Geometry_GeometryCollection:
			if p, ok := geometry.Geometry.(*Geometry_GeometryCollection); ok {
				return g.GeometryCollection.Equal(p.GeometryCollection)
			}
		}
		return false
	}
	return x == nil && geometry == nil
}

func (x *Geometry) CoordTransform(from, to SpatialReference) *Geometry {
	if x == nil {
		return x
	}

	switch g := x.Geometry.(type) {
	case *Geometry_Point:
		return &Geometry{Geometry: &Geometry_Point{Point: g.Point.CoordTransform(from, to)}}
	case *Geometry_MultiPoint:
		return &Geometry{Geometry: &Geometry_MultiPoint{MultiPoint: g.MultiPoint.CoordTransform(from, to)}}
	case *Geometry_LineString:
		return &Geometry{Geometry: &Geometry_LineString{LineString: g.LineString.CoordTransform(from, to)}}
	case *Geometry_MultiLineString:
		return &Geometry{Geometry: &Geometry_MultiLineString{MultiLineString: g.MultiLineString.CoordTransform(from, to)}}
	case *Geometry_Polygon:
		return &Geometry{Geometry: &Geometry_Polygon{Polygon: g.Polygon.CoordTransform(from, to)}}
	case *Geometry_MultiPolygon:
		return &Geometry{Geometry: &Geometry_MultiPolygon{MultiPolygon: g.MultiPolygon.CoordTransform(from, to)}}
	case *Geometry_GeometryCollection:
		return &Geometry{Geometry: &Geometry_GeometryCollection{GeometryCollection: g.GeometryCollection.CoordTransform(from, to)}}
	}
	return x
}
