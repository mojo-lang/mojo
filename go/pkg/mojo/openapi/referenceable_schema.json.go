package openapi

import (
	"io"
	"unsafe"

	jsoniter "github.com/json-iterator/go"
)

func init() {
	jsoniter.RegisterTypeDecoder("openapi.ReferenceableSchema", &ReferenceableSchemaCodec{})
	jsoniter.RegisterTypeEncoder("openapi.ReferenceableSchema", &ReferenceableSchemaCodec{})
}

type ReferenceableSchemaCodec struct {
}

func (codec *ReferenceableSchemaCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	value := iter.ReadAny()
	if iter.Error != nil && iter.Error != io.EOF {
		return
	}
	wrapper := (*ReferenceableSchema)(ptr)
	if value.ValueType() == jsoniter.NilValue {
		wrapper.ReferenceableSchema = nil
		return
	}
	if value.ValueType() != jsoniter.ObjectValue {
		iter.ReportError("ReferenceableSchemaCodec.Decode", "expected an object or null")
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
		inline := &Schema{}
		if decodeJSONValue(iter, value, inline) {
			wrapper.SetSchema(inline)
		}
	default:
		iter.ReportError("ReferenceableSchemaCodec.Decode", "$ref must be a string")
	}
}

func (codec *ReferenceableSchemaCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	wrapper := (*ReferenceableSchema)(ptr)
	if value := wrapper.GetSchema(); value != nil {
		stream.WriteVal(value)
	} else if reference := wrapper.GetReference(); reference != nil {
		stream.WriteVal(reference)
	} else {
		stream.WriteNil()
	}
}

func (codec *ReferenceableSchemaCodec) IsEmpty(ptr unsafe.Pointer) bool {
	wrapper := (*ReferenceableSchema)(ptr)
	return wrapper == nil || wrapper.ReferenceableSchema == nil
}
