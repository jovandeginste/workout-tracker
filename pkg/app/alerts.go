package app

import "github.com/labstack/echo/v4"

func (a *App) addError(data map[string]interface{}, c echo.Context) {
	data["error"] = a.sessionManager.PopString(c.Request().Context(), "error")
}

func (a *App) addNotice(data map[string]interface{}, c echo.Context) {
	data["notice"] = a.sessionManager.PopString(c.Request().Context(), "notice")
}

func (a *App) setNotice(c echo.Context, msg string) {
	if msg == "" {
		return
	}

	a.sessionManager.Put(c.Request().Context(), "notice", msg)
}

func (a *App) setError(c echo.Context, err string) {
	if err == "" {
		return
	}

	a.sessionManager.Put(c.Request().Context(), "error", err)
}
