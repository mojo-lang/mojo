package openapi

import (
	"io"
	"unsafe"

	jsoniter "github.com/json-iterator/go"
)

func init() {
	jsoniter.RegisterTypeDecoder("openapi.ReferenceableLink", &ReferenceableLinkCodec{})
	jsoniter.RegisterTypeEncoder("openapi.ReferenceableLink", &ReferenceableLinkCodec{})
}

type ReferenceableLinkCodec struct {
}

func (codec *ReferenceableLinkCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	value := iter.ReadAny()
	if iter.Error != nil && iter.Error != io.EOF {
		return
	}
	wrapper := (*ReferenceableLink)(ptr)
	if value.ValueType() == jsoniter.NilValue {
		wrapper.ReferenceableLink = nil
		return
	}
	if value.ValueType() != jsoniter.ObjectValue {
		iter.ReportError("ReferenceableLinkCodec.Decode", "expected an object or null")
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
		inline := &Link{}
		if decodeJSONValue(iter, value, inline) {
			wrapper.SetLink(inline)
		}
	default:
		iter.ReportError("ReferenceableLinkCodec.Decode", "$ref must be a string")
	}
}

func (codec *ReferenceableLinkCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	wrapper := (*ReferenceableLink)(ptr)
	if value := wrapper.GetLink(); value != nil {
		stream.WriteVal(value)
	} else if reference := wrapper.GetReference(); reference != nil {
		stream.WriteVal(reference)
	} else {
		stream.WriteNil()
	}
}

func (codec *ReferenceableLinkCodec) IsEmpty(ptr unsafe.Pointer) bool {
	wrapper := (*ReferenceableLink)(ptr)
	return wrapper == nil || wrapper.ReferenceableLink == nil
}
