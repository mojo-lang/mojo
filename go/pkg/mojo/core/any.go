package core

import (
	"errors"
	"reflect"

	"google.golang.org/protobuf/proto"
)

const AnyTypeName = "Any"
const AnyTypeFullName = "mojo.core.Any"

func NewAny(i interface{}) *Any {
	any := &Any{}
	any.Set(i)
	return any
}

func (x *Any) Get() interface{} {
	if x != nil {
		return x.typedVal
	}
	return nil
}

func (x *Any) Set(v interface{}) error {
	if x != nil {
		x.Type = GetMojoTypeName(v)
		if len(x.Type) == 0 {
			return errors.New("unsupported type")
		}
		x.typedVal = v
	}
	return nil
}

func (x *Any) Empty() bool {
	if x != nil {
		return len(x.Type) == 0 && x.typedVal == nil
	}
	return true
}

func (x *Any) GetMessage() proto.Message {
	if msg, ok := x.typedVal.(proto.Message); ok {
		return msg
	}
	return nil
}

func GetMojoTypeName(i interface{}) string {
	switch v := i.(type) {
	case nil:
		return NullTypeName
	case bool:
		return BoolTypeName
	case int8:
		return Int8TypeName
	case uint8:
		return UInt8TypeName
	case int16:
		return Int16TypeName
	case uint16:
		return UInt16TypeName
	case int32:
		return Int32TypeName
	case uint32:
		return UInt32TypeName
	case int64:
		return Int64TypeName
	case uint64:
		return UInt64TypeName
	case int:
		return IntTypeName
	case uint:
		return UIntTypeName
	case float32:
		return Float32TypeName
	case float64:
		return Float64TypeName
	case string:
		return StringTypeName
	case []byte, *[]byte:
		return BytesTypeName
	case proto.Message:
		return string(proto.MessageName(v))
	default:
		value := reflect.ValueOf(i)
		if value.Kind() == reflect.Ptr && value.IsNil() {
			value = reflect.New(value.Type().Elem())
		}

		typ := reflect.Indirect(value).Type()
		if typ.Kind() == reflect.Interface {
			typ = reflect.Indirect(reflect.ValueOf(i)).Elem().Type()
		}
		if typ.Kind() == reflect.Ptr {
			typ = typ.Elem()
		}

		if typ.Kind() == reflect.Slice || typ.Kind() == reflect.Array {
			typ = typ.Elem()
			if typ.Kind() == reflect.Ptr {
				typ = typ.Elem()
			}

			if name := getMojoTypeName(typ.Kind()); len(name) > 0 {
				return "Array<" + name + ">"
			}
			return "Array<" + GetMojoTypeName(reflect.New(typ).Interface()) + ">"
		} else if typ.Kind() == reflect.Map {
			kt := typ.Key()
			vt := typ.Elem()

			if vt.Kind() == reflect.Ptr {
				vt = vt.Elem()
			}

			kn := getMojoTypeName(kt.Kind())
			if len(kn) == 0 {
				kn = GetMojoTypeName(reflect.New(kt).Interface())
			}
			vn := getMojoTypeName(vt.Kind())
			if len(vn) == 0 {
				vn = GetMojoTypeName(reflect.New(vt).Interface())
			}
			return "Map<" + kn + "," + vn + ">"
		} else if typ.Kind() == reflect.Struct {
			panic("not support user defined go struct")
		}
		return ""
	}
}

func getMojoTypeName(kind reflect.Kind) string {
	switch kind {
	case reflect.Bool:
		return BoolTypeName
	case reflect.Int8:
		return Int8TypeName
	case reflect.Uint8:
		return UInt8TypeName
	case reflect.Int16:
		return Int16TypeName
	case reflect.Uint16:
		return UInt16TypeName
	case reflect.Int32:
		return Int32TypeName
	case reflect.Uint32:
		return UInt32TypeName
	case reflect.Int64:
		return Int64TypeName
	case reflect.Uint64:
		return UInt64TypeName
	case reflect.Int:
		return IntTypeName
	case reflect.Uint:
		return UIntTypeName
	case reflect.Float32:
		return Float32TypeName
	case reflect.Float64:
		return Float64TypeName
	case reflect.String:
		return StringTypeName
	default:
		return ""
	}
}
