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

func translator(ctx context.Context) *spreak.Localizer {
	if t := CurrentUser(ctx).GetTranslator(); t != nil {
		return t
	}

	switch t := ctx.Value("translator").(type) {
	case *spreak.Localizer:
		return t
	default:
		return nil
	}
}

func genericTranslator(ctx context.Context) *spreak.Bundle {
	switch v := ctx.Value("generic_translator").(type) {
	case *spreak.Bundle:
		return v
	default:
		return nil
	}
}

func humanizer(ctx context.Context) *humanize.Humanizer {
	if h := CurrentUser(ctx).GetHumanizer(); h != nil {
		return h
	}

	switch v := ctx.Value("humanizer").(type) {
	case *humanize.Humanizer:
		return v
	default:
		return nil
	}
}

func CurrentUser(ctx context.Context) *database.User {
	switch v := ctx.Value("user_info").(type) {
	case *database.User:
		return v
	default:
		return database.AnonymousUser()
	}
}
