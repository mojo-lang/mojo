package geom

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/mojo-lang/mojo/packages/core/go/pkg/mojo/core"
	"unsafe"
)

func init() {
	core.RegisterJSONTypeDecoder("geom.FeatureCollection", &FeatureCollectionCodec{})
	core.RegisterJSONTypeEncoder("geom.FeatureCollection", &FeatureCollectionCodec{})
}

type FeatureCollectionCodec struct {
}

func (codec *FeatureCollectionCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	a := iter.ReadAny()
	if a.ValueType() == jsoniter.ObjectValue {
		t := a.Get("type").ToString()
		if t == "FeatureCollection" {
			features := a.Get("features")
			features.ToVal(&((*FeatureCollection)(ptr)).Features)
		} else {
			iter.ReportError("FeatureCollectionDecode", "the type field is invalid")
		}
	}
}

func (codec *FeatureCollectionCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return len(((*FeatureCollection)(ptr)).Features) == 0
}

func (codec *FeatureCollectionCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	fc := (*FeatureCollection)(ptr)

	stream.WriteObjectStart()

	stream.WriteObjectField("type")
	stream.WriteString("FeatureCollection")

	stream.WriteMore()
	stream.WriteObjectField("features")
	if len(fc.Features) == 0 {
		stream.WriteArrayStart()
		stream.WriteArrayEnd()
	} else {
		stream.WriteVal(fc.Features)
	}

	stream.WriteObjectEnd()
}
