package db

import (
	"unsafe"

	jsoniter "github.com/json-iterator/go"
	"github.com/mojo-lang/mojo/packages/core/go/pkg/mojo/core"
)

func init() {
	jsoniter.RegisterTypeDecoder("core.DeleteTime", &DeleteTimeCodec{})
	jsoniter.RegisterTypeEncoder("core.DeleteTime", &DeleteTimeCodec{})
}

type DeleteTimeCodec struct {
}

func (codec *DeleteTimeCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	a := iter.ReadAny()
	deleteTime := (*DeleteTime)(ptr)
	deleteTime.Val = &core.Timestamp{}
	a.ToVal(deleteTime.Val)
}

func (codec *DeleteTimeCodec) IsEmpty(ptr unsafe.Pointer) bool {
	deleteTime := (*DeleteTime)(ptr)
	return deleteTime == nil || deleteTime.Val == nil || (deleteTime.Val.Seconds == 0 && deleteTime.Val.Nanoseconds == 0)
}

func (codec *DeleteTimeCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	deleteTime := (*DeleteTime)(ptr)
	stream.WriteString(deleteTime.Format())
}
