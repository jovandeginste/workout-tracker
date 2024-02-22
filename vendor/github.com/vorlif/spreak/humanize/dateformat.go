package humanize

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/vorlif/spreak"
	"github.com/vorlif/spreak/internal/calendar"
)

type timeFormatter struct {
	loc  *spreak.Localizer
	data time.Time
}

func newTimeFormatter(h *Humanizer, t time.Time) *timeFormatter {
	return &timeFormatter{
		loc:  h.loc,
		data: t,
	}
}

//nolint:cyclop
func (tf *timeFormatter) format(formatString string) string {
	runes := []rune(formatString)
	var buf strings.Builder
	for i := 0; i < len(runes); i++ {
		switch runes[i] {
		case '\\':
			if i+1 < len(runes) {
				i++
				buf.WriteRune(runes[i])
			} else {
				buf.WriteRune('\\')
			}
		case 'a':
			buf.WriteString(tf.a())
		case 'A':
			buf.WriteString(tf.A())
		case 'e':
			buf.WriteString(tf.e())
		case 'f':
			buf.WriteString(tf.f())
		case 'g':
			buf.WriteString(tf.g())
		case 'G':
			buf.WriteString(tf.G())
		case 'h':
			buf.WriteString(tf.h())
		case 'H':
			buf.WriteString(tf.H())
		case 'i':
			buf.WriteString(tf.i())
		case 'O':
			buf.WriteString(tf.O())
		case 'P':
			buf.WriteString(tf.P())
		case 's':
			buf.WriteString(tf.s())
		case 'T':
			buf.WriteString(tf.T())
		case 'u':
			buf.WriteString(tf.u())
		case 'Z':
			buf.WriteString(tf.Z())
		case 'b':
			buf.WriteString(tf.b())
		case 'c':
			buf.WriteString(tf.c())
		case 'd':
			buf.WriteString(tf.d())
		case 'D':
			buf.WriteString(tf.D())
		case 'E':
			buf.WriteString(tf.E())
		case 'F':
			buf.WriteString(tf.F())
		case 'I':
			buf.WriteString(tf.I())
		case 'j':
			buf.WriteString(tf.j())
		case 'l':
			buf.WriteString(tf.l())
		case 'L':
			buf.WriteString(tf.L())
		case 'm':
			buf.WriteString(tf.m())
		case 'M':
			buf.WriteString(tf.M())
		case 'n':
			buf.WriteString(tf.n())
		case 'N':
			buf.WriteString(tf.N())
		case 'o':
			buf.WriteString(tf.o())
		case 'r':
			buf.WriteString(tf.r())
		case 'S':
			buf.WriteString(tf.S())
		case 't':
			buf.WriteString(tf.t())
		case 'U':
			buf.WriteString(tf.U())
		case 'w':
			buf.WriteString(tf.w())
		case 'W':
			buf.WriteString(tf.W())
		case 'y':
			buf.WriteString(tf.y())
		case 'Y':
			buf.WriteString(tf.Y())
		case 'z':
			buf.WriteString(tf.z())
		default:
			buf.WriteRune(runes[i])
		}
	}
	return buf.String()
}

// 'a.m.' or 'p.m.'.
func (tf *timeFormatter) a() string {
	if tf.data.Hour() > 11 {
		return tf.loc.Get("p.m.")
	}
	return tf.loc.Get("a.m.")
}

// A returns 'AM' or 'PM'.
func (tf *timeFormatter) A() string {
	if tf.data.Hour() > 11 {
		return tf.loc.Get("PM")
	}
	return tf.loc.Get("AM")
}

// Timezone name.
func (tf *timeFormatter) e() string {
	zone, _ := tf.data.Zone()
	return zone
}

// Time, in 12-hour hours and minutes, with minutes left off if they're zero.
// Examples: '1', '1:30', '2:05', '2'
// Proprietary extension.
func (tf *timeFormatter) f() string {
	hour := tf.data.Format("3")
	if tf.data.Minute() > 0 {
		return fmt.Sprintf("%s:%s", hour, tf.data.Format("04"))
	}
	return hour
}

