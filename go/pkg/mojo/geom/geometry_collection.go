package geom

func NewGeometryCollection(geometries ...*Geometry) *GeometryCollection {
	geometryCollection := &GeometryCollection{}
	geometryCollection.Geometries = geometries
	return geometryCollection
}

func (x *GeometryCollection) AddGeometries(geometries ...*Geometry) {
	x.Geometries = append(x.Geometries, geometries...)
}

func (x *GeometryCollection) ToGeometry() *Geometry {
	if x != nil {
		return &Geometry{
			Geometry: &Geometry_GeometryCollection{GeometryCollection: x},
		}
	}
	return nil
}

func (x *GeometryCollection) BoundingBox() *BoundingBox {
	if x != nil && len(x.Geometries) > 0 {
		var box *BoundingBox
		for _, g := range x.Geometries {
			bb := g.BoundingBox()
			if bb != nil {
				if box == nil {
					box = bb
				} else {
					box = box.Extend(bb)
				}
			}
		}
	}
	return nil
}

func (x *GeometryCollection) CoordTransform(from, to SpatialReference) *GeometryCollection {
	if x != nil && len(x.Geometries) > 0 {
		c := NewGeometryCollection()
		for _, g := range x.Geometries {
			c.Geometries = append(c.Geometries, g.CoordTransform(from, to))
		}
		return c
	}
	return x
}
