package app

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/jovandeginste/workout-tracker/pkg/database"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var ErrInvalidAPIKey = errors.New("invalid API key")

type APIResponse struct {
	Errors  []string    `json:"errors"`
	Results interface{} `json:"results"`
}

func (a *App) apiRoutes(e *echo.Group) {
	apiGroup := e.Group("/api/v1")
	apiGroup.Use(a.ValidateAPIKeyMiddleware())

	apiGroup.GET("/whoami", a.apiWhoamiHandler).Name = "api-whoami"
	apiGroup.GET("/workouts", a.apiWorkoutsHandler).Name = "api-workouts"
	apiGroup.GET("/workout/:id", a.apiWorkoutHandler).Name = "api-workout"
}

func (a *App) ValidateAPIKeyMiddleware() echo.MiddlewareFunc {
	return middleware.KeyAuth(func(key string, c echo.Context) (bool, error) {
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
	})
}

func (a *App) apiWhoamiHandler(c echo.Context) error {
	data := a.defaultData(c)

	return c.JSON(http.StatusOK, data)
}

func (a *App) apiWorkoutsHandler(c echo.Context) error {
	resp := APIResponse{}

	w, err := a.getCurrentUser(c).GetWorkouts(a.db)
	if err != nil {
		resp.Errors = append(resp.Errors, err.Error())
	}

	resp.Results = w

	return c.JSON(http.StatusOK, resp)
}

func (a *App) apiWorkoutHandler(c echo.Context) error {
	resp := APIResponse{}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return a.redirectWithError(c, "/workouts", err)
	}

	w, err := a.getCurrentUser(c).GetWorkout(a.db, id)
	if err != nil {
		resp.Errors = append(resp.Errors, err.Error())
	}

	resp.Results = w

	return c.JSON(http.StatusOK, resp)
}
