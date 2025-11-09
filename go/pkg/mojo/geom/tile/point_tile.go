package tile

import (
	geom2 "github.com/mojo-lang/mojo/go/pkg/mojo/geom"
)

func (x *PointTile) PointCount() int {
	return len(x.Xs)
}

func (x *PointTile) lng(index int) float64 {
	return float64(x.Xs[index]) / geom2.E7
}

func (x *PointTile) lat(index int) float64 {
	return float64(x.Ys[index]) / geom2.E7
}

func (x *PointTile) BinLngLat(index int) *geom2.BinLngLat {
	if index >= 0 && index < len(x.Xs) {
		return &geom2.BinLngLat{Longitude: int32(x.Xs[index]), Latitude: int32(x.Ys[index])}
	}
	return nil
}

func (x *PointTile) LngLat(index int) *geom2.LngLat {
	if index >= 0 && index < len(x.Xs) {
		return &geom2.LngLat{Longitude: x.lng(index), Latitude: x.lat(index)}
	}
	return nil
}

func (x *PointTile) Point(index int) *geom2.Point {
	if ll := x.LngLat(index); ll != nil {
		return geom2.NewPoint(ll)
	}
	return nil
}

func (x *PointTile) BinLngLats() []*geom2.BinLngLat {
	var lnglats []*geom2.BinLngLat
	for i := 0; i < len(x.Xs); i++ {
		lnglats = append(lnglats, &geom2.BinLngLat{Longitude: int32(x.Xs[i]), Latitude: int32(x.Ys[i])})
	}
	return lnglats
}

func (x *PointTile) PropertyAt(index int) map[string]*VectorTile_Value {
	if index >= 0 && index < len(x.Tags)/2 {
	}
	return nil
}

//func (x *PointTile) ToMVT(layerName string, compress bool) ([]byte, error) {
//	var buf *geom.GeoJson
//	npoints := len(x.Xs)
//	if len(x.Xs) == len(x.Ids) {
//		features := make([]*geom.Feature, 0, npoints)
//		for i := 0; i < npoints; i++ {
//			point := x.LngLat(i)
//			features = append(features, &geom.Feature{
//				Id:         core.NewIntId(x.Ids[i]),
//				Geometry:   geom.NewPointGeometry(point),
//				Properties: map[string]*core.Value{DefaultLayerTag: core.NewStringValue(layerName)},
//			})
//		}
//		buf = geom.NewFeatureCollectionGeoJson(features...)
//	} else {
//		points := make([]*geom.Point, 0, npoints)
//		for i := 0; i < npoints; i++ {
//			points = append(points, x.Point(i))
//		}
//
//		feature := &geom.Feature{
//			Geometry:   geom.NewMultiPointGeometry(points...),
//			Properties: map[string]*core.Value{DefaultLayerTag: core.NewStringValue(layerName)},
//		}
//		buf = geom.NewGeoJson(feature)
//	}
//
//	vt := &VectorTile{}
//	vt.EncodeFrom(buf, NewTileIdFromQuadKey(x.Id))
//	return vt.MarshalCompressed(compress)
//}

func (x *PointTile) SetBinLngLats(lnglats []*geom2.BinLngLat) {
}

func (x *PointTile) AddLngLats(lngLats ...*geom2.LngLat) {
}
