package http

import (
	"bytes"
	"errors"
	"io"
	"mime"
	"testing"

	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/stretchr/testify/require"
)

type brokenFormWriter struct{ err error }

func (w brokenFormWriter) Write([]byte) (int, error) { return 0, w.err }

func TestMultipartRoundTripAndByteCount(t *testing.T) {
	form := NewMultipartFormData().SetBoundary("test boundary")
	require.NoError(t, form.WriteField("title", "hello"))
	file, err := form.CreateFormFile("upload", "data.bin")
	require.NoError(t, err)
	_, err = file.Write([]byte{0, 1, 2, 255})
	require.NoError(t, err)
	form.Parts = append(form.Parts, &FormDataContent{Disposition: &FormDataContent_Disposition{Name: "count"}, Value: core.NewInt64Value(42)})
	var out bytes.Buffer
	count, err := form.WriteTo(&out)
	require.NoError(t, err)
	require.Equal(t, int64(out.Len()), count)
	typ, params, err := mime.ParseMediaType(form.FormDataContentType())
	require.NoError(t, err)
	require.Equal(t, "multipart/form-data", typ)
	require.Equal(t, form.Boundary, params["boundary"])
	parsed, err := ParseMultipartFormData(out.Bytes(), form.Boundary)
	require.NoError(t, err)
	require.Len(t, parsed.Parts, 3)
	require.Equal(t, "title", parsed.Parts[0].Disposition.Name)
	require.Equal(t, "hello", string(parsed.Parts[0].Value.GetBytes()))
	require.Equal(t, "data.bin", parsed.Parts[1].Disposition.FileName)
	require.Equal(t, []byte{0, 1, 2, 255}, parsed.Parts[1].Value.GetBytes())
	require.Equal(t, "42", string(parsed.Parts[2].Value.GetBytes()))
	encoded, err := form.ToBytes()
	require.NoError(t, err)
	require.Equal(t, out.Bytes(), encoded)
}

func TestMultipartBoundaryAndIOErrors(t *testing.T) {
	form := NewMultipartFormData().SetBoundary("valid")
	require.Same(t, form, form.SetBoundary("\ninvalid"))
	require.Equal(t, "valid", form.Boundary)
	require.NoError(t, form.WriteField("name", "value"))
	sentinel := errors.New("write failed")
	_, err := form.WriteTo(brokenFormWriter{sentinel})
	require.ErrorIs(t, err, sentinel)
	_, err = (&MultipartFormData{Boundary: "\n"}).WriteTo(io.Discard)
	require.Error(t, err)
	parsed, err := ParseMultipartFormData([]byte("--valid\r\nContent-Disposition: form-data; name=\"x\"\r\n\r\nunterminated"), "valid")
	require.Error(t, err)
	require.Nil(t, parsed)
	empty := NewMultipartFormData()
	encoded, err := empty.ToBytes()
	require.NoError(t, err)
	parsed, err = ParseMultipartFormData(encoded, empty.Boundary)
	require.NoError(t, err)
	require.Empty(t, parsed.Parts)
}
