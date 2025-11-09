package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTimeFormat2Layout(t *testing.T) {
	cases := map[string]string{
		"YYYY-MM-DDThh:mm:ssZhh:mm":      "2006-01-02T15:04:05Z07:00",
		"YYYY-MM-DDThh:mm:ss.nanoZhh:mm": "2006-01-02T15:04:05.999999999Z07:00",
		"Week, DD-Mon-YY hh:mm:ss MST":   "Monday, 02-Jan-06 15:04:05 MST",
		"h:mmPM":                         "3:04PM",
	}

	for f, l := range cases {
		layout := TimeFormat2Layout(f)
		assert.Equal(t, l, layout)
	}
}
