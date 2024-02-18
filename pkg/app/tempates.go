package app

import (
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
			return t.Local().Format("2006-01-02 15:04")
		},
		"RelativeDate": humanize.Time,
		"FAIconName": func(wType string) string {
			if wType == "running" {
				return "person-running"
			}

			return "question"
		},
		"FAIconClass": func(wType string) string {
			return "solid"
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
