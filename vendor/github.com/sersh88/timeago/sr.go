package timeago

func srLocale(diff float64, idx int) (ago string, in string) {
	var number = int64(diff)
	switch idx {
	case 0:
		return "малопре", "управо сад"
	case 1:
		return "пре " + srSeconds(number), "за " + srSeconds(number)
	case 2, 3:
		return "пре " + srMinutes(number), "за " + srMinutes(number)
	case 4, 5:
		return "пре " + srHours(number), "за " + srHours(number)
	case 6, 7:
		return "пре " + srDays(number), "за " + srDays(number)
	case 8, 9:
		return "пре " + srWeeks(number), "за " + srWeeks(number)
	case 10, 11:
		return "пре " + srMonths(number), "за " + srMonths(number)
	case 12, 13:
		return "пре " + srYears(number), "за " + srYears(number)
	default:
		return "", ""
	}
}

func srFormatNum(single string, one string, few string, other string, n int64) string {
	var rem10 = n % 10
	var rem100 = n % 100
	if n == 1 {
		return single
	} else if rem10 == 1 && rem100 != 11 {
		return one
	} else if rem10 >= 2 && rem10 <= 4 && !(rem100 >= 12 && rem100 <= 14) {
		return few
	} else {
		return other
	}
}
func srSeconds(num int64) string {
	return srFormatNum("1 секунд", "%d секунд", "%d секунде", "%d секунди", num)
}
func srMinutes(num int64) string {
	return srFormatNum("1 минут", "%d минут", "%d минуте", "%d минута", num)
}
func srHours(num int64) string {
	return srFormatNum("сат времена", "%d сат", "%d сата", "%d сати", num)
}
func srDays(num int64) string {
	return srFormatNum("1 дан", "%d дан", "%d дана", "%d дана", num)
}
func srWeeks(num int64) string {
	return srFormatNum("недељу дана", "%d недељу", "%d недеље", "%d недеља", num)
}
func srMonths(num int64) string {
	return srFormatNum("месец дана", "%d месец", "%d месеца", "%d месеци", num)
}
func srYears(num int64) string {
	return srFormatNum("годину дана", "%d годину", "%d године", "%d година", num)
}
