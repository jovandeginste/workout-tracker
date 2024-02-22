package app

import (
	"html/template"
	"io/fs"
	"strings"

	"github.com/dustin/go-humanize"
	"github.com/jovandeginste/workout-tracker/pkg/templatehelpers"
	"github.com/labstack/gommon/log"
)

func (a *App) viewTemplateFunctions() template.FuncMap {
	return template.FuncMap{
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

		"RelativeDate": humanize.Time,

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
