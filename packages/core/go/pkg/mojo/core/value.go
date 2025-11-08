package core

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
	"unicode/utf8"

	jsoniter "github.com/json-iterator/go"
	"google.golang.org/protobuf/runtime/protoimpl"
)

const ValueTypeName = "Value"
const ValueTypeFullName = "mojo.core.Value"

// NewValue constructs a Value from a general-purpose Go interface.
//
//	╔════════════════════════╤════════════════════════════════════════════╗
//	║ Go type                │ Conversion                                 ║
//	╠════════════════════════╪════════════════════════════════════════════╣
//	║ nil                    │ stored as NullVal                          ║
//	║ bool                   │ stored as BoolVal                          ║
//	║ int, int32, int64      │ stored as NegativeVal/PositiveVal          ║
//	║ uint, uint32, uint64   │ stored as NegativeVal/PositiveVal          ║
//	║ float32, float64       │ stored as DoubleVal                        ║
//	║ string                 │ stored as StringVal; must be valid UTF-8   ║
//	║ []byte                 │ stored as BytesVal; base64-encoded to json ║
//	║ map[string]interface{} │ stored as ObjectVal                        ║
//	║ []interface{}          │ stored as ValuesVal                        ║
//	╚════════════════════════╧════════════════════════════════════════════╝
func NewValue(v interface{}) (*Value, error) {
	switch v := v.(type) {
	case nil:
		return NewNullValue(), nil
	case bool:
		return NewBoolValue(v), nil
	case int:
		return NewIntValue(v), nil
	case int8:
		return NewInt8Value(v), nil
	case int16:
		return NewInt16Value(v), nil
	case int32:
		return NewInt32Value(v), nil
	case int64:
		return NewInt64Value(v), nil
	case uint:
		return NewUIntValue(v), nil
	case uint8:
		return NewUInt8Value(v), nil
	case uint16:
		return NewUInt16Value(v), nil
	case uint32:
		return NewUInt32Value(v), nil
	case uint64:
		return NewUInt64Value(v), nil
	case float32:
		return NewFloat32Value(v), nil
	case float64:
		return NewFloat64Value(v), nil
	case string:
		if !utf8.ValidString(v) {
			return nil, protoimpl.X.NewError("invalid UTF-8 in string: %q", v)
		}
		return NewStringValue(v), nil
	case []byte:
		return NewBytesValue(v), nil
	case map[string]interface{}:
		obj, err := NewObjectFromMap(v)
		if err != nil {
			return nil, err
		}
		return NewObjectValue(obj), nil
	case map[string]*Value:
		obj := &Object{Vals: v}
		return NewObjectValue(obj), nil
	case []interface{}:
		array, err := NewValues(v...)
		if err != nil {
			return nil, err
		}
		return NewValuesValue(array), nil
	case *Value:
		return v, nil
	case *Object:
		return NewObjectValue(v), nil
	case *Values:
		return NewValuesValue(v), nil
	default:
		if sv, ok := v.(StringLike); ok {
			if sv == nil {
				return NewNullValue(), nil
			}
			return NewStringValue(sv.ToString()), nil
		}

		value := reflect.ValueOf(v)
		typ := reflect.Indirect(value).Type()
		if value.Kind() == reflect.Ptr {
			value = value.Elem()
			typ = reflect.Indirect(value).Type()
		}

		if typ.Kind() == reflect.Slice || typ.Kind() == reflect.Array {
			var values []*Value
			for i := 0; i < value.Len(); i++ {
				nv, err := NewValue(value.Index(i).Interface())
				if err != nil {
					return nil, err
				}
				values = append(values, nv)
			}
			if len(values) > 0 {
				return NewArrayValue(values...), nil
			}
		} else if typ.Kind() == reflect.Map {
			values := make(map[string]*Value)
			for iter := value.MapRange(); iter.Next(); {
				val, err := NewValue(iter.Value().Interface())
				if err != nil {
					return nil, err
				}
				values[iter.Key().String()] = val
			}
			if len(values) > 0 {
				return NewMapValue(values), nil
			}
		} else if typ.Kind() == reflect.Struct {
			if _, ok := registeredJsonEncoderTypes[typ.String()]; ok {
				if json, err := jsoniter.ConfigFastest.Marshal(v); err != nil {
					return nil, err
				} else {
					val := &Value{}
					if err = jsoniter.ConfigFastest.Unmarshal(json, val); err != nil {
						return nil, err
					}
					return val, nil
				}
			}

			obj := NewObject()
			if err := obj.From(v); err != nil {
				return nil, err
			}
			return NewObjectValue(obj), nil
		}
		return nil, fmt.Errorf("invalid type: %T", v)
	}
}

