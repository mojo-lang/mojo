package geom

import (
	"fmt"
	"math"
)

func (x *Circle) ToString() string {
	if x != nil && x.Center != nil {
		return fmt.Sprintf("%f,%f,%f", x.Center.Longitude, x.Center.Latitude, x.Radius)
	}
	return ""
}

func (x *Circle) ToPolygon(numberOfSegments int) *Polygon {
	if numberOfSegments == 0 {
		numberOfSegments = 32
	}

	lineString := &LineString{}
	lineString.Coordinates = make([]*LngLat, numberOfSegments)
	for i := 0; i < numberOfSegments; i++ {
		lineString.Coordinates[i] = x.destination(2 * math.Pi * float64(-i) / float64(numberOfSegments))
	}
	lineString.Coordinates = append(lineString.Coordinates, lineString.Coordinates[0])

	polygon := &Polygon{}
	polygon.LineStrings = []*LineString{lineString}
	return polygon
}

func (x *Circle) destination(bearing float64) *LngLat {
	lat := toRadians(x.Center.Latitude)
	lng := toRadians(x.Center.Longitude)
	var dByR = float64(x.Radius / 6378137.0) // distance divided by 6378137 (radius of the earth) wgs84

	lat = math.Asin(math.Sin(lat)*math.Cos(dByR) + math.Cos(lat)*math.Sin(dByR)*math.Cos(bearing))
	lng = lng + math.Atan2(math.Sin(bearing)*math.Sin(dByR)*math.Cos(lat), math.Cos(dByR)-math.Sin(lat)*math.Sin(lat))
	return &LngLat{Longitude: toDegrees(lng), Latitude: toDegrees(lat)}
}

func (x *Circle) CoordTransform(from, to SpatialReference) *Circle {
	if x != nil {
		return &Circle{
			Type:   x.Type,
			Center: x.Center.CoordTransform(from, to),
			Radius: x.Radius,
		}
	}
	return x
}

func toRadians(angleInDegrees float64) float64 {
	return angleInDegrees * math.Pi / 180
}

func toDegrees(angleInRadians float64) float64 {
	return angleInRadians * 180 / math.Pi
}
