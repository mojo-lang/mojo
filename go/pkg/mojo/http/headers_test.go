package http

import (
	nethttp "net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHeaders_AddCookies(t *testing.T) {
	headers := NewHeaders().AddCookies(&Cookie{Name: "id", Value: "0123456"}, &Cookie{Name: "key", Value: "0123456"})

	values := headers.Values("Cookie")
	assert.NotEmpty(t, values)
	assert.Equal(t, "id=0123456", values[0])
	assert.Equal(t, "key=0123456", values[1])
}

func TestHeadersSyncPreservesAllValues(t *testing.T) {
	source := nethttp.Header{"X-Test": {"first", "second"}, "Content-Type": {"application/json"}}
	headers := NewHeadersFrom(source)
	destination := nethttp.Header{"X-Test": {"old"}, "Other": {"keep"}}
	require.Same(t, headers, headers.SyncTo(destination))
	require.Equal(t, source.Values("X-Test"), destination.Values("X-Test"))
	require.Equal(t, "keep", destination.Get("Other"))
	headers.Set("x-test", "replacement").Add("X-TEST", "second")
	require.Equal(t, []string{"replacement", "second"}, headers.Values("x-test"))
	headers.Del("X-test")
	require.Empty(t, headers.Get("X-Test"))
	headers.SetBearToken("token")
	require.Equal(t, "token", headers.GetBearToken())
	media, err := headers.GetContentType()
	require.NoError(t, err)
	require.NotNil(t, media)
}

func TestHeaders_AddSetCookies(t *testing.T) {
	headers := NewHeaders().AddSetCookies(&Cookie{Name: "id", Value: "0123456", Secure: true},
		&Cookie{Name: "key", Value: "0123456", Domain: "company.com", HttpOnly: true})

	values := headers.Values("Set-Cookie")
	assert.NotEmpty(t, values)
	assert.Equal(t, "id=0123456; Secure", values[0])
	assert.Equal(t, "key=0123456; Domain=company.com; HttpOnly", values[1])
}
