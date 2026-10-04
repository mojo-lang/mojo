package openapi

import (
	"io"
	"unsafe"

	jsoniter "github.com/json-iterator/go"
	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
)

func init() {
	jsoniter.RegisterTypeDecoder("openapi.Reference", &ReferenceCodec{})
	jsoniter.RegisterTypeEncoder("openapi.Reference", &ReferenceCodec{})
}

type ReferenceCodec struct {
}

func (codec *ReferenceCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	any := iter.ReadAny()
	if iter.Error != nil && iter.Error != io.EOF {
		return
	}
	r := (*Reference)(ptr)
	if any.ValueType() == jsoniter.NilValue {
		r.Ref, r.Summary, r.Description = nil, "", nil
		return
	}
	if any.ValueType() != jsoniter.ObjectValue || any.Get("$ref").ValueType() != jsoniter.StringValue {
		iter.ReportError("ReferenceCodec.Decode", "expected an object with a string $ref")
		return
	}
	var value struct {
		Ref         string          `json:"$ref"`
		Summary     string          `json:"summary"`
		Description *CachedDocument `json:"description"`
	}
	if !decodeJSONValue(iter, any, &value) {
		return
	}
	url, err := core.NewUrl(value.Ref)
	if err != nil {
		iter.ReportError("ReferenceCodec.Decode", err.Error())
		return
	}
	r.Ref, r.Summary, r.Description = url, value.Summary, value.Description
}

func (codec *ReferenceCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	reference := (*Reference)(ptr)

	if reference != nil && reference.Ref != nil {
		stream.WriteObjectStart()
		stream.WriteObjectField("$ref")
		stream.WriteString(reference.Ref.Format())

		if len(reference.Summary) > 0 {
			stream.WriteMore()
			stream.WriteObjectField("summary")
			stream.WriteVal(reference.Summary)
		}
		if reference.Description != nil {
			stream.WriteMore()
			stream.WriteObjectField("description")
			stream.WriteVal(reference.Description)
		}

		stream.WriteObjectEnd()
	} else {
		stream.WriteNil()
	}
}

func (codec *ReferenceCodec) IsEmpty(ptr unsafe.Pointer) bool {
	reference := (*Reference)(ptr)
	return reference == nil || reference.Ref == nil
}
