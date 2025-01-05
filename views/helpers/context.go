package helpers

import (
	"context"

	"github.com/invopop/ctxi18n"
	"github.com/invopop/ctxi18n/i18n"
	"github.com/jovandeginste/workout-tracker/pkg/database"
	appversion "github.com/jovandeginste/workout-tracker/pkg/version"
	"github.com/labstack/echo/v4"
)

func AppConfig(ctx context.Context) *database.Config {
	switch v := ctx.Value("config").(type) {
	case *database.Config:
		return v
	default:
		return nil
	}
}

func Version(ctx context.Context) *appversion.Version {
	switch v := ctx.Value("version").(type) {
	case *appversion.Version:
		return v
	default:
		return nil
	}
}

func Notices(ctx context.Context) []string {
	switch v := ctx.Value("notices").(type) {
	case string:
		if v == "" {
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
	switch v := ctx.Value("errors").(type) {
	case string:
		if v == "" {
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
	switch v := ctx.Value("echo").(type) {
	case *echo.Echo:
		return v
	default:
		return nil
	}
}

func translator(ctx context.Context) *i18n.Locale {
	if t := ctxi18n.Locale(ctx); t != nil {
		return t
	}

	return ctxi18n.Match(string(ctxi18n.DefaultLocale))
}

func CurrentUser(ctx context.Context) *database.User {
	switch v := ctx.Value("user_info").(type) {
	case *database.User:
		return v
	default:
		return database.AnonymousUser()
	}
}
