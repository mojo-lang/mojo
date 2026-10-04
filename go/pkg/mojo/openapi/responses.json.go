package openapi

import (
	"io"
	"unsafe"

	jsoniter "github.com/json-iterator/go"
)

func init() {
	jsoniter.RegisterTypeDecoder("openapi.Responses", &ResponsesCodec{})
	jsoniter.RegisterTypeEncoder("openapi.Responses", &ResponsesCodec{})
}

type ResponsesCodec struct {
}

func (codec *ResponsesCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	value := iter.ReadAny()
	if iter.Error != nil && iter.Error != io.EOF {
		return
	}
	target := (*Responses)(ptr)
	if value.ValueType() == jsoniter.NilValue {
		target.Vals = nil
		return
	}
	if value.ValueType() != jsoniter.ObjectValue {
		iter.ReportError("ResponsesCodec.Decode", "expected an object or null")
		return
	}
	values := make(map[string]*ReferenceableResponse, len(target.Vals))
	for key, item := range target.Vals {
		values[key] = item
	}
	if decodeJSONValue(iter, value, &values) {
		target.Vals = values
	}
}

func (codec *ResponsesCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	target := (*Responses)(ptr)
	if target != nil && len(target.Vals) > 0 {
		stream.WriteVal(target.Vals)
	} else {
		stream.WriteEmptyObject()
	}
}

func (codec *ResponsesCodec) IsEmpty(ptr unsafe.Pointer) bool {
	target := (*Responses)(ptr)
	return target == nil || len(target.Vals) == 0
}
