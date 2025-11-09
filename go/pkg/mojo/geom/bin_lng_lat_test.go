package geom

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinLngLat(t *testing.T) {
	geoPoint := &LngLat{Longitude: 121.1234567, Latitude: 32.1234567, Altitude: 1.23}

	binPoint := NewBinLngLatWith(geoPoint)

	bp := NewBinLngLat(1211234567, 321234567)
	assert.Equal(t, binPoint.Longitude, bp.Longitude)
	assert.Equal(t, binPoint.Latitude, bp.Latitude)

	assert.Equal(t, int32(321234567), binPoint.Latitude)
	assert.Equal(t, int32(1211234567), binPoint.Longitude)
	assert.Equal(t, int32(123), binPoint.Altitude)

	assert.Equal(t, geoPoint.Longitude, binPoint.Decode().Longitude)
	assert.Equal(t, geoPoint.Latitude, binPoint.Decode().Latitude)
	assert.Equal(t, geoPoint.Altitude, binPoint.Decode().Altitude)

	assert.Equal(t, int32(1211234567), EncodeBinLngLat(geoPoint)[0])
	assert.Equal(t, int32(321234567), EncodeBinLngLat(geoPoint)[1])
	assert.Equal(t, int32(123), EncodeBinLngLat(geoPoint)[2])

	assert.Equal(t, geoPoint.Longitude, DecodeBinLngLat([]int32{1211234567, 321234567, 123}...).Longitude)
	assert.Equal(t, geoPoint.Latitude, DecodeBinLngLat([]int32{1211234567, 321234567, 123}...).Latitude)
	assert.Equal(t, geoPoint.Altitude, DecodeBinLngLat([]int32{1211234567, 321234567, 123}...).Altitude)
}
