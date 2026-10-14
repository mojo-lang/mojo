package http

import (
	nethttp "net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCookieRoundTrip(t *testing.T) {
	for _, mode := range []nethttp.SameSite{nethttp.SameSiteDefaultMode, nethttp.SameSiteLaxMode, nethttp.SameSiteStrictMode, nethttp.SameSiteNoneMode} {
		source := &nethttp.Cookie{Name: "session", Value: "abc", Path: "/", Domain: "example.com", Expires: time.Date(2030, 1, 2, 3, 4, 5, 0, time.UTC), MaxAge: 60, Secure: true, HttpOnly: true, SameSite: mode}
		cookie := FromHttpCookie(source)
		require.Equal(t, SameSite(mode-1), cookie.SameSite)
		actual := cookie.ToHttpCookie()
		actual.Expires = actual.Expires.UTC()
		require.Equal(t, source, actual)
		require.Equal(t, source.String(), cookie.ToString())
	}
}

func TestCookieNilConversions(t *testing.T) {
	var cookie *Cookie
	require.Nil(t, cookie.FromHttpCookie(&nethttp.Cookie{}))
	require.Nil(t, cookie.ToHttpCookie())
	existing := &Cookie{Name: "keep"}
	require.Same(t, existing, existing.FromHttpCookie(nil))
	require.Equal(t, "keep", existing.Name)
}
