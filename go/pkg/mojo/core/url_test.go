package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUrlCodec(t *testing.T) {
	url := Url{
		Scheme: "https",
		Authority: &Url_Authority{
			Host: "mojo-lang.org",
		},
		Path:  "root",
		Query: NewUrlQuery(),
	}

	url.GetQuery().Add("test", "value")
	assert.Equal(t, "https://mojo-lang.org/root?test=value", url.Format())
}

func TestUrlStringCodec_Decode(t *testing.T) {
	url := "http://192.168.5.113:9090//VulnerableApp/PersistentXSSInHTMLTagVulnerability/LEVEL_3?comment=Referer: https://example.com/%0D%0AContent-Length: 0%0D%0A%0D%0A<script>alert('HTTP Header Injection')%3B</script>\\n"

	r, e := ParseUrl(url)
	assert.NoError(t, e)
	assert.NotNil(t, r)
	assert.Equal(t, 1, len(r.Query.Vals))

	url = "http://192.168.5.113:9090//VulnerableApp/XSSInImgTagAttribute/LEVEL_7?src=Set-Cookie: session=1234567890%0D%0AContent-Length: 0%0D%0A%0D%0A<script>alert('CRLF Vulnerability');</script>\\n"
	_, e = ParseUrl(url)
	assert.Error(t, e)

	url = "http://192.168.0.1、80"
	r, e = ParseUrl(url)
	assert.NoError(t, e)
}
