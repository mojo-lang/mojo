package geom

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"unsafe"
)

func init() {
	core.RegisterJSONTypeDecoder("geom.MultiPolygon", &MultiPolygonCodec{})
	core.RegisterJSONTypeEncoder("geom.MultiPolygon", &MultiPolygonCodec{})
}

type MultiPolygonCodec struct {
}

func (codec *MultiPolygonCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	a := iter.ReadAny()
	if a.ValueType() == jsoniter.ObjectValue {
		c := a.Get("coordinates")
		if c.ValueType() == jsoniter.ArrayValue {
			convertMultiPolygon(c, (*MultiPolygon)(ptr))
		}
	} else if a.ValueType() == jsoniter.ArrayValue {
		dim := arrayDim(0, a)
		if dim == 3 {
			polygon := &Polygon{}
			convertPolygon(a, polygon)
			mp := (*MultiPolygon)(ptr)
			mp.Polygons = append(mp.Polygons, polygon)
		} else if dim == 4 {
			convertMultiPolygon(a, (*MultiPolygon)(ptr))
		}
	}
}

func convertMultiPolygon(array jsoniter.Any, multipolygon *MultiPolygon) {
	var coordinates [][][][]float64
	array.ToVal(&coordinates)

	for _, p := range coordinates {
		polygon := &Polygon{}
		for _, line := range p {
			lineString := parseLine(line)
			polygon.LineStrings = append(polygon.LineStrings, lineString)
		}

		multipolygon.Polygons = append(multipolygon.Polygons, polygon)
	}
}

func (codec *MultiPolygonCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return len(((*MultiPolygon)(ptr)).Polygons) == 0
}

func (codec *MultiPolygonCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	mp := (*MultiPolygon)(ptr)
	//if len(mp.Polygons) == 0 {
	//	stream.WriteVal(nil)
	//}

	stream.WriteObjectStart()
	stream.WriteObjectField("type")
	stream.WriteString("MultiPolygon")
	stream.WriteMore()
	stream.WriteObjectField("coordinates")
	stream.WriteArrayStart()
	for i, polygon := range mp.Polygons {

		if i > 0 {
			stream.WriteMore()
		}
		stream.WriteArrayStart()
		for j, line := range polygon.LineStrings {
			if j > 0 {
				stream.WriteMore()
			}
			stream.WriteArrayStart()
			for k, point := range line.Coordinates {
				if k > 0 {
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
		}
		stream.WriteArrayEnd()
	}
	stream.WriteArrayEnd()
	stream.WriteObjectEnd()
}
