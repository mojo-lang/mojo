package tile

import (
	"bytes"
	"compress/gzip"
	"github.com/mojo-lang/mojo/packages/geom/go/pkg/mojo/geom"
	"github.com/mojo-lang/mojo/packages/geom/go/pkg/mojo/geom/converts/orb"
	"github.com/paulmach/orb/encoding/mvt"
	"github.com/paulmach/orb/geojson"
	"github.com/paulmach/orb/maptile"
	"github.com/paulmach/orb/simplify"
)

func UnmarshalGzippedMvt(data []byte) (map[string]*geom.FeatureCollection, error) {
	collection, err := mvt.UnmarshalGzipped(data)
	if err != nil {
		return nil, err
	}

	cs := make(map[string]*geom.FeatureCollection)
	for _, layer := range collection {
		cs[layer.Name] = orb.FeatureCollectionFromFeatures(layer.Features)
	}
	return cs, nil
}

func UnmarshalMvt(data []byte) (map[string]*geom.FeatureCollection, error) {
	collection, err := mvt.Unmarshal(data)
	if err != nil {
		return nil, err
	}

	cs := make(map[string]*geom.FeatureCollection)
	for _, layer := range collection {
		cs[layer.Name] = orb.FeatureCollectionFromFeatures(layer.Features)
	}
	return cs, nil
}

func MarshalMvt(layers map[string]*geom.FeatureCollection, tid *TileId) ([]byte, error) {
	collections := map[string]*geojson.FeatureCollection{}
	for k, v := range layers {
		collections[k] = orb.FeatureCollectionTo(v)
	}

	// Convert to a layers object and project to tile coordinates.
	lyrs := mvt.NewLayers(collections)
	lyrs.ProjectToTile(maptile.New(uint32(tid.X), uint32(tid.Y), maptile.Zoom(tid.Level)))

	// In order to be used as source for MapboxGL geometries need to be clipped
	// to max allowed extent. (uncomment next line)
	lyrs.Clip(mvt.MapboxGLDefaultExtentBound)

	// Simplify the geometry now that it's in tile coordinate space.
	lyrs.Simplify(simplify.DouglasPeucker(1.0))

	// Depending on use-case remove empty geometry, those too small to be
	// represented in this tile space.
	// In this case lines shorter than 1, and areas smaller than 2.
	lyrs.RemoveEmpty(1.0, 2.0)

	// encoding using the Mapbox Vector Tile protobuf encoding.
	return mvt.Marshal(lyrs) // this data is NOT gzipped.
}

func MarshalGzippedMvt(layers map[string]*geom.FeatureCollection, tid *TileId) ([]byte, error) {
	data, err := MarshalMvt(layers, tid)
	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(nil)
	gzwriter := gzip.NewWriter(buf)

	_, err = gzwriter.Write(data)
	if err != nil {
		return nil, err
	}

	err = gzwriter.Close()
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
