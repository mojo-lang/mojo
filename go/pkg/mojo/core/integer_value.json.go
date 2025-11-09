package core

import (
	"fmt"
	"strconv"
	"unsafe"

	jsoniter "github.com/json-iterator/go"
)

func init() {
	RegisterJSONTypeDecoder("core.Int32Value", &Int32ValueCodec{})
	RegisterJSONTypeEncoder("core.Int32Value", &Int32ValueCodec{})
	RegisterJSONTypeDecoder("core.UInt32Value", &UInt32ValueCodec{})
	RegisterJSONTypeEncoder("core.UInt32Value", &UInt32ValueCodec{})
	RegisterJSONTypeDecoder("core.Int64Value", &Int64ValueCodec{})
	RegisterJSONTypeEncoder("core.Int64Value", &Int64ValueCodec{})
	RegisterJSONTypeDecoder("core.UInt64Value", &UInt64ValueCodec{})
	RegisterJSONTypeEncoder("core.UInt64Value", &UInt64ValueCodec{})
}

type Int32ValueCodec struct {
}

func (codec *Int32ValueCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	val := iter.ReadAny()
	if val.ValueType() == jsoniter.NumberValue {
		(*Int32Value)(ptr).Val = val.ToInt32()
	} else {
		iter.ReportError("number", fmt.Sprintf("invalid int32 type value, original type: %d", val.ValueType()))
	}
}

func (codec *Int32ValueCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	v := (*Int32Value)(ptr)
	stream.WriteRaw(strconv.FormatInt(int64(v.Val), 10))
}

func (codec *Int32ValueCodec) IsEmpty(ptr unsafe.Pointer) bool {
	v := (*Int32Value)(ptr)
	return v == nil
}

type UInt32ValueCodec struct {
}

func (codec *UInt32ValueCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	val := iter.ReadAny()
	if val.ValueType() == jsoniter.NumberValue {
		(*UInt32Value)(ptr).Val = val.ToUint32()
	} else {
		iter.ReportError("number", fmt.Sprintf("invalid uint32 type value, original type: %d", val.ValueType()))
	}
}

func (codec *UInt32ValueCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	v := (*UInt32Value)(ptr)
	stream.WriteRaw(strconv.FormatUint(uint64(v.Val), 10))
}

func (codec *UInt32ValueCodec) IsEmpty(ptr unsafe.Pointer) bool {
	v := (*UInt32Value)(ptr)
	return v == nil
}

type Int64ValueCodec struct {
}

func (codec *Int64ValueCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	val := iter.ReadAny()
	if val.ValueType() == jsoniter.NumberValue {
		(*Int64Value)(ptr).Val = val.ToInt64()
	} else {
		iter.ReportError("number", fmt.Sprintf("invalid int64 type value, original type: %d", val.ValueType()))
	}
}

func (codec *Int64ValueCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	v := (*Int64Value)(ptr)
	stream.WriteRaw(strconv.FormatInt(v.Val, 10))
}

func (codec *Int64ValueCodec) IsEmpty(ptr unsafe.Pointer) bool {
	v := (*Int64Value)(ptr)
	return v == nil
}

type UInt64ValueCodec struct {
}

func (codec *UInt64ValueCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	val := iter.ReadAny()
	if val.ValueType() == jsoniter.NumberValue {
		(*UInt64Value)(ptr).Val = val.ToUint64()
	} else {
		iter.ReportError("number", fmt.Sprintf("invalid uint64 type value, original type: %d", val.ValueType()))
	}
}

func (codec *UInt64ValueCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	v := (*UInt64Value)(ptr)
	stream.WriteRaw(strconv.FormatUint(v.Val, 10))
}

func (codec *UInt64ValueCodec) IsEmpty(ptr unsafe.Pointer) bool {
	v := (*UInt64Value)(ptr)
	return v == nil
}
