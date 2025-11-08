package geom

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/mojo-lang/mojo/packages/core/go/pkg/mojo/core"
	"unsafe"
)

func init() {
	core.RegisterJSONTypeDecoder("geom.MultiPoint", &MultiPointCodec{})
	core.RegisterJSONTypeEncoder("geom.MultiPoint", &MultiPointCodec{})
}

type MultiPointCodec struct {
}

func (codec *MultiPointCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	a := iter.ReadAny()
	if a.ValueType() == jsoniter.ObjectValue {
		//t := any.Get("type")
		c := a.Get("coordinates")
		if c.ValueType() == jsoniter.ArrayValue {
			var coordinates [][]float64
			c.ToVal(&coordinates)

			multiPoint := (*MultiPoint)(ptr)
			for index := 0; index < len(coordinates); index++ {
				point := coordinates[index]
				if len(point) > 1 {
					lngLat := &LngLat{}
					lngLat.Longitude = point[0]
					lngLat.Latitude = point[1]

					if len(point) > 2 {
						lngLat.Altitude = point[2]
					}
					multiPoint.Points = append(multiPoint.Points, NewPoint(lngLat))
				} else {
					// skip the invalid point
				}
			}
		}
	}
}

func (codec *MultiPointCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return len(((*MultiLineString)(ptr)).LineStrings) == 0
}

func (codec *MultiPointCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	mp := (*MultiPoint)(ptr)
	//if len(mp.Points) == 0 {
	//	stream.WriteVal(nil)
	//}

	stream.WriteObjectStart()
	stream.WriteObjectField("type")
	stream.WriteString("MultiPoint")
	stream.WriteMore()
	stream.WriteObjectField("coordinates")
	stream.WriteArrayStart()
	for i, point := range mp.Points {
		if i > 0 {
			stream.WriteMore()
		}
		stream.WriteArrayStart()

		stream.WriteFloat64Lossy(point.GetLongitude())
		stream.WriteMore()
		stream.WriteFloat64Lossy(point.GetLatitude())
		if point.GetAltitude() != 0 {
			stream.WriteMore()
			stream.WriteFloat64Lossy(point.GetAltitude())
		}
		stream.WriteArrayEnd()
	}
	stream.WriteArrayEnd()
	stream.WriteObjectEnd()
}
