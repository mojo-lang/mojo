package geom

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLngLat_Parse(t *testing.T) {
	cases := []*LngLatCase{{
		LngLat: &LngLat{Longitude: 123.1, Latitude: 32.2},
		Value:  "123.1,  32.2  ",
	}, {
		LngLat: &LngLat{Longitude: 123.1, Latitude: 32.2},
		Value:  "[123.1,  32.2  ]",
	}, {
		LngLat: &LngLat{Longitude: 123.1, Latitude: 32.2, Altitude: 34},
		Value:  "  [123.1,  32.2 , 34  ] ",
	}, {
		LngLat: &LngLat{Longitude: 123.1, Latitude: 32.2},
		Value:  "   123.1,  32.2  ",
	}}

	for _, c := range cases {
		lngLat := &LngLat{}
		err := lngLat.Parse(c.Value)
		assert.NoError(t, err)
		assert.Equal(t, c.LngLat.Longitude, lngLat.Longitude)
		assert.Equal(t, c.LngLat.Latitude, lngLat.Latitude)
		assert.Equal(t, c.LngLat.Altitude, lngLat.Altitude)
	}
}