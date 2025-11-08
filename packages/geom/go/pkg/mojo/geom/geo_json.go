package geom

import jsoniter "github.com/json-iterator/go"

func NewGeoJson(data interface{}) *GeoJson {
	switch t := data.(type) {
	case *Feature:
		return &GeoJson{GeoJson: &GeoJson_Feature{Feature: t}}
	case *FeatureCollection:
		return &GeoJson{GeoJson: &GeoJson_FeatureCollection{FeatureCollection: t}}
	case *Point:
		return &GeoJson{GeoJson: &GeoJson_Point{Point: t}}
	case *MultiPoint:
		return &GeoJson{GeoJson: &GeoJson_MultiPoint{MultiPoint: t}}
	case *LineString:
		return &GeoJson{GeoJson: &GeoJson_LineString{LineString: t}}
	case *MultiLineString:
		return &GeoJson{GeoJson: &GeoJson_MultiLineString{MultiLineString: t}}
	case *Polygon:
		return &GeoJson{GeoJson: &GeoJson_Polygon{Polygon: t}}
	case *MultiPolygon:
		return &GeoJson{GeoJson: &GeoJson_MultiPolygon{MultiPolygon: t}}
	case *GeometryCollection:
		return &GeoJson{GeoJson: &GeoJson_GeometryCollection{GeometryCollection: t}}
	default:
		return nil
	}
}

func NewGeoJsonFrom(geojson []byte) (*GeoJson, error) {
	g := &GeoJson{}
	err := jsoniter.Unmarshal(geojson, g)
	if err != nil {
		return nil, err
	}
	return g, nil
}

func NewFeatureCollectionGeoJson(features ...*Feature) *GeoJson {
	return NewGeoJson(NewFeatureCollection(features...))
}

func NewPointGeoJson(coordinates *LngLat) *GeoJson {
	return NewGeoJson(NewPoint(coordinates))
}

func NewPointGeoJsonFrom(coordinates ...float64) *GeoJson {
	return NewGeoJson(NewPointFrom(coordinates...))
}

func NewLineStringGeoJson(coordinates ...*LngLat) *GeoJson {
	return NewGeoJson(NewLineString(coordinates...))
}

func NewPolygonGeoJson(lineStrings ...*LineString) *GeoJson {
	return NewGeoJson(NewPolygon(lineStrings...))
}

func NewPolygonGeoJsonFrom(coordinates ...*LngLat) *GeoJson {
	return NewGeoJson(NewPolygonFrom(coordinates...))
}

func (x *GeoJson) ToFeatures() ([]*Feature, error) {
	var feats []*Feature
	if x != nil {
		switch gj := x.GeoJson.(type) {
		case *GeoJson_FeatureCollection:
			if gj.FeatureCollection != nil {
				feats = append(feats, gj.FeatureCollection.Features...)
			}
		case *GeoJson_Feature:
			feats = append(feats, gj.Feature)
		case *GeoJson_Point:
			feats = append(feats, NewFeature(gj.Point.ToGeometry()))
		case *GeoJson_MultiPoint:
			feats = append(feats, NewFeature(gj.MultiPoint.ToGeometry()))
		case *GeoJson_LineString:
			feats = append(feats, NewFeature(gj.LineString.ToGeometry()))
		case *GeoJson_MultiLineString:
			feats = append(feats, NewFeature(gj.MultiLineString.ToGeometry()))
		case *GeoJson_Polygon:
			feats = append(feats, NewFeature(gj.Polygon.ToGeometry()))
		case *GeoJson_MultiPolygon:
			feats = append(feats, NewFeature(gj.MultiPolygon.ToGeometry()))
		case *GeoJson_GeometryCollection:
			feats = append(feats, NewFeature(gj.GeometryCollection.ToGeometry()))
		}
	}

	return feats, nil
}

func (x *GeoJson) ToGeometries() ([]*Geometry, error) {
	var geometries []*Geometry
	if x != nil {
		switch j := x.GeoJson.(type) {
		case *GeoJson_FeatureCollection:
			if j.FeatureCollection != nil {
				for _, f := range j.FeatureCollection.Features {
					geometries = append(geometries, f.Geometry)
				}
			}
		case *GeoJson_Feature:
			geometries = append(geometries, j.Feature.Geometry)
		case *GeoJson_Point:
			geometries = append(geometries, j.Point.ToGeometry())
		case *GeoJson_MultiPoint:
			geometries = append(geometries, j.MultiPoint.ToGeometry())
		case *GeoJson_LineString:
			geometries = append(geometries, j.LineString.ToGeometry())
		case *GeoJson_MultiLineString:
			geometries = append(geometries, j.MultiLineString.ToGeometry())
		case *GeoJson_Polygon:
			geometries = append(geometries, j.Polygon.ToGeometry())
		case *GeoJson_MultiPolygon:
			geometries = append(geometries, j.MultiPolygon.ToGeometry())
		case *GeoJson_GeometryCollection:
			geometries = append(geometries, j.GeometryCollection.ToGeometry())
		}
	}
	return geometries, nil
}

func (x *GeoJson) CoordTransform(from, to SpatialReference) *GeoJson {
	if x != nil {
		switch gj := x.GeoJson.(type) {
		case *GeoJson_FeatureCollection:
			return NewGeoJson(gj.FeatureCollection.CoordTransform(from, to))
		case *GeoJson_Feature:
			return NewGeoJson(gj.Feature.CoordTransform(from, to))
		case *GeoJson_Point:
			return NewGeoJson(gj.Point.CoordTransform(from, to))
		case *GeoJson_MultiPoint:
			return NewGeoJson(gj.MultiPoint.CoordTransform(from, to))
		case *GeoJson_LineString:
			return NewGeoJson(gj.LineString.CoordTransform(from, to))
		case *GeoJson_MultiLineString:
			return NewGeoJson(gj.MultiLineString.CoordTransform(from, to))
		case *GeoJson_Polygon:
			return NewGeoJson(gj.Polygon.CoordTransform(from, to))
		case *GeoJson_MultiPolygon:
			return NewGeoJson(gj.MultiPolygon.CoordTransform(from, to))
		case *GeoJson_GeometryCollection:
			return NewGeoJson(gj.GeometryCollection.CoordTransform(from, to))
		}
	}
	return x
}
