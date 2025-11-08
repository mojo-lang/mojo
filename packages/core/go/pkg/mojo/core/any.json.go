package core

import (
	"fmt"
	"reflect"
	"sort"
	"unsafe"

	"google.golang.org/protobuf/reflect/protoregistry"

	jsoniter "github.com/json-iterator/go"
	"github.com/modern-go/reflect2"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/mojo-lang/mojo/packages/core/go/pkg/logs"
	"github.com/mojo-lang/mojo/packages/core/go/pkg/mojo/core/strcase"
)

const (
	typeFieldName   = "@type"
	valueFieldName  = "value"
	valuesFieldName = "values"
)

func (x *Any) MarshalJSON() ([]byte, error) {
	return jsoniter.ConfigFastest.Marshal(x)
}

func (x *Any) UnmarshalJSON(err []byte) error {
	return jsoniter.ConfigFastest.Unmarshal(err, x)
}

func init() {
	RegisterJSONTypeDecoder("core.Any", &AnyCodec{})
	RegisterJSONTypeEncoder("core.Any", &AnyCodec{})
}

// AnyCodec
// FIXME should implement all the support types codec
type AnyCodec struct {
}

func (codec *AnyCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	any := iter.ReadAny()
	anyVal := (*Any)(ptr)
	if any.ValueType() == jsoniter.ObjectValue {
		anyVal.Type = any.Get(typeFieldName).ToString()
		tn, err := ParseTypeName(anyVal.Type)
		if err != nil {
			iter.ReportError("Any Decode", err.Error())
		}

		switch tn.GetFullName() {
		case Int64TypeName, Int64TypeFullName:
			anyVal.typedVal = any.Get(valueFieldName).ToInt64()
		case StringTypeName:
			anyVal.typedVal = any.Get(valueFieldName).ToString()
		default:
			if msgType, err := protoregistry.GlobalTypes.FindMessageByName(protoreflect.FullName(anyVal.Type)); err != nil {
				iter.ReportError("Any Decode", err.Error())
				return
			} else {
				msg := msgType.New().Interface()
				if _, ok := msg.(ScalarLike); ok {
					if parser, ok := msg.(Parser); ok {
						if err = parser.Parse(any.Get(valueFieldName).ToString()); err != nil {
							iter.ReportError("Any Decode", err.Error())
						}
					} else if float64Convert, ok := msg.(FromFloat64Converter); ok {
						if err = float64Convert.FromFloat64(any.Get(valueFieldName).ToFloat64()); err != nil {
							iter.ReportError("Any Decode", err.Error())
						}
					} else {
						iter.ReportError("Any Decode", fmt.Sprintf("not support type for %s", anyVal.Type))
					}
				} else if _, ok = msg.(ArrayLike); ok {
					any.Get(valuesFieldName).ToVal(msg)
				} else {
					any.ToVal(msg)
				}

				anyVal.typedVal = msg
			}
		}
	}
}

func (codec *AnyCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*Any)(ptr).Empty()
}

func (codec *AnyCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	kvWriter := func(t string, v interface{}) {
		stream.WriteObjectStart()
		stream.WriteObjectField(typeFieldName)
		stream.WriteVal(t)
		stream.WriteMore()
		if _, ok := v.(ArrayLike); ok {
			stream.WriteObjectField(valuesFieldName)
		} else {
			stream.WriteObjectField(valueFieldName)
		}
		stream.WriteVal(v)
		stream.WriteObjectEnd()
	}

	any := (*Any)(ptr)
	t := reflect2.TypeOf(any.typedVal)
	switch v := any.typedVal.(type) {
	case proto.Message:
		reflectMsg := v.ProtoReflect()

		if value, ok := any.typedVal.(ScalarLike); ok {
			kvWriter(string(reflectMsg.Descriptor().FullName()), value.ToScalar())
		} else {
			fullName := string(reflectMsg.Descriptor().FullName())

			stream.WriteObjectStart()
			stream.WriteObjectField(typeFieldName)
			stream.WriteVal(fullName)

			if t.Kind() == reflect.Ptr {
				t = t.(*reflect2.UnsafePtrType).Elem()
			}
			if obj, ok := t.(*reflect2.UnsafeStructType); ok {
				var fields Fields
				reflectMsg.Range(func(descriptor protoreflect.FieldDescriptor, value protoreflect.Value) bool {
					fields = append(fields, descriptor)
					return true
				})

				sort.Sort(fields)
				for _, descriptor := range fields {
					stream.WriteMore()
					stream.WriteObjectField(descriptor.JSONName())

					fieldName := strcase.ToCamel(string(descriptor.Name()))
					field := obj.FieldByName(fieldName)
					if encoder, ok := registeredJsonEncoderTypeFields[t.String()+"."+fieldName]; ok {
						f := field.UnsafeGet(reflect2.PtrOf(any.typedVal))
						if !encoder.IsEmpty(f) {
							encoder.Encode(f, stream)
						} else {
							stream.WriteVal(nil)
						}
					} else {
						stream.WriteVal(field.Get(any.typedVal))
					}
				}
			} else {
				logs.Errorw("invalid the struct type for any object", "type", fullName)
			}

			stream.WriteObjectEnd()
		}
	default:
		if t.Kind() == reflect.Map {
			if um, ok := t.(*reflect2.UnsafeMapType); ok && !um.IsNil(any.typedVal) {
				itr := um.Iterate(um.Indirect(any.typedVal))
				stream.WriteObjectStart()
				stream.WriteObjectField(typeFieldName)
				stream.WriteVal(GetMojoTypeName(any.typedVal))

				for itr.HasNext() {
					stream.WriteMore()

					k, value := itr.Next()
					stream.WriteObjectField(ToString(k))
					stream.WriteVal(value)
				}
				stream.WriteObjectEnd()
			}
		} else {
			kvWriter(GetMojoTypeName(any.typedVal), v)
		}
	}
}

type Fields []protoreflect.FieldDescriptor

func (f Fields) Len() int {
	return len(f)
}
func (f Fields) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}
func (f Fields) Less(i, j int) bool {
	return f[i].JSONName() < f[j].JSONName()
}
