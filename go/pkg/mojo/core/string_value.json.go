package core

import (
	"unsafe"

	jsoniter "github.com/json-iterator/go"
)

func init() {
	RegisterJSONTypeDecoder("core.StringValue", &StringValueCodec{})
	RegisterJSONTypeEncoder("core.StringValue", &StringValueCodec{})
}

type StringValueCodec struct {
}

func (codec *StringValueCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	(*StringValue)(ptr).Val = iter.ReadAny().ToString()
}

func (codec *StringValueCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	stream.WriteString((*StringValue)(ptr).Val)
}

func (codec *StringValueCodec) IsEmpty(ptr unsafe.Pointer) bool {
	e := (*StringValue)(ptr)
	return e == nil && len(e.Val) > 0
}
