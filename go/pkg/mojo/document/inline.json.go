package document

import (
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
	inline := (*Inline)(ptr)
	switch any.ValueType() {
	case jsoniter.StringValue:
	case jsoniter.BoolValue:
	case jsoniter.NumberValue:
	case jsoniter.ArrayValue:
	case jsoniter.ObjectValue:
		t := any.Get("@type").ToString()
		switch t {
		case "Link", "mojo.document.Link":
			val := &Link{}
			any.ToVal(val)
			inline.Inline = &Inline_Link{val}
		}
	}
}

func (codec *InlineCodec) IsEmpty(ptr unsafe.Pointer) bool {
	inline := (*Inline)(ptr)
	return inline == nil || inline.Inline == nil
}

func (codec *InlineCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	inline := (*Inline)(ptr)
	switch x := inline.Inline.(type) {
	case *Inline_Text:
		stream.WriteVal(x.Text.Val)
	}
}
