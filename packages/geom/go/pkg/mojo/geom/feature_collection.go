package geom

func (x *FeatureCollection) AddFeatures(features ...*Feature) {
	x.Features = append(x.Features, features...)
}

func NewFeatureCollection(features ...*Feature) *FeatureCollection {
	fc := &FeatureCollection{}
	fc.AddFeatures(features...)
	return fc
}

func (x *FeatureCollection) CoordTransform(from, to SpatialReference) *FeatureCollection {
	if x == nil {
		return x
	}
	fc := &FeatureCollection{
		Type:   x.Type,
		Keys:   x.Keys,
		Values: x.Values,
	}
	for _, f := range x.Features {
		fc.Features = append(fc.Features, f.CoordTransform(from, to))
	}
	return fc
}
