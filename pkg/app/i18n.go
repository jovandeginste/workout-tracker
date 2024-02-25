package app

import (
	"github.com/labstack/echo/v4"
	"github.com/vorlif/spreak"
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

func langFromContext(ctx echo.Context) []interface{} {
	return []interface{}{
		ctx.QueryParam("lang"),
		ctx.Get("user_language"),
		ctx.Request().Header.Get("Accept-Language"),
	}
}

func (a *App) i18n(ctx echo.Context, message string, vars ...interface{}) string {
	return a.translatorFromContext(ctx).Getf(message, vars...)
}

func (a *App) translatorFromContext(ctx echo.Context) *spreak.Localizer {
	return spreak.NewLocalizer(a.translator, langFromContext(ctx)...)
}

func (a *App) humanizerFromContext(ctx echo.Context) *humanize.Humanizer {
	return a.humanizer.CreateHumanizer(langFromContext(ctx)...)
}
