package app

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/labstack/gommon/log"
)

func viewTemplateFunctions() template.FuncMap {
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

func parseViewTemplates() *template.Template {
	templ := template.New("views").Funcs(viewTemplateFunctions())

	err := filepath.Walk("./views", func(path string, _ os.FileInfo, err error) error {
		if strings.Contains(path, ".html") {
			if _, myErr := templ.ParseFiles(path); err != nil {
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
