package orb

import (
	"github.com/mojo-lang/mojo/packages/core/go/pkg/mojo/core"
	"github.com/mojo-lang/mojo/packages/geom/go/pkg/mojo/geom"
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/geojson"
)

func FeatureCollectionFrom(collection *geojson.FeatureCollection) *geom.FeatureCollection {
	if collection != nil {
		fc := FeatureCollectionFromFeatures(collection.Features)
		if fc == nil {
			fc = &geom.FeatureCollection{}
		}

		fc.Type = collection.Type
		return fc
	}
	return nil
}

func FeatureCollectionFromFeatures(features []*geojson.Feature) *geom.FeatureCollection {
	if len(features) > 0 {
		fc := &geom.FeatureCollection{}
		for _, feat := range features {
			fc.Features = append(fc.Features, FeatureFrom(feat))
		}
		return fc
	}
	return nil
}

func FeatureCollectionTo(collection *geom.FeatureCollection) *geojson.FeatureCollection {
	if collection != nil {
		fc := geojson.NewFeatureCollection()
		for _, feat := range collection.Features {
			fc.Features = append(fc.Features, FeatureTo(feat))
		}
	}
	return nil
}

func FeatureFrom(feature *geojson.Feature) *geom.Feature {
	if feature != nil {
		feat := geom.NewFeature(nil)
		if feature.ID != nil {
			switch id := feature.ID.(type) {
			case float32:
				feat.Id = core.NewIntId(uint64(id))
			case float64:
				feat.Id = core.NewIntId(uint64(id))
			case int:
				feat.Id = core.NewIntId(uint64(id))
			case int32:
				feat.Id = core.NewIntId(uint64(id))
			case int64:
				feat.Id = core.NewIntId(uint64(id))
			case string:
				feat.Id = core.NewStringId(id)
			}
		}

		feat.Bbox = BoundingBoxFrom(feature.BBox)
		feat.Geometry = GeometryFrom(feature.Geometry)
		for k, v := range feature.Properties {
			switch value := v.(type) {
			case bool:
				feat.Properties[k] = core.NewBoolValue(value)
			case int:
				feat.Properties[k] = core.NewInt32Value(int32(value))
			case int32:
				feat.Properties[k] = core.NewInt32Value(value)
			case int64:
				feat.Properties[k] = core.NewInt64Value(value)
			case float32:
				feat.Properties[k] = core.NewFloat32Value(value)
			case float64:
				feat.Properties[k] = core.NewFloat64Value(value)
			case string:
				feat.Properties[k] = core.NewStringValue(value)
			}
		}
		return feat
	}

	return nil
}

func FeatureTo(feature *geom.Feature) *geojson.Feature {
	if feature != nil {
		feat := geojson.NewFeature(nil)
		if feature.Id != nil {
			switch id := feature.Id.Id.(type) {
			case *core.Id_StringVal:
				feat.ID = id.StringVal
			case *core.Id_Uint64Val:
				feat.ID = id.Uint64Val
			case *core.Id_UuidVal:
				feat.ID = id.UuidVal.Format()
			}
		}
		if feature.Bbox != nil {
			feat.BBox = BoundingBoxTo(feature.Bbox)
		}
		feat.Geometry = GeometryTo(feature.Geometry)
		for k, v := range feature.Properties {
			switch v.GetKind() {
			case core.ValueKind_VALUE_KIND_INTEGER:
				feat.Properties[k] = v.GetInt64()
			case core.ValueKind_VALUE_KIND_NUMBER:
				feat.Properties[k] = v.GetFloat64()
			case core.ValueKind_VALUE_KIND_BOOLEAN:
				feat.Properties[k] = v.GetBool()
			case core.ValueKind_VALUE_KIND_STRING:
				feat.Properties[k] = v.GetString()
			}
		}
	}
	return nil
}

