package app

import (
	"github.com/labstack/echo/v4"
)

func (a *App) addErrorN(c echo.Context, msg string, count int, vars ...any) {
	c.Logger().Errorf("%s[%d]: %v", msg, count, vars)
	a.addError(c, a.i18nN(c, msg, count, vars...))
}

func (a *App) addErrorT(c echo.Context, msg string, vars ...any) {
	c.Logger().Errorf("%s: %v", msg, vars)
	a.addError(c, a.i18nT(c, msg, vars...))
}

func (a *App) addError(c echo.Context, e string) {
	var v []string

	switch msges := a.sessionManager.Get(c.Request().Context(), "errors").(type) {
	case []string:
		v = msges
		v = append(v, e)
	default:
		v = []string{e}
	}

	a.sessionManager.Put(c.Request().Context(), "errors", v)
}

func (a *App) addNoticeNRaw(c echo.Context, msg string, count int, vars ...any) {
	// TODO: Convert the notification system to structs instead of strings
	a.addNotice(c, "<!-- raw -->"+a.i18nN(c, msg, count, vars...))
}

func (a *App) addNoticeN(c echo.Context, msg string, count int, vars ...any) {
	a.addNotice(c, a.i18nN(c, msg, count, vars...))
}

func (a *App) addNoticeT(c echo.Context, msg string, vars ...any) {
	a.addNotice(c, a.i18nT(c, msg, vars...))
}

func (a *App) addNotice(c echo.Context, e string) {
	var v []string

	switch msges := a.sessionManager.Get(c.Request().Context(), "notices").(type) {
	case []string:
		v = msges
		v = append(v, e)
	default:
		v = []string{e}
	}

	a.sessionManager.Put(c.Request().Context(), "notices", v)
}
