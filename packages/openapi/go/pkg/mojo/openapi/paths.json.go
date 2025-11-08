package openapi

import (
	"unsafe"

	jsoniter "github.com/json-iterator/go"
)

func init() {
	jsoniter.RegisterTypeDecoder("openapi.Paths", &PathsCodec{})
	jsoniter.RegisterTypeEncoder("openapi.Paths", &PathsCodec{})
}

type PathsCodec struct {
}

func (codec *PathsCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	any := iter.ReadAny()
	paths := (*Paths)(ptr)
	if any.ValueType() == jsoniter.ObjectValue {
		if paths.Vals == nil {
			paths.Vals = make(map[string]*PathItem)
		}
		any.ToVal(&paths.Vals)
	}
}

func (codec *PathsCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	paths := (*Paths)(ptr)

	if paths != nil && len(paths.Vals) > 0 {
		stream.WriteVal(paths.Vals)
	}
}

func (codec *PathsCodec) IsEmpty(ptr unsafe.Pointer) bool {
	paths := (*Paths)(ptr)
	return paths == nil || len(paths.Vals) == 0
}
