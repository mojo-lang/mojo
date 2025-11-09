package geom

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"unsafe"
)

func init() {
	core.RegisterJSONTypeDecoder("geom.GeoJson", &GeoJsonCodec{})
	core.RegisterJSONTypeEncoder("geom.GeoJson", &GeoJsonCodec{})
}

type GeoJsonCodec struct {
}

func (codec *GeoJsonCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	a := iter.ReadAny()
	if a.ValueType() == jsoniter.ObjectValue {
		t := a.Get("type").ToString()
		switch t {
		case "Feature":
			val := &Feature{}
			a.ToVal(val)
			((*GeoJson)(ptr)).GeoJson = &GeoJson_Feature{Feature: val}
		case "FeatureCollection":
			val := &FeatureCollection{}
			a.ToVal(val)
			((*GeoJson)(ptr)).GeoJson = &GeoJson_FeatureCollection{FeatureCollection: val}
		case "Point":
			val := &Point{}
			a.ToVal(val)
			((*GeoJson)(ptr)).GeoJson = &GeoJson_Point{Point: val}
		}
	}
}

func (codec *GeoJsonCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return ((*GeoJson)(ptr)).GeoJson == nil
}

func (codec *GeoJsonCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	geoJson := (*GeoJson)(ptr)

	switch x := geoJson.GeoJson.(type) {
	case *GeoJson_Feature:
		stream.WriteVal(x.Feature)
	case *GeoJson_FeatureCollection:
		stream.WriteVal(x.FeatureCollection)
	case *GeoJson_Point:
		stream.WriteVal(x.Point)
	case nil:
		stream.WriteVal(nil)
	default:
	}
}
