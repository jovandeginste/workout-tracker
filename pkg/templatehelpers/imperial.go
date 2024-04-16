package templatehelpers

import (
	"fmt"
	"math"
)

const (
	milesPerKM   = 0.621371192
	feetPerMeter = 3.2808399
)

func HumanDistanceMile(d float64) string {
	return fmt.Sprintf("%.2f", milesPerKM*d/1000)
}

func HumanSpeedMilePH(mps float64) string {
	if mps == 0 {
		return InvalidValue
	}

	kmph := 3.6 * mps

	return fmt.Sprintf("%.2f", milesPerKM*kmph)
}

func HumanTempoMile(mps float64) string {
	if mps == 0 || math.IsNaN(mps) {
		return InvalidValue
	}

	mpm := 1000 / (60 * mps) / milesPerKM

	wholeMinutes := math.Floor(mpm)
	seconds := (mpm - wholeMinutes) * 60

	return fmt.Sprintf("%d:%02d", int(wholeMinutes), int(seconds))
}

func HumanElevationFt(m float64) string {
	return fmt.Sprintf("%.2f", feetPerMeter*m)
}
