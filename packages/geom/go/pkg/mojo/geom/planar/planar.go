package planar

import "math"

// come from https://github.com/go-spatial/geom/blob/master/planar/planar.go

// PointLineDistance is the abstract method to get the distance from point
// to a line depending on projection
type PointLineDistance func(line [2][2]float64, point [2]float64) float64

// PerpendicularDistance  provides the distance between a line and a point in Euclidean space.
// ref: https://en.wikipedia.org/wiki/Distance_from_a_point_to_a_line#Line_defined_by_two_points
func PerpendicularDistance(line [2][2]float64, point [2]float64) float64 {

	deltaX := line[1][0] - line[0][0]
	deltaY := line[1][1] - line[0][1]
	deltaXSq := deltaX * deltaX
	deltaYSq := deltaY * deltaY

	num := math.Abs((deltaY * point[0]) - (deltaX * point[1]) + (line[1][0] * line[0][1]) - (line[1][1] * line[0][0]))
	denom := math.Sqrt(deltaYSq + deltaXSq)
	if denom == 0 {
		return 0
	}
	return num / denom
}
