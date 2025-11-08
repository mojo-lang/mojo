package core

const (
	// Int8ValuesTypeName     = "Int8Values"
	// Int8ValuesTypeFullName = "mojo.core.Int8Values"
	//
	// Int16ValuesTypeName     = "Int16Values"
	// Int16ValuesTypeFullName = "mojo.core.Int16Values"

	Int32ValuesTypeName     = "Int32Values"
	Int32ValuesTypeFullName = "mojo.core.Int32Values"

	Int64ValuesTypeName     = "Int64Values"
	Int64ValuesTypeFullName = "mojo.core.Int64Values"

	IntValuesTypeName     = "IntValues"
	IntValuesTypeFullName = "mojo.core.IntValues"

	// UInt8ValuesTypeName     = "UInt8Values"
	// UInt8ValuesTypeFullName = "mojo.core.UInt8Values"
	//
	// UInt16ValuesTypeName     = "UInt16Values"
	// UInt16ValuesTypeFullName = "mojo.core.UInt16Values"

	UInt32ValuesTypeName     = "UInt32Values"
	UInt32ValuesTypeFullName = "mojo.core.UInt32Values"

	UInt64ValuesTypeName     = "UInt64Values"
	UInt64ValuesTypeFullName = "mojo.core.UInt64Values"

	UIntValuesTypeName     = "UIntValues"
	UIntValuesTypeFullName = "mojo.core.UIntValues"
)

func (x *Int32Values) ToArray() interface{} {
	if x != nil {
		return x.Vals
	}
	return []int32{}
}

func (x *Int64Values) ToArray() interface{} {
	if x != nil {
		return x.Vals
	}
	return []int64{}
}

func (x *UInt32Values) ToArray() interface{} {
	if x != nil {
		return x.Vals
	}
	return []uint32{}
}

func (x *UInt64Values) ToArray() interface{} {
	if x != nil {
		return x.Vals
	}
	return []uint64{}
}
