package http

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/stretchr/testify/assert"
)

var getGoReq = &http.Request{
	Method: "GET",
	URL: &url.URL{
		Scheme:   "https",
		Host:     "localhost:8080",
		Path:     "/path/to/service",
		RawQuery: "foo=bar",
	},
	ProtoMajor: 1,
	ProtoMinor: 1,
	Header: map[string][]string{
		"Accept-Encoding": {"gzip, deflate"},
		"Accept-Language": {"en-us"},
		"Foo":             {"Bar", "two"},
	},
}

var getReq = &Request{
	Method: Method_METHOD_GET,
	Url: &core.Url{
		Scheme: "https",
		Authority: &core.Url_Authority{
			Host: "localhost",
			Port: 8080,
		},
		Path: "/path/to/service",
		Query: core.NewUrlQueryFrom(map[string][]string{
			"foo": {"bar"},
		}),
	},
	Version: &Version{Major: 1, Minor: 1},
	Headers: NewHeadersFrom(map[string][]string{
		"Accept-Encoding": {"gzip, deflate"},
		"Accept-Language": {"en-us"},
		"Foo":             {"Bar", "two"},
	}),
}

var obj, _ = core.NewObjectFrom(&FormDataContent{
	Disposition: &FormDataContent_Disposition{
		Name:     "foo",
		FileName: "",
	},
	Value: core.NewStringValue("bar"),
})

var multipartReq = &Request{
	Method: Method_METHOD_POST,
	Url: &core.Url{
		Scheme: "https",
		Authority: &core.Url_Authority{
			Host: "localhost",
			Port: 8080,
		},
		Path: "/path/to/service",
		Query: core.NewUrlQueryFrom(map[string][]string{
			"foo": {"bar"},
		}),
	},
	Version: &Version{Major: 1, Minor: 1},
	Headers: NewHeadersFrom(map[string][]string{
		ContentTypeHeaderName: {core.MultipartFormData + ";boundary=" + "000"},
	}),
	Body: core.NewArrayValue(core.NewObjectValue(obj)),
}

func TestRequest_SyncTo(t *testing.T) {
	req, err := getReq.SyncTo()
	assert.NoError(t, err)
	assert.Equal(t, getGoReq.Method, req.Method)
	assert.Equal(t, getGoReq.URL, req.URL)
	assert.Equal(t, getGoReq.ProtoMajor, req.ProtoMajor)
	assert.Equal(t, getGoReq.ProtoMinor, req.ProtoMinor)
	assert.Equal(t, getGoReq.Header["Accept-Encoding"], req.Header["Accept-Encoding"])
	assert.Equal(t, getGoReq.Header["Accept-Language"], req.Header["Accept-Language"])
	assert.Equal(t, getGoReq.Header["Foo"], req.Header["Foo"])
}

func TestRequest_SyncTo2(t *testing.T) {

}

func TestRequest_SyncFrom(t *testing.T) {
	req := &Request{}
	err := req.SyncFrom(getGoReq)

	assert.NoError(t, err)
	assert.Equal(t, getReq.Method, req.Method)
	assert.Equal(t, getReq.Url.Format(), req.Url.Format())
	assert.Equal(t, getReq.Version.Major, req.Version.Major)
	assert.Equal(t, getReq.Version.Minor, req.Version.Minor)
	assert.Equal(t, getReq.Headers.Values("Accept-Encoding"), req.Headers.Values("Accept-Encoding"))
	assert.Equal(t, getReq.Headers.Values("Accept-Language"), req.Headers.Values("Accept-Language"))
	assert.Equal(t, getReq.Headers.Values("Foo"), req.Headers.Values("Foo"))
}
