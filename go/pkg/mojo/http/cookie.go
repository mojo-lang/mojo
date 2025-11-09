package http

import (
	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"net/http"
)

func FromHttpCookie(cookie *http.Cookie) *Cookie {
	return (&Cookie{}).FromHttpCookie(cookie)
}

func (x *Cookie) FromHttpCookie(cookie *http.Cookie) *Cookie {
	if x != nil && cookie != nil {
		x.Name = cookie.Name
		x.Value = cookie.Value
		x.Path = cookie.Path
		x.Domain = cookie.Domain
		x.Expires = core.FromTime(cookie.Expires)
		x.MaxAge = int32(cookie.MaxAge)
		x.Secure = cookie.Secure
		x.HttpOnly = cookie.HttpOnly
		x.SameSite = x.SameSite - 1
	}
	return x
}

func (x *Cookie) ToHttpCookie() *http.Cookie {
	if x != nil {
		return &http.Cookie{
			Name:     x.Name,
			Value:    x.Value,
			Path:     x.Path,
			Domain:   x.Domain,
			Expires:  x.Expires.ToTime(),
			MaxAge:   int(x.MaxAge),
			Secure:   x.Secure,
			HttpOnly: x.HttpOnly,
			SameSite: http.SameSite(x.SameSite + 1),
		}
	}
	return nil
}

func (x *Cookie) ToString() string {
	return x.ToHttpCookie().String()
}
