package nl

import (
	"embed"

	"golang.org/x/text/language"

	"github.com/vorlif/spreak/humanize"
)

//go:embed *.po
var fsys embed.FS

func New() *humanize.LocaleData {
	return &humanize.LocaleData{
		Lang: language.MustParse("nl"),
		Fs:   fsys,
		Format: &humanize.FormatData{
			DateFormat:          "j F Y",
			TimeFormat:          "H:i",
			DateTimeFormat:      "j F Y H:i",
			YearMonthFormat:     "F Y",
			MonthDayFormat:      "j F",
			ShortDateFormat:     "j-n-Y",
			ShortDatetimeFormat: "j-n-Y H:i",
			FirstDayOfWeek:      1,
		},
	}
}
