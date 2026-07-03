package tile

import (
	"github.com/mojo-lang/mojo/go/pkg/mojo/geom"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTileId_Polygon(t *testing.T) {
	tid := NewTileId(109767, 53549, 17)
	polygon := tid.Polygon()
	assert.NotNil(t, polygon)

	polygon = polygon.CoordTransform(geom.SpatialReference_SPATIAL_REFERENCE_WGS84, geom.SpatialReference_SPATIAL_REFERENCE_GCJ02)
	assert.NotNil(t, polygon)
}
