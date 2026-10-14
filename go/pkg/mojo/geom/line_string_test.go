package geom

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLinearRingClosureAndOrientation(t *testing.T) {
	require.Nil(t, NewLinearRing())
	require.Nil(t, NewLinearRing(NewLngLat(0, 0)))
	ring := NewLinearRing(NewLngLat(0, 0), NewLngLat(2, 0), NewLngLat(2, 2), NewLngLat(0, 2))
	require.Len(t, ring.Coordinates, 5)
	require.True(t, ring.IsClosed())
	orientation := ring.Clockwise()
	require.NotZero(t, orientation)
	ring.Invert()
	require.Equal(t, -orientation, ring.Clockwise())
	require.True(t, ring.IsClosed())
	closed := NewLinearRing(ring.Coordinates...)
	require.Len(t, closed.Coordinates, 5)
	rectangle := NewLinearRing(NewLngLat(3, 4), NewLngLat(1, 2))
	require.Len(t, rectangle.Coordinates, 5)
	require.True(t, rectangle.IsClosed())
	require.True(t, NewBoundingBox(NewLngLat(1, 2), NewLngLat(3, 4)).Equal(rectangle.BoundingBox()))
}

func TestLineStringEqualityBoundsAndNil(t *testing.T) {
	a, b := NewLngLat(-10, 20), NewLngLat(30, -5)
	line := NewLineString(a, b)
	require.True(t, line.Equal(NewLineString(a, b)))
	require.False(t, line.Equal(NewLineString(b, a)))
	require.False(t, line.Equal(NewLineString(a)))
	require.False(t, line.Equal(nil))
	box := line.BoundingBox()
	require.True(t, box.LeftBottom.Equal(NewLngLat(-10, -5)))
	require.True(t, box.RightTop.Equal(NewLngLat(30, 20)))
	require.True(t, a.Equal(NewLngLat(-10, 20)))
	require.Same(t, line, line.ToGeometry().GetLineString())
	line.Add(NewLngLat(50, 40))
	require.Len(t, line.Coordinates, 3)
	var empty *LineString
	require.True(t, empty.Equal(nil))
	require.False(t, empty.IsClosed())
	require.Zero(t, empty.Clockwise())
	require.Nil(t, empty.BoundingBox())
	require.Nil(t, empty.ToGeometry())
	require.NotPanics(t, func() { empty.Invert() })
	require.Nil(t, NewLineString().BoundingBox())
}
