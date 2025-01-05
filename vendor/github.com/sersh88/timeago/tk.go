package timeago

import (
	"fmt"
	"math"
	"regexp"
)

var tk = []string{"sekunt", "minut", "sagat", "gün", "hepde", "aý", "ýyl"}

func tkGetSuffix(unit string) string {
	if regexp.MustCompile(`[aouy]`).MatchString(unit) {
		return "dan"
	}
	return "den"
}
func tkLocale(diff float64, idx int) (ago string, in string) {
	if idx == 0 {
		return "biraz öň", "şuwagt"
	}
	var unit = tk[int(math.Floor(float64(idx)/2))]
	return fmt.Sprintf("%d %s öň", int(diff), unit),
		fmt.Sprintf("%d %s%s", int(diff), unit, tkGetSuffix(unit))
}
