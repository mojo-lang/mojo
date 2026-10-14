package tile

import (
	"testing"

	"github.com/mojo-lang/mojo/go/pkg/mojo/geom"
	"github.com/stretchr/testify/require"
)

func TestTileProjectionKnownCoordinatesAndClamping(t *testing.T) {
	for _, level := range []int32{0, 1, 8, 16} {
		n := int32(1) << level
		require.Equal(t, int32(0), GetTileX(-200, level))
		require.Equal(t, n-1, GetTileX(200, level))
		require.Equal(t, int32(0), GetTileY(90, level))
		require.Equal(t, n-1, GetTileY(-90, level))
		require.InDelta(t, -180, GetTileLng(0, level), 1e-9)
		require.InDelta(t, 180, GetTileLng(float64(n), level), 1e-9)
		require.InDelta(t, 0, GetTileLat(float64(n)/2, level), 1e-9)
		x, y := LngLatToPixelXY(geom.NewLngLat(0, 0), level)
		require.Equal(t, MapSize(level)/2, x)
		require.Equal(t, MapSize(level)/2, y)
		coord := PixelXYToLngLat(x, y, level)
		require.InDelta(t, 0, coord.Longitude, 1e-9)
		require.InDelta(t, 0, coord.Latitude, 1e-9)
	}
	require.Equal(t, 0, Clip(-1, 0, 10))
	require.Equal(t, 10, Clip(11, 0, 10))
	require.Equal(t, 4, Clip(4, 0, 10))
	require.InDelta(t, GroundResolution(0, 4)/2, GroundResolution(0, 5), 1e-9)
	require.InDelta(t, GroundResolution(0, 5)*96/0.0254, MapScale(0, 5, 96), 1e-9)
}

func TestTileLocalCoordinatesRoundTrip(t *testing.T) {
	tid := &TileId{X: 4, Y: 5, Level: 4}
	point := TileXY2LonLat(tid, 128, 128)
	x, y := LngLat2TileXY(point, tid)
	require.InDelta(t, 128, x, 1)
	require.InDelta(t, 128, y, 1)
	require.Nil(t, TileXY2LonLat(nil, 0, 0))
	x, y = LngLat2TileXY(nil, tid)
	require.Zero(t, x)
	require.Zero(t, y)
}
