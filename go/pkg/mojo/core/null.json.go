package core

import (
	"unsafe"

	jsoniter "github.com/json-iterator/go"
)

func init() {
	RegisterJSONTypeEncoder("core.Null", &NullCodec{})
}

type NullCodec struct {
}

func (codec *NullCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
}

func (codec *NullCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return true
}