func NewNullValue() *Value {
	return &Value{Val: &Value_NullVal{NullVal: &Null{}}}
}

func NewBoolValue(val bool) *Value {
	return &Value{Val: &Value_BoolVal{BoolVal: val}}
}

func NewIntValue(val int) *Value {
	return NewInt64Value(int64(val))
}

func NewInt8Value(val int8) *Value {
	return NewInt64Value(int64(val))
}

func NewInt16Value(val int16) *Value {
	return NewInt64Value(int64(val))
}

func NewInt32Value(val int32) *Value {
	return NewInt64Value(int64(val))
}

func NewInt64Value(val int64) *Value {
	if val >= 0 {
		return &Value{Val: &Value_PositiveVal{PositiveVal: uint64(val)}}
	}
	value := uint64(0)
	if val == math.MinInt64 {
		value = uint64(math.MaxInt64) + 1
	} else {
		value = uint64(-val)
	}

	return &Value{Val: &Value_NegativeVal{NegativeVal: value}}
}

func NewUIntValue(val uint) *Value {
	return NewUInt64Value(uint64(val))
}

func NewUInt8Value(val uint8) *Value {
	return NewUInt64Value(uint64(val))
}

func NewUInt16Value(val uint16) *Value {
	return NewUInt64Value(uint64(val))
}

func NewUInt32Value(val uint32) *Value {
	return NewUInt64Value(uint64(val))
}

func NewUInt64Value(val uint64) *Value {
	return &Value{Val: &Value_PositiveVal{PositiveVal: val}}
}

func NewFloat32Value(val float32) *Value {
	return NewFloat64Value(float64(val))
}

func NewFloat64Value(val float64) *Value {
	return &Value{Val: &Value_DoubleVal{DoubleVal: val}}
}

func NewStringValue(val string) *Value {
	return &Value{Val: &Value_StringVal{StringVal: val}}
}

func NewBytesValue(val []byte) *Value {
	return &Value{Val: &Value_BytesVal{BytesVal: val}}
}

func NewObjectValue(obj *Object) *Value {
	return &Value{Val: &Value_ObjectVal{ObjectVal: obj}}
}

func NewValuesValue(array *Values) *Value {
	return &Value{Val: &Value_ValuesVal{ValuesVal: array}}
}

func NewArrayValue(values ...*Value) *Value {
	return &Value{Val: &Value_ValuesVal{ValuesVal: &Values{Vals: values}}}
}

func NewMapValue(values map[string]*Value) *Value {
	return &Value{Val: &Value_ObjectVal{ObjectVal: &Object{Vals: values}}}
}

func NewIntArrayValue(ints ...int) *Value {
	values := make([]*Value, 0, len(ints))
	for _, v := range ints {
		values = append(values, NewIntValue(v))
	}
	return &Value{Val: &Value_ValuesVal{ValuesVal: &Values{Vals: values}}}
}

func NewInt32ArrayValue(int32s ...int32) *Value {
	values := make([]*Value, 0, len(int32s))
	for _, v := range int32s {
		values = append(values, NewInt32Value(v))
	}
	return &Value{Val: &Value_ValuesVal{ValuesVal: &Values{Vals: values}}}
}

