package humanize

import (
	"math"

	"golang.org/x/text/number"

	"github.com/vorlif/spreak/internal/util"
)

var apnumbers = []string{
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}

// Apnumber returns for numbers 1-9, the number spelled out. Otherwise, return the  number.
// This follows Associated Press style.
//
// Valid inputs are all values that can be interpreted as a number.
//
// Examples:
//
// - 1 becomes one.
//
// - 2 becomes two.
//
// - 10 becomes 10.
func (h *Humanizer) Apnumber(val interface{}) string {
	fl, err := util.ToNumber(val)
	if err != nil {
		return formatErrorMessage(val)
	}

	num := int64(fl)
	if num <= 0 || num >= 10 {
		return h.loc.Print("%d", num)
	}

	return h.loc.Get(apnumbers[num-1])
}

// A tuple of standard large number to their converters.
var intwordConverters = []struct {
	exponent int
	singular string
	plural   string
}{
	{6, "%[1]s million", "%[1]s million"},
	{9, "%[1]s billion", "%[1]s billion"},
	{12, "%[1]s trillion", "%[1]s trillion"},
	{15, "%[1]s quadrillion", "%[1]s quadrillion"},
	{18, "%[1]s quintillion", "%[1]s quintillion"},
	{21, "%[1]s sextillion", "%[1]s sextillion"},
	{24, "%[1]s septillion", "%[1]s septillion"},
	{27, "%[1]s octillion", "%[1]s octillion"},
	{30, "%[1]s nonillion", "%[1]s nonillion"},
	{33, "%[1]s decillion", "%[1]s decillion"},
	{100, "%[1]s googol", "%[1]s googol"},
}

// Intword convert a large integer to a friendly text representation. Works best
// for numbers over 1 million. For example, 1000000 becomes '1.0 million',
// 1200000 becomes '1.2 million' and '1200000000' becomes '1.2 billion'.
// Values up to 10^100 (Googol) are supported.
//
// Translates 1.0 as a singular phrase and all other numeric values as plural, this may be incorrect for some languages.
// Works best for numbers over 1 million.
//
// Valid inputs are all values that can be interpreted as a number.
func (h *Humanizer) Intword(i interface{}) string {
	value, err := util.ToNumber(i)
	if err != nil {
		return formatErrorMessage(i)
	}

	absValue := math.Abs(value)
	if absValue < 1_000_000 {
		return h.Intcomma(value)
	}

	for _, converter := range intwordConverters {
		largeNumber := math.Pow10(converter.exponent)
		if absValue < (largeNumber * float64(1000)) {
			newValue := value / largeNumber
			roundedValue := int64(math.Ceil(newValue-1)) + 1
			if roundedValue < 0 {
				roundedValue = int64(math.Abs(math.Floor(newValue)))
			}
			formattedNumber := h.loc.Print("%.1f", newValue)
			return h.loc.NGetf(converter.singular, converter.plural, roundedValue, formattedNumber)
		}
	}
	return h.Intcomma(value)
}

// Intcomma converts an integer to a string containing commas every three digits.
// For example, 3000 becomes '3,000' and 45000 becomes '45,000'.
//
// Valid inputs are all values that can be interpreted as a number.
func (h *Humanizer) Intcomma(i interface{}) string {
	value, err := util.ToNumber(i)
	if err != nil {
		return formatErrorMessage(i)
	}

	return h.loc.Print("%v", number.Decimal(value, number.MaxFractionDigits(6)))
}

var ordinalTemplates = []struct {
	Context string
	ID      string
}{
	// Translators: Ordinal format when value ends with 0, e.g. 80th.
	{"ordinal 0", "%vth"},
	// Translators: Ordinal format when value ends with 1, e.g. 81st, except 11.
	{"ordinal 1", "%vst"},
	// Translators: Ordinal format when value ends with 2, e.g. 82nd, except 12.
	{"ordinal 2", "%vnd"},
	// Translators: Ordinal format when value ends with 3, e.g. 83rd, except 13.
	{"ordinal 3", "%vrd"},
	// Translators: Ordinal format when value ends with 4, e.g. 84th.
	{"ordinal 4", "%vth"},
	// Translators: Ordinal format when value ends with 5, e.g. 85th.
	{"ordinal 5", "%vth"},
	// Translators: Ordinal format when value ends with 6, e.g. 86th.
	{"ordinal 6", "%vth"},
	// Translators: Ordinal format when value ends with 7, e.g. 87th.
	{"ordinal 7", "%vth"},
	// Translators: Ordinal format when value ends with 8, e.g. 88th.
	{"ordinal 8", "%vth"},
	// Translators: Ordinal format when value ends with 9, e.g. 89th.
	{"ordinal 9", "%vth"},
}

// Ordinal converts an integer to its ordinal as a string. 1 is '1st', 2 is '2nd',
// 3 is '3rd', etc. Works for any integer.
//
// Valid inputs are all values that can be interpreted as a number.
func (h *Humanizer) Ordinal(i interface{}) string {
	floatValue, err := util.ToNumber(i)
	if err != nil {
		return formatErrorMessage(i)
	}

	value := int64(floatValue)
	switch value % 100 {
	case 11, 12, 13:
		return h.loc.PGetf("ordinal 11, 12, 13", "%vth", value)
	}

	template := ordinalTemplates[value%10]
	return h.loc.PGetf(template.Context, template.ID, value)
}
