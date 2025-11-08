package openapi

import (
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
	any := iter.ReadAny()
	requestBody := (*ReferenceableRequestBody)(ptr)
	if any.ValueType() == jsoniter.ObjectValue {
		ref := any.Get("$ref")
		if ref.ValueType() == jsoniter.StringValue {
			r := &Reference{}
			any.ToVal(r)
			requestBody.SetReference(r)
		} else {
			r := &RequestBody{}
			any.ToVal(r)
			requestBody.SetRequestBody(r)
		}
	}
}

func (codec *ReferenceableRequestBodyCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	requestBody := (*ReferenceableRequestBody)(ptr)
	if r := requestBody.GetRequestBody(); r != nil {
		stream.WriteVal(r)
	} else if r := requestBody.GetReference(); r != nil {
		stream.WriteVal(r)
	}
}

func (codec *ReferenceableRequestBodyCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return ((*ReferenceableRequestBody)(ptr)).ReferenceableRequestBody == nil
}
