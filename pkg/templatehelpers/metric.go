package templatehelpers

import (
	"fmt"
	"math"
)

func HumanDistanceKM(d float64) string {
	return fmt.Sprintf("%.2f", d/1000)
}

func HumanSpeedKPH(mps float64) string {
	if mps == 0 {
		return InvalidValue
	}

	kmph := 3.6 * mps

	return fmt.Sprintf("%.2f", kmph)
}

func HumanTempoKM(mps float64) string {
	if mps == 0 {
		return InvalidValue
	}

	mpk := 1000 / (60 * mps)

	wholeMinutes := math.Floor(mpk)
	seconds := (mpk - wholeMinutes) * 60

	return fmt.Sprintf("%d:%02d", int(wholeMinutes), int(seconds))
}

func HumanElevationM(m float64) string {
	return fmt.Sprintf("%.2f", m)
}
