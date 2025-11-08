package simplify

import "github.com/mojo-lang/mojo/packages/geom/go/pkg/mojo/geom/planar"

func New() planar.Simplifier {
	return &DouglasPeucker{
		Tolerance: 0.00002,
	}
}

func NewBy(tolerance float64, distance planar.PointLineDistance) planar.Simplifier {
	return &DouglasPeucker{
		Tolerance: tolerance,
		Dist:      distance,
	}
}
