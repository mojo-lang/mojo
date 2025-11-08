package openapi

import (
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
	any := iter.ReadAny()
	link := (*ReferenceableLink)(ptr)
	if any.ValueType() == jsoniter.ObjectValue {
		ref := any.Get("$ref")
		if ref.ValueType() == jsoniter.StringValue {
			r := &Reference{}
			any.ToVal(r)
			link.SetReference(r)
		} else {
			s := &Link{}
			any.ToVal(s)
			link.SetLink(s)
		}
	}
}

func (codec *ReferenceableLinkCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	link := (*ReferenceableLink)(ptr)
	if s := link.GetLink(); s != nil {
		stream.WriteVal(s)
	} else if r := link.GetReference(); r != nil {
		stream.WriteVal(r)
	}
}

func (codec *ReferenceableLinkCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return ((*ReferenceableLink)(ptr)).ReferenceableLink == nil
}
