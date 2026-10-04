package openapi

import (
	"io"
	"unsafe"

	jsoniter "github.com/json-iterator/go"
	"github.com/mojo-lang/mojo/go/pkg/mojo/document"
)

func init() {
	jsoniter.RegisterTypeDecoder("openapi.CachedDocument", &CachedDocumentCodec{})
	jsoniter.RegisterTypeEncoder("openapi.CachedDocument", &CachedDocumentCodec{})
}

type CachedDocumentCodec struct {
}

func (codec *CachedDocumentCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	any := iter.ReadAny()
	if iter.Error != nil && iter.Error != io.EOF {
		return
	}
	doc := (*CachedDocument)(ptr)
	switch any.ValueType() {
	case jsoniter.ObjectValue:
		value := &document.Document{}
		if decodeJSONValue(iter, any, value) {
			doc.Val, doc.Cache = value, ""
		}
	case jsoniter.StringValue:
		doc.Val, doc.Cache = nil, any.ToString()
	case jsoniter.NilValue:
		doc.Val, doc.Cache = nil, ""
	default:
		iter.ReportError("CachedDocumentCodec.Decode", "expected a string, document object, or null")
	}
}

func (codec *CachedDocumentCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	doc := (*CachedDocument)(ptr)
	stream.WriteVal(doc.GetStringDocument())
}

func (codec *CachedDocumentCodec) IsEmpty(ptr unsafe.Pointer) bool {
	doc := (*CachedDocument)(ptr)
	return doc == nil || (len(doc.Cache) == 0 && doc.Val == nil)
}
