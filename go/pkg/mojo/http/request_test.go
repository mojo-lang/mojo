package http

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

func TestRequestBodyRoundTrip(t *testing.T) {
	url, err := core.NewUrl("https://example.com/item")
	require.NoError(t, err)
	for _, value := range []*core.Value{core.NewStringValue("hello"), core.NewBytesValue([]byte("hello"))} {
		req := &Request{Method: Method_METHOD_POST, Url: url, Body: value}
		standard, err := req.SyncTo()
		require.NoError(t, err)
		content, err := io.ReadAll(standard.Body)
		require.NoError(t, err)
		require.Equal(t, "hello", string(content))
		standard.Body = io.NopCloser(strings.NewReader("reply"))
		var restored Request
		require.NoError(t, restored.SyncFrom(standard))
		require.Equal(t, []byte("reply"), restored.Body.GetBytes())
		require.Equal(t, Method_METHOD_POST, restored.Method)
	}
}

func TestRequestDoAndCancellation(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		_, _ = io.WriteString(w, "created")
	}))
	defer server.Close()
	url, err := core.NewUrl(server.URL)
	require.NoError(t, err)
	req := &Request{Method: Method_METHOD_POST, Url: url}
	response, err := req.Do(context.Background())
	require.NoError(t, err)
	require.Equal(t, int32(201), response.Status.Code)
	require.Equal(t, []byte("created"), response.Body.GetBytes())
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	_, err = req.Do(ctx)
	require.ErrorIs(t, err, context.Canceled)
	_, err = (&Request{}).SyncTo()
	require.Error(t, err)
	_, err = (&Request{Url: url}).SyncTo()
	require.Error(t, err)
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
