package timeago

import (
	"math"
	"strconv"
	"strings"
	"time"
)

type TimeAgo struct {
	Time       time.Time
	RelativeTo time.Time
	Locale     string
}

func New(t time.Time) *TimeAgo {
	n := new(TimeAgo)
	n.Time = t
	n.RelativeTo = time.Now()
	n.Locale = "en"
	return n
}

func (t *TimeAgo) WithLocale(locale string) *TimeAgo {
	t.Locale = locale
	return t
}

func (t *TimeAgo) WithRelativeTime(tt time.Time) *TimeAgo {
	t.RelativeTo = tt
	return t
}

func (t TimeAgo) Format() string {
	var localeFunc LocaleFunction
	var ok bool
	// if locale not exists, use default en locale
	if localeFunc, ok = localeFunctions[t.Locale]; !ok {
		localeFunc = localeFunctions["en"]
	}
	diff := -t.Time.Sub(t.RelativeTo).Round(time.Second).Seconds()
	agoIn := diff < 0
	diff = math.Abs(diff)
	// Unit of time
	var idx = 0
	for ; idx < len(secArray) && diff >= secArray[idx]; idx++ {
		diff /= secArray[idx]
	}
	diff = math.Floor(diff)
	idx *= 2
	if idx == 0 {
		if diff > 9 {
			idx += 1
		}
	} else {
		if diff > 1 {
			idx += 1
		}
	}
	ago, in := localeFunc(diff, idx)
	if agoIn {
		return strings.ReplaceAll(in, "%d", strconv.FormatInt(int64(diff), 10))
	} else {
		return strings.ReplaceAll(ago, "%d", strconv.FormatInt(int64(diff), 10))
	}
}

func RegisterLocale(localeName string, f LocaleFunction) {
	localeFunctions[localeName] = f
}
