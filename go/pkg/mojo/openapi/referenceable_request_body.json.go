package openapi

import (
	"io"
	"unsafe"

	jsoniter "github.com/json-iterator/go"
)

func init() {
	jsoniter.RegisterTypeDecoder("openapi.ReferenceableRequestBody", &ReferenceableRequestBodyCodec{})
	jsoniter.RegisterTypeEncoder("openapi.ReferenceableRequestBody", &ReferenceableRequestBodyCodec{})
}

type ReferenceableRequestBodyCodec struct {
}

func (codec *ReferenceableRequestBodyCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	value := iter.ReadAny()
	if iter.Error != nil && iter.Error != io.EOF {
		return
	}
	wrapper := (*ReferenceableRequestBody)(ptr)
	if value.ValueType() == jsoniter.NilValue {
		wrapper.ReferenceableRequestBody = nil
		return
	}
	if value.ValueType() != jsoniter.ObjectValue {
		iter.ReportError("ReferenceableRequestBodyCodec.Decode", "expected an object or null")
		return
	}
	ref := value.Get("$ref")
	switch ref.ValueType() {
	case jsoniter.StringValue:
		reference := &Reference{}
		if decodeJSONValue(iter, value, reference) {
			wrapper.SetReference(reference)
		}
	case jsoniter.InvalidValue:
		inline := &RequestBody{}
		if decodeJSONValue(iter, value, inline) {
			wrapper.SetRequestBody(inline)
		}
	default:
		iter.ReportError("ReferenceableRequestBodyCodec.Decode", "$ref must be a string")
	}
}

func (codec *ReferenceableRequestBodyCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	wrapper := (*ReferenceableRequestBody)(ptr)
	if value := wrapper.GetRequestBody(); value != nil {
		stream.WriteVal(value)
	} else if reference := wrapper.GetReference(); reference != nil {
		stream.WriteVal(reference)
	} else {
		stream.WriteNil()
	}
}

func (codec *ReferenceableRequestBodyCodec) IsEmpty(ptr unsafe.Pointer) bool {
	wrapper := (*ReferenceableRequestBody)(ptr)
	return wrapper == nil || wrapper.ReferenceableRequestBody == nil
}
