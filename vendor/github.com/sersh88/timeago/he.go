package timeago

func heLocale(number float64, index int) (ago string, in string) {
	switch index {
	case 0:
		return "זה עתה", "עכשיו"
	case 1:
		return "לפני %d שניות", "בעוד %d שניות"
	case 2:
		return "לפני דקה", "בעוד דקה"
	case 3:
		return "לפני %d דקות", "בעוד %d דקות"
	case 4:
		return "לפני שעה", "בעוד שעה"
	case 5:
		if number == 2 {
			return "לפני שעתיים", "בעוד שעתיים"
		} else {
			return "לפני %d שעות", "בעוד %d שעות"
		}
	case 6:
		return "אתמול", "מחר"
	case 7:
		if number == 2 {
			return "לפני יומיים", "בעוד יומיים"
		} else {
			return "לפני %d ימים", "בעוד %d ימים"
		}
	case 8:
		return "לפני שבוע", "בעוד שבוע"
	case 9:
		if number == 2 {
			return "לפני שבועיים", "בעוד שבועיים"
		} else {
			return "לפני %d שבועות", "בעוד %d שבועות"
		}
	case 10:
		return "לפני חודש", "בעוד חודש"
	case 11:
		if number == 2 {
			return "לפני חודשיים", "בעוד חודשיים"
		} else {
			return "לפני %d חודשים", "בעוד %d חודשים"
		}
	case 12:
		return "לפני שנה", "בעוד שנה"
	case 13:
		if number == 2 {
			return "לפני שנתיים", "בעוד שנתיים"
		} else {
			return "לפני %d שנים", "בעוד %d שנים"
		}
	default:
		return "", ""
	}
}