func NewInt64ArrayValue(int64s ...int64) *Value {
	values := make([]*Value, 0, len(int64s))
	for _, v := range int64s {
		values = append(values, NewInt64Value(v))
	}
	return &Value{Val: &Value_ValuesVal{ValuesVal: &Values{Vals: values}}}
}

func NewUInt32ArrayValue(uint32s ...uint32) *Value {
	values := make([]*Value, 0, len(uint32s))
	for _, v := range uint32s {
		values = append(values, NewUInt32Value(v))
	}
	return &Value{Val: &Value_ValuesVal{ValuesVal: &Values{Vals: values}}}
}

func NewUInt64ArrayValue(uint64s ...uint64) *Value {
	values := make([]*Value, 0, len(uint64s))
	for _, v := range uint64s {
		values = append(values, NewUInt64Value(v))
	}
	return &Value{Val: &Value_ValuesVal{ValuesVal: &Values{Vals: values}}}
}

func NewFloat32ArrayValue(float32s ...float32) *Value {
	values := make([]*Value, 0, len(float32s))
	for _, v := range float32s {
		values = append(values, NewFloat32Value(v))
	}
	return &Value{Val: &Value_ValuesVal{ValuesVal: &Values{Vals: values}}}
}

func NewFloat64ArrayValue(float64s ...float64) *Value {
	values := make([]*Value, 0, len(float64s))
	for _, v := range float64s {
		values = append(values, NewFloat64Value(v))
	}
	return &Value{Val: &Value_ValuesVal{ValuesVal: &Values{Vals: values}}}
}

func NewStringArrayValue(strings ...string) *Value {
	values := make([]*Value, 0, len(strings))
	for _, v := range strings {
		values = append(values, NewStringValue(v))
	}
	return &Value{Val: &Value_ValuesVal{ValuesVal: &Values{Vals: values}}}
}

func NewObjectArrayValue(objects ...*Object) *Value {
	values := make([]*Value, 0, len(objects))
	for _, v := range objects {
		values = append(values, NewObjectValue(v))
	}
	return &Value{Val: &Value_ValuesVal{ValuesVal: &Values{Vals: values}}}
}

// ToInterface converts x to a general-purpose Go interface.
//
// Calling Value.MarshalJSON and "encoding/json".Marshal on this output produce
// semantically equivalent JSON (assuming no errors occur).
//
// Floating-point values (i.e., "NaN", "Infinity", and "-Infinity") are
// converted as strings to remain compatible with MarshalJSON.
func (x *Value) ToInterface() interface{} {
	switch v := x.GetVal().(type) {
	case *Value_BoolVal:
		if v != nil {
			return v.BoolVal
		}
	case *Value_NegativeVal:
		if v != nil {
			return -v.NegativeVal
		}
	case *Value_PositiveVal:
		if v != nil {
			return v.PositiveVal
		}
	case *Value_DoubleVal:
		if v != nil {
			switch {
			case math.IsNaN(v.DoubleVal):
				return "NaN"
			case math.IsInf(v.DoubleVal, +1):
				return "Infinity"
			case math.IsInf(v.DoubleVal, -1):
				return "-Infinity"
			default:
				return v.DoubleVal
			}
		}
	case *Value_StringVal:
		if v != nil {
			return v.StringVal
		}
	case *Value_BytesVal:
		if v != nil {
			return v.BytesVal
		}
	case *Value_ObjectVal:
		if v != nil {
			return v.ObjectVal.ToMapInterface()
		}
	case *Value_ValuesVal:
		if v != nil {
			return v.ValuesVal.ToSlice()
		}
	}
	return nil
}

