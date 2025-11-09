package geom

import (
	"github.com/golang/geo/s2"
	"math"
)

func NewPolygon(lineStrings ...*LineString) *Polygon {
	polygon := &Polygon{}
	polygon.LineStrings = append(polygon.LineStrings, lineStrings...)
	return polygon
}

func NewPolygonFrom(coordinates ...*LngLat) *Polygon {
	polygon := &Polygon{}
	polygon.LineStrings = []*LineString{NewLinearRing(coordinates...)}
	return polygon
}

func NewPolygonFromCell(cell s2.Cell) *Polygon {
	points := CellVertexesFromCell(cell)
	return NewPolygonFrom(points[0], points[1], points[2], points[3], points[0])
}

func NewPolygonFromCellID(cid s2.CellID) *Polygon {
	return NewPolygonFromCell(s2.CellFromCellID(cid))
}

func NewPolygonFromGeoHash(geohash string) *Polygon {
	points := GeoHashVertexes(geohash)
	return NewPolygonFrom(points[0], points[1], points[2], points[3], points[0])
}

// Add Appends the passed in contour to the current Polygon.
// Notice: point is add to first LineString of the Polygon.
func (x *Polygon) Add(point *LngLat) *Polygon {
	if x != nil && len(x.LineStrings) > 0 {
		x.LineStrings[0].Coordinates = append(x.LineStrings[0].Coordinates, point)
	}

	return x
}

// Invert all LineStrings in the Polygon.
func (x *Polygon) Invert() {
	// For non-special loops, reverse the slice of vertices.
	for _, line := range x.LineStrings {
		for i := len(line.Coordinates)/2 - 1; i >= 0; i-- {
			opp := len(line.Coordinates) - 1 - i
			line.Coordinates[i], line.Coordinates[opp] = line.Coordinates[opp], line.Coordinates[i]
		}
	}
}

// IsClosed returns whether or not the polygon is closed.
// TODO:  This can obviously be improved, but for now,
//
//	this should be sufficient for detecting if points
//	are contained using the raycast algorithm.
func (x *Polygon) IsClosed() bool {
	for _, line := range x.LineStrings {
		if len(line.Coordinates) < 3 {
			return false
		}
	}

	return true
}

// Contains returns whether the current Polygon contains the passed in Point.
// If the Point is contained by only one LineString, it's contained by the Polygon.
func (x *Polygon) Contains(point *LngLat) bool {
	if !x.IsClosed() {
		return false
	}

	contains := false
	for _, line := range x.LineStrings {
		start := len(line.Coordinates) - 1
		end := 0

		if x.intersectsWithRaycast(point, line.Coordinates[start], line.Coordinates[end]) {
			contains = !contains
		}

		for i := 1; i < len(line.Coordinates); i++ {
			if x.intersectsWithRaycast(point, line.Coordinates[i-1], line.Coordinates[i]) {
				contains = !contains
			}
		}
	}

	return contains
}

// Area Assumption: The holes are not intersected with each other.
func (x *Polygon) Area() float64 {
	var area float64
	if len(x.LineStrings) == 1 && x.LineStrings[0].Clockwise() == 1 {
		x.Invert()
	}

	for idx, line := range x.LineStrings {
		var pts []s2.Point
		for _, pt := range line.Coordinates {
			pts = append(pts, s2.PointFromLatLng(s2.LatLngFromDegrees(pt.Latitude, pt.Longitude)))
		}
		if len(pts) > 0 && pts[0] == pts[len(pts)-1] {
			pts = pts[:len(pts)-1]
		}
		loop := s2.LoopFromPoints(pts)

		if idx == 0 {
			area += loop.Area()
		} else {
			area -= loop.Area()
		}
	}
	return area
}

// intersectsWithRaycast Using the raycast algorithm, this returns whether or not the passed in point
// Intersects with the edge drawn by the passed in start and end points.
// Original implementation: http://rosettacode.org/wiki/Ray-casting_algorithm#Go
func (x *Polygon) intersectsWithRaycast(point *LngLat, start *LngLat, end *LngLat) bool {
	// Always ensure that the the first point
	// has a y coordinate that is less than the second point
	if start.Longitude > end.Longitude {

		// Switch the points if otherwise.
		start, end = end, start

	}

	// Move the point's y coordinate
	// outside of the bounds of the testing region
	// so we can start drawing a ray
	for point.Longitude == start.Longitude || point.Longitude == end.Longitude {
		newLng := math.Nextafter(point.Longitude, math.Inf(1))
		point = NewLngLat(newLng, point.Latitude)
	}

	// If we are outside of the polygon, indicate so.
	if point.Longitude < start.Longitude || point.Longitude > end.Longitude {
		return false
	}

	if start.Latitude > end.Latitude {
		if point.Latitude > start.Latitude {
			return false
		}
		if point.Latitude < end.Latitude {
			return true
		}

	} else {
		if point.Latitude > end.Latitude {
			return false
		}
		if point.Latitude < start.Latitude {
			return true
		}
	}

	raySlope := (point.Longitude - start.Longitude) / (point.Latitude - start.Latitude)
	diagSlope := (end.Longitude - start.Longitude) / (end.Latitude - start.Latitude)

	return raySlope >= diagSlope
}

//func (x *Polygon) ToArray() interface{} {
//	if x != nil {
//		if len(x.LineStrings) == 1 {
//			return x.LineStrings[0].Coordinates
//		} else if len(x.LineStrings) > 1 {
//			return x.LineStrings
//		}
//	}
//	return nil
//}

func (x *Polygon) ToGeometry() *Geometry {
	if x != nil {
		return &Geometry{
			Geometry: &Geometry_Polygon{Polygon: x},
		}
	}
	return nil
}

func (x *Polygon) BoundingBox() *BoundingBox {
	if x != nil && len(x.LineStrings) > 0 {
		return x.LineStrings[0].BoundingBox()
	}
	return nil
}

func (x *Polygon) CoordTransform(from, to SpatialReference) *Polygon {
	if x != nil && len(x.LineStrings) > 0 {
		p := &Polygon{}
		for _, l := range x.LineStrings {
			p.LineStrings = append(p.LineStrings, l.CoordTransform(from, to))
		}
		return p
	}
	return x
}
