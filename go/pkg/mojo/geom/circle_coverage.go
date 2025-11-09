package geom

import "github.com/golang/geo/s2"

func (x *Circle) CellCovering(minLevel, maxLevel, maxCells int) s2.CellUnion {
	return x.ToPolygon(64).CellCovering(minLevel, maxLevel, maxCells)
}

func (x *Circle) GeoHashCovering(minLevel, maxLevel, maxHashes int) []string {
	return x.ToPolygon(64).GeoHashCovering(minLevel, maxLevel, maxHashes)
}

func (x *Circle) GeoHashFractionCovering(minLevel, maxLevel, maxHashes int, fraction float32) []string {
	return x.ToPolygon(64).GeoHashFractionCovering(minLevel, maxLevel, maxHashes, fraction)
}
