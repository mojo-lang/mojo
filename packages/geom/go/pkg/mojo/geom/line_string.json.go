package geom

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/mojo-lang/mojo/packages/core/go/pkg/mojo/core"
	"unsafe"
)

func init() {
	core.RegisterJSONTypeDecoder("geom.LineString", &LineStringCodec{})
	core.RegisterJSONTypeEncoder("geom.LineString", &LineStringCodec{})
}

type LineStringCodec struct {
}

func (codec *LineStringCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	a := iter.ReadAny()
	if a.ValueType() == jsoniter.ObjectValue {
		c := a.Get("coordinates")
		if c.ValueType() == jsoniter.ArrayValue {
			var coordinates [][]float64
			c.ToVal(&coordinates)

			line := (*LineString)(ptr)
			for _, point := range coordinates {
				if len(point) > 1 {
					lngLat := &LngLat{}
					lngLat.Longitude = point[0]
					lngLat.Latitude = point[1]

					if len(point) > 2 {
						lngLat.Altitude = point[2]
					}
					line.Coordinates = append(line.Coordinates, lngLat)
				}
			}
		}
	}
}

func (codec *LineStringCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return len(((*LineString)(ptr)).Coordinates) == 0
}

func (codec *LineStringCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	line := (*LineString)(ptr)

	//if len(line.Coordinates) == 0 {
	//	stream.WriteVal(nil)
	//}

	stream.WriteObjectStart()
	stream.WriteObjectField("type")
	stream.WriteString("LineString")
	stream.WriteMore()
	stream.WriteObjectField("coordinates")

	stream.WriteArrayStart()
	for idx, point := range line.Coordinates {
		if idx > 0 {
			stream.WriteMore()
		}

		stream.WriteArrayStart()
		stream.WriteFloat64Lossy(point.Longitude)

		stream.WriteMore()
		stream.WriteFloat64Lossy(point.Latitude)
		if point.Altitude != 0 {
			stream.WriteMore()
			stream.WriteFloat64Lossy(point.Altitude)
		}
		stream.WriteArrayEnd()
	}
	stream.WriteArrayEnd()
	stream.WriteObjectEnd()
}
