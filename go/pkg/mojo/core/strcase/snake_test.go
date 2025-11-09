package strcase

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var cases = [][]string{
	{"Int8", "int8"},
	{"Int16", "int16"},
	{"Int32", "int32"},
	{"Int64", "int64"},
	{"UInt8", "uint8"},
	{"UInt16", "uint16"},
	{"UInt32", "uint32"},
	{"UInt64", "uint64"},
	{"UInt", "uint"},
	{"Float32", "float32"},
	{"Float64", "float64"},
	{"IPV4", "ipv4"},
	{"B2B", "b2b"},
}

func TestToSnake(t *testing.T) {
	for _, c := range cases {
		assert.Equal(t, c[1], ToSnake(c[0]))
		assert.Equal(t, c[1], ToKebab(c[0]))
	}
}
