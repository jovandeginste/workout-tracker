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

	e.Renderer = &Template{template.Must(template.ParseGlob("views/*.html"))}

	e.Static("/assets", "assets")
	e.GET("/user/signin", a.loginHandler)
	e.POST("/user/signin", a.SignIn)
	e.POST("/user/register", a.Register)
	e.GET("/user/signout", a.SignOut)

	a.addSecureRoutes(e)

	a.echo = e

	return nil
}

func (a *App) addSecureRoutes(e *echo.Echo) {
	secureGroup := e.Group("")

	secureGroup.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  a.jwtSecret,
		TokenLookup: "cookie:token",
		SuccessHandler: func(c echo.Context) {
			a.setUser(c)
		},
		ErrorHandler: func(c echo.Context, err error) error {
			log.Warn(err.Error())
			return c.Redirect(http.StatusMovedPermanently, "/user/signin")
		},
	}))

	secureGroup.GET("/", a.dashboardHandler)
	secureGroup.GET("/workouts", a.workoutsHandler)
	secureGroup.GET("/workouts/statistics", a.workoutsStatisticsHandler)
	secureGroup.GET("/workouts/add", a.workoutsAddHandler)
	secureGroup.POST("/workouts/add", a.addWorkout)
	secureGroup.POST("/map", a.workoutsShowHandler)
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
