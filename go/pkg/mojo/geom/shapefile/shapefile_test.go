package shapefile

import (
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"testing"
)

func TestShapefile_ReadAll(t *testing.T) {
	features, err := Shapefile{}.ReadAll("./test_data/example.shp")
	assert.NoError(t, err)
	assert.NotEmpty(t, features)
}

func TestShapefile_WriteAll(t *testing.T) {
	features, err := Shapefile{}.ReadAll("./test_data/example.shp")
	assert.NoError(t, err)
	assert.NotEmpty(t, features)

	output := filepath.Join(t.TempDir(), "example_copy.shp")
	err = Shapefile{}.WriteAll(features, output)
	assert.NoError(t, err)

	newFeatures, err := Shapefile{}.ReadAll(output)
	assert.NoError(t, err)
	assert.NotEmpty(t, newFeatures)
	assert.Len(t, newFeatures, len(features))
}