func (x *Value) GetKind() ValueKind {
	if x != nil {
		switch x.Val.(type) {
		case *Value_NullVal:
			return ValueKind_VALUE_KIND_NULL
		case *Value_BoolVal:
			return ValueKind_VALUE_KIND_BOOLEAN
		case *Value_NegativeVal, *Value_PositiveVal:
			return ValueKind_VALUE_KIND_INTEGER
		case *Value_DoubleVal:
			return ValueKind_VALUE_KIND_NUMBER
		case *Value_StringVal:
			return ValueKind_VALUE_KIND_STRING
		case *Value_BytesVal:
			return ValueKind_VALUE_KIND_BYTES
		case *Value_ValuesVal:
			return ValueKind_VALUE_KIND_ARRAY
		case *Value_ObjectVal:
			return ValueKind_VALUE_KIND_OBJECT
		default:
			return ValueKind_VALUE_KIND_UNSPECIFIED
		}
	}
	return ValueKind_VALUE_KIND_UNSPECIFIED
}

func (x *Value) GetBool() bool {
	if v, ok := x.GetVal().(*Value_BoolVal); ok {
		return v.BoolVal
	} else if s, ok := x.GetVal().(*Value_StringVal); ok && len(s.StringVal) > 0 {
		return strings.ToLower(s.StringVal) == "true"
	}
	return false
}

func (x *Value) GetInt() int {
	return int(x.GetInt64())
}

func (x *Value) GetIntVal() int {
	return int(x.GetInt64Val())
}

func (x *Value) GetInt32() int32 {
	return int32(x.GetInt64())
}

func (x *Value) GetInt32Val() int32 {
	return int32(x.GetInt64Val())
}

func (x *Value) GetInt64Val() int64 {
	if negative, ok := x.GetVal().(*Value_NegativeVal); ok {
		if negative.NegativeVal == uint64(math.MaxInt64)+1 {
			return math.MinInt64
		}
		return -int64(negative.NegativeVal)
	} else if positive, ok := x.GetVal().(*Value_PositiveVal); ok {
		return int64(positive.PositiveVal)
	}
	return 0
}

func (x *Value) GetInt64() int64 {
	if v := x.GetInt64Val(); v != 0 {
		return v
	} else if str, ok := x.GetVal().(*Value_StringVal); ok && len(str.StringVal) > 0 {
		v, _ = strconv.ParseInt(str.StringVal, 10, 64)
		return v
	}
	return 0
}

func (x *Value) GetUint() uint {
	return uint(x.GetPositiveVal())
}

func (x *Value) GetUint32() uint32 {
	return uint32(x.GetPositiveVal())
}

func (x *Value) GetUint64() uint64 {
	return x.GetPositiveVal()
}

func (x *Value) GetFloat32() float32 {
	return float32(x.GetFloat64())
}

func (x *Value) GetFloat64() float64 {
	if v, ok := x.GetVal().(*Value_DoubleVal); ok {
		return v.DoubleVal
	} else {
		if  str, ok := x.GetVal().(*Value_StringVal); ok && len(str.StringVal) > 0 {
			f, _ := strconv.ParseFloat(str.StringVal, 64)
			return f
		} else {
			return float64(x.GetInt64Val())
		}
	}
}

func (x *Value) GetDouble() float64 {
	return x.GetFloat64()
}

func (x *Value) GetString() string {
	if str, ok := x.GetVal().(*Value_StringVal); ok {
		return str.StringVal
	} else if negative, ok := x.GetVal().(*Value_NegativeVal); ok {
		return "-" + strconv.FormatUint(negative.NegativeVal, 10)
	} else if positive, ok := x.GetVal().(*Value_PositiveVal); ok {
		return strconv.FormatUint(positive.PositiveVal, 10)
	}
	return ""
}

func (x *Value) GetBytes() []byte {
	return x.GetBytesVal()
}

func (x *Value) GetObject() *Object {
	return x.GetObjectVal()
}

func (x *Value) GetValues() []*Value {
	if values := x.GetValuesVal(); values != nil {
		return values.Vals
	}
	return nil
}

func (x *Value) GetObjectArray() []*Object {
	if values := x.GetValues(); len(values) > 0 {
		objects := make([]*Object, 0, len(values))
		for _, value := range values {
			objects = append(objects, value.GetObject())
		}
		return objects
	}
	return nil
}

