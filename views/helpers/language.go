package helpers

import (
	"context"
	"strings"
	"time"

	emojiflag "github.com/jayco/go-emoji-flag"
	"github.com/vorlif/spreak/humanize"
	"github.com/vorlif/spreak/humanize/locale/de"
	"github.com/vorlif/spreak/humanize/locale/fa"
	"github.com/vorlif/spreak/humanize/locale/fr"
	"github.com/vorlif/spreak/humanize/locale/id"
	"github.com/vorlif/spreak/humanize/locale/it"
	"github.com/vorlif/spreak/humanize/locale/nb"
	"github.com/vorlif/spreak/humanize/locale/nl"
	"github.com/vorlif/spreak/humanize/locale/ru"
	"golang.org/x/text/language"
	"golang.org/x/text/language/display"
)

var (
	englishTag   = display.English.Languages()
	humanLocales = []*humanize.LocaleData{
		de.New(),
		fa.New(),
		fr.New(),
		id.New(),
		it.New(),
		nb.New(),
		nl.New(),
		ru.New(),
	}

	humanizer = humanize.MustNew(humanize.WithLocale(humanLocales...))
)

type LanguageInformation struct {
	Code        string
	EnglishName string
	LocalName   string
	Flag        string
}

func Language(ctx context.Context) string {
	return translator(ctx).Code().String()
}

func SupportedLanguages(ctx context.Context) []language.Tag {
	return []language.Tag{
		language.Dutch,
		language.English,
		language.French,
		language.German,
		language.Indonesian,
		language.Italian,
		language.Norwegian,
		language.Persian,
		language.Russian,
	}
}

func ToLanguageInformation(code string) LanguageInformation {
	cc := code
	if strings.Contains(cc, "-") {
		cc = strings.Split(cc, "-")[1]
	}

	if cc == "en" {
		cc = "us"
	}

	l := LanguageInformation{
		Code: code,
		Flag: emojiflag.GetFlag(cc),
	}

	if l.Flag == "" {
		l.Flag = "👽"
	}

	localTag := language.MustParse(code)
	l.LocalName = display.Self.Name(localTag)
	l.EnglishName = englishTag.Name(localTag)

	return l
}

func RelativeDate(ctx context.Context, t *time.Time) string {
	return humanizer.CreateHumanizer(Language(ctx)).NaturalTime(t)
}
