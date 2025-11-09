package core

import (
	"unsafe"

	jsoniter "github.com/json-iterator/go"
)

func init() {
	RegisterJSONTypeDecoder("core.Regex", &RegexStringCodec{})
	RegisterJSONTypeEncoder("core.Regex", &RegexStringCodec{})
}

// BareRegex will be jsonify to raw, without any codec
type BareRegex FieldMask

type RegexStringCodec struct {
	IsFieldPointer bool
}

func (codec *RegexStringCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	s := iter.ReadString()
	regex := codec.regex(ptr)
	if regex == nil {
		regex = &Regex{}
		*(**Regex)(ptr) = regex
	}

	if err := regex.Parse(s); err != nil {
		iter.ReportError("RegexStringCodec", err.Error())
	}
}

func (codec *RegexStringCodec) IsEmpty(ptr unsafe.Pointer) bool {
	regex := codec.regex(ptr)
	if regex != nil {
		if checker, ok := interface{}(regex).(EmptyChecker); ok {
			return checker.IsEmpty()
		}
		return false
	}
	return true
}

func (codec *RegexStringCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	regex := codec.regex(ptr)
	stream.WriteString(regex.Format())
}

func (codec *RegexStringCodec) regex(ptr unsafe.Pointer) *Regex {
	if codec.IsFieldPointer {
		return *(**Regex)(ptr)
	}
	return (*Regex)(ptr)
}
