package geom

import (
	"github.com/mmcloughlin/geohash"
)

func GeoHashCenter(hash string) *LngLat {
	lat, lng := geohash.BoundingBox(hash).Center()
	return &LngLat{Longitude: lng, Latitude: lat}
}

func GeoHashVertex(hash string, k int) *LngLat {
	return GeoHashBox(hash).Vertex(k)
}

func GeoHashVertexes(hash string) []*LngLat {
	return GeoHashBox(hash).Vertexes()
}

func GeoHashBox(hash string) *BoundingBox {
	box := geohash.BoundingBox(hash)
	return &BoundingBox{LeftBottom: &LngLat{Longitude: box.MinLng, Latitude: box.MinLat},
		RightTop: &LngLat{Longitude: box.MaxLng, Latitude: box.MaxLat}}
}
