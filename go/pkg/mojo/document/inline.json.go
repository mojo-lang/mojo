package document

import (
	"fmt"
	"io"
	"unsafe"

	jsoniter "github.com/json-iterator/go"
)

func init() {
	jsoniter.RegisterTypeDecoder("document.Inline", &InlineCodec{})
	jsoniter.RegisterTypeEncoder("document.Inline", &InlineCodec{})
}

type InlineCodec struct {
}

func (codec *InlineCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	any := iter.ReadAny()
	if iter.Error != nil && iter.Error != io.EOF {
		return
	}
	inline := (*Inline)(ptr)
	switch any.ValueType() {
	case jsoniter.NilValue:
		inline.Inline = nil
	case jsoniter.StringValue:
		inline.SetText(any.ToString())
	case jsoniter.ObjectValue:
		t := any.Get("@type").ToString()
		switch t {
		case "Link", "mojo.document.Link":
			val := &Link{}
			// Any.ToVal discards decoding errors, including invalid nested inlines.
			decoder := iter.Pool().BorrowIterator([]byte(any.ToString()))
			defer iter.Pool().ReturnIterator(decoder)
			decoder.ReadVal(val)
			if decoder.Error != nil {
				iter.ReportError("decode document.Inline", decoder.Error.Error())
				return
			}
			inline.Inline = &Inline_Link{val}
		default:
			iter.ReportError("decode document.Inline", fmt.Sprintf("unsupported @type %q", t))
		}
	default:
		iter.ReportError("decode document.Inline", "expected a string, a supported @type object, or null")
	}
}

func (codec *InlineCodec) IsEmpty(ptr unsafe.Pointer) bool {
	inline := (*Inline)(ptr)
	return inline == nil || inline.Inline == nil
}

func (codec *InlineCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	inline := (*Inline)(ptr)
	if inline == nil || inline.Inline == nil {
		stream.WriteNil()
		return
	}
	switch x := inline.Inline.(type) {
	case *Inline_Text:
		if x == nil || x.Text == nil {
			stream.WriteNil()
			return
		}
		stream.WriteVal(x.Text.Val)
	case *Inline_Link:
		if x == nil || x.Link == nil {
			stream.WriteNil()
			return
		}
		stream.WriteVal(struct {
			Type string `json:"@type"`
			*Link
		}{Type: "Link", Link: x.Link})
	default:
		stream.Error = fmt.Errorf("encode document.Inline: unsupported variant %T", inline.Inline)
	}
}
