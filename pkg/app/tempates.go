package app

import (
	"html/template"
	"io"
	"io/fs"
	"strings"

	legHumanize "github.com/dustin/go-humanize"
	"github.com/jovandeginste/workout-tracker/pkg/database"
	"github.com/jovandeginste/workout-tracker/pkg/templatehelpers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/vorlif/spreak"
	"github.com/vorlif/spreak/humanize"
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

	viewContext, isMap := data.(map[string]interface{})
	if !isMap {
		return r.ExecuteTemplate(w, name, data)
	}

	userLang := ""

	u, ok := viewContext["currentUser"].(*database.User)
	if ok {
		userLang = u.Profile.Language
	}

	viewContext["context"] = ctx
	clientLanguages := []interface{}{
		ctx.QueryParam("lang"),
		userLang,
		ctx.Request().Header.Get("Accept-Language"),
	}

	tr := spreak.NewLocalizer(t.app.translator, clientLanguages...)
	r.Funcs(template.FuncMap{
		"i18n":     tr.Getf,
		"language": tr.Language().String,
	})

	return r.ExecuteTemplate(w, name, data)
}

func echoFunc(key string, _ ...interface{}) string {
	return key
}

func (a *App) viewTemplateFunctions() template.FuncMap {
	return template.FuncMap{
		"i18n":               echoFunc,
		"language":           func() string { return DefaultLanguage },
		"supportedLanguages": a.translator.SupportedLanguages,
		"humanizer":          func() *humanize.Collection { return a.humanizer },

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

		"RelativeDate": legHumanize.Time,

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

	err := fs.WalkDir(a.Views, ".", func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return err
		}

		if strings.HasSuffix(path, ".html") {
			if _, myErr := templ.ParseFS(a.Views, path); err != nil {
				log.Warn(myErr)
				return myErr
			}
		}

		return err
	})
	if err != nil {
		log.Warn(err)
	}

	return templ
}
