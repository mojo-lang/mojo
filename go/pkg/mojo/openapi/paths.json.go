package openapi

import (
	"io"
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
	value := iter.ReadAny()
	if iter.Error != nil && iter.Error != io.EOF {
		return
	}
	target := (*Paths)(ptr)
	if value.ValueType() == jsoniter.NilValue {
		target.Vals = nil
		return
	}
	if value.ValueType() != jsoniter.ObjectValue {
		iter.ReportError("PathsCodec.Decode", "expected an object or null")
		return
	}
	values := make(map[string]*PathItem, len(target.Vals))
	for key, item := range target.Vals {
		values[key] = item
	}
	if decodeJSONValue(iter, value, &values) {
		target.Vals = values
	}
}

func (codec *PathsCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	target := (*Paths)(ptr)
	if target != nil && len(target.Vals) > 0 {
		stream.WriteVal(target.Vals)
	} else {
		stream.WriteEmptyObject()
	}
}

func (codec *PathsCodec) IsEmpty(ptr unsafe.Pointer) bool {
	target := (*Paths)(ptr)
	return target == nil || len(target.Vals) == 0
}
