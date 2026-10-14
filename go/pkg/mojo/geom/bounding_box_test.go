package geom

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBox_Extend(t *testing.T) {
	box1 := BoundingBox{LeftBottom: &LngLat{Longitude: 121.1234, Latitude: 31.11},
		RightTop: &LngLat{Longitude: 122.1234, Latitude: 32.11}}

	box2 := BoundingBox{LeftBottom: &LngLat{Longitude: 121.1238, Latitude: 31.22},
		RightTop: &LngLat{Longitude: 123.1234, Latitude: 33.11}}

	extend := box1.Extend(&box2)
	assert.Equal(t, 121.1234, extend.LeftBottom.Longitude)
	assert.Equal(t, 123.1234, extend.RightTop.Longitude)
}

func TestBoundingBoxParsingRejectsInvalidInput(t *testing.T) {
	for _, input := range []string{"", "garbage", "[]", "[1,2,3]", "[[1,2],[3]]"} {
		box, err := NewBoundingBoxFromString(input)
		require.Error(t, err, input)
		require.Nil(t, box)
	}
	box, err := NewBoundingBoxFromString("[3,4,1,2]")
	require.NoError(t, err)
	require.True(t, box.LeftBottom.Equal(NewLngLat(1, 2)))
	require.True(t, box.RightTop.Equal(NewLngLat(3, 4)))
	require.True(t, box.Center().Equal(NewLngLat(2, 3)))
	require.Len(t, box.Vertexes(), 4)
	require.Nil(t, box.Vertex(-1))
	require.Nil(t, box.Vertex(4))
	require.True(t, box.ToPolygon().LineStrings[0].IsClosed())
}

func TestBox_GeoHashCovering(t *testing.T) {
	box := BoundingBox{LeftBottom: &LngLat{Longitude: 121.47789, Latitude: 31.223224},
		RightTop: &LngLat{Longitude: 121.491537, Latitude: 31.233279}}

	hashes := box.GeoHashCovering(5, 5, 100)
	print(hashes)
}

func TestBoundingBox_ExtendToPoint(t *testing.T) {
	box := &BoundingBox{}

	box.ExtendToPoint(NewLngLat(121.1234, 31.11))
	assert.EqualValues(t, 121.1234, box.LeftBottom.Longitude)
	assert.EqualValues(t, 121.1234, box.RightTop.Longitude)

	box.ExtendToPoint(NewLngLat(122.1234, 32.11))
	assert.EqualValues(t, 122.1234, box.RightTop.Longitude)
	assert.EqualValues(t, 121.1234, box.LeftBottom.Longitude)
}

func TestNewBoundingBoxFromString(t *testing.T) {
	box := BoundingBox{LeftBottom: &LngLat{Longitude: 121.1234, Latitude: 31.11},
		RightTop: &LngLat{Longitude: 122.1234, Latitude: 32.11}}

	b1, err := NewBoundingBoxFromString("[[121.1234, 31.11],[122.1234, 32.11]]")
	assert.NoError(t, err)
	assert.True(t, box.Equal(b1))

	b2, err := NewBoundingBoxFromString("[121.1234, 31.11,122.1234, 32.11]")
	assert.NoError(t, err)
	assert.True(t, box.Equal(b2))

	b3, err := NewBoundingBoxFromString("121.1234, 31.11,122.1234, 32.11")
	assert.NoError(t, err)
	assert.True(t, box.Equal(b3))
}
