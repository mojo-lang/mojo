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
	var content []byte
	var err error

	if req != nil {
		content, err = jsoniter.Marshal(req)
		if err != nil {
			return err
		}
	}

	r, err := http.Post(url, core.ApplicationJson, bytes.NewReader(content))
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
		content, err = io.ReadAll(r.Body)
		if err != nil {
			return err
		}

		if len(content) > 0 {
			if err = jsoniter.Unmarshal(content, resp); err != nil {
				logs.Warnf("the response is: %s", logContent(content))
				return err
			}
		}
	}

	return nil
}

func Get(url string, resp interface{}) error {
	content, err := GetContent(url)
	if err != nil {
		return err
	}

	if resp != nil && len(content) > 0 {
		if err = jsoniter.Unmarshal(content, resp); err != nil {
			logs.Warnf("the response is: %s", logContent(content))
			return err
		}
	}

	return nil
}

func GetContent(url string) ([]byte, error) {
	r, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = r.Body.Close()
	}()

	if r.StatusCode >= 400 {
		return nil, errors.New("http return error " + r.Status)
	}

	var reader io.ReadCloser
	switch r.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(r.Body)
		defer func() {
			_ = reader.Close()
		}()
		if err != nil {
			return nil, errors.New("failed to read the gzip content")
		}
	default:
		reader = r.Body
	}

	content, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func Delete(url string) error {
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}
	_, err = http.DefaultClient.Do(req)
	return err
}
