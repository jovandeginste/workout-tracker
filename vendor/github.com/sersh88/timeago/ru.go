package timeago

func ruLocale(diff float64, idx int) (ago string, in string) {
	var num = int64(diff)
	switch idx {
	case 0:
		return "только что", "через несколько секунд"
	case 1:
		return ruFormatNum("секунду", "%d секунду", "%d секунды", "%d секунд", num) + " назад",
			"через " + ruFormatNum("секунду", "%d секунду", "%d секунды", "%d секунд", num)
	case 2, 3:
		return ruFormatNum("минуту", "%d минуту", "%d минуты", "%d минут", num) + " назад",
			"через " + ruFormatNum("минуту", "%d минуту", "%d минуты", "%d минут", num)
	case 4, 5:
		return ruFormatNum("час", "%d час", "%d часа", "%d часов", num) + " назад",
			"через " + ruFormatNum("час", "%d час", "%d часа", "%d часов", num)
	case 6:
		return "вчера", "завтра"
	case 7:
		return ruFormatNum("день", "%d день", "%d дня", "%d дней", num) + " назад",
			"через " + ruFormatNum("день", "%d день", "%d дня", "%d дней", num)
	case 8, 9:
		return ruFormatNum("неделю", "%d неделю", "%d недели", "%d недель", num) + " назад",
			"через " + ruFormatNum("неделю", "%d неделю", "%d недели", "%d недель", num)
	case 10, 11:
		return ruFormatNum("месяц", "%d месяц", "%d месяца", "%d месяцев", num) + " назад",
			"через " + ruFormatNum("месяц", "%d месяц", "%d месяца", "%d месяцев", num)
	case 12, 13:
		return ruFormatNum("год", "%d год", "%d года", "%d лет", num) + " назад",
			"через " + ruFormatNum("год", "%d год", "%d года", "%d лет", num)
	default:
		return "", ""
	}
}

func ruFormatNum(f1 string, f string, s string, t string, n int64) string {
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
