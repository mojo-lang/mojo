package geom

import "github.com/mmcloughlin/geohash"

func (x *BoundingBox) GeoHashCovering(minLevel, maxLevel, maxHashes int) []string {
	hash := geohash.EncodeWithPrecision(x.LeftBottom.Latitude, x.LeftBottom.Longitude, uint(maxLevel))
	box := geohash.BoundingBox(hash)
	width := box.MaxLng - box.MinLng
	height := box.MaxLat - box.MinLat

	hashes := make([]string, 0)
	for lat := x.LeftBottom.Latitude; lat <= x.RightTop.Latitude; lat += height {
		for lng := x.LeftBottom.Longitude; lng <= x.RightTop.Longitude; lng += width {
			hashes = append(hashes, geohash.EncodeWithPrecision(lat, lng, uint(maxLevel)))
		}
	}

	// TODO
	// ensure have the borders covered
	//for (double lat = bottomRightLat; lat <= topLeftLat; lat += actualHeightDegreesPerHash) {
	//	hashes.add(encodeHashToLong(lat, maxLon, length));
	//}
	//for (double lon = topLeftLon; lon <= maxLon; lon += actualWidthDegreesPerHash) {
	//	hashes.add(encodeHashToLong(topLeftLat, lon, length));
	//}
	// ensure that the topRight corner is covered
	//hashes.add(encodeHashToLong(topLeftLat, maxLon, length));

	return hashes
}
