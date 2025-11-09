package geom

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	emptyGeometryCollection = `{"type":"GeometryCollection","geometries":[]}`
)

func TestGeometryCollectionCodec_Encode(t *testing.T) {
	gc := &GeometryCollection{}

	str, _ := jsoniter.MarshalToString(gc)
	assert.Equal(t, emptyGeometryCollection, str)
}
