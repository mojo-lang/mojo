package geom

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
	"testing"
)

type LngLatCase struct {
	*LngLat
	Value string
}

func TestLngLatCodec_Decode(t *testing.T) {
	cases := []*LngLatCase{{
		LngLat: &LngLat{Longitude: 123.1, Latitude: 32.2},
		Value:  `"123.1,  32.2  "`,
	}, {
		LngLat: &LngLat{Longitude: 123.1, Latitude: 32.2},
		Value:  `"[123.1,  32.2  ]"`,
	}, {
		LngLat: &LngLat{Longitude: 123.1, Latitude: 32.2, Altitude: 34},
		Value:  `"  [123.1,  32.2 , 34  ] "`,
	}, {
		LngLat: &LngLat{Longitude: 123.1, Latitude: 32.2},
		Value:  `"   123.1,  32.2  "`,
	}, {
		LngLat: &LngLat{Longitude: 123.1, Latitude: 32.2},
		Value:  `{"longitude": 123.1, "latitude": 32.2  }`,
	}, {
		LngLat: &LngLat{Longitude: 123.1, Latitude: 32.2},
		Value:  `{"lng": 123.1, "lat": 32.2  }`,
	}}

	for _, c := range cases {
		lngLat := &LngLat{}
		err := jsoniter.ConfigFastest.UnmarshalFromString(c.Value, lngLat)
		assert.NoError(t, err)
		assert.Equal(t, c.LngLat.Longitude, lngLat.Longitude)
		assert.Equal(t, c.LngLat.Latitude, lngLat.Latitude)
		assert.Equal(t, c.LngLat.Altitude, lngLat.Altitude)
	}
}
