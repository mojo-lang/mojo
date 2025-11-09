package planar

import "context"

// Simplifier is an interface for Simplifying geometries.
type Simplifier interface {
	Simplify(ctx context.Context, linestring [][2]float64, isClosed bool) ([][2]float64, error)
}
