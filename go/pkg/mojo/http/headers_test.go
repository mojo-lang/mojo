package http

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeaders_AddCookies(t *testing.T) {
	headers := NewHeaders().AddCookies(&Cookie{Name: "id", Value: "0123456"}, &Cookie{Name: "key", Value: "0123456"})

	values := headers.Values("Cookie")
	assert.NotEmpty(t, values)
	assert.Equal(t, "id=0123456", values[0])
	assert.Equal(t, "key=0123456", values[1])
}

func TestHeaders_AddSetCookies(t *testing.T) {
	headers := NewHeaders().AddSetCookies(&Cookie{Name: "id", Value: "0123456", Secure: true},
		&Cookie{Name: "key", Value: "0123456", Domain: "company.com", HttpOnly: true})

	values := headers.Values("Set-Cookie")
	assert.NotEmpty(t, values)
	assert.Equal(t, "id=0123456; Secure", values[0])
	assert.Equal(t, "key=0123456; Domain=company.com; HttpOnly", values[1])
}
