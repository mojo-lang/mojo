package geom

// NewLinearRing make sure LineString is a closed ring
func NewLinearRing(coordinates ...*LngLat) *LineString {
	coordsLen := len(coordinates)
	if coordsLen == 2 {
		box := NewBoundingBox(coordinates[0], coordinates[1])
		points := box.Vertexes()
		points = append(points, box.Vertex(0))
		return NewLineString(points...)
	} else if coordsLen > 2 {
		if !coordinates[0].Equal(coordinates[len(coordinates)-1]) {
			coordinates = append(coordinates, coordinates[0])
		}
		return NewLineString(coordinates...)
	}
	return nil
}

func NewLineString(coordinates ...*LngLat) *LineString {
	line := &LineString{}
	line.Coordinates = append(line.Coordinates, coordinates...)
	return line
}

func (x *LineString) Add(point *LngLat) {
	x.Coordinates = append(x.Coordinates, point)
}

func (x *LineString) IsClosed() bool {
	if x != nil && len(x.Coordinates) >= 3 {
		return x.Coordinates[0].Equal(x.Coordinates[len(x.Coordinates)-1])
	}
	return false
}

// Clockwise check the closed LineString is clockwise or not
// 0: unknown
// 1: clockwise
// -1: anti-clockwise
func (x *LineString) Clockwise() int {
	if x.IsClosed() {
		mi := 0
		maxLat := float64(-90)
		for i, ll := range x.Coordinates {
			if ll.Latitude > maxLat {
				maxLat = ll.Latitude
				mi = i
			}
		}

		cp := float64(0)
		if mi == 0 {
			cp = x.Coordinates[0].CrossProduct(x.Coordinates[len(x.Coordinates)-2], x.Coordinates[1])
		} else if mi == len(x.Coordinates)-1 {
			cp = x.Coordinates[len(x.Coordinates)-1].CrossProduct(x.Coordinates[1], x.Coordinates[len(x.Coordinates)-2])
		} else {
			cp = x.Coordinates[mi].CrossProduct(x.Coordinates[mi-1], x.Coordinates[mi+1])
		}
		if cp < 0 {
			return 1
		} else {
			return -1
		}
	}
	return 0
}

func (x *LineString) ToGeometry() *Geometry {
	if x != nil {
		return &Geometry{
			Geometry: &Geometry_LineString{LineString: x},
		}
	}
	return nil
}

func (x *LineString) BoundingBox() *BoundingBox {
	if x != nil && len(x.Coordinates) > 0 {
		box := &BoundingBox{}
		for _, ll := range x.Coordinates {
			box = box.ExtendToPoint(ll)
		}
		return box
	}
	return nil
}

func (x *LineString) CoordTransform(from, to SpatialReference) *LineString {
	if x != nil && len(x.Coordinates) > 0 {
		ls := &LineString{}
		for _, c := range x.Coordinates {
			ls.Coordinates = append(ls.Coordinates, c.CoordTransform(from, to))
		}
		return ls
	}
	return x
}
