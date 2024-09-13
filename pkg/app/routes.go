package app

import (
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

	e.Renderer = &Template{
		app:       a,
		templates: a.parseViewTemplates(),
	}

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

func (a *App) ValidateUserMiddleware(ctx echo.Context) {
	if err := a.setUser(ctx); err != nil {
		log.Warn(err.Error())
	}

	u := a.getCurrentUser(ctx)
	if u.IsAnonymous() {
		panic("User is not found")
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
		SuccessHandler: a.ValidateUserMiddleware,
	}))

	secureGroup.GET("/", a.dashboardHandler).Name = "dashboard"
	secureGroup.GET("/statistics", a.statisticsHandler).Name = "statistics"
	secureGroup.POST("/lookup-address", a.lookupAddressHandler).Name = "lookup-address"

	a.addRoutesSelf(secureGroup)
	a.addRoutesUsers(secureGroup)
	a.addRoutesEquipment(secureGroup)
	a.addRoutesWorkouts(secureGroup)
	a.addRoutesSegments(secureGroup)

	return secureGroup
}
