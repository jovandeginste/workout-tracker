package timeago

func ukLocale(diff float64, idx int) (ago string, in string) {
	var number = int64(diff)
	switch idx {
	case 0:
		return "щойно", "через декілька секунд"
	case 1:
		return ukSeconds(number) + " тому", "через " + ukSeconds(number)
	case 2, 3:
		return ukMinutes(number) + " тому", "через " + ukMinutes(number)
	case 4, 5:
		return ukHours(number) + " тому", "через " + ukHours(number)
	case 6, 7:
		return ukDays(number) + " тому", "через " + ukDays(number)
	case 8, 9:
		return ukWeeks(number) + " тому", "через " + ukWeeks(number)
	case 10, 11:
		return ukMonths(number) + " тому", "через " + ukMonths(number)
	case 12, 13:
		return ukYears(number) + " тому", "через " + ukYears(number)
	default:
		return "", ""
	}
}

func ukFormatNum(f1 string, f string, s string, t string, n int64) string {
	var n10 = n % 10
	var str = t

	if n == 1 {
		str = f1
	} else if n10 == 1 && n > 20 {
		str = f
	} else if n10 > 1 && n10 < 5 && (n > 20 || n < 10) {
		str = s
	}
	return str
}

func ukSeconds(num int64) string {
	return ukFormatNum("секунду", "%d секунду", "%d секунди", "%d секунд", num)
}
func ukMinutes(num int64) string {
	return ukFormatNum("хвилину", "%d хвилину", "%d хвилини", "%d хвилин", num)
}
func ukHours(num int64) string {
	return ukFormatNum("годину", "%d годину", "%d години", "%d годин", num)
}
func ukDays(num int64) string {
	return ukFormatNum("день", "%d день", "%d дні", "%d днів", num)
}
func ukWeeks(num int64) string {
	return ukFormatNum("тиждень", "%d тиждень", "%d тиждні", "%d тижднів", num)
}
func ukMonths(num int64) string {
	return ukFormatNum("місяць", "%d місяць", "%d місяці", "%d місяців", num)
}
func ukYears(num int64) string {
	return ukFormatNum("рік", "%d рік", "%d роки", "%d років", num)
}
