package openapi

import (
	"io"
	"unsafe"

	jsoniter "github.com/json-iterator/go"
)

func init() {
	jsoniter.RegisterTypeDecoder("openapi.ReferenceableResponse", &ReferenceableResponseCodec{})
	jsoniter.RegisterTypeEncoder("openapi.ReferenceableResponse", &ReferenceableResponseCodec{})
}

type ReferenceableResponseCodec struct {
}

func (codec *ReferenceableResponseCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	value := iter.ReadAny()
	if iter.Error != nil && iter.Error != io.EOF {
		return
	}
	wrapper := (*ReferenceableResponse)(ptr)
	if value.ValueType() == jsoniter.NilValue {
		wrapper.ReferenceableResponse = nil
		return
	}
	if value.ValueType() != jsoniter.ObjectValue {
		iter.ReportError("ReferenceableResponseCodec.Decode", "expected an object or null")
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
		inline := &Response{}
		if decodeJSONValue(iter, value, inline) {
			wrapper.SetResponse(inline)
		}
	default:
		iter.ReportError("ReferenceableResponseCodec.Decode", "$ref must be a string")
	}
}

func (codec *ReferenceableResponseCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	wrapper := (*ReferenceableResponse)(ptr)
	if value := wrapper.GetResponse(); value != nil {
		stream.WriteVal(value)
	} else if reference := wrapper.GetReference(); reference != nil {
		stream.WriteVal(reference)
	} else {
		stream.WriteNil()
	}
}

func (codec *ReferenceableResponseCodec) IsEmpty(ptr unsafe.Pointer) bool {
	wrapper := (*ReferenceableResponse)(ptr)
	return wrapper == nil || wrapper.ReferenceableResponse == nil
}
