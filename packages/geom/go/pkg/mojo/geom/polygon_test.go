package geom

import (
	"encoding/json"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

const (
	emptyPolygon    = `{"type":"Polygon","coordinates":[]}`
	polygonStr      = `{"type":"Polygon","coordinates":[[[113.167297,23.522705],[113.166503,23.515686],[113.166503,23.515686]],[[113.167297,23.522705],[113.167297,23.522705],[113.166503,23.515686],[113.166503,23.515686]]]}`
	polygonWithHole = `{"type":"Polygon","coordinates":[[[121.768212,31.029051],[121.767461,31.023057],[121.775722,31.022487],[121.777288,31.027544]],[[121.769006,31.028206],[121.768383,31.023958],[121.772289,31.023774],[121.772718,31.027801]]]}`
)

var (
	pointInPolygonWithHole = &LngLat{Longitude: 121.773898, Latitude: 31.02701}
	pointInHole            = &LngLat{Longitude: 121.769435, Latitude: 31.025834}
)

func TestNewPolygon(t *testing.T) {
	lls := make([]*LngLat, 0, 4)
	lls = append(lls, &LngLat{
		Longitude: 120,
		Latitude:  30,
	})
	lls = append(lls, &LngLat{
		Longitude: 121,
		Latitude:  31,
	})
	polygon := NewPolygonFrom(lls...)
	s, _ := jsoniter.MarshalToString(polygon)

	fmt.Println("s:", s)
	fmt.Println("area:", EarthArea(polygon.Area()).Km2())
}

func TestPolygonArea(t *testing.T) {
	polygon := NewPolygonFromGeoHash("wtkzm")
	fmt.Println("area:", EarthArea(polygon.Area()).Km2())
}

func TestPointInPolygon(t *testing.T) {
	//brunei, err := polygonFromFile("test/data/brunei.json")
	//if err != nil {
	//	t.Error("brunei json file failed to parse: ", err)
	//}
	//
	//point := LngLat{Longitude: 114.9480600, Latitude: 4.9402900}
	//if !brunei.Contains(&point) {
	//	t.Error("Expected the capital of Brunei to be in Brunei, but it wasn't.")
	//}
}

// Ensures that the polygon logic can correctly identify if a polygon does not contain a point.
// Uses Brunei, Seattle, and a point directly outside of Brunei limits as test points.
func TestPointNotInPolygon(t *testing.T) {
	//brunei, err := polygonFromFile("test/data/brunei.json")
	//if err != nil {
	//	t.Error("brunei json file failed to parse: ", err)
	//}
	//
	//// Seattle, WA should not be inside of Brunei
	//point := NewLngLat(122.30, 47.45)
	//if brunei.Contains(point) {
	//	t.Error("Seattle, WA [47.45, 122.30] should not be inside of Brunei")
	//}
	//
	//// A point just outside of the successful bounds in Brunei
	//// Should not be contained in the Polygon
	//precision := NewLngLat(4.007636, 114.659596)
	//if brunei.Contains(precision) {
	//	t.Error("A point just outside of Brunei should not be contained in the Polygon")
	//}
}

// Ensures that a point can be contained in a complex polygon (e.g. a donut)
// This particular Polygon has a hole in it.
func TestPointInPolygonWithHole(t *testing.T) {
	//nsw, err := polygonFromFile("test/data/nsw.json")
	//if err != nil {
	//	t.Error("nsw json file failed to parse: ", err)
	//}
	//
	//act, err := polygonFromFile("test/data/act.json")
	//if err != nil {
	//	t.Error("act json file failed to parse: ", err)
	//}
	//
	//// Look at two contours
	//canberra := LngLat{Longitude: 149.128684300000030000, Latitude: -35.2819998}
	//isnsw := nsw.Contains(&canberra)
	//isact := act.Contains(&canberra)
	//if !isnsw && !isact {
	//	t.Error("Canberra should be in NSW and also in the sub-contour ACT state")
	//}
	//
	//// Using NSW as a multi-contour polygon
	//nswmulti := &Polygon{}
	//for _, p := range nsw.LineStrings[0].Coordinates {
	//	nswmulti.Add(p)
	//}
	//
	//for _, p := range act.LineStrings[0].Coordinates {
	//	nswmulti.Add(p)
	//}
	//
	//isnsw = nswmulti.Contains(&canberra)
	//if isnsw {
	//	t.Error("Canberra should not be in NSW as it falls in the donut contour of the ACT")
	//}
	//
	//sydney := LngLat{Longitude: 151.209, Latitude: -33.866}
	//
	//if !nswmulti.Contains(&sydney) {
	//	t.Error("Sydney should be in NSW")
	//}
	//
	//losangeles := LngLat{Longitude: 118.28333, Latitude: 34.01667}
	//isnsw = nswmulti.Contains(&losangeles)
	//
	//if isnsw {
	//	t.Error("Los Angeles should not be in NSW")
	//}
}

func TestPointNotInPolygonWithHole(t *testing.T) {
	p := &Polygon{}
	err := jsoniter.UnmarshalFromString(polygonWithHole, p)
	assert.NoError(t, err)

	in := p.Contains(pointInPolygonWithHole)
	assert.Equal(t, true, in)
	in = p.Contains(pointInHole)
	assert.Equal(t, false, in)
}

// Ensures that jumping over the equator and the greenwich meridian
// Doesn't give us any false positives or false negatives
func TestEquatorGreenwichContains(t *testing.T) {
	//point1 := NewLngLat(0.0, 0.0)
	//point2 := NewLngLat(0.1, 0.1)
	//point3 := NewLngLat(-0.1, 0.1)
	//point4 := NewLngLat(-0.1, -0.1)
	//point5 := NewLngLat(0.1, -0.1)
	//polygon, err := polygonFromFile("test/data/equator_greenwich.json")
	//
	//if err != nil {
	//	t.Error("error parsing polygon", err)
	//}
	//
	//if !polygon.Contains(point1) {
	//	t.Error("Should contain middle point of earth")
	//}
	//
	//if !polygon.Contains(point2) {
	//	t.Errorf("Should contain point %v", point2)
	//}
	//
	//if !polygon.Contains(point3) {
	//	t.Errorf("Should contain point %v", point3)
	//}
	//
	//if !polygon.Contains(point4) {
	//	t.Errorf("Should contain point %v", point4)
	//}
	//
	//if !polygon.Contains(point5) {
	//	t.Errorf("Should contain point %v", point5)
	//}
}

func TestPolygonCodec_EncodeEmpty(t *testing.T) {
	p := &Polygon{}

	str, _ := jsoniter.ConfigFastest.MarshalToString(p)
	assert.Equal(t, emptyPolygon, str)
}

// A test struct used to encapsulate and
// Unmarshal JSON into.
type testPoints struct {
	Points []*LngLat
}

// Opens a JSON file and unmarshal the data into a Polygon
func polygonFromFile(filename string) (*Polygon, error) {
	p := &Polygon{}
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	points := new(testPoints)
	jsonParser := json.NewDecoder(file)
	if err = jsonParser.Decode(&points); err != nil {
		return nil, err
	}

	for _, point := range points.Points {
		p.Add(point)
	}

	return p, nil
}
