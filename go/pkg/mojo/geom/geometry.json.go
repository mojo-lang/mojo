package geom

import (
	"unsafe"

	jsoniter "github.com/json-iterator/go"
)

func init() {
	jsoniter.RegisterTypeDecoder("geom.Geometry", &GeometryCodec{})
	jsoniter.RegisterTypeEncoder("geom.Geometry", &GeometryCodec{})
}

type GeometryCodec struct {
}

func (codec *GeometryCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	a := iter.ReadAny()
	if a.ValueType() == jsoniter.ObjectValue {
		t := a.Get("type").ToString()
		switch t {
		case "Point":
			val := &Point{}
			a.ToVal(val)
			((*Geometry)(ptr)).Geometry = &Geometry_Point{Point: val}
		case "MultiPoint":
			val := &MultiPoint{}
			a.ToVal(val)
			((*Geometry)(ptr)).Geometry = &Geometry_MultiPoint{MultiPoint: val}
		case "LineString":
			val := &LineString{}
			a.ToVal(val)
			((*Geometry)(ptr)).Geometry = &Geometry_LineString{LineString: val}
		case "MultiLineString":
			val := &MultiLineString{}
			a.ToVal(val)
			((*Geometry)(ptr)).Geometry = &Geometry_MultiLineString{MultiLineString: val}
		case "Polygon":
			val := &Polygon{}
			a.ToVal(val)
			((*Geometry)(ptr)).Geometry = &Geometry_Polygon{Polygon: val}
		case "MultiPolygon":
			val := &MultiPolygon{}
			a.ToVal(val)
			((*Geometry)(ptr)).Geometry = &Geometry_MultiPolygon{val}
		case "GeometryCollection":
			val := &GeometryCollection{}
			a.ToVal(val)
			((*Geometry)(ptr)).Geometry = &Geometry_GeometryCollection{GeometryCollection: val}
		}
	}
}

func (codec *GeometryCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return ((*Geometry)(ptr)).Geometry == nil
}

func (codec *GeometryCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	geometry := (*Geometry)(ptr)

	switch x := geometry.Geometry.(type) {
	case *Geometry_Point:
		x.Point.Type = "Point"
		stream.WriteVal(x.Point)
	case *Geometry_MultiPoint:
		x.MultiPoint.Type = "MultiPoint"
		stream.WriteVal(x.MultiPoint)
	case *Geometry_LineString:
		x.LineString.Type = "LineString"
		stream.WriteVal(x.LineString)
	case *Geometry_MultiLineString:
		x.MultiLineString.Type = "MultiLineString"
		stream.WriteVal(x.MultiLineString)
	case *Geometry_Polygon:
		x.Polygon.Type = "Polygon"
		stream.WriteVal(x.Polygon)
	case *Geometry_MultiPolygon:
		x.MultiPolygon.Type = "MultiPolygon"
		stream.WriteVal(x.MultiPolygon)
	case *Geometry_GeometryCollection:
		x.GeometryCollection.Type = "GeometryCollection"
		stream.WriteVal(x.GeometryCollection)
	default:
		stream.WriteObjectStart()
		stream.WriteObjectEnd()
	}
}
