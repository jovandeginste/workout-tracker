package templatehelpers

import (
	"fmt"
	"math"

	"github.com/bcicen/go-units"
)

var impOpts = units.FmtOptions{
	Label:     true, // append unit name/symbol
	Short:     true, // use unit symbol
	Precision: 2,
}

func ImperialDistance(d float64) string {
	value := units.NewValue(d, units.Meter)

	var f units.Value

	if d > 2000 {
		f = value.MustConvert(units.Mile)
	} else {
		f = value.MustConvert(units.Foot)
	}

	return f.Fmt(impOpts)
}

func ImperialSpeed(mps float64) string {
	if mps == 0 {
		return InvalidValid
	}

	mph := mps * 3600
	value := units.NewValue(mph, units.Meter)

	var f units.Value

	if mph > 2000 {
		f = value.MustConvert(units.Mile)
	} else {
		f = value.MustConvert(units.Foot)
	}

	return f.Fmt(impOpts) + "/h"
}

func ImperialTempo(mps float64) string {
	if mps == 0 {
		return InvalidValid
	}

	mph := mps * 3600

	value := units.NewValue(mph, units.Meter)
	mileph := value.MustConvert(units.Mile)

	wholeMinutes := math.Floor(mileph.Float())
	seconds := (mileph.Float() - wholeMinutes) * 60

	return fmt.Sprintf("%d:%02d min/%s", int(wholeMinutes), int(seconds), mileph.Unit().Symbol)
}
