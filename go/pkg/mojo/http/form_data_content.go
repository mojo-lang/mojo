package http

import (
	"bytes"
	"errors"
	"mime"
	"net/textproto"
	"strings"

	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
)

var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

func escapeQuotes(s string) string {
	return quoteEscaper.Replace(s)
}

func ParseFormDataContentFrom(obj *core.Object) (*FormDataContent, error) {
	content := &FormDataContent{}
	if err := content.ParseFrom(obj); err != nil {
		return nil, err
	}
	return content, nil
}

func (x *FormDataContent) MIMEHeader() textproto.MIMEHeader {
	if x == nil {
		return nil
	}

	headers := make(textproto.MIMEHeader)
	if x.Disposition != nil {
		value := "form-data"
		if len(x.Disposition.Value) > 0 {
			value = x.Disposition.Value
		}
		if len(x.Disposition.Name) > 0 {
			value += `; name="` + escapeQuotes(x.Disposition.Name) + `"`
		}
		if len(x.Disposition.FileName) > 0 {
			value += `; filename="` + escapeQuotes(x.Disposition.FileName) + `"`
		}
		headers["Content-Disposition"] = []string{value}
	}

	if x.Type != nil {
		headers[ContentTypeHeaderName] = []string{x.Type.Format()}
	}
	if len(x.TransferEncoding) > 0 {
		headers[ContentTransferEncodingHeaderName] = []string{x.TransferEncoding}
	}

	if x.Headers != nil {
		for k, v := range x.Headers.GetVals() {
			headers[k] = v.GetVals()
		}
	}
	return headers
}

func (x *FormDataContent) SetMIMEHeader(header textproto.MIMEHeader) {
	for k, v := range header {
		switch k {
		case "Content-Disposition":
			if len(v) > 0 {
				disposition, params, err := mime.ParseMediaType(v[0])
				if err != nil {
					params = make(map[string]string)
				}

				x.Disposition = &FormDataContent_Disposition{
					Value:    disposition,
					Name:     params["name"],
					FileName: params["filename"],
				}
			}
		case ContentTypeHeaderName:
			if len(v) > 0 {
				mt, _ := core.ParseMediaType(v[0])
				x.Type = mt
			}
		case ContentTransferEncodingHeaderName:
			if len(v) > 0 {
				x.TransferEncoding = v[0]
			}
		default:
			x.Headers.SetValues(k, v...)
		}
	}
}

func (x *FormDataContent) Write(p []byte) (int, error) {
	if x != nil {
		buffer := bytes.Buffer{}
		cnt, err := buffer.Write(p)
		if err != nil {
			return 0, err
		}

		bs := append(x.Value.GetBytes(), buffer.Bytes()...)
		x.Value = core.NewBytesValue(bs)
		return cnt, nil
	}

	return 0, nil
}

func (x *FormDataContent) Valid() bool {
	return x != nil && x.Value != nil
}

func (x *FormDataContent) ParseFrom(obj *core.Object) error {
	if obj != nil {
		if err := obj.To(x); err != nil {
			return err
		}
		if !x.Valid() {
			return errors.New("invalid form data content")
		}
		return nil
	}
	return errors.New("input obj is null")
}
