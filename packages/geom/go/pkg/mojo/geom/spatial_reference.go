package geom

func CoordTransformSupported(sr SpatialReference) bool {
	switch sr {
	case SpatialReference_SPATIAL_REFERENCE_WGS84,
		SpatialReference_SPATIAL_REFERENCE_GCJ02,
		SpatialReference_SPATIAL_REFERENCE_BD09:
		return true
	}
	return false
}
