package core

import (
	"errors"
	"strconv"
	"unsafe"

	jsoniter "github.com/json-iterator/go"
)

func init() {
	RegisterJSONTypeDecoder("core.Id", &IdCodec{})
	RegisterJSONTypeEncoder("core.Id", &IdCodec{})
}

type IdCodec struct {
}

func NewIdCodec() *IdCodec {
	return &IdCodec{}
}

func (codec *IdCodec) DecodeAny(any jsoniter.Any) (*Id, error) {
	switch any.ValueType() {
	case jsoniter.NumberValue:
		value, err := strconv.ParseUint(any.ToString(), 10, 64)
		if err != nil {
			return nil, err
		}
		return NewIntId(value), nil
	case jsoniter.StringValue:
		return ParseId(any.ToString())
	}
	return nil, errors.New("type is invalid for ID")
}

func (codec *IdCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	id, err := codec.DecodeAny(iter.ReadAny())
	if err != nil {
		iter.ReportError("Id Decode", err.Error())
		return
	}
	(*Id)(ptr).Id = id.Id
}

func (codec *IdCodec) IsEmpty(ptr unsafe.Pointer) bool {
	id := (*Id)(ptr)
	return id == nil || id.Id == nil
}

func (codec *IdCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	id := (*Id)(ptr)
	stream.WriteVal(id.Format())
}
