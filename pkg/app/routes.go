package app

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/alexedwards/scs/gormstore"
	"github.com/alexedwards/scs/v2"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"

	slogecho "github.com/samber/slog-echo"

	session "github.com/spazzymoto/echo-scs-session"
)

func newEcho(logger *slog.Logger) *echo.Echo {
	e := echo.New()

	e.HideBanner = true
	e.HidePort = true

	e.Use(slogecho.New(logger.With("module", "webserver")))
	e.Use(middleware.Recover())
	e.Use(middleware.Secure())
	e.Use(middleware.CORS())
	e.Use(middleware.Gzip())
	e.Pre(middleware.RemoveTrailingSlash())

	return e
}

func (a *App) ConfigureWebserver() error {
	var err error

	e := newEcho(a.rawLogger)
	e.Debug = a.Config.Debug

	a.sessionManager = scs.New()
	a.sessionManager.Cookie.Path = "/"
	a.sessionManager.Cookie.HttpOnly = true
	a.sessionManager.Lifetime = 24 * time.Hour

	if a.sessionManager.Store, err = gormstore.New(a.db); err != nil {
		return err
	}

	e.Use(session.LoadAndSave(a.sessionManager))
	e.Use(a.ContextValueMiddleware)
	e.Use(func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
		return func(context echo.Context) error {
			a.setContext(context)
			return handlerFunc(context)
		}
	})

	publicGroup := e.Group("")

	a.apiRoutes(publicGroup)

	publicGroup.StaticFS("/assets", a.Assets)

	publicGroup.GET("/assets", func(c echo.Context) error {
		return c.Redirect(http.StatusFound, a.echo.Reverse("dashboard"))
	}).Name = "assets"
	publicGroup.GET("/share/:uuid", a.workoutShowShared).Name = "share"

	userGroup := publicGroup.Group("/user")
	userGroup.GET("/signin", a.userLoginHandler).Name = "user-login"
	userGroup.POST("/signin", a.userSigninHandler).Name = "user-signin"
	userGroup.POST("/register", a.userRegisterHandler).Name = "user-register"
	userGroup.GET("/signout", a.userSignoutHandler).Name = "user-signout"

	sec := a.addRoutesSecure(publicGroup)
	a.adminRoutes(sec)

	a.echo = e

	return nil
}

func (a *App) ValidateAdminMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		u := a.getCurrentUser(ctx)
		if u.IsAnonymous() || !u.IsActive() {
			log.Warn("User is not found")
			return ctx.Redirect(http.StatusFound, a.echo.Reverse("user-signout"))
		}

		if !u.Admin {
			log.Warn("User is not an admin")
			return ctx.Redirect(http.StatusFound, a.echo.Reverse("dashboard"))
		}

		return next(ctx)
	}
}

func (a *App) ValidateUserMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		if err := a.setUserFromContext(ctx); err != nil {
			a.logger.Warn("error validating user", "error", err.Error())
			return ctx.Redirect(http.StatusFound, a.echo.Reverse("user-signout"))
		}

		return next(ctx)
	}
}

func (a *App) addRoutesSecure(e *echo.Group) *echo.Group {
	secureGroup := e.Group("")

	secureGroup.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  a.jwtSecret(),
		TokenLookup: "cookie:token",
		ErrorHandler: func(c echo.Context, err error) error {
			log.Warn(err.Error())
			return c.Redirect(http.StatusFound, a.echo.Reverse("user-signout"))
		},
	}))
	secureGroup.Use(a.ValidateUserMiddleware)

	secureGroup.GET("/", a.dashboardHandler).Name = "dashboard"
	secureGroup.GET("/daily", a.dailyHandler).Name = "daily"
	secureGroup.POST("/daily", a.dailyUpdateHandler).Name = "daily-update"
	secureGroup.DELETE("/daily/:date", a.dailyDeleteHandler).Name = "daily-delete"
	secureGroup.GET("/statistics", a.statisticsHandler).Name = "statistics"
	secureGroup.GET("/heatmap", a.heatmapHandler).Name = "heatmap"
	secureGroup.POST("/lookup-address", a.lookupAddressHandler).Name = "lookup-address"

	a.addRoutesSelf(secureGroup)
	a.addRoutesUsers(secureGroup)
	a.addRoutesEquipment(secureGroup)
	a.addRoutesWorkouts(secureGroup)
	a.addRoutesSegments(secureGroup)

	return secureGroup
}

// extend echo.Context
type contextValue struct {
	echo.Context
}

func (c contextValue) Get(key string) any {
	if val := c.Context.Get(key); val != nil {
		return val
	}

	return c.Request().Context().Value(key)
}

func (c contextValue) Set(key string, val any) {
	// we're replacing the whole Request in echo.Context
	// with a copied request that has the updated context value
	c.SetRequest(
		c.Request().WithContext(
			context.WithValue(c.Request().Context(), key, val), //nolint:staticcheck
		),
	)
	c.Context.Set(key, val)
}

func (a *App) ContextValueMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// instead of passing next(c) as you usually would,
		// you return it with the extended version
		return next(contextValue{c})
	}
}
