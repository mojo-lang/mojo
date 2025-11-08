package core

import "math"

const (
	Float32TypeName     = "Float32"
	Float32TypeFullName = "mojo.core.Float32"

	Float64TypeName     = "Float64"
	Float64TypeFullName = "mojo.core.Float64"

	FloatTypeName     = "Float"
	FloatTypeFullName = "mojo.core.Float"

	DoubleTypeName     = "Double"
	DoubleTypeFullName = "mojo.core.Double"
)

// TruncateNaive
// unit: represents the number of decimal places.
//       The format bit 0.000001 is specified as several bits.
func TruncateNaive(f float64, unit float64) float64 {
	return math.Trunc(f/unit) * unit
}

// RoundNaive
// unit: represents the number of decimal places.
//       The format bit 0.000001 is specified as several bits.
func RoundNaive(f float64, unit float64) float64 {
	return math.Round(f/unit) * unit
}
