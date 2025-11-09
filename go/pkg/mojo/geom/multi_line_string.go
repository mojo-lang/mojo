package geom

func NewMultiLineString(linestring ...*LineString) *MultiLineString {
	multi := &MultiLineString{}
	multi.LineStrings = append(multi.LineStrings, linestring...)
	return multi
}

func NewMultiLineStringFrom(coordinates ...[]*LngLat) *MultiLineString {
	multi := &MultiLineString{}
	for _, c := range coordinates {
		multi.LineStrings = append(multi.LineStrings, &LineString{Coordinates: c})
	}

	return multi
}

func (x *MultiLineString) ToGeometry() *Geometry {
	if x != nil {
		return &Geometry{
			Geometry: &Geometry_MultiLineString{MultiLineString: x},
		}
	}
	return nil
}

func (x *MultiLineString) BoundingBox() *BoundingBox {
	if x != nil && len(x.LineStrings) > 0 {
		var box *BoundingBox
		for _, ls := range x.LineStrings {
			if bb := ls.BoundingBox(); bb != nil {
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

func (x *MultiLineString) CoordTransform(from, to SpatialReference) *MultiLineString {
	if x != nil && len(x.LineStrings) > 0 {
		ml := &MultiLineString{}
		for _, l := range x.LineStrings {
			ml.LineStrings = append(ml.LineStrings, l.CoordTransform(from, to))
		}
		return ml
	}
	return x
}
