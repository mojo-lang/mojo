package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUrl_FormatWithoutSchema(t *testing.T) {
	const URL = "apis.company.com/path/to/resource#anchor"
	url, _ := ParseUrl("https://" + URL)
	assert.Equal(t, URL, url.FormatWithoutSchema())

	url, _ = ParseUrl("http://" + URL)
	assert.Equal(t, URL, url.FormatWithoutSchema())
}

func TestUrl_Format(t *testing.T) {
	const URL = "https://apis.company.com/path/to/resource#anchor"
	url, err := ParseUrl(URL)
	assert.NoError(t, err)
	assert.Equal(t, URL, url.Format())
	assert.Equal(t, "https", url.Scheme)
}

func TestUrl_Parse(t *testing.T) {
	const URL = "apis.company.com:8080/path/to/resource#anchor"
	url, err := ParseUrl(URL)
	assert.NoError(t, err)
	assert.Equal(t, "apis.company.com", url.Authority.Host)
	assert.Equal(t, int64(8080), url.Authority.Port)
	assert.Equal(t, URL, url.Format())

	const URL2 = "apis.company.com/path/to/resource#anchor"
	url, err = ParseUrl(URL2)
	assert.NoError(t, err)
	assert.Equal(t, URL2, url.Format())
}

func TestUrl_Parse2(t *testing.T) {
	const URL = "path/to/resource#anchor"
	url, err := ParseUrl(URL)
	assert.NoError(t, err)
	assert.Equal(t, URL, url.Format())

	const URL2 = "#anchor/name"
	url, err = ParseUrl(URL2)
	assert.NoError(t, err)
	assert.NotEmpty(t, url.Fragment)
	assert.Equal(t, URL2, url.Format())
}
