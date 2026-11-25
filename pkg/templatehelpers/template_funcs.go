package templatehelpers

import (
	"math"
	"strings"

	"github.com/biter777/countries"
	"github.com/spf13/cast"
	"golang.org/x/text/language"
)

const InvalidValue = "N/A"

func HumanCadence(cad float64) string {
	return RoundFloat64(cad)
}

func HumanPower(pow float64) string {
	return RoundFloat64(pow)
}

func HumanCalories(cal float64) string {
	return RoundFloat64(cal)
}

var languageToCountryCodes = map[language.Tag][]countries.CountryCode{
	language.Dutch:               {countries.Netherlands, countries.Belgium},
	language.English:             {countries.UnitedStatesOfAmerica, countries.UnitedKingdom},
	language.French:              {countries.France, countries.Belgium},
	language.German:              {countries.Germany, countries.Austria, countries.Switzerland},
	language.Spanish:             {countries.Spain},
	language.Italian:             {countries.Italy},
	language.Polish:              {countries.Poland},
	language.Portuguese:          {countries.Portugal},
	language.BrazilianPortuguese: {countries.Brazil},
	language.Russian:             {countries.Russia},
	language.Swedish:             {countries.Sweden},
	language.Turkish:             {countries.Turkey},
	language.Finnish:             {countries.Finland},
	language.Persian:             {countries.Iran},
	language.Indonesian:          {countries.Indonesia},
	language.MustParse("nb-NO"):  {countries.Norway},
	language.SimplifiedChinese:   {countries.China},
}

func LanguageToFlag(code string) string {
	tag, err := language.Parse(code)
	if err != nil {
		return "👽"
	}

	ccs, found := languageToCountryCodes[tag]
	if !found {
		return "👽"
	}

	flags := make([]string, 0, len(ccs))
	for _, cc := range ccs {
		flags = append(flags, cc.Emoji())
	}

	result := strings.Join(flags, "")

	if code == "en" {
		result = "🌐" + result
	}

	return result
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
	case "nm":
		return HumanDistanceNM
	default:
		return HumanDistanceKM
	}
}

func HumanSpeedFor(unit string) func(float64) string {
	switch unit {
	case "mph":
		return HumanSpeedMilePH
	case "kn":
		return HumanSpeedKnots
	default:
		return HumanSpeedKPH
	}
}

func HumanTempoFor(unit string) func(float64) string {
	switch unit {
	case "min/mi", "mi":
		return HumanTempoMile
	case "min/nm", "nm":
		return HumanTempoNM
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
	case "nm":
		return v * MeterPerNM
	default:
		return v * MeterPerKM
	}
}

// RoundFloat64 rounds a float64 to 2 decimal places and returns it as a string.
func RoundFloat64(f float64) string {
	return cast.ToString(math.Round(f*100) / 100)
}
