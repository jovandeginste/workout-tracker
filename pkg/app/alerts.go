package app

import (
	"github.com/labstack/echo/v4"
)

func (a *App) addError(c echo.Context, msg string, vars ...any) {
	switch msges := c.Get("errors").(type) {
	case []string:
		c.Set("errors", append(msges, a.i18n(c, msg, vars...)))
	default:
		c.Set("errors", []string{a.i18n(c, msg, vars...)})
	}
}

func (a *App) addNotice(c echo.Context, msg string, vars ...any) {
	switch msges := c.Get("notices").(type) {
	case []string:
		c.Set("notices", append(msges, a.i18n(c, msg, vars...)))
	default:
		c.Set("notices", []string{a.i18n(c, msg, vars...)})
	}
}
