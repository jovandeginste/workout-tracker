package app

import (
	"github.com/labstack/echo/v4"
)

func (a *App) addError(c echo.Context, msg string, vars ...any) {
	var v []string

	e := a.i18n(c, msg, vars...)
	c.Logger().Errorf("%s: %v", msg, vars)

	switch msges := a.sessionManager.Get(c.Request().Context(), "errors").(type) {
	case []string:
		v = msges
		v = append(v, e)
	default:
		v = []string{e}
	}

	a.sessionManager.Put(c.Request().Context(), "errors", v)
}

func (a *App) addNotice(c echo.Context, msg string, vars ...any) {
	var v []string

	e := a.i18n(c, msg, vars...)

	switch msges := a.sessionManager.Get(c.Request().Context(), "notices").(type) {
	case []string:
		v = msges
		v = append(v, e)
	default:
		v = []string{e}
	}

	a.sessionManager.Put(c.Request().Context(), "notices", v)
}
