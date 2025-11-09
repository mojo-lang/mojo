package openapi

import (
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
	any := iter.ReadAny()
	response := (*ReferenceableResponse)(ptr)
	if any.ValueType() == jsoniter.ObjectValue {
		ref := any.Get("$ref")
		if ref.ValueType() == jsoniter.StringValue {
			r := &Reference{}
			any.ToVal(r)
			response.SetReference(r)
		} else {
			s := &Response{}
			any.ToVal(s)
			response.SetResponse(s)
		}
	}
}

func (codec *ReferenceableResponseCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	response := (*ReferenceableResponse)(ptr)
	if s := response.GetResponse(); s != nil {
		stream.WriteVal(s)
	} else if r := response.GetReference(); r != nil {
		stream.WriteVal(r)
	}
}

func (codec *ReferenceableResponseCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return ((*ReferenceableResponse)(ptr)).ReferenceableResponse == nil
}
