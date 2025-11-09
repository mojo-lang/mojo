package core

const Float32ValueTypeName = "Float32Value"
const Float32ValueTypeFullName = "mojo.core.Float32Value"

const Float64ValueTypeName = "Float64Value"
const Float64ValueTypeFullName = "mojo.core.Float64Value"

const FloatValueTypeName = "FloatValue"
const FloatValueTypeFullName = "mojo.core.FloatValue"

const DoubleValueTypeName = "DoubleValue"
const DoubleValueTypeFullName = "mojo.core.DoubleValue"

type FloatValue struct {
	Float32Value
}

type DoubleValue struct {
	Float64Value
}
