package openapi

import (
	"io"
	"unsafe"

	jsoniter "github.com/json-iterator/go"
)

func init() {
	jsoniter.RegisterTypeDecoder("openapi.ReferenceableExample", &ReferenceableExampleCodec{})
	jsoniter.RegisterTypeEncoder("openapi.ReferenceableExample", &ReferenceableExampleCodec{})
}

type ReferenceableExampleCodec struct {
}

func (codec *ReferenceableExampleCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	value := iter.ReadAny()
	if iter.Error != nil && iter.Error != io.EOF {
		return
	}
	wrapper := (*ReferenceableExample)(ptr)
	if value.ValueType() == jsoniter.NilValue {
		wrapper.ReferenceableExample = nil
		return
	}
	if value.ValueType() != jsoniter.ObjectValue {
		iter.ReportError("ReferenceableExampleCodec.Decode", "expected an object or null")
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
		inline := &Example{}
		if decodeJSONValue(iter, value, inline) {
			wrapper.SetExample(inline)
		}
	default:
		iter.ReportError("ReferenceableExampleCodec.Decode", "$ref must be a string")
	}
}

func (codec *ReferenceableExampleCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	wrapper := (*ReferenceableExample)(ptr)
	if value := wrapper.GetExample(); value != nil {
		stream.WriteVal(value)
	} else if reference := wrapper.GetReference(); reference != nil {
		stream.WriteVal(reference)
	} else {
		stream.WriteNil()
	}
}

func (codec *ReferenceableExampleCodec) IsEmpty(ptr unsafe.Pointer) bool {
	wrapper := (*ReferenceableExample)(ptr)
	return wrapper == nil || wrapper.ReferenceableExample == nil
}
