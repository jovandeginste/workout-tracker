package helpers

import (
	"cmp"
	"context"
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/invopop/ctxi18n/i18n"
	emojiflag "github.com/jayco/go-emoji-flag"
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
		language.Russian,
	}
)

func THas(ctx context.Context, key string, args ...any) string {
	if i18n.Has(ctx, key) {
		return i18n.T(ctx, key, args...)
	}

	return fmt.Sprintf(key, args...)
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
	cc := code.String()
	if strings.Contains(cc, "-") {
		cc = strings.Split(cc, "-")[1]
	}

	if cc == "en" {
		cc = "us"
	}

	l := LanguageInformation{
		Code: code.String(),
		Flag: emojiflag.GetFlag(cc),
	}

	if l.Flag == "" {
		l.Flag = "👽"
	}

	localTag := language.MustParse(code.String())
	l.LocalName = display.Self.Name(localTag)
	l.EnglishName = englishTag.Name(localTag)

	return l
}

func RelativeDate(ctx context.Context, t *time.Time) string {
	return timeago.New(*t).WithLocale(Language(ctx)).Format()
}
