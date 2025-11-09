package openapi

import (
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
	any := iter.ReadAny()
	example := (*ReferenceableExample)(ptr)
	if any.ValueType() == jsoniter.ObjectValue {
		ref := any.Get("$ref")
		if ref.ValueType() == jsoniter.StringValue {
			r := &Reference{}
			any.ToVal(r)
			example.SetReference(r)
		} else {
			s := &Example{}
			any.ToVal(s)
			example.SetExample(s)
		}
	}
}

func (codec *ReferenceableExampleCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	example := (*ReferenceableExample)(ptr)
	if s := example.GetExample(); s != nil {
		stream.WriteVal(s)
	} else if r := example.GetReference(); r != nil {
		stream.WriteVal(r)
	}
}

func (codec *ReferenceableExampleCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return ((*ReferenceableExample)(ptr)).ReferenceableExample == nil
}
