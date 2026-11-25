package templatehelpers

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
	return humanSpeed(mps, 1.94384)
}

// HumanTempoNM calculates tempo in min/nm
// mps is meters per second
func HumanTempoNM(mps float64) string {
	return humanTempo(mps, MeterPerNM)
}
