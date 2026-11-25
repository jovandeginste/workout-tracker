package app

import (
	"fmt"
	"slices"

	"github.com/labstack/echo/v5"
)

func (a *App) addErrorN(c *echo.Context, msg string, count int, vars ...any) {
	c.Logger().Error(fmt.Sprintf("%s[%d]: %v", msg, count, vars), "remote_ip", c.RealIP())
	a.addError(c, a.i18nN(c, msg, count, vars...))
}

func (a *App) addErrorT(c *echo.Context, msg string, vars ...any) {
	c.Logger().Error(fmt.Sprintf("%s: %v", msg, vars), "remote_ip", c.RealIP())
	a.addError(c, a.i18nT(c, msg, vars...))
}

func (a *App) addError(c *echo.Context, e string) {
	a.appendSessionMsg(c, "errors", e)
}

func (a *App) appendSessionMsg(c *echo.Context, key, msg string) {
	var v []string

	switch msgs := a.sessionManager.Get(c.Request().Context(), key).(type) {
	case []string:
		v = slices.Concat(msgs, []string{msg})
	default:
		v = []string{msg}
	}

	a.sessionManager.Put(c.Request().Context(), key, v)
}

func (a *App) addNoticeNRaw(c *echo.Context, msg string, count int, vars ...any) {
	// TODO: Convert the notification system to structs instead of strings
	a.addNotice(c, "<!-- raw -->"+a.i18nN(c, msg, count, vars...))
}

func (a *App) addNoticeN(c *echo.Context, msg string, count int, vars ...any) {
	a.addNotice(c, a.i18nN(c, msg, count, vars...))
}

func (a *App) addNoticeT(c *echo.Context, msg string, vars ...any) {
	a.addNotice(c, a.i18nT(c, msg, vars...))
}

func (a *App) addNotice(c *echo.Context, e string) {
	a.appendSessionMsg(c, "notices", e)
}
