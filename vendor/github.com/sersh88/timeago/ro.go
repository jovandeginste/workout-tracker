package timeago

import "strings"

func roLocale(diff float64, index int) (ago string, in string) {
	var langTable = [][]string{
		{"chiar acum", "chiar acum"},
		{"acum %d secunde", "peste %d secunde"},
		{"acum un minut", "peste un minut"},
		{"acum %d minute", "peste %d minute"},
		{"acum o oră", "peste o oră"},
		{"acum %d ore", "peste %d ore"},
		{"acum o zi", "peste o zi"},
		{"acum %d zile", "peste %d zile"},
		{"acum o săptămână", "peste o săptămână"},
		{"acum %d săptămâni", "peste %d săptămâni"},
		{"acum o lună", "peste o lună"},
		{"acum %d luni", "peste %d luni"},
		{"acum un an", "peste un an"},
		{"acum %d ani", "peste %d ani"},
	}
	var number = int64(diff)
	var res = langTable[index]
	if number < 20 {
		return res[0], res[1]
	}
	// A `de` preposition must be added between the number and the adverb
	// if the number is greater than 20.
	return strings.ReplaceAll(res[0], "%d", "%d de"),
		strings.ReplaceAll(res[1], "%d", "%d de")
}
