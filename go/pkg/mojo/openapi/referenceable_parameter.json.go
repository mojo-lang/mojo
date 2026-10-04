package openapi

import (
	"io"
	"unsafe"

	jsoniter "github.com/json-iterator/go"
)

func init() {
	jsoniter.RegisterTypeDecoder("openapi.ReferenceableParameter", &ReferenceableParameterCodec{})
	jsoniter.RegisterTypeEncoder("openapi.ReferenceableParameter", &ReferenceableParameterCodec{})
}

type ReferenceableParameterCodec struct {
}

func (codec *ReferenceableParameterCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	value := iter.ReadAny()
	if iter.Error != nil && iter.Error != io.EOF {
		return
	}
	wrapper := (*ReferenceableParameter)(ptr)
	if value.ValueType() == jsoniter.NilValue {
		wrapper.ReferenceableParameter = nil
		return
	}
	if value.ValueType() != jsoniter.ObjectValue {
		iter.ReportError("ReferenceableParameterCodec.Decode", "expected an object or null")
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
		inline := &Parameter{}
		if decodeJSONValue(iter, value, inline) {
			wrapper.SetParameter(inline)
		}
	default:
		iter.ReportError("ReferenceableParameterCodec.Decode", "$ref must be a string")
	}
}

func (codec *ReferenceableParameterCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	wrapper := (*ReferenceableParameter)(ptr)
	if value := wrapper.GetParameter(); value != nil {
		stream.WriteVal(value)
	} else if reference := wrapper.GetReference(); reference != nil {
		stream.WriteVal(reference)
	} else {
		stream.WriteNil()
	}
}

func (codec *ReferenceableParameterCodec) IsEmpty(ptr unsafe.Pointer) bool {
	wrapper := (*ReferenceableParameter)(ptr)
	return wrapper == nil || wrapper.ReferenceableParameter == nil
}
