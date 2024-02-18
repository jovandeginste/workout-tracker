package app

import (
	"html/template"
	"net/http"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/labstack/echo/v4"
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

func (a *App) ValidateUserMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		if err := a.setUser(ctx); err != nil {
			log.Warn(err.Error())
			return ctx.Redirect(http.StatusMovedPermanently, "/user/signout")
		}

		return next(ctx)
	}
}
