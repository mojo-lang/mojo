package core

const Float32ValuesTypeName = "Float32Values"
const Float32ValuesTypeFullName = "mojo.core.Float32Values"

const Float64ValuesTypeName = "Float64Values"
const Float64ValuesTypeFullName = "mojo.core.Float64Values"

const FloatValuesTypeName = "FloatValues"
const FloatValuesTypeFullName = "mojo.core.FloatValues"

const DoubleValuesTypeName = "DoubleValues"
const DoubleValuesTypeFullName = "mojo.core.DoubleValues"

func (x *Float32Values) ToArray() interface{} {
	if x != nil {
		return x.Vals
	}
	return []float32{}
}

func (x *Float64Values) ToArray() interface{} {
	if x != nil {
		return x.Vals
	}
	return []float64{}
}
