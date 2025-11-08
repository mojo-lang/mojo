package core

func GetProperBoxedScalarName(scalar string) string {
	toBoxed := func(name string) string { return name + "Value" }
	switch scalar {
	case BoolTypeName, BoolTypeFullName,
		Int8TypeName, Int8TypeFullName, Int16TypeName, Int16TypeFullName, Int32TypeName, Int32TypeFullName,
		UInt8TypeName, UInt8TypeFullName, UInt16TypeName, UInt16TypeFullName, UInt32TypeName, UInt32TypeFullName,
		IntTypeName, IntTypeFullName, Int64TypeName, Int64TypeFullName,
		UIntTypeName, UIntTypeFullName, UInt64TypeName, UInt64TypeFullName,
		Float32TypeName, Float32TypeFullName, FloatTypeName, FloatTypeFullName,
		Float64TypeName, Float64TypeFullName, DoubleTypeName, DoubleTypeFullName,
		StringTypeName, StringTypeFullName,
		BytesTypeName, BytesTypeFullName:
		return toBoxed(scalar)
	}
	return ""
}
