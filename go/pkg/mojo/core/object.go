package core

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"unicode/utf8"

	jsoniter "github.com/json-iterator/go"
	"github.com/modern-go/reflect2"
	"google.golang.org/protobuf/runtime/protoimpl"

	"github.com/mojo-lang/mojo/go/pkg/mojo/core/strcase"
)

const ObjectTypeName = "Object"
const ObjectTypeFullName = "mojo.core.Object"

func NewObject() *Object {
	return &Object{
		Vals: make(map[string]*Value),
	}
}

// NewObjectFromMap constructs a Struct from a general-purpose Go map.
// The map keys must be valid UTF-8.
// The map values are converted using NewValue.
func NewObjectFromMap(v map[string]interface{}) (*Object, error) {
	x := &Object{Vals: make(map[string]*Value, len(v))}
	for k, v := range v {
		if !utf8.ValidString(k) {
			return nil, protoimpl.X.NewError("invalid UTF-8 in string: %q", k)
		}
		var err error
		x.Vals[k], err = NewValue(v)
		if err != nil {
			return nil, err
		}
	}
	return x, nil
}

func NewObjectFromKeyValues(kvs ...interface{}) (*Object, error) {
	if len(kvs)%2 != 0 {
		return nil, fmt.Errorf("invalid key value pairs")
	}

	m := make(map[string]interface{})
	for i := 0; i < len(kvs); i += 2 {
		key, ok := kvs[i].(string)
		if !ok {
			return nil, fmt.Errorf("invalid key type, must be string. key: %v", kvs[i])
		}
		m[key] = kvs[i+1]
	}
	return NewObjectFromMap(m)
}

func NewObjectFrom(value interface{}) (*Object, error) {
	obj := &Object{}
	return obj, obj.From(value)
}

func MergeObjects(objs ...*Object) *Object {
	obj := &Object{
		Vals: make(map[string]*Value),
	}
	for _, o := range objs {
		obj.Merge(o)
	}
	return obj
}

func (x *Object) ToMap() interface{} {
	if x != nil && x.Vals != nil {
		return x.Vals
	}
	return nil
}

func (x *Object) ToMapInterface() map[string]interface{} {
	if x != nil && x.Vals != nil && len(x.Vals) > 0 {
		vs := make(map[string]interface{})
		for k, v := range x.Vals {
			vs[k] = v.ToInterface()
		}
		return vs
	}
	return nil
}

func (x *Object) To(value interface{}) error {
	if x != nil {
		if json, err := jsoniter.ConfigFastest.Marshal(x); err != nil {
			return err
		} else {
			return jsoniter.ConfigFastest.Unmarshal(json, value)
		}
	}
	return nil
}

