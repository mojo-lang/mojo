package xlsx

import (
	"github.com/mojo-lang/mojo/go/pkg/mojo/geom/shapefile"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFile_WriteAll(t *testing.T) {
	features, err := shapefile.Shapefile{}.ReadAll("../shapefile/test_data/example.shp")
	assert.NoError(t, err)
	assert.NotEmpty(t, features)

	err = File{}.WriteAll(features, "./test_data/example_copy.xlsx")
	assert.NoError(t, err)
}
