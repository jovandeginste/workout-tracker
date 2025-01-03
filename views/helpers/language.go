package helpers

import (
	"context"
	"strings"

	emojiflag "github.com/jayco/go-emoji-flag"
	"golang.org/x/text/language"
	lng "golang.org/x/text/language"
	"golang.org/x/text/language/display"
)

var englishTag = display.English.Languages()

type LanguageInformation struct {
	Code        string
	EnglishName string
	LocalName   string
	Flag        string
}

func Language(ctx context.Context) string {
	return translator(ctx).Language().String()
}

func SupportedLanguages(ctx context.Context) []lng.Tag {
	return genericTranslator(ctx).SupportedLanguages()
}

func I18n(ctx context.Context, msg string, params ...any) string {
	return translator(ctx).Getf(msg, params...)
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