func (x *Object) From(value interface{}) error {
	if x != nil {
		if value == nil {
			return nil
		}
		if reflect.ValueOf(value).IsZero() {
			return nil
		}

		if x.Vals == nil {
			x.Vals = make(map[string]*Value)
		}

		switch value := value.(type) {
		case nil, bool, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, string, []byte, []interface{}, *Values:
			return errors.New(fmt.Sprintf("invalid input object format: %T", value))
		case map[string]interface{}:
			if obj, err := NewObjectFromMap(value); err != nil {
				return err
			} else {
				x.Vals = obj.Vals
			}
		case *Value:
			if obj := value.GetObject(); obj != nil {
				x.Vals = obj.Vals
			} else {
				return errors.New(fmt.Sprintf("invalid input object format: %T", value))
			}
		case *Object:
			x.Vals = value.Vals
		case map[string]*Value:
			x.Vals = value
		default:
			v := reflect.ValueOf(value)
			typ := reflect.Indirect(v).Type()
			if v.Kind() == reflect.Ptr {
				v = v.Elem()
				typ = reflect.Indirect(v).Type()
			}
			if typ.Kind() != reflect.Struct {
				return errors.New(fmt.Sprintf("invalid input object format: %T", value))
			}

			if _, ok := registeredJsonEncoderTypes[typ.String()]; ok {
				if json, err := jsoniter.ConfigFastest.Marshal(value); err != nil {
					return err
				} else {
					if err = jsoniter.ConfigFastest.Unmarshal(json, &x.Vals); err != nil {
						return err
					}
				}
			} else {
				for i := 0; i < typ.NumField(); i++ {
					key := typ.Field(i).Name
					if (key[0] >= 'a' && key[0] <= 'z') || strings.HasPrefix(key, "XXX") {
						continue
					}
					if v.Field(i).IsZero() {
						continue
					}

					val := &Value{}
					if encoder, ok := registeredJsonEncoderTypeFields[typ.String()+"."+key]; ok {
						rt := reflect2.TypeOf(value)
						if rt.Kind() == reflect.Ptr {
							rt = rt.(*reflect2.UnsafePtrType).Elem()
						}
						if obj, ok := rt.(*reflect2.UnsafeStructType); ok {
							field := obj.FieldByName(key)
							f := field.UnsafeGet(reflect2.PtrOf(value))

							buf := &strings.Builder{}
							stream := jsoniter.NewStream(jsoniter.ConfigFastest, buf, 1024)
							encoder.Encode(f, stream)
							if err := jsoniter.ConfigFastest.Unmarshal(stream.Buffer(), &val); err != nil {
								return err
							}
						}
					} else {
						var err error
						if val, err = NewValue(v.Field(i).Interface()); err != nil {
							return err
						}
					}
					// FIXME should using the json tag name in the struct
					x.Vals[strcase.ToLowerCamel(key)] = val
				}
			}
		}
	}
	return nil
}

func (x *Object) IsEmpty() bool {
	if x != nil {
		return len(x.Vals) == 0
	}
	return true
}

func (x *Object) GetValue(key string) *Value {
	if x != nil && x.Vals != nil {
		return x.Vals[key]
	}
	return nil
}

func (x *Object) GetString(key string) string {
	return x.GetValue(key).GetString()
}

func (x *Object) GetBool(key string) bool {
	return x.GetValue(key).GetBool()
}

func (x *Object) GetInt32(key string) int32 {
	return x.GetValue(key).GetInt32()
}

func (x *Object) GetInt64(key string) int64 {
	return x.GetValue(key).GetInt64()
}

func (x *Object) GetUint64(key string) uint64 {
	return x.GetValue(key).GetUint64()
}

func (x *Object) GetFloat32(key string) float32 {
	return x.GetValue(key).GetFloat32()
}

func (x *Object) GetFloat64(key string) float64 {
	return x.GetValue(key).GetFloat64()
}

func (x *Object) GetBoolArray(key string) []bool {
	return x.GetValue(key).GetBoolArray()
}

func (x *Object) GetInt32Array(key string) []int32 {
	return x.GetValue(key).GetInt32Array()
}

func (x *Object) GetInt64Array(key string) []int64 {
	return x.GetValue(key).GetInt64Array()
}

func (x *Object) GetUint32Array(key string) []uint32 {
	return x.GetValue(key).GetUint32Array()
}

func (x *Object) GetUint64Array(key string) []uint64 {
	return x.GetValue(key).GetUint64Array()
}

func (x *Object) GetFloat32Array(key string) []float32 {
	return x.GetValue(key).GetFloat32Array()
}

func (x *Object) GetFloat64Array(key string) []float64 {
	return x.GetValue(key).GetFloat64Array()
}

func (x *Object) GetStringArray(key string) []string {
	return x.GetValue(key).GetStringArray()
}

func (x *Object) GetObjectArray(key string) []*Object {
	return x.GetValue(key).GetObjectArray()
}

func (x *Object) init() {
	if x != nil && x.Vals == nil {
		x.Vals = make(map[string]*Value)
	}
}

func (x *Object) SetValue(key string, value *Value) *Object {
	if x != nil {
		x.init()
		x.Vals[key] = value
	}

	return x
}

func (x *Object) SetBool(key string, value bool) *Object {
	if x != nil {
		x.init()
		x.Vals[key] = NewBoolValue(value)
	}
	return x
}

