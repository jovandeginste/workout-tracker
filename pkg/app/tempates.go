package app

import (
	"fmt"
	"html/template"
	"io/fs"
	"strings"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/labstack/gommon/log"
)

func (a *App) viewTemplateFunctions() template.FuncMap {
	return template.FuncMap{
		"LocalDate": func(t time.Time) string {
			return t.Local().Format("2006-01-02 15:04") //nolint:gosmopolitan
		},
		"RelativeDate": humanize.Time,
		"HumanDistance": func(d float64) string {
			value, prefix := humanize.ComputeSI(d)

			return fmt.Sprintf("%.2f %sm", value, prefix)
		},
		"HumanSpeed": func(mps float64) string {
			mph := mps * 3600
			value, prefix := humanize.ComputeSI(mph)

			return fmt.Sprintf("%.2f %sm/h", value, prefix)
		},
		"HumanTempo": func(mps float64) string {
			mpk := 1000000 / (mps * 60)
			value, prefix := humanize.ComputeSI(mpk)

			return fmt.Sprintf("%.2f min/%sm", value, prefix)
		},
		"FAIconName": func(wType string) string {
			if wType == "running" {
				return "person-running"
			}

			return "question"
		},
		"FAIconClass": func(_ string) string {
			return "solid"
		},
		"BoolToHTML": func(b bool) template.HTML {
			if b {
				return `<i class="text-green-500 fas fa-check"></i>`
			}

			return `<i class="text-rose-500 fas fa-times"></i>`
		},
		"BoolToCheckbox": func(b bool) template.HTML {
			if b {
				return "checked"
			}

			return ""
		},
		"RouteFor": func(name string, params ...interface{}) string {
			rev := a.echo.Reverse(name, params...)
			if rev == "" {
				return "/invalid/route/#" + name
			}

			return rev
		},
		"BuildDecoratedAttribute": func(icon, name string, value interface{}) interface{} {
			return struct {
				Icon  string
				Name  string
				Value interface{}
			}{
				Icon:  icon,
				Name:  name,
				Value: value,
			}
		},
	}
}

func (a *App) parseViewTemplates() *template.Template {
	templ := template.New("").Funcs(a.viewTemplateFunctions())

	err := fs.WalkDir(a.Views, ".", func(path string, d fs.DirEntry, err error) error {
		if strings.Contains(path, ".html") {
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