// Hour, 12-hour format without leading zeros; i.e. '1' to '12'.
func (tf *timeFormatter) g() string {
	return tf.data.Format("3")
}

// G - Hour, 24-hour format without leading zeros; i.e. '0' to '23'.
func (tf *timeFormatter) G() string {
	return strconv.Itoa(tf.data.Hour())
}

// h - Hour, 12-hour format; i.e. '01' to '12'.
func (tf *timeFormatter) h() string {
	return tf.data.Format("03")
}

// H - Hour, 24-hour format; i.e. '00' to '23'.
func (tf *timeFormatter) H() string {
	return fmt.Sprintf("%02d", tf.data.Hour())
}

// i - Minutes; i.e. '00' to '59'.
func (tf *timeFormatter) i() string {
	return fmt.Sprintf("%02d", tf.data.Minute())
}

// O - Difference to Greenwich time in hours; e.g. '+0200', '-0430'.
func (tf *timeFormatter) O() string {
	return tf.data.Format("-0700")
}

// P - Time, in 12-hour hours, minutes and 'a.m.'/'p.m.', with minutes left off
// if they're zero and the strings 'midnight' and 'noon' if appropriate.
// Examples: '1 a.m.', '1:30 p.m.', 'midnight', 'noon', '12:30 p.m.'
// Proprietary extension.
func (tf *timeFormatter) P() string {
	if tf.data.Minute() == 0 && tf.data.Hour() == 0 {
		return tf.loc.Get("midnight")
	}

	if tf.data.Minute() == 0 && tf.data.Hour() == 12 {
		return tf.loc.Get("noon")
	}

	return fmt.Sprintf("%s %s", tf.f(), tf.a())
}

// s - Seconds; i.e. '00' to '59'.
func (tf *timeFormatter) s() string {
	return fmt.Sprintf("%02d", tf.data.Second())
}

// T - Time zone of this machine; e.g. 'EST' or 'MDT'.
func (tf *timeFormatter) T() string {
	return time.Local.String()
}

// u return microseconds; i.e. '000000' to '999999'.
func (tf *timeFormatter) u() string {
	tmp := tf.data.Nanosecond() / 1e3
	return fmt.Sprintf("%.6d", tmp)
}

// Z - Time zone offset in seconds (i.e. '-43200' to '43200'). The offset for
// timezones west of UTC is always negative, and for those east of UTC is
// always positive.
func (tf *timeFormatter) Z() string {
	_, offset := tf.data.Zone()
	return strconv.Itoa(offset)
}

// b - Month, textual, 3 letters, lowercase; e.g. 'jan'.
func (tf *timeFormatter) b() string {
	monthShortName := months3[int(tf.data.Month())]
	return tf.loc.Get(monthShortName)
}

// ISO 8601 Format
// Example : '2008-01-02T10:30:00.000123'.
func (tf *timeFormatter) c() string {
	return tf.data.Format("2006-01-02T15:04:05.") + strconv.Itoa(tf.data.Nanosecond())
}

// d - Day of the month, 2 digits with leading zeros; i.e. '01' to '31'.
func (tf *timeFormatter) d() string {
	return fmt.Sprintf("%02d", tf.data.Day())
}

// D - Day of the week, textual, 3 letters; e.g. 'Fri'.
func (tf *timeFormatter) D() string {
	weekdayIndex := (6 + int(tf.data.Weekday())) % 7
	weekdayName := weekdaysAbbr[weekdayIndex]
	return tf.loc.Get(weekdayName)
}

// E - Alternative month names as required by some locales. Proprietary extension.
func (tf *timeFormatter) E() string {
	entry := monthsAlt[int(tf.data.Month())]
	return tf.loc.PGet(entry.context, entry.singular)
}

// F - Month, textual, long; e.g. 'January'.
func (tf *timeFormatter) F() string {
	monthName := months[int(tf.data.Month())]
	return tf.loc.Get(monthName)
}