func (x *Object) SetInt(key string, value int) *Object {
	if x != nil {
		x.init()
		x.Vals[key] = NewIntValue(value)
	}

	return x
}

func (x *Object) SetInt32(key string, value int32) *Object {
	if x != nil {
		x.init()
		x.Vals[key] = NewInt32Value(value)
	}
	return x
}

func (x *Object) SetInt64(key string, value int64) *Object {
	if x != nil {
		x.init()
		x.Vals[key] = NewInt64Value(value)
	}
	return x
}

func (x *Object) SetUint32(key string, value uint32) *Object {
	if x != nil {
		x.init()
		x.Vals[key] = NewUInt32Value(value)
	}
	return x
}

func (x *Object) SetUint64(key string, value uint64) *Object {
	if x != nil {
		x.init()
		x.Vals[key] = NewUInt64Value(value)
	}
	return x
}

func (x *Object) SetFloat32(key string, value float32) *Object {
	if x != nil {
		x.init()
		x.Vals[key] = NewFloat32Value(value)
	}
	return x
}

func (x *Object) SetFloat64(key string, value float64) *Object {
	if x != nil {
		x.init()
		x.Vals[key] = NewFloat64Value(value)
	}
	return x
}

func (x *Object) SetString(key string, value string) *Object {
	if x != nil {
		x.init()
		x.Vals[key] = NewStringValue(value)
	}
	return x
}

func (x *Object) SetObject(key string, value *Object) *Object {
	if x != nil {
		x.init()
		x.Vals[key] = NewObjectValue(value)
	}
	return x
}

func (x *Object) SetInt32Array(key string, vals ...int32) *Object {
	if x != nil {
		x.init()
		x.Vals[key] = NewInt32ArrayValue(vals...)
	}
	return x
}

func (x *Object) SetInt64Array(key string, vals ...int64) *Object {
	if x != nil {
		x.init()
		x.Vals[key] = NewInt64ArrayValue(vals...)
	}
	return x
}

func (x *Object) SetUint32Array(key string, vals ...uint32) *Object {
	if x != nil {
		x.init()
		x.Vals[key] = NewUInt32ArrayValue(vals...)
	}
	return x
}

func (x *Object) SetUint64Array(key string, vals ...uint64) *Object {
	if x != nil {
		x.init()
		x.Vals[key] = NewUInt64ArrayValue(vals...)
	}
	return x
}

func (x *Object) SetFloat32Array(key string, vals ...float32) *Object {
	if x != nil {
		x.init()
		x.Vals[key] = NewFloat32ArrayValue(vals...)
	}
	return x
}

func (x *Object) SetFloat64Array(key string, vals ...float64) *Object {
	if x != nil {
		x.init()
		x.Vals[key] = NewFloat64ArrayValue(vals...)
	}
	return x
}

func (x *Object) SetStringArray(key string, vals ...string) *Object {
	if x != nil {
		x.init()
		x.Vals[key] = NewStringArrayValue(vals...)
	}
	return x
}

func (x *Object) SetObjectArray(key string, vals ...*Object) *Object {
	if x != nil {
		x.init()
		x.Vals[key] = NewObjectArrayValue(vals...)
	}
	return x
}

func (x *Object) Delete(key string) *Object {
	if x != nil && len(x.Vals) > 0 {
		delete(x.Vals, key)
	}
	return x
}

func (x *Object) Merge(obj *Object) *Object {
	if x != nil {
		for k, v := range obj.GetVals() {
			x.Vals[k] = v
		}
	}
	return x
}

func (x *Object) ToLowerCamelKeys() *Object {
	if x != nil {
		obj := NewObject()
		for k, v := range x.GetVals() {
			obj.SetValue(strcase.ToLowerCamel(k), v)
		}
		return obj
	}
	return x
}

func (x *Object) ToSnakeKeys() *Object {
	if x != nil {
		obj := NewObject()
		for k, v := range x.GetVals() {
			obj.SetValue(strcase.ToSnake(k), v)
		}
		return obj
	}
	return x
}
