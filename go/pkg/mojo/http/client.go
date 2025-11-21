package http

import (
	"bytes"
	"compress/gzip"
	"errors"
	jsoniter "github.com/json-iterator/go"
	"github.com/mojo-lang/mojo/go/pkg/logs"
	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"io"
	"net/http"
)

func logContent(content []byte) string {
	const MaxLogLength = 1024
	if len(content) > MaxLogLength {
		return string(content[0:MaxLogLength]) + "..."
	} else {
		return string(content)
	}
}

func Post(url string, req interface{}, resp interface{}) error {
	return Do(Method_METHOD_POST, url, nil, req, resp)
}

func Put(url string, req interface{}, resp interface{}) error {
	return Do(Method_METHOD_PUT, url, nil, req, resp)
}

func Do(method Method, url string, headers *Headers, req interface{}, resp interface{}) error {
	var content []byte
	var err error

	if req != nil {
		content, err = jsoniter.Marshal(req)
		if err != nil {
			return err
		}
	}

	hr, err := http.NewRequest(method.Format(), url, bytes.NewReader(content))
	if err != nil {
		return err
	}
	if headers != nil {
		hr.Header = make(http.Header)
		headers.SyncTo(hr.Header)
	}
	if req != nil && len(hr.Header.Get("Content-Type")) == 0 {
		hr.Header.Set("Content-Type", core.ApplicationJson)
	}

	r, err := http.DefaultClient.Do(hr)
	if err != nil {
		return err
	}
	defer func() {
		_ = r.Body.Close()
	}()

	if r.StatusCode >= 400 {
		return errors.New("http return error " + r.Status)
	}

	if resp != nil {
		var reader io.ReadCloser
		switch r.Header.Get("Content-Encoding") {
		case "gzip":
			reader, err = gzip.NewReader(r.Body)
			defer func() {
				_ = reader.Close()
			}()
			if err != nil {
				return errors.New("failed to read the gzip content")
			}
		default:
			reader = r.Body
		}

		content, err = io.ReadAll(reader)
		if err != nil {
			return err
		}

		if len(content) > 0 && resp != nil {
			if _, ok := resp.(*[]byte); ok {
				resp = content
			} else {
				if err = jsoniter.Unmarshal(content, resp); err != nil {
					logs.Warnf("the response is: %s", logContent(content))
					return err
				}
			}
		}
	}

	return nil
}

func Get(url string, resp interface{}) error {
	return Do(Method_METHOD_GET, url, nil, nil, resp)
}

func GetContent(url string) ([]byte, error) {
	var bs []byte
	err := Do(Method_METHOD_GET, url, nil, nil, &bs)
	if err != nil {
		return nil, err
	}
	return bs, nil
}

func Delete(url string) error {
	return Do(Method_METHOD_DELETE, url, nil, nil, nil)
}