// I - '1' if daylight saving time, '0' otherwise.
func (tf *timeFormatter) I() string {
	_, timeOffset := tf.data.Zone()
	_, winterOffset := time.Date(tf.data.Year(), 1, 1, 0, 0, 0, 0, tf.data.Location()).Zone()
	_, summerOffset := time.Date(tf.data.Year(), 7, 1, 0, 0, 0, 0, tf.data.Location()).Zone()

	if winterOffset > summerOffset {
		winterOffset, summerOffset = summerOffset, winterOffset
	}

	if winterOffset != summerOffset { // the location has daylight saving
		if timeOffset != winterOffset {
			return "1"
		}
	}
	return "0"
}

// j: Day of the month without leading zeros; i.e. '1' to '31'
func (tf *timeFormatter) j() string {
	return strconv.Itoa(tf.data.Day())
}

// Day of the week, textual, long; e.g. 'Friday'.
func (tf *timeFormatter) l() string {
	weekdayIndex := (6 + int(tf.data.Weekday())) % 7
	weekdayName := weekdays[weekdayIndex]
	return tf.loc.Get(weekdayName)
}

// L return the string representation of boolean for whether it is a leap year;
// i.e. True or False.
func (tf *timeFormatter) L() string {
	return fmt.Sprintf("%t", calendar.IsLeap(tf.data.Year()))
}

// m returns the Month; i.e. '01' to '12'.
func (tf *timeFormatter) m() string {
	return fmt.Sprintf("%02d", tf.data.Month())
}

// M return the month, textual, 3 letters; e.g. 'Jan'.
func (tf *timeFormatter) M() string {
	text := months3[int(tf.data.Month())]
	r := []rune(tf.loc.Get(text))
	return string(append([]rune{unicode.ToUpper(r[0])}, r[1:]...))
}

// n returns the month without leading zeros; i.e. '1' to '12'.
func (tf *timeFormatter) n() string {
	return strconv.Itoa(int(tf.data.Month()))
}

// N return the month abbreviation in Associated Press style. Proprietary extension.
func (tf *timeFormatter) N() string {
	entry := monthsAp[int(tf.data.Month())]
	return tf.loc.PGet(entry.context, entry.singular)
}

// o return the ISO 8601 year number matching the ISO week number (W).
func (tf *timeFormatter) o() string {
	year, _ := tf.data.ISOWeek()
	return strconv.Itoa(year)
}

// r return an RFC 5322 formatted date; e.g. 'Thu, 21 Dec 2000 16:01:07 +0200'.
func (tf *timeFormatter) r() string {
	// is this the same?
	return tf.data.Format(time.RFC1123Z)
}

// S return the english ordinal suffix for the day of the month, 2 characters; i.e.
// 'st', 'nd', 'rd' or 'th'.
func (tf *timeFormatter) S() string {
	switch tf.data.Day() {
	case 11, 12, 13: // Special case
		return "th"
	}
	switch tf.data.Day() % 10 {
	case 1:
		return "st"
	case 2:
		return "nd"
	case 3:
		return "rd"
	default:
		return "th"
	}
}

// t return the number of days in the given month; i.e. '28' to '31'.
func (tf *timeFormatter) t() string {
	return strconv.Itoa(calendar.DaysInMonth(tf.data.Year(), tf.data.Month()))
}

// U return seconds since the Unix epoch (January 1 1970 00:00:00 GMT).
func (tf *timeFormatter) U() string {
	return strconv.FormatInt(tf.data.Unix(), 10)
}

// w return day of the week, numeric, i.e. '0' (Sunday) to '6' (Saturday).
func (tf *timeFormatter) w() string {
	return strconv.Itoa(int(tf.data.Weekday()))
}

// W return ISO-8601 week number of year, weeks starting on Monday.
func (tf *timeFormatter) W() string {
	_, wn := tf.data.ISOWeek()
	return strconv.Itoa(wn)
}

// y return year, 2 digits with leading zeros; e.g. '99'.
func (tf *timeFormatter) y() string {
	return fmt.Sprintf("%02d", tf.data.Year()%100)
}

// Y returns year, 4 digits with leading zeros; e.g. '1999'.
func (tf *timeFormatter) Y() string {
	return fmt.Sprintf("%04d", tf.data.Year())
}

// Day of the year, i.e. 1 to 366.
func (tf *timeFormatter) z() string {
	return strconv.Itoa(tf.data.YearDay())
}
