package simplify

import (
	planar2 "github.com/mojo-lang/mojo/go/pkg/mojo/geom/planar"
)

func New() planar2.Simplifier {
	return &DouglasPeucker{
		Tolerance: 0.00002,
	}
}

func NewBy(tolerance float64, distance planar2.PointLineDistance) planar2.Simplifier {
	return &DouglasPeucker{
		Tolerance: tolerance,
		Dist:      distance,
	}
}
