package templatehelpers

import (
	"fmt"
	"math"
)

const (
	NauticalMilePerKM = 0.539957
	MeterPerNM        = 1852
)

// HumanDistanceNM converts meters to nautical miles
func HumanDistanceNM(d float64) string {
	return RoundFloat64(d / MeterPerNM)
}

// HumanSpeedKnots converts m/s to knots
func HumanSpeedKnots(mps float64) string {
	if mps == 0 {
		return InvalidValue
	}

	knots := mps * 1.94384
	return RoundFloat64(knots)
}

// HumanTempoNM calculates tempo in min/nm
// mps is meters per second
func HumanTempoNM(mps float64) string {
	if mps == 0 || math.IsNaN(mps) {
		return InvalidValue
	}

	mpnm := MeterPerNM / (60 * mps)

	wholeMinutes := math.Floor(mpnm)
	seconds := (mpnm - wholeMinutes) * 60

	return fmt.Sprintf("%d:%02d", int(wholeMinutes), int(seconds))
}
