package geom

const (
	E7 = 10000000.0
	E5 = 100000.0
	E2 = 100.0
)

func NewBinLngLat(lng, lat int32) *BinLngLat {
	return &BinLngLat{Longitude: lng, Latitude: lat}
}

func NewBinLngLatWith(point *LngLat) *BinLngLat {
	bin := &BinLngLat{}
	bin.Encode(point)
	return bin
}

func (x *BinLngLat) Encode(lnglat *LngLat) *BinLngLat {
	if x != nil {
		x.Longitude = int32(lnglat.Longitude * E7)
		x.Latitude = int32(lnglat.Latitude * E7)
		x.Altitude = int32(lnglat.Altitude * E2)
	}

	return x
}

func (x *BinLngLat) Decode() *LngLat {
	if x != nil {
		return &LngLat{Longitude: float64(x.Longitude) / E7, Latitude: float64(x.Latitude) / E7, Altitude: float64(x.Altitude) / E2}
	}
	return nil
}

func (x *BinLngLat) CoordTransform(from, to SpatialReference) *BinLngLat {
	if x != nil {
		return NewBinLngLatWith(x.Decode().CoordTransform(from, to))
	}
	return nil
}

func DecodeBinLngLat(coordinates ...int32) *LngLat {
	if len(coordinates) == 2 {
		bin := &BinLngLat{Longitude: coordinates[0], Latitude: coordinates[1]}
		return bin.Decode()
	} else if len(coordinates) == 3 {
		bin := &BinLngLat{Longitude: coordinates[0], Latitude: coordinates[1], Altitude: coordinates[2]}
		return bin.Decode()
	}
	return nil
}

func EncodeBinLngLat(point *LngLat) []int32 {
	lnglat := NewBinLngLatWith(point)
	return []int32{lnglat.Longitude, lnglat.Latitude, lnglat.Altitude}
}
