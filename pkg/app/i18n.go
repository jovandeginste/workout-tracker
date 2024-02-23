package app

import (
	"github.com/vorlif/spreak/humanize"
	"github.com/vorlif/spreak/humanize/locale/nl"
	"golang.org/x/text/language"
)

var (
	DefaultLanguage = "browser"

	DefaultTheme = Theme{Name: "System default", Code: "browser", Icon: "🌐"}
	DarkTheme    = Theme{Name: "Dark theme", Code: "dark", Icon: "🌑"}
)

func translations() []interface{} {
	return []interface{}{
		language.English,
		language.Dutch,
	}
}

func humanLocales() []*humanize.LocaleData {
	return []*humanize.LocaleData{
		nl.New(),
	}
}

func themes() []Theme {
	return []Theme{
		DefaultTheme,
		DarkTheme,
	}
}

type Theme struct {
	Code string
	Icon string
	Name string
}
