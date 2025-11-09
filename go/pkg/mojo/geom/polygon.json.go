package geom

import (
	jsoniter "github.com/json-iterator/go"
	"unsafe"
)

func init() {
	jsoniter.RegisterTypeDecoder("geom.Polygon", &PolygonCodec{})
	jsoniter.RegisterTypeEncoder("geom.Polygon", &PolygonCodec{})
}

type PolygonCodec struct {
}

func arrayDim(dim int, a jsoniter.Any) int {
	if a.ValueType() == jsoniter.ArrayValue {
		if a.Size() > 0 {
			return arrayDim(dim+1, a.Get(0))
		} else {
			return dim + 1
		}
	}
	return dim
}

func parseLine(line [][]float64) *LineString {
	lineString := &LineString{}
	for _, point := range line {
		if len(point) > 1 {
			lngLat := &LngLat{}
			lngLat.Longitude = point[0]
			lngLat.Latitude = point[1]

			if len(point) > 2 {
				lngLat.Altitude = point[2]
			}
			lineString.Add(lngLat)
		}
	}
	return lineString
}

func convertPolygon(array jsoniter.Any, polygon *Polygon) {
	var coordinates [][][]float64
	array.ToVal(&coordinates)

	for _, line := range coordinates {
		lineString := parseLine(line)
		polygon.LineStrings = append(polygon.LineStrings, lineString)
	}
}

func (codec *PolygonCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	a := iter.ReadAny()
	if a.ValueType() == jsoniter.ObjectValue {
		c := a.Get("coordinates")
		if c.ValueType() == jsoniter.ArrayValue {
			convertPolygon(c, (*Polygon)(ptr))
		}
	} else if a.ValueType() == jsoniter.ArrayValue {
		dim := arrayDim(0, a)
		if dim == 2 {
			var coordinates [][]float64
			a.ToVal(&coordinates)
			lineString := parseLine(coordinates)
			polygon := (*Polygon)(ptr)
			polygon.LineStrings = append(polygon.LineStrings, lineString)
		} else if dim == 3 {
			convertPolygon(a, (*Polygon)(ptr))
		}
	}
}

func (codec *PolygonCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return len(((*Polygon)(ptr)).LineStrings) == 0
}

func (codec *PolygonCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	polygon := (*Polygon)(ptr)

	//if len(polygon.Coordinates) == 0 {
	//	stream.WriteVal(nil)
	//}

	stream.WriteObjectStart()
	stream.WriteObjectField("type")
	stream.WriteString("Polygon")
	stream.WriteMore()
	stream.WriteObjectField("coordinates")

	stream.WriteArrayStart()
	for i, line := range polygon.LineStrings {
		if i > 0 {
			stream.WriteMore()
		}
		stream.WriteArrayStart()
		for j, point := range line.Coordinates {
			if j > 0 {
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
	}

	stream.WriteArrayEnd()
	stream.WriteObjectEnd()
}
