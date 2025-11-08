package core

import (
	"encoding/base64"
	"unsafe"

	jsoniter "github.com/json-iterator/go"
)

func init() {
	RegisterJSONTypeDecoder("core.BytesValue", &BytesValueCodec{})
	RegisterJSONTypeEncoder("core.BytesValue", &BytesValueCodec{})
}

type BytesValueCodec struct {
}

func (codec *BytesValueCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	if bytes, err := base64.StdEncoding.DecodeString(iter.ReadAny().ToString()); err != nil {
		iter.ReportError("BytesValueCodec:Decode", err.Error())
	} else {
		(*BytesValue)(ptr).Val = bytes
	}
}

func (codec *BytesValueCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	stream.WriteString(base64.StdEncoding.EncodeToString((*BytesValue)(ptr).Val))
}

func (codec *BytesValueCodec) IsEmpty(ptr unsafe.Pointer) bool {
	e := (*BytesValue)(ptr)
	return e == nil && len(e.Val) > 0
}
