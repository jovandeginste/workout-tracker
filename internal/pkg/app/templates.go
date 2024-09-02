package app

import (
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"strings"
	"time"

	"github.com/Masterminds/sprig/v3"
	"github.com/jovandeginste/workout-tracker/internal/database"
	"github.com/jovandeginste/workout-tracker/internal/pkg/templatehelpers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/vorlif/spreak/humanize"
	"golang.org/x/text/language"
)

type Template struct {
	app       *App
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data any, ctx echo.Context) error {
	r, err := t.templates.Clone()
	if err != nil {
		return err
	}

	tr := t.app.translatorFromContext(ctx)
	h := t.app.humanizerFromContext(ctx)
	u := t.app.getCurrentUser(ctx)

	r.Funcs(template.FuncMap{
		"i18n":         tr.Getf,
		"language":     tr.Language().String,
		"humanizer":    func() *humanize.Humanizer { return h },
		"RelativeDate": h.NaturalTime,
		"CurrentUser":  func() *database.User { return u },
		"LocalTime":    func(t time.Time) time.Time { return t.In(u.Timezone()) },
		"LocalDate":    func(t time.Time) string { return t.In(u.Timezone()).Format("2006-01-02 15:04") },

		"HumanElevation": templatehelpers.HumanElevationFor(u.PreferredUnits().Elevation()),
		"HumanDistance":  templatehelpers.HumanDistanceFor(u.PreferredUnits().Distance()),
		"HumanSpeed":     templatehelpers.HumanSpeedFor(u.PreferredUnits().Speed()),
		"HumanTempo":     templatehelpers.HumanTempoFor(u.PreferredUnits().Distance()),
	})

	return r.ExecuteTemplate(w, name, data)
}

func echoFunc(key string, _ ...any) string {
	return key
}

func (a *App) viewTemplateFunctions() template.FuncMap {
	h := a.humanizer.CreateHumanizer(language.English)

	return template.FuncMap{
		"i18n":        echoFunc,
		"Version":     func() *Version { return &a.Version },
		"AppConfig":   func() *database.Config { return &a.Config },
		"language":    func() string { return BrowserLanguage },
		"humanizer":   func() *humanize.Humanizer { return h },
		"CurrentUser": func() *database.User { return nil },
		"LocalTime":   func(t time.Time) time.Time { return t.UTC() },
		"LocalDate":   func(t time.Time) string { return t.UTC().Format("2006-01-02 15:04") },

		"supportedLanguages":    a.translator.SupportedLanguages,
		"workoutTypes":          database.WorkoutTypes,
		"statisticSinceOptions": statisticSinceOptions,
		"statisticPerOptions":   statisticPerOptions,

		"NumericDuration":         templatehelpers.NumericDuration,
		"CountryCodeToFlag":       templatehelpers.CountryCodeToFlag,
		"HumanDuration":           templatehelpers.HumanDuration,
		"IconFor":                 templatehelpers.IconFor,
		"BoolToHTML":              templatehelpers.BoolToHTML,
		"BoolToCheckbox":          templatehelpers.BoolToCheckbox,
		"BuildDecoratedAttribute": templatehelpers.BuildDecoratedAttribute,
		"ToLanguageInformation":   templatehelpers.ToLanguageInformation,
		"Timezones":               templatehelpers.Timezones,
		"SelectIf":                templatehelpers.SelectIf,

		"HumanElevation": templatehelpers.HumanElevationM,
		"HumanDistance":  templatehelpers.HumanDistanceKM,
		"HumanSpeed":     templatehelpers.HumanSpeedKPH,
		"HumanTempo":     templatehelpers.HumanTempoKM,
		"HumanCalories":  templatehelpers.HumanCaloriesKcal,

		"RelativeDate": h.NaturalTime,

		"RouteFor": func(name string, params ...any) string {
			rev := a.echo.Reverse(name, params...)
			if rev == "" {
				return "/invalid/route/#" + name
			}

			return rev
		},
	}
}

func (a *App) parseViewTemplates() *template.Template {
	templ := template.New("").Funcs(sprig.FuncMap()).Funcs(a.viewTemplateFunctions())
	if a.Views == nil {
		return templ
	}

	err := fs.WalkDir(a.Views, ".", func(path string, d fs.DirEntry, err error) error {
		if d != nil && d.IsDir() {
			return err
		}

		if strings.HasSuffix(path, ".html") {
			if _, myErr := templ.ParseFS(a.Views, path); err != nil {
				a.logger.Warn(fmt.Sprintf("Error loading template: %v", myErr))
				return myErr
			}
		}

		return err
	})
	if err != nil {
		a.logger.Warn(fmt.Sprintf("Error loading template: %v", err))
		log.Warn(fmt.Sprintf("Error loading template: %v", err))
	}

	return templ
}

func statisticSinceOptions() []string {
	return []string{
		"3 months",
		"6 months",
		"1 year",
		"2 years",
		"5 years",
		"10 years",
	}
}

func statisticPerOptions() []string {
	return []string{
		"day",
		"7 days",
		"15 days",
		"month",
	}
}
