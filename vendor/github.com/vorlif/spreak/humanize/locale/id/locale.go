package id

import (
	"embed"

	"golang.org/x/text/language"

	"github.com/vorlif/spreak/humanize"
)

//go:embed *.po
var fsys embed.FS

func New() *humanize.LocaleData {
	return &humanize.LocaleData{
		Lang: language.MustParse("id"),
		Fs:   fsys,
		Format: &humanize.FormatData{
			DateFormat:          "j N Y",
			DateTimeFormat:      "j N Y, G.i",
			TimeFormat:          "G.i",
			YearMonthFormat:     "F Y",
			MonthDayFormat:      "j F",
			ShortDateFormat:     "d-m-Y",
			ShortDatetimeFormat: "d-m-Y G.i",
			FirstDayOfWeek:      1,
		},
	}
}
