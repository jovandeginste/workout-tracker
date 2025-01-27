package fitbit

// Locale is used to specify the language and units of API responses.
type Locale string

const (
	LocaleAustralia     Locale = "en_AU"
	LocaleFrance        Locale = "fr_FR"
	LocaleGermany       Locale = "de_DE"
	LocaleJapan         Locale = "ja_JP"
	LocaleNewZealand    Locale = "en_NZ"
	LocaleSpain         Locale = "es_ES"
	LocaleUnitedKingdom Locale = "en_GB"
	LocaleUnitedStates  Locale = "en_US"
)

func (l *Locale) asString() string {
	if l == nil {
		return ""
	}
	return string(*l)
}

// Unit represents a list of units used in API responses.
type Unit struct {
	Distance         string
	Elevation        string
	Height           string
	Weight           string
	BodyMeasurements string
	Liquids          string
	BloodGlucose     string
}

var (
	// UnitedStatesUnit represents a list of units that is used
	// when the language is set to LocaleUnitedStates.
	UnitedStatesUnit = &Unit{
		Distance:         "mile",
		Elevation:        "ft",
		Height:           "in",
		Weight:           "lb",
		BodyMeasurements: "in",
		Liquids:          "fl oz",
		BloodGlucose:     "mg/dL",
	}

	// UnitedKingdomUnit represents a list of units that is used
	// when the language is set to LocaleUnitedStates.
	UnitedKingdomUnit = &Unit{
		Distance:         "km",
		Elevation:        "m",
		Height:           "cm",
		Weight:           "st",
		BodyMeasurements: "cm",
		Liquids:          "ml",
		BloodGlucose:     "mmol/l",
	}

	// MetricUnit represents a list of units that is used
	// when the language is set to neither LocaleUnitedStates nor LocaleUnitedStates.
	MetricUnit = &Unit{
		Distance:         "km",
		Elevation:        "m",
		Height:           "cm",
		Weight:           "kg",
		BodyMeasurements: "cm",
		Liquids:          "ml",
		BloodGlucose:     "mmol/l",
	}
)

func getCorrespondingUnit(locale Locale) *Unit {
	switch locale {
	case LocaleUnitedStates:
		return UnitedStatesUnit
	case LocaleUnitedKingdom:
		return UnitedKingdomUnit
	default:
		return MetricUnit
	}
}
