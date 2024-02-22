package util

import "math"

func FloatEqual(a, b float64) bool {
	return math.Abs(a-b) <= 1e-6
}
