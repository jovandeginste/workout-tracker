package app

import (
	"errors"
	"net/http"

	"github.com/jovandeginste/workout-tracker/pkg/database"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var ErrInvalidAPIKey = errors.New("invalid API key")

func (a *App) apiRoutes(e *echo.Group) {
	apiGroup := e.Group("/api/v1")
	apiGroup.Use(middleware.KeyAuth(func(key string, c echo.Context) (bool, error) {
		u, err := database.GetUserByAPIKey(a.db, key)
		if err != nil {
			return false, ErrInvalidAPIKey
		}

		if !u.IsActive() {
			return false, ErrInvalidAPIKey
		}

		c.Set("user_info", u)
		c.Set("user_language", u.Profile.Language)
		c.Set("user_totals_show", u.Profile.TotalsShow)

		return true, nil
	}))

	apiGroup.GET("/whoami", a.apiWhoamiHandler).Name = "api-whoami"
}

func (a *App) apiWhoamiHandler(c echo.Context) error {
	data := a.defaultData(c)

	return c.JSON(http.StatusOK, data)
}
