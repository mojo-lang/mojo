package tile

import (
	"github.com/mojo-lang/mojo/packages/geom/go/pkg/mojo/geom"
	"golang.org/x/exp/constraints"
	"math"
)

const (
	DefaultTilePixels = 256
)

// GetTileX gets the tile x of the z level from the longitude.
func GetTileX(lng float64, z int32) int32 {
	n := 1 << z // 2^z
	x := int32(math.Floor(float64(n) * (lng + 180) / 360))
	return Clip(x, 0, int32(n-1))
}

// GetTileY gets the tile y of the z level from the latitude.
func GetTileY(lat float64, z int32) int32 {
	lat = Clip(lat, LatitudeMin, LatitudeMax)
	n := 1 << z // 2^z
	rad := lat * math.Pi / 180.0
	y := int32(math.Floor((1 - math.Log(math.Tan(rad)+1/math.Cos(rad))/math.Pi) / 2.0 * float64(n)))
	return Clip(y, 0, int32(n-1))
}

// GetTileLng get the longitude of the x tile
func GetTileLng(x float64, z int32) float64 {
	n := 1 << z // 2^z
	return x/float64(n)*360 - 180
}

// GetTileLat get the latitude of the y tile
func GetTileLat(y float64, z int32) float64 {
	n := 1 << z // 2^z
	return math.Atan(math.Sinh(math.Pi*(1-2*y/float64(n)))) * 180 / math.Pi
}

// Clip
// Clips a number to the specified minimum and maximum values.
func Clip[T constraints.Integer | constraints.Float](n, minValue, maxValue T) T {
	if n < minValue {
		return minValue
	}
	if n > maxValue {
		return maxValue
	}
	return n
}

// GroundResolution
// Determines the ground resolution (in meters per pixel) at a specified
// latitude and level of detail.
func GroundResolution(latitude float64, level int32) float64 {
	latitude = Clip(latitude, LatitudeMin, LatitudeMax)
	return math.Cos(latitude*math.Pi/180) * 2 * math.Pi * EarthRadius / float64(MapSize(level))
}

// MapSize
// Determines the map width and height (in pixels) at a specified level of detail.
func MapSize(level int32) int32 {
	return 256 << level
}

// MapScale
// Determines the map scale at a specified latitude, level of detail,
// and screen resolution.
func MapScale(latitude float64, levelOfDetail int32, screenDpi int32) float64 {
	return GroundResolution(latitude, levelOfDetail) * float64(screenDpi) / 0.0254
}

// LngLatToPixelXY
// Converts a point from latitude/longitude WGS-84 coordinates (in degrees)
// into pixel XY coordinates at a specified level of detail.
func LngLatToPixelXY(lnglat *geom.LngLat, level int32) (int32, int32) {
	longitude := Clip(lnglat.Longitude, LongitudeMin, LongitudeMax)
	latitude := Clip(lnglat.Latitude, LatitudeMin, LatitudeMax)

	lx := (longitude + 180) / 360
	sinLat := math.Sin(latitude * math.Pi / 180)
	ly := 0.5 - math.Log((1+sinLat)/(1-sinLat))/(4*math.Pi)
	mapSize := float64(uint64(DefaultTilePixels) << uint(level))
	pixelX := Clip(lx*mapSize, -mapSize/2, mapSize+mapSize/2-1)
	pixelY := Clip(ly*mapSize, -mapSize/2, mapSize+mapSize/2-1)
	return int32(pixelX), int32(pixelY)
}

func LngLat2TileXY(lnglat *geom.LngLat, tid *TileId) (int32, int32) {
	if lnglat != nil && tid != nil {
		px, py := LngLatToPixelXY(lnglat, tid.Level)
		return px - tid.X*DefaultTilePixels, py - tid.Y*DefaultTilePixels
	}
	return 0, 0
}

// PixelXYToLngLat
// Converts a pixel from pixel XY coordinates at a specified level of detail
// into latitude/longitude WGS-84 coordinates (in degrees).
//
// pixelX: X coordinate of the point, in pixels.
// pixelY: Y coordinates of the point, in pixels.
// level: Level of detail, from 1 (lowest detail) to 23
func PixelXYToLngLat(pixelX, pixelY int32, level int32) *geom.LngLat {
	lng := GetTileLng(float64(pixelX)/DefaultTilePixels, level)
	lat := GetTileLat(float64(pixelY)/DefaultTilePixels, level)
	return geom.NewLngLat(lng, lat)
}

// TileXY2LonLat converts a point x/y in a tile to lng/lat
func TileXY2LonLat(tid *TileId, x, y float64) *geom.LngLat {
	if tid != nil {
		lng := GetTileLng(float64(tid.X)+x/DefaultTilePixels, tid.Level)
		lat := GetTileLat(float64(tid.Y)+y/DefaultTilePixels, tid.Level)
		return geom.NewLngLat(lng, lat)
	}
	return nil
}
