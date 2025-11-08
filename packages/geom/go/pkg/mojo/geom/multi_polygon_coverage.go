package geom

import "github.com/golang/geo/s2"

func (x *MultiPolygon) CellCovering(minLevel, maxLevel, maxCells int) s2.CellUnion {
	cells := s2.CellUnion{}
	for _, polygon := range x.Polygons {
		cells = append(cells, polygon.CellCovering(minLevel, maxLevel, maxCells)...)
		if len(cells) >= maxCells {
			return cells
		}
	}
	return cells
}

func (x *MultiPolygon) GeoHashCovering(minLevel, maxLevel, maxHashes int) []string {
	var hashes []string
	for _, polygon := range x.Polygons {
		hashes = append(hashes, polygon.GeoHashCovering(minLevel, maxLevel, maxHashes)...)
		if len(hashes) >= maxHashes {
			return hashes
		}
	}
	return hashes
}
