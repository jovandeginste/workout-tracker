package humanize

import (
	"fmt"
	"math"
	"reflect"

	"golang.org/x/text/language"
	"golang.org/x/text/number"

	"github.com/vorlif/spreak/internal/util"
)

// LanguageName returns the name of the spoken language as called by the languages used.
//
// If no translation exists for the name of the language, the input is returned.
func (h *Humanizer) LanguageName(lang string) string {
	return h.loc.Get(lang)
}

// LanguageNameByCode returns the name of a language for a language code
// in the national language of the Humanizer.
//
// If no language is known for the code, an empty string is returned.
// If a language is known, but no translations are available, the English name is returned.
//
// For example, if a Humanizer was created for French, 'Anglais' is returned for the code 'en'
// and the value 'Allemand' for the code 'de'.
func (h *Humanizer) LanguageNameByCode(code string) string {
	info, ok := LocaleInfos[code]
	if !ok {
		return ""
	}
	return h.loc.Get(info.Name)
}

// Language returns the currently used language.
func (h *Humanizer) Language() language.Tag {
	return h.loc.Language()
}

const (
	kb = int64(1) << 10
	mb = int64(1) << 20
	gb = int64(1) << 30
	tb = int64(1) << 40
	pb = int64(1) << 50
)

// FilesizeFormat format the value like a 'human-readable' file size (i.e. 13 KB, 4.1 MB, 102 bytes, etc.).
//
// Valid inputs are byte arrays or any numeric value.
// For all other inputs, a string is returned with an error message in fmt style.
func (h *Humanizer) FilesizeFormat(v interface{}) string {
	var count int64
	switch val := v.(type) {
	case []byte:
		count = int64(len(val))
	default:
		value, err := util.ToNumber(v)
		if err != nil {
			return formatErrorMessage(v)
		}
		count = int64(value)
	}

	isNegative := count < 0
	if isNegative {
		count = -count
	}

	var result string
	if count < kb {
		result = h.loc.NGetf("%[1]d byte", "%[1]d bytes", count, count)
	} else if count < mb {
		formatted := h.loc.Print("%v", number.Decimal(float64(count)/float64(kb), number.MaxFractionDigits(1)))
		result = h.loc.Getf("%s KB", formatted)
	} else if count < gb {
		formatted := h.loc.Print("%v", number.Decimal(float64(count)/float64(mb), number.MaxFractionDigits(1)))
		result = h.loc.Getf("%s MB", formatted)
	} else if count < tb {
		formatted := h.loc.Print("%v", number.Decimal(float64(count)/float64(gb), number.MaxFractionDigits(1)))
		result = h.loc.Getf("%s GB", formatted)
	} else if count < pb {
		formatted := h.loc.Print("%v", number.Decimal(float64(count)/float64(tb), number.MaxFractionDigits(1)))
		result = h.loc.Getf("%s TB", formatted)
	} else {
		formatted := h.loc.Print("%v", number.Decimal(float64(count)/float64(pb), number.MaxFractionDigits(1)))
		result = h.loc.Getf("%s PB", formatted)
	}

	if isNegative {
		result = "-" + result
	}
	return result
}

func floorDivision(a, b float64) int64 {
	return int64(math.Floor(toFixed(a/b, 3)))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return math.Round(num*output) / output
}

func formatErrorMessage(i interface{}) string {
	t := reflect.TypeOf(i)
	if t == nil {
		return "<nil>"
	}

	return fmt.Sprintf("%%!(%s=%v)", t.String(), i)
}
