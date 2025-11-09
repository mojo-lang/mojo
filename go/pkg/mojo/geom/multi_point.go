package geom

func NewMultiPoint(points ...*Point) *MultiPoint {
	multipoint := &MultiPoint{}
	multipoint.Points = append(multipoint.Points, points...)
	return multipoint
}

func NewMultiPointFrom(points ...*LngLat) *MultiPoint {
	multipoint := &MultiPoint{}
	for _, p := range points {
		multipoint.Points = append(multipoint.Points, NewPoint(p))
	}
	return multipoint
}

func (x *MultiPoint) GetLngLats() []*LngLat {
	var lngLats []*LngLat
	if x != nil {
		for _, p := range x.Points {
			lngLats = append(lngLats, p.Coordinate)
		}
	}
	return lngLats
}

func (x *MultiPoint) ToGeometry() *Geometry {
	if x != nil {
		return &Geometry{
			Geometry: &Geometry_MultiPoint{MultiPoint: x},
		}
	}
	return nil
}

func (x *MultiPoint) BoundingBox() *BoundingBox {
	if x != nil && len(x.Points) > 0 {
		box := &BoundingBox{}
		for _, pt := range x.Points {
			box.ExtendToPoint(pt.Coordinate)
		}
		return box
	}
	return nil
}

func (x *MultiPoint) CoordTransform(from, to SpatialReference) *MultiPoint {
	if x != nil && len(x.Points) > 0 {
		mp := &MultiPoint{}
		for _, p := range x.Points {
			mp.Points = append(mp.Points, p.CoordTransform(from, to))
		}
		return mp
	}
	return x
}
