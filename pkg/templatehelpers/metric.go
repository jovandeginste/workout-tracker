package templatehelpers

import (
	"fmt"
	"math"

	"github.com/dustin/go-humanize"
)

func MetricDistance(d float64) string {
	if d < 1000 {
		return fmt.Sprintf("%.2f m", d)
	}

	return fmt.Sprintf("%.2f km", d/1000)
}

func MetricSpeed(mps float64) string {
	if mps == 0 {
		return InvalidValid
	}

	mph := mps * 3600
	value, prefix := humanize.ComputeSI(mph)

	return fmt.Sprintf("%.2f %sm/h", value, prefix)
}

func MetricTempo(mps float64) string {
	if mps == 0 {
		return InvalidValid
	}

	mpk := 1000000 / (mps * 60)
	value, prefix := humanize.ComputeSI(mpk)

	wholeMinutes := math.Floor(value)
	seconds := (value - wholeMinutes) * 60

	return fmt.Sprintf("%d:%02d min/%sm", int(wholeMinutes), int(seconds), prefix)
}
