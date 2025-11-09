package core

import (
	"fmt"
	"strconv"
	"unsafe"

	jsoniter "github.com/json-iterator/go"
)

func init() {
	RegisterJSONTypeDecoder("core.BoolValue", &BoolValueCodec{})
	RegisterJSONTypeEncoder("core.BoolValue", &BoolValueCodec{})
}

type BoolValueCodec struct {
}

func (codec *BoolValueCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	val := iter.ReadAny()
	if val.ValueType() == jsoniter.BoolValue {
		(*BoolValue)(ptr).Val = val.ToBool()
	} else {
		iter.ReportError("bool", fmt.Sprintf("invalid bool type value, original type: %d", val.ValueType()))
	}
}

func (codec *BoolValueCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	e := (*BoolValue)(ptr)
	stream.WriteRaw(strconv.FormatBool(e.Val))
}

func (codec *BoolValueCodec) IsEmpty(ptr unsafe.Pointer) bool {
	e := (*BoolValue)(ptr)
	return e == nil
}
