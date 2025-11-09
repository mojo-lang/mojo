package geom

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"unsafe"
)

func init() {
	core.RegisterJSONTypeDecoder("geom.GeometryCollection", &GeometryCollectionCodec{})
	core.RegisterJSONTypeEncoder("geom.GeometryCollection", &GeometryCollectionCodec{})
}

type GeometryCollectionCodec struct {
}

func (codec *GeometryCollectionCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	a := iter.ReadAny()
	if a.ValueType() == jsoniter.ObjectValue {
		t := a.Get("type").ToString()
		if t != "GeometryCollection" {

		}

		geometries := a.Get("geometries")
		geometries.ToVal(&((*GeometryCollection)(ptr)).Geometries)
	}
}

func (codec *GeometryCollectionCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return len(((*GeometryCollection)(ptr)).Geometries) == 0
}

func (codec *GeometryCollectionCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	gc := (*GeometryCollection)(ptr)

	//if len(gc.Geometries) == 0 {
	//	stream.WriteVal(nil)
	//}

	stream.WriteObjectStart()

	stream.WriteObjectField("type")
	stream.WriteString("GeometryCollection")

	stream.WriteMore()
	stream.WriteObjectField("geometries")
	if len(gc.Geometries) == 0 {
		stream.WriteArrayStart()
		stream.WriteArrayEnd()
	} else {
		stream.WriteVal(gc.Geometries)
	}

	stream.WriteObjectEnd()
}
