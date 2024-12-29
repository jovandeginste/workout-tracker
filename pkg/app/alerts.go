package app

import (
	"github.com/labstack/echo/v4"
)

func (a *App) addError(data map[string]any, c echo.Context) {
	data["error"] = a.sessionManager.PopString(c.Request().Context(), "error")
	c.Set("errors", data["error"])
}

func (a *App) addNotice(data map[string]any, c echo.Context) {
	data["notice"] = a.sessionManager.PopString(c.Request().Context(), "notice")
	c.Set("notices", data["notice"])
}

func (a *App) setNotice(c echo.Context, msg string, vars ...any) {
	if msg == "" {
		return
	}

	theMsg := a.i18n(c, msg, vars...)

	c.Set("notices", []string{theMsg})
	a.sessionManager.Put(c.Request().Context(), "notice", theMsg)
}

func (a *App) setError(c echo.Context, msg string, vars ...any) {
	if msg == "" {
		return
	}

	theMsg := a.i18n(c, msg, vars...)

	c.Set("errors", []string{theMsg})
	a.sessionManager.Put(c.Request().Context(), "error", theMsg)
}
