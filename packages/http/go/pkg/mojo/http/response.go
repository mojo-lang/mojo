package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/mojo-lang/mojo/packages/core/go/pkg/mojo/core"
)

func NewResponseFrom(response *http.Response) (*Response, error) {
	res := &Response{}
	if err := res.SyncFrom(response); err != nil {
		return nil, err
	}
	return res, nil
}

// Cookies parses and returns the cookies set in the Set-Cookie headers.
func (x *Response) Cookies() []*Cookie {
	if x != nil {
		return x.Headers.GetSetCookies()
	}
	return nil
}

func (x *Response) Cookie(name string) *Cookie {
	for _, cookie := range x.Cookies() {
		if cookie.Name == name {
			return cookie
		}
	}

	return nil
}

func (x *Response) SyncFrom(response *http.Response) error {
	if x != nil && response != nil {
		if response.Body != nil {
			body, err := ioutil.ReadAll(response.Body)
			if err != nil {
				return err
			}
			x.Body = core.NewBytesValue(body)
		}

		x.Version = &Version{
			Major: int32(response.ProtoMajor),
			Minor: int32(response.ProtoMinor),
		}
		x.Status = &Status{}
		if response.StatusCode > 0 {
			x.Status.Code = int32(response.StatusCode)
			x.Status.Reason = strings.TrimSpace(strings.TrimPrefix(response.Status, strconv.Itoa(response.StatusCode)))
		} else {
			segments := strings.Split(response.Status, " ")
			if len(segments) > 0 {
				code, err := strconv.ParseInt(strings.TrimSpace(segments[0]), 10, 32)
				if err != nil {
					return fmt.Errorf("invalid Status: %s, error: %w", response.Status, err)
				}
				x.Status.Code = int32(code)
				x.Status.Reason = strings.Join(segments[1:], " ")
			} else {
				return fmt.Errorf("invalid Status: %s", response.Status)
			}
		}

		if len(response.Header) > 0 {
			x.Headers = NewHeadersFrom(response.Header)
		}
	}

	return nil
}