func (x *Value) GetStringArray() []string {
	values := x.GetValues()
	array := make([]string, 0, len(values))
	for _, value := range values {
		if value.GetKind() == ValueKind_VALUE_KIND_STRING {
			array = append(array, value.GetString())
		} else {
			return nil
		}
	}
	return array
}

func (x *Value) GetStringValues() *StringValues {
	vals := x.GetStringArray()
	if len(vals) > 0 {
		return &StringValues{Vals: vals}
	}
	return nil
}

func (x *Value) GetBoolArray() []bool {
	values := x.GetValues()
	array := make([]bool, 0, len(values))
	for _, value := range values {
		array = append(array, value.GetBool())
	}
	return array
}

func (x *Value) GetIntArray() []int {
	values := x.GetValues()
	array := make([]int, 0, len(values))
	for _, value := range values {
		array = append(array, value.GetInt())
	}
	return array
}

func (x *Value) GetInt32Array() []int32 {
	values := x.GetValues()
	array := make([]int32, 0, len(values))
	for _, value := range values {
		array = append(array, value.GetInt32())
	}
	return array
}

func (x *Value) GetInt64Array() []int64 {
	values := x.GetValues()
	array := make([]int64, 0, len(values))
	for _, value := range values {
		array = append(array, value.GetInt64())
	}
	return array
}

func (x *Value) GetUintArray() []uint {
	values := x.GetValues()
	array := make([]uint, 0, len(values))
	for _, value := range values {
		array = append(array, value.GetUint())
	}
	return array
}

func (x *Value) GetUint32Array() []uint32 {
	values := x.GetValues()
	array := make([]uint32, 0, len(values))
	for _, value := range values {
		array = append(array, value.GetUint32())
	}
	return array
}

func (x *Value) GetUint64Array() []uint64 {
	values := x.GetValues()
	array := make([]uint64, 0, len(values))
	for _, value := range values {
		array = append(array, value.GetUint64())
	}
	return array
}

func (x *Value) GetFloat32Array() []float32 {
	values := x.GetValues()
	array := make([]float32, 0, len(values))
	for _, value := range values {
		array = append(array, value.GetFloat32())
	}
	return array
}

func (x *Value) GetFloat64Array() []float64 {
	values := x.GetValues()
	array := make([]float64, 0, len(values))
	for _, value := range values {
		array = append(array, value.GetFloat64())
	}
	return array
}

// type UrlStringOptions struct {
// 	Explode   bool
// 	Separator string // default: <||>
// }
//
// func (x *Value) ToUrlString(options *UrlStringOptions) string {
// 	if x != nil {
// 		builder := &strings.Builder{}
// 		if err := x.toUrlString(builder, options); err != nil {
// 			return ""
// 		}
// 		return builder.String()
// 	}
// 	return ""
// }
//
// func (x *Value) toUrlString(buf *strings.Builder, options *UrlStringOptions) error {
// 	if x != nil {
// 		switch x.GetKind() {
// 		case ValueKind_VALUE_KIND_NULL:
// 			buf.WriteString("")
// 		case ValueKind_VALUE_KIND_STRING:
// 			buf.WriteString(x.GetString())
// 		case ValueKind_VALUE_KIND_BOOLEAN:
// 			if val := x.GetBool(); val {
// 				buf.WriteString("true")
// 			} else {
// 				buf.WriteString("false")
// 			}
// 		case ValueKind_VALUE_KIND_INTEGER:
// 			val := x.GetInt64()
// 			buf.WriteString(strconv.FormatInt(val, 10))
// 		case ValueKind_VALUE_KIND_NUMBER:
// 			val := x.GetFloat64()
// 			buf.WriteString(strconv.FormatFloat(val, 'g', 12, 64))
// 		case ValueKind_VALUE_KIND_ARRAY:
// 			// values := x.GetValues()
// 			// for _, val := range values {
// 			//
// 			// }
// 		case ValueKind_VALUE_KIND_OBJECT:
// 		case ValueKind_VALUE_KIND_BYTES:
// 			return errors.New("not supported bytes type in converting from the Value to safe string")
// 		}
// 	}
// 	return nil
// }
