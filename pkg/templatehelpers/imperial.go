package templatehelpers

import (
	"fmt"
	"math"
)

const (
	MilesPerKM   = 0.621371192
	FeetPerMeter = 3.2808399
	MeterPerMile = 1609.344
	PoundsPerKG  = 2.20462262
	CmPerInch    = 2.54
	InchPerFoot  = 12
)

func HumanHeightInch(d float64) string {
	in := d / CmPerInch

	return fmt.Sprintf("%.0f", in)
}

func HumanHeightFeetInch(d float64) string {
	h := d / CmPerInch
	ft := math.Floor(h / InchPerFoot)
	in := math.Mod(h, InchPerFoot)

	return fmt.Sprintf("%.0f ft %.0f in", ft, in)
}

func HumanWeightPounds(d float64) string {
	return RoundFloat64(PoundsPerKG * d)
}

func HumanDistanceMile(d float64) string {
	return RoundFloat64(MilesPerKM * d / MeterPerKM)
}

func HumanSpeedMilePH(mps float64) string {
	return humanSpeed(mps, 3.6*MilesPerKM)
}

func HumanTempoMile(mps float64) string {
	return humanTempo(mps, MeterPerMile)
}

func HumanElevationFt(m float64) string {
	return RoundFloat64(FeetPerMeter * m)
}
