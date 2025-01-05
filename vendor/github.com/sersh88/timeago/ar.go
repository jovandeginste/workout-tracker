package timeago

import "math"

var arTimeTypes = [][]string{
	{"ثانية", "ثانيتين", "%d ثوان", "%d ثانية"},     // Seconds
	{"دقيقة", "دقيقتين", "%d دقائق", "%d دقيقة"},    // Minutes
	{"ساعة", "ساعتين", "%d ساعات", "%d ساعة"},       // Hours
	{"يوم", "يومين", "%d أيام", "%d يوماً"},         // Days
	{"أسبوع", "أسبوعين", "%d أسابيع", "%d أسبوعاً"}, // Weeks
	{"شهر", "شهرين", "%d أشهر", "%d شهراً"},         // Months
	{"عام", "عامين", "%d أعوام", "%d عاماً"},        // Years
}

func arLocale(diff float64, idx int) (ago string, in string) {
	if idx == 0 {
		return "منذ لحظات", "بعد لحظات"
	}
	var timeStr = arFormatTime(int64(math.Floor(float64(idx)/2)), int64(diff))
	return "منذ" + " " + timeStr, "بعد" + " " + timeStr
}

func arFormatTime(tp int64, n int64) string {
	if n < 3 {
		return arTimeTypes[tp][n-1]
	}
	if n >= 3 && n <= 10 {
		return arTimeTypes[tp][2]
	}
	return arTimeTypes[tp][3]
}
