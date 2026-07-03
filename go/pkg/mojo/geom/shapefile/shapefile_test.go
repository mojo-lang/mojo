package shapefile

import (
	"github.com/stretchr/testify/assert"
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

	err = Shapefile{}.WriteAll(features, "./test_data/example_copy.shp")
	assert.NoError(t, err)

	newFeatures, err := Shapefile{}.ReadAll("./test_data/example_copy.shp")
	assert.NoError(t, err)
	assert.NotEmpty(t, newFeatures)
}
