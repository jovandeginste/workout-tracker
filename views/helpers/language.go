package helpers

import (
	"cmp"
	"context"
	"slices"
	"time"

	"github.com/invopop/ctxi18n/i18n"
	"github.com/jovandeginste/workout-tracker/v2/pkg/templatehelpers"
	"github.com/sersh88/timeago"
	"golang.org/x/text/language"
	"golang.org/x/text/language/display"
)

var (
	englishTag = display.English.Languages()

	languagesSorted = false
	languages       = []language.Tag{
		language.Dutch,
		language.English,
		language.Finnish,
		language.French,
		language.German,
		language.Indonesian,
		language.Italian,
		language.Norwegian,
		language.Persian,
		language.Polish,
		language.Russian,
		language.SimplifiedChinese,
	}
)

func THas(ctx context.Context, key string) string {
	if i18n.Has(ctx, key) {
		return i18n.T(ctx, key)
	}

	return key
}

type LanguageInformation struct {
	Code        string
	EnglishName string
	LocalName   string
	Flag        string
}

func Language(ctx context.Context) string {
	return translator(ctx).Code().String()
}

func SupportedLanguages() []language.Tag {
	sortLanguages()

	return languages
}

func sortLanguages() {
	if languagesSorted {
		return
	}

	slices.SortFunc(languages, func(a, b language.Tag) int {
		return cmp.Compare(a.String(), b.String())
	})

	languagesSorted = true
}

func ToLanguageInformation(code language.Tag) LanguageInformation {
	f := templatehelpers.LanguageToFlag(code.String())

	l := LanguageInformation{
		Code: code.String(),
		Flag: f,
	}

	if l.Flag == "" {
		l.Flag = "👽"
	}

	localTag := language.MustParse(code.String())
	l.LocalName = display.Self.Name(localTag)
	l.EnglishName = englishTag.Name(localTag)

	return l
}

func RelativeDate(ctx context.Context, t time.Time) string {
	return timeago.New(t).WithLocale(Language(ctx)).Format()
}
