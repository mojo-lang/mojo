package core

import (
	"unsafe"

	jsoniter "github.com/json-iterator/go"
)

func init() {
	RegisterJSONTypeDecoder("core.Object", &ObjectCodec{})
	RegisterJSONTypeEncoder("core.Object", &ObjectCodec{})
}

type ObjectCodec struct {
}

func (codec *ObjectCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	any := iter.ReadAny()
	obj := (*Object)(ptr)
	if any.ValueType() == jsoniter.ObjectValue && any.Size() > 0 {
		obj.Vals = make(map[string]*Value)
		for _, key := range any.Keys() {
			val, err := NewValueCodec().DecodeAny(any.Get(key))
			if err != nil {
				iter.ReportError("ObjectCodec.Decode", err.Error())
			}
			obj.Vals[key] = val
		}
	}
}

func (codec *ObjectCodec) IsEmpty(ptr unsafe.Pointer) bool {
	object := (*Object)(ptr)
	return object == nil || len(object.Vals) == 0
}

func (codec *ObjectCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	object := (*Object)(ptr)
	stream.WriteVal(object.Vals)
}
