package mpm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinaryFile_Decode(t *testing.T) {
	file := &BinaryFile{Parts: map[string][]byte{
		"foo": []byte("foo-bar"),
		"bar": []byte("bar-foo"),
	}}

	bytes := file.Encode()

	f, err := DecodeBinaryFile(bytes)
	assert.NoError(t, err)
	assert.Equal(t, len(file.Parts), len(f.Parts))
	assert.Equal(t, file.Parts["foo"], f.Parts["foo"])
	assert.Equal(t, file.Parts["bar"], f.Parts["bar"])
}

func TestGenerateMojoPackages(t *testing.T) {
	// err := GenerateMojoPackages("../../../")
	// assert.NoError(t, err)
}
