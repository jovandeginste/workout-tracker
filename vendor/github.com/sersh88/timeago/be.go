package timeago

func beLocale(diff float64, idx int) (ago string, in string) {
	var num = int64(diff)
	switch idx {
	case 0:
		return "толькі што", "праз некалькі секунд"
	case 1:
		return ruFormatNum("секунду", "%d секунду", "%d секунды", "%d секунд", num) + " таму",
			"праз  " + ruFormatNum("секунду", "%d секунду", "%d секунды", "%d секунд", num)
	case 2, 3:
		return ruFormatNum("хвіліну", "%d хвіліну", "%d хвіліны", "%d хвілін", num) + " таму",
			"праз  " + ruFormatNum("хвіліну", "%d хвіліну", "%d хвіліны", "%d хвілін", num)
	case 4, 5:
		return ruFormatNum("гадзіну", "%d гадзіну", "%d гадзіны", "%d гадзін", num) + " таму",
			"праз  " + ruFormatNum("гадзіну", "%d гадзіну", "%d гадзіны", "%d гадзін", num)
	case 6, 7:
		return ruFormatNum("дзень", "%d дзень", "%d дні", "%d дзён", num) + " таму",
			"праз  " + ruFormatNum("дзень", "%d дзень", "%d дні", "%d дзён", num)
	case 8, 9:
		return ruFormatNum("тыдзень", "%d тыдзень", "%d тыдні", "%d тыдняў", num) + " таму",
			"праз  " + ruFormatNum("тыдзень", "%d тыдзень", "%d тыдні", "%d тыдняў", num)
	case 10, 11:
		return ruFormatNum("месяц", "%d месяц", "%d месяцы", "%d месяцаў", num) + " таму",
			"праз  " + ruFormatNum("месяц", "%d месяц", "%d месяцы", "%d месяцаў", num)
	case 12, 13:
		return ruFormatNum("год", "%d год", "%d гады", "%d гадоў", num) + " таму",
			"праз  " + ruFormatNum("год", "%d год", "%d гады", "%d гадоў", num)
	default:
		return "", ""
	}
}
