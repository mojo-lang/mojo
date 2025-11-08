package geom

func NewMultiPolygon(coordinates ...*Polygon) *MultiPolygon {
	multipolygon := &MultiPolygon{}
	multipolygon.Polygons = append(multipolygon.Polygons, coordinates...)
	return multipolygon
}

func (x *MultiPolygon) Contains(point *LngLat) bool {
	for _, polygon := range x.Polygons {
		if polygon.Contains(point) {
			return true
		}
	}
	return false
}

func (x *MultiPolygon) Invert() {
	for i := 0; i < len(x.Polygons); i++ {
		x.Polygons[i].Invert()
	}
}

func (x *MultiPolygon) Area() float64 {
	if x != nil {
		area := 0.0
		for _, p := range x.Polygons {
			area += p.Area()
		}
		return area
	}
	return 0
}

func (x *MultiPolygon) ToGeometry() *Geometry {
	if x != nil {
		return &Geometry{
			Geometry: &Geometry_MultiPolygon{MultiPolygon: x},
		}
	}
	return nil
}

func (x *MultiPolygon) BoundingBox() *BoundingBox {
	if x != nil && len(x.Polygons) > 0 {
		var box *BoundingBox
		for _, pg := range x.Polygons {
			if bb := pg.BoundingBox(); bb != nil {
				if box == nil {
					box = bb
				} else {
					box = box.Extend(bb)
				}
			}
		}
		return box
	}
	return nil
}

func (x *MultiPolygon) CoordTransform(from, to SpatialReference) *MultiPolygon {
	if x != nil && len(x.Polygons) > 0 {
		mp := &MultiPolygon{}
		for _, p := range x.Polygons {
			mp.Polygons = append(mp.Polygons, p.CoordTransform(from, to))
		}
		return mp
	}
	return x
}
