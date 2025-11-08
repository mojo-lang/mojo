package geom

import (
	jsoniter "github.com/json-iterator/go"
	"unsafe"
)

func init() {
	jsoniter.RegisterTypeDecoder("geom.LngLat", &LngLatCodec{})
}

type LngLatCodec struct {
}

func (codec *LngLatCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	a := iter.ReadAny()
	lngLat := (*LngLat)(ptr)
	if a.ValueType() == jsoniter.ArrayValue {
		var points []float64
		a.ToVal(&points)
		if len(points) < 2 {
			iter.ReportError("LngLatCodec.Decode", "invalid format of the lngLat, at least two element")
			return
		}
		lngLat.Longitude = points[0]
		lngLat.Latitude = points[1]
		if len(points) > 2 {
			lngLat.Altitude = points[2]
		}
	} else if a.ValueType() == jsoniter.ObjectValue {
		lng := a.Get("longitude")
		if lng.ValueType() == jsoniter.NumberValue {
			lngLat.Longitude = lng.ToFloat64()
		} else {
			lngLat.Longitude = a.Get("lng").ToFloat64()
		}
		lat := a.Get("latitude")
		if lat.ValueType() == jsoniter.NumberValue {
			lngLat.Latitude = lat.ToFloat64()
		} else {
			lngLat.Latitude = a.Get("lat").ToFloat64()
		}
		alt := a.Get("altitude")
		if lng.ValueType() == jsoniter.NumberValue {
			lngLat.Altitude = alt.ToFloat64()
		} else {
			lngLat.Altitude = a.Get("alt").ToFloat64()
		}
	} else if a.ValueType() == jsoniter.StringValue {
		err := lngLat.Parse(a.ToString())
		if err != nil {
			iter.ReportError("LngLatCodec.Decode", err.Error())
		}
	}
}

func (codec *LngLatCodec) IsEmpty(ptr unsafe.Pointer) bool {
	lngLat := (*LngLat)(ptr)
	return lngLat == nil || lngLat.IsEmpty()
}
