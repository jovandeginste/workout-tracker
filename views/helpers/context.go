package helpers

import (
	"context"

	"github.com/jovandeginste/workout-tracker/pkg/database"
	appversion "github.com/jovandeginste/workout-tracker/pkg/version"
	"github.com/labstack/echo/v4"
	"github.com/vorlif/spreak"
	"github.com/vorlif/spreak/humanize"
)

func AppConfig(ctx context.Context) *database.Config {
	if v := ctx.Value("config"); v != nil {
		return v.(*database.Config)
	}

	return nil
}

func Version(ctx context.Context) *appversion.Version {
	if v := ctx.Value("version"); v != nil {
		return v.(*appversion.Version)
	}

	return nil
}

func Notices(ctx context.Context) []string {
	e := ctx.Value("notices")
	if e == nil {
		return nil
	}

	switch v := e.(type) {
	case string:
		if len(v) == 0 {
			return nil
		}

		return []string{v}
	case []string:
		return v
	default:
		return nil
	}
}

func Errors(ctx context.Context) []string {
	e := ctx.Value("errors")
	if e == nil {
		return nil
	}

	switch v := e.(type) {
	case string:
		if len(v) == 0 {
			return nil
		}

		return []string{v}
	case []string:
		return v
	default:
		return nil
	}
}

func appEcho(ctx context.Context) *echo.Echo {
	if e := ctx.Value("echo"); e != nil {
		return e.(*echo.Echo)
	}

	return nil
}

func translator(ctx context.Context) *spreak.Localizer {
	if t := CurrentUser(ctx).GetTranslator(); t != nil {
		return t
	}

	if t := ctx.Value("translator"); t != nil {
		return t.(*spreak.Localizer)
	}

	return nil
}

func genericTranslator(ctx context.Context) *spreak.Bundle {
	if t := ctx.Value("generic_translator"); t != nil {
		return t.(*spreak.Bundle)
	}

	return nil
}

func humanizer(ctx context.Context) *humanize.Humanizer {
	if h := CurrentUser(ctx).GetHumanizer(); h != nil {
		return h
	}

	if h := ctx.Value("humanizer"); h != nil {
		return h.(*humanize.Humanizer)
	}

	return nil
}

func CurrentUser(ctx context.Context) *database.User {
	if dbUser := ctx.Value("user_info"); dbUser != nil {
		return dbUser.(*database.User)
	}

	return database.AnonymousUser()
}
