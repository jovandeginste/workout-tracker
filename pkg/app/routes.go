package app

import (
	"html/template"
	"io"
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

func newEcho() *echo.Echo {
	e := echo.New()

	e.Debug = true
	e.HideBanner = true
	e.HidePort = true

	e.Use(middleware.Recover())
	e.Use(middleware.Secure())
	e.Use(middleware.CORS())
	e.Use(middleware.Gzip())

	return e
}

func (a *App) Configure() error {
	err := a.Connect()
	if err != nil {
		return err
	}

	e := newEcho()
	e.Use(slogecho.New(a.log))

	a.sessionManager = scs.New()
	a.sessionManager.Cookie.Path = "/"
	a.sessionManager.Cookie.HttpOnly = true
	a.sessionManager.Lifetime = 24 * time.Hour

	if a.sessionManager.Store, err = gormstore.New(a.db); err != nil {
		return err
	}

	e.Use(session.LoadAndSave(a.sessionManager))

	e.Renderer = &Template{parseViewTemplates()}

	e.Static("/assets", "assets")
	e.GET("/user/signin", a.loginHandler)
	e.POST("/user/signin", a.SignIn)
	e.POST("/user/register", a.Register)
	e.GET("/user/signout", a.SignOut)

	a.addSecureRoutes(e)
	a.addAdminRoutes(e)

	a.echo = e

	return nil
}

func (a *App) ValidateAdminMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		u := a.getUser(ctx)
		if u == nil || !u.IsActive() {
			log.Warn("User is not found")
			return ctx.Redirect(http.StatusFound, "/user/signout")
		}

		if !u.Admin {
			log.Warn("User is not an admin")
			return ctx.Redirect(http.StatusFound, "/")
		}

		return next(ctx)
	}
}

func (a *App) ValidateUserMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		if err := a.setUser(ctx); err != nil {
			log.Warn(err.Error())
			return ctx.Redirect(http.StatusFound, "/user/signout")
		}

		return next(ctx)
	}
}

func (a *App) addSecureRoutes(e *echo.Echo) {
	secureGroup := e.Group("")

	secureGroup.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  a.jwtSecret,
		TokenLookup: "cookie:token",
		ErrorHandler: func(c echo.Context, err error) error {
			log.Warn(err.Error())
			return c.Redirect(http.StatusFound, "/user/signout")
		},
	}))
	secureGroup.Use(a.ValidateUserMiddleware)

	secureGroup.GET("/", a.dashboardHandler)
	secureGroup.GET("/workouts", a.workoutsHandler)
	secureGroup.GET("/workouts/:id", a.workoutsShowHandler)
	secureGroup.GET("/workouts/edit/:id", a.workoutsEditHandler)
	secureGroup.POST("/workouts/:id", a.workoutsPostHandler)
	secureGroup.GET("/workouts/add", a.workoutsAddHandler)
	secureGroup.GET("/user/profile", a.userProfileHandler)
	secureGroup.POST("/workouts/add", a.addWorkout)
}

func (a *App) addAdminRoutes(e *echo.Echo) {
	adminGroup := e.Group("/admin")
	adminGroup.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  a.jwtSecret,
		TokenLookup: "cookie:token",
		ErrorHandler: func(c echo.Context, err error) error {
			log.Warn(err.Error())
			return c.Redirect(http.StatusFound, "/user/signout")
		},
	}))
	adminGroup.Use(a.ValidateUserMiddleware)
	adminGroup.Use(a.ValidateAdminMiddleware)

	adminGroup.GET("/", a.adminRootHandler)
}

func (a *App) Serve() error {
	return a.echo.Start(":8080")
}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, _ echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
