package http

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mojo-lang/mojo/go/pkg/mojo/http/testdata"
)

func TestMultipartFormData_FromBytes(t *testing.T) {
	form := NewMultipartFormData().SetBoundary("boundary")
	err := form.FromBytes(testdata.MineData)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(form.Parts))
	assert.Equal(t, "field1", form.Parts[0].Disposition.Name)
	assert.Equal(t, "example.txt", form.Parts[1].Disposition.FileName)

	bytes, err := form.ToBytes()
	assert.NoError(t, err)
	assert.NotEmpty(t, bytes)
}

func TestMultipartFormData_FromBytes_1(t *testing.T) {
	form := NewMultipartFormData().SetBoundary("AaB03x")
	err := form.FromBytes(testdata.CharsetMineData)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(form.Parts))
	assert.Equal(t, "_charset_", form.Parts[0].Disposition.Name)
}
