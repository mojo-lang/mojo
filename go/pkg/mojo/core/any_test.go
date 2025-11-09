package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
)

var anyScalarValues = []interface{}{
	int64(100),
	"this is a string",
}

func TestAny_Marshal_Scalar(t *testing.T) {
	for _, v := range anyScalarValues {
		any := NewAny(v)
		b, err := proto.Marshal(any)
		assert.NoError(t, err)

		a := &Any{}
		err = proto.Unmarshal(b, a)
		assert.NoError(t, err)

		assert.Equal(t, v, a.Get())
	}
}

func TestAny_Marshal_Struct(t *testing.T) {
	any := NewAny(&Checksum{
		Algorithm: Checksum_ALGORITHM_MD5,
		Val:       "md5",
	})

	b, err := proto.Marshal(any)
	assert.NoError(t, err)

	a := &Any{}
	err = proto.Unmarshal(b, a)
	assert.NoError(t, err)

	v, ok := a.Get().(*Checksum)
	assert.True(t, ok)
	assert.Equal(t, Checksum_ALGORITHM_MD5, v.Algorithm)
	assert.Equal(t, "md5", v.Val)
}

func TestGetMojoTypeName(t *testing.T) {
	assert.Equal(t, "Int", GetMojoTypeName(1))
	assert.Equal(t, "Int32", GetMojoTypeName(int32(1)))
	assert.Equal(t, "Int64", GetMojoTypeName(int64(1)))
	assert.Equal(t, "UInt", GetMojoTypeName(uint(1)))
	assert.Equal(t, "UInt32", GetMojoTypeName(uint32(1)))
	assert.Equal(t, "UInt64", GetMojoTypeName(uint64(1)))
	assert.Equal(t, "Float32", GetMojoTypeName(float32(1.0)))
	assert.Equal(t, "Float64", GetMojoTypeName(1.0))

	assert.Equal(t, "String", GetMojoTypeName(""))
	assert.Equal(t, "Bytes", GetMojoTypeName([]byte{}))

	assert.Equal(t, "Array<Int32>", GetMojoTypeName([]int32{}))
	assert.Equal(t, "Array<Bytes>", GetMojoTypeName([][]byte{}))
	assert.Equal(t, "Array<mojo.core.Checksum>", GetMojoTypeName([]*Checksum{}))
	assert.Equal(t, "Array<Array<mojo.core.Checksum>>", GetMojoTypeName([][]*Checksum{}))
	assert.Equal(t, "Map<String,Int32>", GetMojoTypeName(map[string]int32{}))
	assert.Equal(t, "Map<String,mojo.core.Checksum>", GetMojoTypeName(map[string]*Checksum{}))
	assert.Equal(t, "Map<String,Array<mojo.core.Checksum>>", GetMojoTypeName(map[string][]*Checksum{}))
	assert.Equal(t, "mojo.core.Checksum", GetMojoTypeName(&Checksum{}))

	type UserDefined struct {
		Field int
	}

	assert.Panics(t, func() {
		GetMojoTypeName(&UserDefined{})
	})
}
