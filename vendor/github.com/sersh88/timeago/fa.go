package timeago

import (
	"strconv"
	"strings"
)

func faLocale(number float64, index int) (ago string, in string) {
	var formattedString = [][]string{
		{"لحظاتی پیش", "همین حالا"},
		{"%d ثانیه پیش", "%d ثانیه دیگر"},
		{"۱ دقیقه پیش", "۱ دقیقه دیگر"},
		{"%d دقیقه پیش", "%d دقیقه دیگر"},
		{"۱ ساعت پیش", "۱ ساعت دیگر"},
		{"%d ساعت پیش", "%d ساعت دیگر"},
		{"۱ روز پیش", "۱ روز دیگر"},
		{"%d روز پیش", "%d روز دیگر"},
		{"۱ هفته پیش", "۱ هفته دیگر"},
		{"%d هفته پیش", "%d هفته دیگر"},
		{"۱ ماه پیش", "۱ ماه دیگر"},
		{"%d ماه پیش", "%d ماه دیگر"},
		{"۱ سال پیش", "۱ سال دیگر"},
		{"%d سال پیش", "%d سال دیگر"},
	}[index]
	return strings.ReplaceAll(formattedString[0], "%d", toPersianNumber(int64(number))),
		strings.ReplaceAll(formattedString[1], "%d", toPersianNumber(int64(number)))
}

// As persian language has different number symbols we need to replace regular numbers
// to standard persian numbers.
func toPersianNumber(number int64) string {
	//List of standard persian numbers from 0 to 9
	var persianDigits = []rune{'۰', '۱', '۲', '۳', '۴', '۵', '۶', '۷', '۸', '۹'}
	var res = []rune(strconv.FormatInt(number, 10))
	for k, d := range res {
		digit, _ := strconv.ParseInt(string(d), 10, 64)
		res[k] = persianDigits[digit]
	}
	return string(res)
}
