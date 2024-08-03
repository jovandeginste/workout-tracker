package templatehelpers

import (
	"fmt"
	"math"
)

const (
	MilesPerKM   = 0.621371192
	FeetPerMeter = 3.2808399
	MeterPerMile = 1609.344
)

func HumanDistanceMile(d float64) string {
	return fmt.Sprintf("%.2f", MilesPerKM*d/MeterPerKM)
}

func HumanSpeedMilePH(mps float64) string {
	if mps == 0 {
		return InvalidValue
	}

	kmph := 3.6 * mps

	return fmt.Sprintf("%.2f", MilesPerKM*kmph)
}

func HumanTempoMile(mps float64) string {
	if mps == 0 || math.IsNaN(mps) {
		return InvalidValue
	}

	mpm := MeterPerMile / (60 * mps)

	wholeMinutes := math.Floor(mpm)
	seconds := (mpm - wholeMinutes) * 60

	return fmt.Sprintf("%d:%02d", int(wholeMinutes), int(seconds))
}

func HumanElevationFt(m float64) string {
	return fmt.Sprintf("%.2f", FeetPerMeter*m)
}
