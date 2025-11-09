package geom

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"testing"
)

const JSON = `{"type":"Polygon","coordinates":[[[121.012886,30.909123],[121.013049,30.909658],[121.013334,30.910787],[121.015296,30.91049],[121.014918,30.909632],[121.014864,30.909495],[121.014739,30.9088],[121.012886,30.909123]]]}`

const JSON2 = `{"type":"Polygon","coordinates":[[[121.677339,31.141503],[121.678937,31.142036],[121.680246,31.142421],[121.681137,31.142651],[121.681351,31.142159],[121.681839,31.140915],[121.682488,31.139333],[121.683092,31.137811],[121.683336,31.137201],[121.683379,31.137049],[121.683379,31.136999],[121.683379,31.136944],[121.683341,31.136879],[121.68097,31.136034],[121.680788,31.135952],[121.680611,31.135883],[121.68008,31.135731],[121.678583,31.135199],[121.67743,31.13479],[121.676346,31.134386],[121.67625,31.134469],[121.675617,31.13609],[121.675659,31.136158],[121.675702,31.1362],[121.676008,31.136315],[121.67611,31.136406],[121.676459,31.136342],[121.677489,31.13625],[121.677585,31.13625],[121.677988,31.136356],[121.678615,31.136558],[121.679908,31.137026],[121.680267,31.137237],[121.680477,31.137361],[121.680589,31.137439],[121.680632,31.137531],[121.680568,31.137664],[121.680525,31.137761],[121.680471,31.138215],[121.680284,31.138496],[121.679876,31.138698],[121.679645,31.138776],[121.679302,31.138817],[121.679018,31.138711],[121.678739,31.138617],[121.678591,31.138603],[121.678444,31.138592],[121.678261,31.138578],[121.677805,31.13872],[121.6778,31.139267],[121.677596,31.139928],[121.677285,31.140534],[121.677151,31.140819],[121.677086,31.141058],[121.677097,31.141099],[121.677253,31.141305],[121.677339,31.141503]]]}`

func PolygonJson(p *Polygon) string {
	hashes := p.GeoHashCovering(7, 7, 100)

	mp := MultiPolygon{}
	mp.Polygons = append(mp.Polygons, p)

	for _, hash := range hashes {
		mp.Polygons = append(mp.Polygons, NewPolygonFromGeoHash(hash))
	}

	geojson, _ := jsoniter.ConfigFastest.MarshalToString(mp)
	return geojson
}

func TestCircle_GeoHashCovering(t *testing.T) {
	circle := &Circle{Center: &LngLat{Longitude: 121.482885, Latitude: 31.236148}, Radius: 75}

	hashes := circle.GeoHashFractionCovering(7, 7, 10000, 0.5)
	fmt.Print(hashes)

	mp := MultiPolygon{}
	mp.Polygons = append(mp.Polygons, circle.ToPolygon(64))

	for _, hash := range hashes {
		mp.Polygons = append(mp.Polygons, NewPolygonFromGeoHash(hash))
	}

	geojson, _ := jsoniter.ConfigFastest.MarshalToString(mp)
	fmt.Print(geojson)

	poly := &Polygon{}
	jsoniter.ConfigFastest.Unmarshal([]byte(JSON), poly)

	poly.GeoHashCovering(7, 7, 100)

	jsoniter.ConfigFastest.Unmarshal([]byte(JSON2), poly)
	fmt.Print(PolygonJson(poly))
}
