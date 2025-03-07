package app

import (
	"fmt"
	"strings"

	"github.com/invopop/ctxi18n"
	"github.com/invopop/ctxi18n/i18n"
	"github.com/jovandeginste/workout-tracker/v2/pkg/database"
	"github.com/labstack/echo/v4"
)

const (
	BrowserLanguage   = "browser"
	BrowserTheme      = "browser"
	DefaultTotalsShow = database.WorkoutTypeRunning
)

func (a *App) ConfigureLocalizer() error {
	if err := ctxi18n.LoadWithDefault(a.Translations, "en"); err != nil {
		return err
	}

	a.translator = ctxi18n.Match(string(ctxi18n.DefaultLocale))

	return nil
}

func langFromContextString(ctx echo.Context) string {
	langs := langFromContext(ctx)
	res := []string{}

	for _, lang := range langs {
		if l, ok := lang.(string); ok {
			if l != "" {
				res = append(res, lang.(string))
			}
		}
	}

	return strings.Join(res, ";")
}

func langFromContext(ctx echo.Context) []any {
	return []any{
		ctx.QueryParam("lang"),
		ctx.Get("user_language"),
		ctx.Request().Header.Get("Accept-Language"),
	}
}

func (a *App) i18n(ctx echo.Context, message string, vars ...any) string {
	t := a.translatorFromContext(ctx)
	if t.Has(message) {
		return t.T(message, vars...)
	}

	return fmt.Sprintf(message, vars...)
}

func (a *App) translatorFromContext(ctx echo.Context) *i18n.Locale {
	if l := ctxi18n.Locale(ctx.Request().Context()); l != nil {
		return l
	}

	return a.translator
}
