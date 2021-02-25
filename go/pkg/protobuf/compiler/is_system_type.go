package compiler

func IsSystemScalarType(name string) bool {
	switch name {
	case "Bool", "Int8", "UInt8", "Int16", "UInt16", "Int32", "UInt32", "Int64", "UInt64", "Int", "UInt",
		"Float32", "Float64", "Float", "Double", "String", "Bytes",
		"mojo.core.Bool", "mojo.core.Int8", "mojo.core.UInt8", "mojo.core.Int16", "mojo.core.UInt16",
		"mojo.core.Int32", "mojo.core.UInt32", "mojo.core.Int64", "mojo.core.UInt64", "mojo.core.Int", "mojo.core.UInt",
		"mojo.core.Float32", "mojo.core.Float64", "mojo.core.Float", "mojo.core.Double",
		"mojo.core.String", "mojo.core.Bytes":
		return true
	default:
		return false
	}
}
