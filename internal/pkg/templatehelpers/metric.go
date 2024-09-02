package templatehelpers

import (
	"fmt"
	"math"
)

const MeterPerKM = 1000

func HumanDistanceKM(d float64) string {
	return fmt.Sprintf("%.2f", d/MeterPerKM)
}

func HumanSpeedKPH(mps float64) string {
	if mps == 0 {
		return InvalidValue
	}

	kmph := 3.6 * mps

	return fmt.Sprintf("%.2f", kmph)
}

func HumanTempoKM(mps float64) string {
	if mps == 0 || math.IsNaN(mps) {
		return InvalidValue
	}

	mpk := MeterPerKM / (60 * mps)

	wholeMinutes := math.Floor(mpk)
	seconds := (mpk - wholeMinutes) * 60

	return fmt.Sprintf("%d:%02d", int(wholeMinutes), int(seconds))
}

func HumanElevationM(m float64) string {
	return fmt.Sprintf("%.2f", m)
}
