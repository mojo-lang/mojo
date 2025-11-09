package geom

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"strconv"
	"strings"
	"unsafe"
)

func init() {
	jsoniter.RegisterTypeDecoder("geom.Circle", &CircleCodec{})
	jsoniter.RegisterTypeEncoder("geom.Circle", &CircleCodec{})
}

type CircleCodec struct {
}

func (codec *CircleCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	a := iter.ReadAny()
	if a.ValueType() == jsoniter.ObjectValue {
		if a.Get("properties", "subType").ToString() != "Circle" {
			iter.ReportError("Circle Decode", fmt.Sprintf("NOT circle feature"))
		} else {
			circle := (*Circle)(ptr)

			c := a.Get("geometry", "coordinates")
			if c.ValueType() == jsoniter.ArrayValue {
				var coordinates []float64
				c.ToVal(&coordinates)
				if len(coordinates) == 2 {
					circle.Center = &LngLat{
						Longitude: coordinates[0],
						Latitude:  coordinates[1],
					}
				} else {
					iter.ReportError("Circle Decode", fmt.Sprintf("invalid circle coordinates"))
					return
				}
			}

			r := a.Get("properties", "radius")
			circle.Radius = r.ToFloat32()
		}
	} else if a.ValueType() == jsoniter.StringValue {
		segments := strings.Split(a.ToString(), ",")
		if len(segments) == 3 {
			circle := (*Circle)(ptr)
			lng, err := strconv.ParseFloat(segments[0], 64)
			if err != nil {
				iter.ReportError("Circle Decode", fmt.Sprintf("invalid longitude of circle center %s, error: %s", segments[0], err.Error()))
				return
			}
			lat, err := strconv.ParseFloat(segments[1], 64)
			if err != nil {
				iter.ReportError("Circle Decode", fmt.Sprintf("invalid latitude of circle center %s, error: %s", segments[1], err.Error()))
				return
			}
			r, err := strconv.ParseFloat(segments[2], 64)
			if err != nil {
				iter.ReportError("Circle Decode", fmt.Sprintf("invalid radius of circle %s, error: %s", segments[2], err.Error()))
				return
			}
			circle.Center = &LngLat{
				Longitude: lng,
				Latitude:  lat,
			}
			circle.Radius = float32(r)
		} else {
			iter.ReportError("Circle Decode", fmt.Sprintf("invalid circle string %s", a.ToString()))
		}
	}
}

func (codec *CircleCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return ((*Circle)(ptr)).Radius == 0 || ((*Circle)(ptr)).Center == nil
}

func (codec *CircleCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	circle := (*Circle)(ptr)

	stream.WriteObjectStart()
	stream.WriteObjectField("type")
	stream.WriteString("Feature")

	stream.WriteObjectField("geometry")
	stream.WriteObjectStart()
	stream.WriteObjectField("type")
	stream.WriteString("Point")
	stream.WriteMore()
	stream.WriteObjectField("coordinates")

	stream.WriteArrayStart()
	stream.WriteFloat64Lossy(circle.Center.Longitude)
	stream.WriteMore()
	stream.WriteFloat64Lossy(circle.Center.Latitude)
	stream.WriteArrayEnd()

	stream.WriteObjectEnd()

	stream.WriteObjectField("properties")
	stream.WriteObjectStart()
	stream.WriteObjectField("subType")
	stream.WriteString("Circle")
	stream.WriteObjectField("radius")
	stream.WriteFloat32Lossy(circle.Radius)
	stream.WriteObjectEnd()
}
