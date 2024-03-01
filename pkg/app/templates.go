package app

import (
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"strings"

	"github.com/jovandeginste/workout-tracker/pkg/database"
	"github.com/jovandeginste/workout-tracker/pkg/templatehelpers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/vorlif/spreak/humanize"
	"golang.org/x/text/language"
)

type Template struct {
	app       *App
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, ctx echo.Context) error {
	r, err := t.templates.Clone()
	if err != nil {
		return err
	}

	tr := t.app.translatorFromContext(ctx)
	h := t.app.humanizerFromContext(ctx)

	r.Funcs(template.FuncMap{
		"i18n":         tr.Getf,
		"language":     tr.Language().String,
		"humanizer":    func() *humanize.Humanizer { return h },
		"RelativeDate": h.NaturalTime,
		"CurrentUser":  func() *database.User { return t.app.getCurrentUser(ctx) },
	})

	return r.ExecuteTemplate(w, name, data)
}

func echoFunc(key string, _ ...interface{}) string {
	return key
}

func (a *App) viewTemplateFunctions() template.FuncMap {
	h := a.humanizer.CreateHumanizer(language.English)

	return template.FuncMap{
		"i18n":        echoFunc,
		"language":    func() string { return BrowserLanguage },
		"humanizer":   func() *humanize.Humanizer { return h },
		"CurrentUser": func() *database.User { return nil },

		"supportedLanguages": a.translator.SupportedLanguages,
		"workoutTypes":       database.WorkoutTypes,

		"NumericDuration":         templatehelpers.NumericDuration,
		"CountryCodeToFlag":       templatehelpers.CountryCodeToFlag,
		"LocalDate":               templatehelpers.LocalDate,
		"ToKilometer":             templatehelpers.ToKilometer,
		"HumanDistance":           templatehelpers.HumanDistance,
		"HumanSpeed":              templatehelpers.HumanSpeed,
		"HumanTempo":              templatehelpers.HumanTempo,
		"HumanDuration":           templatehelpers.HumanDuration,
		"IconFor":                 templatehelpers.IconFor,
		"BoolToHTML":              templatehelpers.BoolToHTML,
		"BoolToCheckbox":          templatehelpers.BoolToCheckbox,
		"BuildDecoratedAttribute": templatehelpers.BuildDecoratedAttribute,
		"ToLanguageInformation":   templatehelpers.ToLanguageInformation,

		"RelativeDate": h.NaturalTime,

		"RouteFor": func(name string, params ...interface{}) string {
			rev := a.echo.Reverse(name, params...)
			if rev == "" {
				return "/invalid/route/#" + name
			}

			return rev
		},
	}
}

func (a *App) parseViewTemplates() *template.Template {
	templ := template.New("").Funcs(a.viewTemplateFunctions())
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
