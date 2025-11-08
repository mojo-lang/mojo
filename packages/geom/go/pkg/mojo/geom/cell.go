package geom

import (
	"github.com/golang/geo/s1"
	"github.com/golang/geo/s2"
)

func LatLngToDegrees(latLng s2.LatLng) s2.LatLng {
	return s2.LatLng{Lat: latLng.Lat / s1.Degree, Lng: latLng.Lng / s1.Degree}
}

func CellCenter(cid s2.CellID) *LngLat {
	point := LatLngToDegrees(cid.LatLng())
	return &LngLat{Longitude: float64(point.Lng), Latitude: float64(point.Lat)}
}

func CellVertex(cid s2.CellID, k int) *LngLat {
	cell := s2.CellFromCellID(cid)
	point := LatLngToDegrees(s2.LatLngFromPoint(cell.Vertex(k)))
	return &LngLat{Longitude: float64(point.Lng), Latitude: float64(point.Lat)}
}

func CellVertexes(cid s2.CellID) []*LngLat {
	cell := s2.CellFromCellID(cid)
	return CellVertexesFromCell(cell)
}

func CellVertexesFromCell(cell s2.Cell) []*LngLat {
	points := make([]*LngLat, 4)
	for i := 0; i < 4; i++ {
		point := LatLngToDegrees(s2.LatLngFromPoint(cell.Vertex(i)))
		points[i] = &LngLat{Longitude: float64(point.Lng), Latitude: float64(point.Lat)}
	}

	return points
}

func CellTokens(cells s2.CellUnion) [][]byte {
	tokens := make([][]byte, len(cells))
	for i, id := range cells {
		tokens[i] = []byte(id.ToToken())
	}
	return tokens
}

func CellToPolygon(cid s2.CellID) *Geometry {
	points := CellVertexes(cid)
	points = append(points, points[0])
	return NewPolygonGeometryFrom(points...)
}