func GeometryFrom(geometry orb.Geometry) *geom.Geometry {
	newPoints := func(points []orb.Point) []*geom.LngLat {
		var ps []*geom.LngLat
		for _, p := range points {
			ps = append(ps, geom.NewLngLat(p[0], p[1]))
		}
		return ps
	}

	switch g := geometry.(type) {
	case orb.Point:
		return geom.NewPointGeometryFrom(g[0], g[1])
	case orb.MultiPoint:
		return geom.NewMultiPointGeometryFrom(newPoints(g)...)
	case orb.LineString:
		return geom.NewLineStringGeometry(newPoints(g)...)
	case orb.MultiLineString:
		return geom.NewMultiLineStringGeometry(newLines(g)...)
	case orb.Ring:
		return geom.NewLineStringGeometry(newPoints(g)...)
	case orb.Polygon:
		return geom.NewPolygonGeometry(newLines(g)...)
	case orb.MultiPolygon:
		var ps []*geom.Polygon
		for _, polygon := range g {
			ps = append(ps, GeometryFrom(polygon).GetPolygon())
		}
		return geom.NewMultiPolygonGeometry(ps...)
	case orb.Collection:
		var collection []*geom.Geometry
		for _, gm := range g {
			collection = append(collection, GeometryFrom(gm))
		}
		return geom.NewGeometryCollectionGeometry(collection...)
	}
	return nil
}

func newLines[T orb.Geometry](lines []T) []*geom.LineString {
	var ls []*geom.LineString
	for _, line := range lines {
		ls = append(ls, GeometryFrom(line).GetLineString())
	}
	return ls
}

func newOrbLines[T orb.Geometry](lines []*geom.LineString) []T {
	var ls []T
	for _, line := range lines {
		ls = append(ls, GeometryTo(geom.NewGeometry(line)).(T))
	}
	return ls
}

func GeometryTo(geometry *geom.Geometry) orb.Geometry {
	newPoint := func(point *geom.LngLat) orb.Point {
		return orb.Point{point.GetLongitude(), point.GetLatitude()}
	}

	newPoints := func(points []*geom.LngLat) []orb.Point {
		var ps []orb.Point
		for _, p := range points {
			ps = append(ps, newPoint(p))
		}
		return ps
	}

	if geometry != nil {
		switch g := geometry.Geometry.(type) {
		case *geom.Geometry_Point:
			return newPoint(g.Point.Coordinate)
		case *geom.Geometry_LineString:
			return orb.LineString(newPoints(g.LineString.Coordinates))
		case *geom.Geometry_Polygon:
			return orb.Polygon(newOrbLines[orb.Ring](g.Polygon.LineStrings))
		case *geom.Geometry_MultiPoint:
			ps := orb.MultiPoint{}
			for _, point := range g.MultiPoint.Points {
				ps = append(ps, newPoint(point.Coordinate))
			}
			return ps
		case *geom.Geometry_MultiLineString:
			return orb.MultiLineString(newOrbLines[orb.LineString](g.MultiLineString.LineStrings))
		case *geom.Geometry_MultiPolygon:
			mp := orb.MultiPolygon{}
			for _, polygon := range g.MultiPolygon.Polygons {
				mp = append(mp, newOrbLines[orb.Ring](polygon.LineStrings))
			}
			return mp
		case *geom.Geometry_GeometryCollection:
			collection := orb.Collection{}
			for _, gm := range g.GeometryCollection.Geometries {
				collection = append(collection, GeometryTo(gm))
			}
			return collection
		}
	}
	return nil
}

func BoundingBoxFrom(bbx geojson.BBox) *geom.BoundingBox {
	if bbx != nil && len(bbx) == 4 {
		return geom.NewBoundingBox(geom.NewLngLat(bbx[0], bbx[1]), geom.NewLngLat(bbx[2], bbx[3]))
	}
	return nil
}

func BoundingBoxTo(bbx *geom.BoundingBox) geojson.BBox {
	if bbx != nil && bbx.LeftBottom != nil && bbx.RightTop != nil {
		return geojson.BBox{bbx.LeftBottom.Longitude, bbx.LeftBottom.Latitude, bbx.RightTop.Longitude, bbx.RightTop.Latitude}
	}
	return nil
}
