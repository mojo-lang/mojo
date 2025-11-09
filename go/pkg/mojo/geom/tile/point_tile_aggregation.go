package tile

import (
	"errors"
	geom2 "github.com/mojo-lang/mojo/go/pkg/mojo/geom"

	"math"

	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
)

var (
	DefaultSamplingRate = math.Sqrt(math.Ln2)
)

func AggregatePointTiles(subTiles []*PointTile, quadKey string, options core.Options) (*PointTile, error) {
	if len(subTiles) == 0 {
		return nil, errors.New("no sub tiles")
	}

	xs := make([]int32, 0, 8*1024)
	ys := make([]int32, 0, 8*1024)

	var (
		minPoints  = 1024
		extent     = float64(1024)
		sampleRate = DefaultSamplingRate
	)
	if v, ok := options["min_point_count"]; ok {
		minPoints = v.(int)
	}

	if v, ok := options["extent"]; ok {
		extent = v.(float64)
	}

	if v, ok := options["sampling_rate"]; ok {
		sampleRate = v.(float64)
	}

	zxy := NewTileIdFromQuadKey(quadKey)

	for _, sub := range subTiles {
		cnt := len(sub.Xs)
		if cnt <= minPoints {
			xs = append(xs, sub.Xs...)
			ys = append(ys, sub.Ys...)
			continue
		}
		for i := 0; i < cnt; i++ {
			x, y := sub.Xs[i], sub.Ys[i]
			lon, lat := float64(x)/geom2.E7, float64(y)/geom2.E7
			u, v := zxy.LngLat2XY(geom2.NewLngLat(lon, lat))
			if sampled(extent, float64(u)/DefaultTilePixels, float64(v)/DefaultTilePixels, sampleRate) {
				xs = append(xs, x)
				ys = append(ys, y)
			}
		}
	}

	return &PointTile{
		Id:   quadKey,
		Type: subTiles[0].Type,
		Xs:   xs,
		Ys:   ys,
	}, nil
}

// sample rate: (0, 1]
func sampled(extent float64, u, v float64, sampleRate float64) bool {
	halfSR := 0.5 * sampleRate
	u = u * extent
	v = v * extent
	if math.Abs(u-math.Round(u)) <= halfSR &&
		math.Abs(v-math.Round(v)) <= halfSR {
		return true
	}
	return false
}
