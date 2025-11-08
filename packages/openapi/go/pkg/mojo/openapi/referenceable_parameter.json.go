package openapi

import (
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
	any := iter.ReadAny()
	parameter := (*ReferenceableParameter)(ptr)
	if any.ValueType() == jsoniter.ObjectValue {
		ref := any.Get("$ref")
		if ref.ValueType() == jsoniter.StringValue {
			r := &Reference{}
			any.ToVal(r)
			parameter.SetReference(r)
		} else {
			s := &Parameter{}
			any.ToVal(s)
			parameter.SetParameter(s)
		}
	}
}

func (codec *ReferenceableParameterCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	parameter := (*ReferenceableParameter)(ptr)
	if s := parameter.GetParameter(); s != nil {
		stream.WriteVal(s)
	} else if r := parameter.GetReference(); r != nil {
		stream.WriteVal(r)
	}
}

func (codec *ReferenceableParameterCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return ((*ReferenceableParameter)(ptr)).ReferenceableParameter == nil
}
