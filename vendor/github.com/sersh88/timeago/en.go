package timeago

import (
	"fmt"
	"math"
)

var EN = []string{"second", "minute", "hour", "day", "week", "month", "year"}

func enLocale(diff float64, idx int) (ago string, in string) {
	if idx == 0 {
		return "just now", "right now"
	}
	var unit = EN[int(math.Floor(float64(idx)/2))]
	if diff > 1 {
		unit += "s"
	}
	return fmt.Sprintf("%d %s ago", int(diff), unit), fmt.Sprintf("in %d %s", int(diff), unit)
}
