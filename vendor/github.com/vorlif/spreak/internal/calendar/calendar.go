package calendar

import (
	"math"
	"time"
)

// IsLeap return true for leap years, false for non-leap years.
//
// Adapted from https://github.com/python/cpython/blob/main/Lib/calendar.py
func IsLeap(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

// LeapDays return number of leap years in range [y1, y2). Assume y1 <= y2.
//
// Adapted from https://github.com/python/cpython/blob/main/Lib/calendar.py
func LeapDays(year1, year2 int) int {
	y1 := float64(year1 - 1)
	y2 := float64(year2 - 1)
	tmp := math.Floor(y2/4) - math.Floor(y1/4)
	tmp -= math.Floor(y2/100) - math.Floor(y1/100)
	tmp += math.Floor(y2/400) - math.Floor(y1/400)
	return int(tmp)
}

func DaysInMonth(year int, month time.Month) int {
	return time.Date(year, month, 1, 0, 0, 0, 0, time.UTC).AddDate(0, 1, -1).Day()
}
