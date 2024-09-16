package app

import (
	"github.com/jovandeginste/workout-tracker/internal/database"
	"github.com/labstack/echo/v4"
	"github.com/vorlif/spreak"
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
)

const (
	BrowserLanguage   = "browser"
	DefaultTotalsShow = database.WorkoutTypeRunning
)

func (a *App) ConfigureLocalizer() error {
	var domain spreak.FsOption

	if a.Translations != nil {
		domain = spreak.WithFs(a.Translations)
	} else {
		domain = spreak.WithPath(".")
	}

	bundle, err := spreak.NewBundle(
		// Set the language used in the program code/templates
		spreak.WithSourceLanguage(language.English),
		// Set the path from which the translations should be loaded
		spreak.WithFilesystemLoader(spreak.NoDomain, domain),
		// Specify the languages you want to load
		spreak.WithLanguage(translations()...),
	)
	if err != nil {
		return err
	}

	a.translator = bundle

	a.humanizer = humanize.MustNew(
		humanize.WithLocale(humanLocales()...),
	)

	return nil
}

func translations() []any {
	return []any{
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

func humanLocales() []*humanize.LocaleData {
	return []*humanize.LocaleData{
		de.New(),
		fa.New(),
		fr.New(),
		id.New(),
		it.New(),
		nb.New(),
		nl.New(),
		ru.New(),
	}
}

func langFromContext(ctx echo.Context) []any {
	return []any{
		ctx.QueryParam("lang"),
		ctx.Get("user_language"),
		ctx.Request().Header.Get("Accept-Language"),
	}
}

func (a *App) i18n(ctx echo.Context, message string, vars ...any) string {
	return a.translatorFromContext(ctx).Getf(message, vars...)
}

func (a *App) translatorFromContext(ctx echo.Context) *spreak.Localizer {
	return spreak.NewLocalizer(a.translator, langFromContext(ctx)...)
}

func (a *App) humanizerFromContext(ctx echo.Context) *humanize.Humanizer {
	return a.humanizer.CreateHumanizer(langFromContext(ctx)...)
}
