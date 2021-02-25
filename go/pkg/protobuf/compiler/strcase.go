package compiler

import (
	"github.com/iancoleman/strcase"
)

var specialIndex = map[string]string{
	"Int8":    "int8",
	"Int16":   "int16",
	"Int32":   "int32",
	"Int64":   "int64",
	"UInt8":   "uint8",
	"UInt16":  "uint16",
	"UInt32":  "uint32",
	"UInt64":  "uint64",
	"UInt":    "uint",
	"Float32": "float32",
	"Float64": "float64",
}

func ToSnake(name string) string {
	str := specialIndex[name]
	if len(str) > 0 {
		return str
	}
	return strcase.ToSnake(name)
}
