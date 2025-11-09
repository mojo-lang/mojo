package http

import (
	"bytes"
	"errors"
	"io"
	"mime/multipart"
	"strings"

	jsoniter "github.com/json-iterator/go"
	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
)

func ParseMultipartFormData(bytes []byte, boundary string) (*MultipartFormData, error) {
	data := NewMultipartFormData().SetBoundary(boundary)

	err := data.FromBytes(bytes)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func NewMultipartFormData() *MultipartFormData {
	return &MultipartFormData{
		Boundary: multipart.NewWriter(nil).Boundary(),
		Parts:    nil,
	}
}

func (x *MultipartFormData) SetBoundary(boundary string) *MultipartFormData {
	if x != nil {
		w := multipart.NewWriter(nil)
		if err := w.SetBoundary(boundary); err == nil && len(w.Boundary()) > 0 {
			x.Boundary = w.Boundary()
		}
	}
	return x
}

func (x *MultipartFormData) FormDataContentType() string {
	b := x.Boundary
	// We must quote the boundary if it contains any of the
	// specials characters defined by RFC 2045, or space.
	if strings.ContainsAny(b, `()<>@,;:\"/[]?= `) {
		b = `"` + b + `"`
	}
	return "multipart/form-data; boundary=" + b
}

func (x *MultipartFormData) CreatePart(headers *Headers) (io.Writer, error) {
	part := &FormDataContent{
		Headers: headers,
	}
	x.Parts = append(x.Parts, part)
	return part, nil
}

func (x *MultipartFormData) CreateFormFile(fieldName, filename string) (io.Writer, error) {
	part := &FormDataContent{
		Disposition: &FormDataContent_Disposition{
			Name:     fieldName,
			FileName: filename,
		},
		Type: core.NewApplicationOctetStream(),
	}
	x.Parts = append(x.Parts, part)
	return part, nil
}

func (x *MultipartFormData) CreateFormField(fieldName string) (io.Writer, error) {
	part := &FormDataContent{
		Disposition: &FormDataContent_Disposition{
			Name: fieldName,
		},
	}
	x.Parts = append(x.Parts, part)
	return part, nil
}

func (x *MultipartFormData) WriteField(fieldName, value string) error {
	p, err := x.CreateFormField(fieldName)
	if err != nil {
		return err
	}

	_, err = p.Write([]byte(value))
	return err
}

func (x *MultipartFormData) WriteTo(writer io.Writer) (int64, error) {
	counter := core.NewCountingWriter(writer)
	w := multipart.NewWriter(counter)
	if err := w.SetBoundary(x.Boundary); err != nil {
		return 0, err
	}

	for _, part := range x.Parts {
		wr, err := w.CreatePart(part.MIMEHeader())
		if err != nil {
			return 0, err
		}

		if part.Value != nil {
			var err error
			switch part.Value.GetKind() {
			case core.ValueKind_VALUE_KIND_BYTES:
				_, err = wr.Write(part.Value.GetBytes())
			case core.ValueKind_VALUE_KIND_STRING:
				_, err = wr.Write([]byte(part.Value.GetString()))
			default:
				if bs, e := jsoniter.Marshal(part.Value); e != nil {
					return 0, e
				} else {
					_, err = writer.Write(bs)
				}
			}
			if err != nil {
				return 0, err
			}
		}
	}

	if err := w.Close(); err != nil {
		return 0, err
	}

	return counter.Count(), nil
}

func (x *MultipartFormData) ReadFrom(r io.Reader) (int64, error) {
	counting := core.NewCountingReader(r)
	reader := multipart.NewReader(counting, x.Boundary)
	for {
		part, err := reader.NextPart()
		if err != nil {
			if errors.Is(err, io.EOF) {
				return counting.Count(), nil
			}

			return 0, err
		}

		value, err := io.ReadAll(part)
		if err != nil {
			return 0, err
		}

		content := &FormDataContent{
			Value: core.NewBytesValue(value),
		}
		content.SetMIMEHeader(part.Header)
		if name := part.FormName(); len(name) > 0 {
			content.Disposition.Name = name
		}
		if filename := part.FileName(); len(filename) > 0 {
			content.Disposition.FileName = filename
		}
		x.Parts = append(x.Parts, content)
	}
}

func (x *MultipartFormData) ToBytes() ([]byte, error) {
	buffer := &bytes.Buffer{}
	if _, err := x.WriteTo(buffer); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func (x *MultipartFormData) FromBytes(bs []byte) error {
	reader := bytes.NewReader(bs)
	_, err := x.ReadFrom(reader)
	return err
}
