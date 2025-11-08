package core

import (
	"unsafe"

	jsoniter "github.com/json-iterator/go"
)

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

func init() {
	RegisterJSONTypeDecoder("core.Timestamp", &TimestampCodec{})
	RegisterJSONTypeEncoder("core.Timestamp", &TimestampCodec{})
}

type TimestampCodec struct {
}

func (codec *TimestampCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	any := iter.ReadAny()
	ts := (*Timestamp)(ptr)
	if any.ValueType() == jsoniter.NumberValue {
		number := any.ToInt64()

		if number < int64(MaxInt) {
			ts.Seconds = number
		} else {
			ts.Seconds = number / 1000
			ts.Nanoseconds = int32((number - ts.Seconds*1000) * 1000000)
		}
	} else if any.ValueType() == jsoniter.StringValue {
		if err := ts.Parse(any.ToString()); err != nil {
			iter.ReportError("Decode Timestamp", err.Error())
		}
	}
}

func (codec *TimestampCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return ((*Timestamp)(ptr)).Seconds == 0
}

func (codec *TimestampCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	stream.WriteString((*Timestamp)(ptr).Format())
}
