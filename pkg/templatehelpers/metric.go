package templatehelpers

import (
	"fmt"
	"math"
)

const MeterPerKM = 1000

func HumanHeightCMNoSuffix(d float64) string {
	return RoundFloat64(d)
}

func HumanHeightCM(d float64) string {
	return fmt.Sprintf("%.0f cm", d)
}

func HumanWeightKG(d float64) string {
	return RoundFloat64(d)
}

func HumanDistanceKM(d float64) string {
	return RoundFloat64(d / MeterPerKM)
}

func HumanSpeedKPH(mps float64) string {
	return humanSpeed(mps, 3.6)
}

func HumanTempoKM(mps float64) string {
	return humanTempo(mps, MeterPerKM)
}

func humanSpeed(mps, factor float64) string {
	if mps == 0 {
		return InvalidValue
	}

	return RoundFloat64(factor * mps)
}

func humanTempo(mps, metersPerUnit float64) string {
	if mps == 0 || math.IsNaN(mps) {
		return InvalidValue
	}

	minutesPerUnit := metersPerUnit / (60 * mps)

	wholeMinutes := math.Floor(minutesPerUnit)
	seconds := (minutesPerUnit - wholeMinutes) * 60

	return fmt.Sprintf("%d:%02d", int(wholeMinutes), int(seconds))
}

func HumanElevationM(m float64) string {
	return RoundFloat64(m)
}
