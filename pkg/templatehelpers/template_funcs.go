package templatehelpers

import (
	"fmt"
	"time"

	emojiflag "github.com/jayco/go-emoji-flag"
)

const InvalidValue = "N/A"

func HumanCaloriesKcal(cal float64) string {
	return fmt.Sprintf("%.2f kcal", cal)
}

func NumericDuration(d time.Duration) float64 {
	return d.Seconds()
}

func CountryCodeToFlag(cc string) string {
	return emojiflag.GetFlag(cc)
}

func HumanElevationFor(unit string) func(float64) string {
	switch unit {
	case "ft":
		return HumanElevationFt
	default:
		return HumanElevationM
	}
}

func HumanHeightSingleFor(unit string) func(float64) string {
	switch unit {
	case "in":
		return HumanHeightInch
	default:
		return HumanHeightCMNoSuffix
	}
}

func HumanHeightFor(unit string) func(float64) string {
	switch unit {
	case "in":
		return HumanHeightFeetInch
	default:
		return HumanHeightCM
	}
}

func HumanWeightFor(unit string) func(float64) string {
	switch unit {
	case "lbs":
		return HumanWeightPounds
	default:
		return HumanWeightKG
	}
}

func HumanDistanceFor(unit string) func(float64) string {
	switch unit {
	case "mi":
		return HumanDistanceMile
	default:
		return HumanDistanceKM
	}
}

func HumanSpeedFor(unit string) func(float64) string {
	switch unit {
	case "mph":
		return HumanSpeedMilePH
	default:
		return HumanSpeedKPH
	}
}

func HumanTempoFor(unit string) func(float64) string {
	switch unit {
	case "min/mi", "mi":
		return HumanTempoMile
	default:
		return HumanTempoKM
	}
}

func HeightToDatabase(v float64, unit string) float64 {
	switch unit {
	case "in":
		return v * CmPerInch
	default:
		return v
	}
}

func WeightToDatabase(v float64, unit string) float64 {
	switch unit {
	case "lbs":
		return v / PoundsPerKG
	default:
		return v
	}
}

func DistanceToDatabase(v float64, unit string) float64 {
	switch unit {
	case "mi":
		return v * MeterPerMile
	case "m":
		return v
	default:
		return v * MeterPerKM
	}
}
