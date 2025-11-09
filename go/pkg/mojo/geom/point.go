package geom

func NewPoint(coordinate *LngLat) *Point {
	return &Point{Coordinate: coordinate}
}

func NewPointFrom(coordinates ...float64) *Point {
	return &Point{Coordinate: NewLngLat(coordinates...)}
}

func (x *Point) GetLongitude() float64 {
	return x.GetCoordinate().GetLongitude()
}

func (x *Point) GetLatitude() float64 {
	return x.GetCoordinate().GetLatitude()
}

func (x *Point) GetAltitude() float64 {
	return x.GetCoordinate().GetAltitude()
}

func (x *Point) BoundingBox() *BoundingBox {
	if x != nil {
		return &BoundingBox{
			LeftBottom: x.Coordinate,
			RightTop:   x.Coordinate,
		}
	}
	return nil
}

func (x *Point) ToGeometry() *Geometry {
	if x != nil {
		return &Geometry{
			Geometry: &Geometry_Point{Point: x},
		}
	}
	return nil
}

func (x *Point) CoordTransform(from, to SpatialReference) *Point {
	if x != nil {
		return &Point{
			Coordinate: x.Coordinate.CoordTransform(from, to),
		}
	}
	return x
}
