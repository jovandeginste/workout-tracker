package app

import (
	"context"
	"log/slog"
	"net/http"
	"path"
	"strings"
	"time"

	"github.com/alexedwards/scs/gormstore"
	"github.com/alexedwards/scs/v2"
	"github.com/invopop/ctxi18n"
	echojwt "github.com/labstack/echo-jwt/v5"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"github.com/labstack/gommon/log"
)

func (a *App) WebRoot() string {
	root := path.Join("/", a.Config.WebRoot)
	return strings.TrimSuffix(root, "/")
}

func slogLogger(logger *slog.Logger) echo.MiddlewareFunc {
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogRemoteIP: true,
		LogMethod:   true,
		LogStatus:   true,
		LogURI:      true,
		LogLatency:  true,
		HandleError: true, // forwards the error to the global error handler so it can pick the status code

		LogValuesFunc: func(c *echo.Context, v middleware.RequestLoggerValues) error {
			l := slog.LevelInfo
			n := "REQUEST"
			a := []slog.Attr{
				slog.Int("status", v.Status),
				slog.Duration("latency", v.Latency),
				slog.String("uri", v.URI),
				slog.String("method", v.Method),
				slog.String("remote_ip", v.RemoteIP),
			}

			if v.Error != nil {
				l = slog.LevelError
				n = "REQUEST_ERROR"
				a = append(
					a,
					slog.String("err", v.Error.Error()),
				)
			}

			logger.LogAttrs(context.Background(), l, n, a...)

			return nil
		},
	})
}

func sessionLoadAndSave(sm *scs.SessionManager) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			var err error
			sm.LoadAndSave(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				c.SetRequest(r)
				err = next(c)
			})).ServeHTTP(c.Response(), c.Request())
			return err
		}
	}
}

func newEcho(logger *slog.Logger) *echo.Echo {
	e := echo.New()

	e.Use(slogLogger(logger.With("module", "webserver")))
	e.Use(middleware.Recover())
	e.Use(middleware.Secure())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))
	e.Pre(middleware.RemoveTrailingSlash())
	e.Pre(middleware.MethodOverrideWithConfig(middleware.MethodOverrideConfig{
		Getter: middleware.MethodFromHeader(echo.HeaderXHTTPMethodOverride),
	}))

	return e
}

func (a *App) ConfigureWebserver() error {
	var err error

	e := newEcho(a.rawLogger)

	a.sessionManager = scs.New()
	a.sessionManager.Cookie.Path = "/"
	a.sessionManager.Cookie.HttpOnly = true
	a.sessionManager.Lifetime = 24 * time.Hour

	if a.sessionManager.Store, err = gormstore.New(a.db); err != nil {
		return err
	}

	if a.Config.Compress {
		e.Use(middleware.Gzip())
	}

	e.Use(sessionLoadAndSave(a.sessionManager))
	e.Use(func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
		return func(context *echo.Context) error {
			a.setContext(context)
			return handlerFunc(context)
		}
	})

	publicGroup := e.Group(a.WebRoot())

	a.GET(publicGroup, "/health", func(c *echo.Context) error {
		return c.String(http.StatusOK, "OK")
	}, "health")

	a.apiRoutes(publicGroup)

	if a.AssetDir != "" {
		publicGroup.Static("/assets", a.AssetDir)
	} else {
		publicGroup.StaticFS("/assets", a.Assets)
	}

	a.GET(publicGroup, "/assets", func(c *echo.Context) error {
		return c.Redirect(http.StatusFound, a.Reverse("dashboard"))
	}, "assets")
	a.GET(publicGroup, "/share/:uuid", a.workoutShowShared, "share")

	userGroup := publicGroup.Group("/user")
	a.GET(userGroup, "/signin", a.userLoginHandler, "user-login")
	a.POST(userGroup, "/signin", a.userSigninHandler, "user-signin")
	a.POST(userGroup, "/register", a.userRegisterHandler, "user-register")
	a.GET(userGroup, "/signout", a.userSignoutHandler, "user-signout")

	sec := a.addRoutesSecure(publicGroup)
	a.adminRoutes(sec)

	a.echo = e

	return nil
}

func (a *App) ValidateAdminMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx *echo.Context) error {
		u := a.getCurrentUser(ctx)
		if u.IsAnonymous() || !u.IsActive() {
			log.Warn("User is not found")
			return ctx.Redirect(http.StatusFound, a.Reverse("user-signout"))
		}

		if !u.Admin {
			log.Warn("User is not an admin")
			return ctx.Redirect(http.StatusFound, a.Reverse("dashboard"))
		}

		return next(ctx)
	}
}

func (a *App) ValidateUserMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx *echo.Context) error {
		if err := a.setUserFromContext(ctx); err != nil {
			a.logger.Warn("error validating user", "error", err.Error())
			return ctx.Redirect(http.StatusFound, a.Reverse("user-signout"))
		}

		lctx, _ := ctxi18n.WithLocale(ctx.Request().Context(), a.langFromContextString(ctx))
		if lctx != nil {
			ctx.SetRequest(ctx.Request().WithContext(lctx))
		}

		return next(ctx)
	}
}

func (a *App) addRoutesSecure(g *echo.Group) *echo.Group {
	secureGroup := g.Group("")

	secureGroup.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  a.jwtSecret(),
		TokenLookup: "cookie:token",
		ErrorHandler: func(c *echo.Context, err error) error {
			log.Warn(err.Error())
			return c.Redirect(http.StatusFound, a.Reverse("user-signout"))
		},
	}))
	secureGroup.Use(a.ValidateUserMiddleware)

	dashboardPath := "/"
	if a.WebRoot() != "" {
		// Match the prefixed root after RemoveTrailingSlash has run.
		dashboardPath = ""
	}
	a.GET(secureGroup, dashboardPath, a.dashboardHandler, "dashboard")
	a.GET(secureGroup, "/daily", a.dailyHandler, "daily")
	a.POST(secureGroup, "/daily", a.dailyUpdateHandler, "daily-update")
	a.DELETE(secureGroup, "/daily/:date", a.dailyDeleteHandler, "daily-delete")
	a.GET(secureGroup, "/statistics", a.statisticsHandler, "statistics")
	a.GET(secureGroup, "/heatmap", a.heatmapHandler, "heatmap")
	a.POST(secureGroup, "/lookup-address", a.lookupAddressHandler, "lookup-address")

	a.addRoutesSelf(secureGroup)
	a.addRoutesUsers(secureGroup)
	a.addRoutesEquipment(secureGroup)
	a.addRoutesWorkouts(secureGroup)
	a.addRoutesSegments(secureGroup)

	return secureGroup
}

func (a *App) GET(g *echo.Group, p string, h echo.HandlerFunc, name string, m ...echo.MiddlewareFunc) {
	_, _ = g.AddRoute(echo.Route{Method: http.MethodGet, Path: p, Handler: h, Name: name, Middlewares: m})
}

func (a *App) POST(g *echo.Group, p string, h echo.HandlerFunc, name string, m ...echo.MiddlewareFunc) {
	_, _ = g.AddRoute(echo.Route{Method: http.MethodPost, Path: p, Handler: h, Name: name, Middlewares: m})
}

func (a *App) DELETE(g *echo.Group, p string, h echo.HandlerFunc, name string, m ...echo.MiddlewareFunc) {
	_, _ = g.AddRoute(echo.Route{Method: http.MethodDelete, Path: p, Handler: h, Name: name, Middlewares: m})
}

func isHtmx(c *echo.Context) bool {
	return c.Request().Header.Get("HX-Request") == "true"
}
