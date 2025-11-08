package core

const ValuesTypeName = "Values"
const ValuesTypeFullName = "mojo.core.Values"

func NewValues(values ...interface{}) (*Values, error) {
	vs := make([]*Value, 0, len(values))
	for _, v := range values {
		value, err := NewValue(v)
		if err != nil {
			return nil, err
		}
		vs = append(vs, value)
	}
	return &Values{Vals: vs}, nil
}

func (x *Values) ToSlice() []interface{} {
	if x != nil && len(x.Vals) > 0 {
		vs := make([]interface{}, len(x.Vals))
		for i, v := range x.Vals {
			vs[i] = v.ToInterface()
		}
		return vs
	}
	return []interface{}{}
}

func (x *Values) Len() int {
	if x != nil {
		return len(x.Vals)
	}
	return 0
}

func (x *Values) GetValue(index int) *Value {
	if x != nil && index >= 0 && index < len(x.Vals) {
		return x.Vals[index]
	}
	return nil
}

func (x *Values) GetString(index int) string {
	return x.GetValue(index).GetString()
}

func (x *Values) GetBool(index int) bool {
	return x.GetValue(index).GetBool()
}

func (x *Values) GetInt32(index int) int32 {
	return x.GetValue(index).GetInt32()
}

func (x *Values) GetInt64(index int) int64 {
	return x.GetValue(index).GetInt64()
}

func (x *Values) GetUint64(index int) uint64 {
	return x.GetValue(index).GetUint64()
}

func (x *Values) GetInt32Array(index int) []int32 {
	return x.GetValue(index).GetInt32Array()
}

func (x *Values) GetInt64Array(index int) []int64 {
	return x.GetValue(index).GetInt64Array()
}

func (x *Values) GetUint32Array(index int) []uint32 {
	return x.GetValue(index).GetUint32Array()
}

func (x *Values) GetUint64Array(index int) []uint64 {
	return x.GetValue(index).GetUint64Array()
}

func (x *Values) GetStringArray(index int) []string {
	return x.GetValue(index).GetStringArray()
}

func (x *Values) GetObjectArray(index int) []*Object {
	return x.GetValue(index).GetObjectArray()
}

func (x *Values) AppendValue(value *Value) *Values {
	if x != nil {
		x.Vals = append(x.Vals, value)
	}
	return x
}

func (x *Values) AppendValues(values ...*Value) *Values {
	if x != nil {
		x.Vals = append(x.Vals, values...)
	}
	return x
}

func (x *Values) AppendBool(value bool) *Values {
	if x != nil {
		return x.AppendValue(NewBoolValue(value))
	}
	return x
}

func (x *Values) AppendInt(value int) *Values {
	if x != nil {
		return x.AppendValue(NewIntValue(value))
	}
	return x
}

func (x *Values) AppendInt32(value int32) *Values {
	if x != nil {
		return x.AppendValue(NewInt32Value(value))
	}
	return x
}

func (x *Values) AppendInt64(value int64) *Values {
	if x != nil {
		return x.AppendValue(NewInt64Value(value))
	}
	return x
}

func (x *Values) AppendUint32(value uint32) *Values {
	if x != nil {
		return x.AppendValue(NewUInt32Value(value))
	}
	return x
}

func (x *Values) AppendUint64(value uint64) *Values {
	if x != nil {
		return x.AppendValue(NewUInt64Value(value))
	}
	return x
}

func (x *Values) AppendFloat32(value float32) *Values {
	if x != nil {
		return x.AppendValue(NewFloat32Value(value))
	}
	return x
}

func (x *Values) AppendFloat64(value float64) *Values {
	if x != nil {
		return x.AppendValue(NewFloat64Value(value))
	}
	return x
}

func (x *Values) AppendString(value string) *Values {
	if x != nil {
		return x.AppendValue(NewStringValue(value))
	}
	return x
}

func (x *Values) AppendObject(value *Object) *Values {
	if x != nil {
		return x.AppendValue(NewObjectValue(value))
	}
	return x
}
