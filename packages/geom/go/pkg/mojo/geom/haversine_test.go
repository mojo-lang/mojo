package geom

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHaversineDistance(t *testing.T) {
	var tests = []struct {
		from     *LngLat
		to       *LngLat
		expected float64
	}{
		{
			&LngLat{Latitude: 22.55, Longitude: 43.12},
			&LngLat{Latitude: 13.45, Longitude: 100.28},
			6094544.408786774,
		},
		{
			&LngLat{Latitude: 20.10, Longitude: 57.30},
			&LngLat{Latitude: 0.57, Longitude: 100.21},
			5145525.771394785,
		},
		{
			&LngLat{Latitude: 51.45, Longitude: 1.15},
			&LngLat{Latitude: 41.54, Longitude: 12.27},
			1389179.3118293067,
		},
		{
			&LngLat{Latitude: 22.34, Longitude: 17.05},
			&LngLat{Latitude: 51.56, Longitude: 4.29},
			3429893.10043882,
		},
		{
			&LngLat{Latitude: 63.24, Longitude: 56.59},
			&LngLat{Latitude: 8.50, Longitude: 13.14},
			6996185.95539861,
		},
		{
			&LngLat{Latitude: 90.00, Longitude: 0.00},
			&LngLat{Latitude: 48.51, Longitude: 2.21},
			4613477.506482742,
		},
		{
			&LngLat{Latitude: 45.04, Longitude: 7.42},
			&LngLat{Latitude: 3.09, Longitude: 101.42},
			10078111.954385415,
		},
	}

	for _, input := range tests {
		d1 := input.from.Distance(input.to)
		assert.InDelta(t, input.expected, d1, 1e-5)

		d2 := input.from.GreatCircleDistance(input.to)
		assert.InDelta(t, input.expected, d2, 1e-5)

		assert.InDelta(t, d1, d2, 1e-5)
	}
}
