package core

import (
	"unsafe"

	jsoniter "github.com/json-iterator/go"
)

func init() {
	RegisterJSONTypeDecoder("core.Duration", &DurationCodec{})
	RegisterJSONTypeEncoder("core.Duration", &DurationCodec{})
}

type DurationCodec struct {
}

func (codec *DurationCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	any := iter.ReadAny()
	duration := (*Duration)(ptr)
	if any.ValueType() == jsoniter.NumberValue {
		number := any.ToFloat64()
		duration.FromSeconds(number)
	} else if any.ValueType() == jsoniter.StringValue {
		str := any.ToString()
		err := duration.Parse(str)
		if err != nil {
			iter.ReportError("Decode Duration", err.Error())
		}
	}
}

func (codec *DurationCodec) IsEmpty(ptr unsafe.Pointer) bool {
	duration := (*Duration)(ptr)
	return duration == nil || (duration.Seconds == 0 && duration.Nanoseconds == 0)
}

func (codec *DurationCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	duration := (*Duration)(ptr)
	stream.WriteString(duration.Format())
}
