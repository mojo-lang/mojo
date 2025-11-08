package tile

import (
	"math"

	"github.com/golang/geo/s2"
	"github.com/mojo-lang/mojo/packages/geom/go/pkg/mojo/geom"
)

const (
	EarthRadius = 6378137.0
	LatitudeMax = 85.051128779806604
	LatitudeMin = -85.051128779806604

	LongitudeMax = 180.0
	LongitudeMin = -180.0

	DegreeMax = 360.0

	Deg2Rad = math.Pi / 180.0
	Rad2Deg = 180.0 / math.Pi
)

func NewTileId(x, y, level int32) *TileId {
	return &TileId{
		X:     x,
		Y:     y,
		Level: level,
	}
}

func NewTileIdFromQuadKey(s string) *TileId {
	z := int32(len(s))
	id := &TileId{Level: z}

	for i := id.Level; i > 0; i-- {
		mask := int32(1) << uint(i-1)
		switch s[z-i] {
		case '0':
		case '1':
			id.X |= mask
		case '2':
			id.Y |= mask
		case '3':
			id.X |= mask
			id.Y |= mask
		default:
			panic("invalid string " + s)
		}
	}

	return id
}

func (x *TileId) Vertex(index int32) *geom.LngLat {
	switch index {
	case 0:
		return TopLeftPoint(x)
	case 1:
		return TopLeftPoint(&TileId{X: x.X, Y: x.Y + 1, Level: x.Level})
	case 2:
		return TopLeftPoint(&TileId{X: x.X + 1, Y: x.Y + 1, Level: x.Level})
	case 3:
		return TopLeftPoint(&TileId{X: x.X + 1, Y: x.Y, Level: x.Level})
	default:
		return nil
	}
}

func (x *TileId) Polygon() *geom.Polygon {
	polygon := &geom.Polygon{}
	line := &geom.LineString{}
	line.Coordinates = append(line.Coordinates, x.Vertex(0))
	line.Coordinates = append(line.Coordinates, x.Vertex(1))
	line.Coordinates = append(line.Coordinates, x.Vertex(2))
	line.Coordinates = append(line.Coordinates, x.Vertex(3))
	line.Coordinates = append(line.Coordinates, line.Coordinates[0])
	polygon.LineStrings = []*geom.LineString{line}
	return polygon
}

func (x *TileId) CellCovering(minLevel, maxLevel, maxCells int) s2.CellUnion {
	return x.Polygon().CellCovering(minLevel, maxLevel, maxCells)
}

func (x *TileId) SubTiles() []TileId {
	return x.SubLevelTiles(1)
}

func (x *TileId) SubLevelTiles(sublevel int) []TileId {
	total := int(math.Pow(2, float64(sublevel)))

	tiles := make([]TileId, total*total)
	originX := x.X * int32(total)
	originY := x.Y * int32(total)

	for i := 0; i < total; i++ {
		for j := 0; j < total; j++ {
			tiles[total*j+i].X = originX + int32(i)
			tiles[total*j+i].Y = originY + int32(j)
			tiles[total*j+i].Level = x.Level + int32(sublevel)
		}
	}

	return tiles
}

/*
func (t TileId) HexKey() string {
	if t.Level == 0 {

	} else {
		for i := t.Level; i > 0; i = i - 2 {

		}
	}
}*/

func (x *TileId) QuadKey() string {
	key := ""
	for i := x.Level; i > 0; i-- {
		digit := '0'
		mask := int32(1) << uint(i-1)
		if x.X&mask != 0 {
			digit++
		}
		if x.Y&mask != 0 {
			digit += 2
		}
		key += string(digit)
	}

	return key
}

func (x *TileId) ResetQuadKey(key string) {
	id := NewTileIdFromQuadKey(key)
	x.Level = id.Level
	x.X = id.X
	x.Y = id.Y
}

func (x *TileId) LngLat2XY(lnglat *geom.LngLat) (int32, int32) {
	return LngLat2TileXY(lnglat, x)
}

func (x *TileId) XY2LonLat(tx float64, ty float64) *geom.LngLat {
	return TileXY2LonLat(x, tx, ty)
}

func QuadKeyLevel(key string) int {
	return len(key)
}

func GetTileId(longitude, latitude float64, level int32) *TileId {
	lat := Clip(latitude, LatitudeMin, LatitudeMax)
	lng := Clip(longitude, LongitudeMin, LongitudeMax)

	x := (lng + 180) / 360
	r := lat * Deg2Rad
	y := 0.5 - math.Log(math.Tan(r)+1/math.Cos(r))/(2.0*math.Pi)

	size := float64(uint32(1) << uint32(level))
	tileX := int32(Clip(x*size, 0, size-1))
	tileY := int32(Clip(y*size, 0, size-1))

	return &TileId{X: tileX, Y: tileY, Level: level}
}

func TopLeftPoint(id *TileId) *geom.LngLat {
	if id != nil {
		n := math.Pi - 2.0*math.Pi*float64(id.Y)/math.Pow(2.0, float64(id.Level))
		longitude := float64(id.X)/math.Pow(2.0, float64(id.Level))*DegreeMax - LongitudeMax
		latitude := Rad2Deg * math.Atan(0.5*(math.Exp(n)-math.Exp(-n)))
		return &geom.LngLat{Longitude: longitude, Latitude: latitude}
	}
	return nil
}

func MinX(lb *TileId, rt *TileId) int32 {
	if lb == nil {
		if rt == nil {
			return 0
		} else {
			return rt.X
		}
	} else {
		if rt == nil {
			return lb.X
		} else {
			return min(lb.X, rt.X)
		}
	}
}

func MaxX(lb *TileId, rt *TileId) int32 {
	if lb == nil {
		if rt == nil {
			return 0
		} else {
			return rt.X
		}
	} else {
		if rt == nil {
			return lb.X
		} else {
			return max(lb.X, rt.X)
		}
	}
}

func MinY(lb *TileId, rt *TileId) int32 {
	if lb == nil {
		if rt == nil {
			return 0
		} else {
			return rt.Y
		}
	} else {
		if rt == nil {
			return lb.Y
		} else {
			return min(lb.Y, rt.Y)
		}
	}
}

func MaxY(lb *TileId, rt *TileId) int32 {
	if lb == nil {
		if rt == nil {
			return 0
		} else {
			return rt.Y
		}
	} else {
		if rt == nil {
			return lb.Y
		} else {
			return max(lb.Y, rt.Y)
		}
	}
}
