package geom

import (
	"github.com/golang/geo/s2"
	"github.com/mmcloughlin/geohash"
)

func (x *Point) CellCovering(minLevel, maxLevel, maxCells int) s2.CellUnion {
	_ = minLevel
	_ = maxCells

	cell := s2.CellFromLatLng(s2.LatLngFromDegrees(x.Coordinate.Latitude, x.Coordinate.Longitude))
	cellID := cell.ID()
	if cellID.Level() > maxLevel {
		cellID = cellID.Parent(maxLevel)
	}

	cellUnion := s2.CellUnion{}
	cellUnion = append(cellUnion, cellID)
	return cellUnion
}

func (x *Point) GeoHashCovering(minLevel, maxLevel, maxHashes int) []string {
	_ = minLevel
	_ = maxHashes

	hash := geohash.EncodeWithPrecision(x.Coordinate.Latitude, x.Coordinate.Longitude, uint(maxLevel))
	return []string{hash}
}
