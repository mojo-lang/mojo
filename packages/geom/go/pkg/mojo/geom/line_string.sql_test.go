package geom

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseLineString(t *testing.T) {
	polygon, err := parseLineString("  [  (  123.12 , 32.22  )  ,  ( 123.12,  32.12 )  ] ")
	assert.NoError(t, err)

	assert.Equal(t, 123.12, polygon[0].Longitude)
	assert.Equal(t, 32.22, polygon[0].Latitude)
}
