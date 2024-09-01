package fa

import (
	"embed"

	"golang.org/x/text/language"

	"github.com/vorlif/spreak/humanize"
)

//go:embed *.po
var fsys embed.FS

func New() *humanize.LocaleData {
	return &humanize.LocaleData{
		Lang: language.MustParse("fa"),
		Fs:   fsys,
		Format: &humanize.FormatData{
			DateFormat:          "j F Y",
			TimeFormat:          "G:i",
			DateTimeFormat:      "j F Y، ساعت G:i",
			YearMonthFormat:     "F Y",
			MonthDayFormat:      "j F",
			ShortDateFormat:     "Y/n/j",
			ShortDatetimeFormat: "Y/n/j،‏ G:i",
			FirstDayOfWeek:      6,
		},
	}
}
