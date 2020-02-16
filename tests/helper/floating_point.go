package helper

import (
	"math"
)

// Floating-Point comparison guide from
// https://floating-point-gui.de/errors/comparison

func FloatNearlyEqual(a float64, b float64, epsilon float64) bool {
	absA := math.Abs(a)
	absB := math.Abs(b)
	diff := math.Abs(a - b)

	if a == b { // shortcut, handles infinities
		return true
	} else if a == 0 || b == 0 || absA+absB < math.SmallestNonzeroFloat64 {
		// a or b is zero or both are extremely close to it
		// relative error is less meaningful here
		return diff < (epsilon * math.MaxFloat64)
	} else { // use relative error
		return diff/math.Min((absA+absB), math.MaxFloat64) < epsilon
	}
}
