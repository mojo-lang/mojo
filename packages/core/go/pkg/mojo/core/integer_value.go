package core

const (
	// Int8ValueTypeName     = "Int8Value"
	// Int8ValueTypeFullName = "mojo.core.Int8Value"
	//
	// Int16ValueTypeName     = "Int16Value"
	// Int16ValueTypeFullName = "mojo.core.Int16Value"

	Int32ValueTypeName     = "Int32Value"
	Int32ValueTypeFullName = "mojo.core.Int32Value"

	Int64ValueTypeName     = "Int64Value"
	Int64ValueTypeFullName = "mojo.core.Int64Value"

	IntValueTypeName     = "IntValue"
	IntValueTypeFullName = "mojo.core.IntValue"

	// UInt8ValueTypeName     = "UInt8Value"
	// UInt8ValueTypeFullName = "mojo.core.UInt8Value"
	//
	// UInt16ValueTypeName     = "UInt16Value"
	// UInt16ValueTypeFullName = "mojo.core.UInt16Value"

	UInt32ValueTypeName     = "UInt32Value"
	UInt32ValueTypeFullName = "mojo.core.UInt32Value"

	UInt64ValueTypeName     = "UInt64Value"
	UInt64ValueTypeFullName = "mojo.core.UInt64Value"

	UIntValueTypeName     = "UIntValue"
	UIntValueTypeFullName = "mojo.core.UIntValue"
)

type IntValue struct {
	Int64Value
}

type UIntValue struct {
	UInt64Value
}
