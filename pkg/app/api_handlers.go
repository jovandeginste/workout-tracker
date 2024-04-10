package app

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/jovandeginste/workout-tracker/pkg/database"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

var ErrInvalidAPIKey = errors.New("invalid API key")

type APIResponse struct {
	Errors  []string    `json:"errors"`
	Results interface{} `json:"results"`
}

func (a *App) apiRoutes(e *echo.Group) {
	apiGroup := e.Group("/api/v1")
	apiGroup.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  a.jwtSecret(),
		TokenLookup: "cookie:token",
		ErrorHandler: func(c echo.Context, err error) error {
			log.Warn(err.Error())
			return c.JSON(http.StatusForbidden, "Not authorized")
		},
		Skipper: func(ctx echo.Context) bool {
			return ctx.Request().Header.Get("Authorization") != ""
		},
		SuccessHandler: a.ValidateUserMiddleware,
	}))
	apiGroup.Use(a.ValidateAPIKeyMiddleware())

	apiGroup.GET("/whoami", a.apiWhoamiHandler).Name = "api-whoami"
	apiGroup.GET("/workouts", a.apiWorkoutsHandler).Name = "api-workouts"
	apiGroup.GET("/workout/:id", a.apiWorkoutHandler).Name = "api-workout"
	apiGroup.GET("/statistics", a.apiStatisticsHandler).Name = "api-statistics"
	apiGroup.GET("/totals", a.apiTotalsHandler).Name = "api-totals"
	apiGroup.GET("/records", a.apiRecordsHandler).Name = "api-records"
}

func (a *App) ValidateAPIKeyMiddleware() echo.MiddlewareFunc {
	return middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		Validator: func(key string, c echo.Context) (bool, error) {
			u, err := database.GetUserByAPIKey(a.db, key)
			if err != nil {
				return false, ErrInvalidAPIKey
			}

			if !u.IsActive() || !u.Profile.APIActive {
				return false, ErrInvalidAPIKey
			}

			c.Set("user_info", u)
			c.Set("user_language", u.Profile.Language)
			c.Set("user_totals_show", u.Profile.TotalsShow)

			return true, nil
		},
		Skipper: func(ctx echo.Context) bool {
			return ctx.Request().Header.Get("Authorization") == ""
		},
	})
}

func (a *App) apiWhoamiHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, a.getCurrentUser(c))
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

func (a *App) apiRecordsHandler(c echo.Context) error {
	resp := APIResponse{}

	var workoutType string

	if err := echo.QueryParamsBinder(c).String("type", &workoutType).BindError(); err != nil {
		return a.renderAPIError(c, resp, err)
	}

	s, err := a.getCurrentUser(c).GetRecords(a.db, database.AsWorkoutType(workoutType))
	if err != nil {
		resp.Errors = append(resp.Errors, err.Error())
	}

	resp.Results = s

	return c.JSON(http.StatusOK, resp)
}

func (a *App) apiTotalsHandler(c echo.Context) error {
	resp := APIResponse{}

	var workoutType string

	if err := echo.QueryParamsBinder(c).String("type", &workoutType).BindError(); err != nil {
		return a.renderAPIError(c, resp, err)
	}

	s, err := a.getCurrentUser(c).GetTotals(database.AsWorkoutType(workoutType))
	if err != nil {
		resp.Errors = append(resp.Errors, err.Error())
	}

	resp.Results = s

	return c.JSON(http.StatusOK, resp)
}

func (a *App) apiStatisticsHandler(c echo.Context) error {
	resp := APIResponse{}

	var statConfig database.StatConfig

	if err := c.Bind(&statConfig); err != nil {
		return a.renderAPIError(c, resp, err)
	}

	s, err := a.getCurrentUser(c).GetStatistics(statConfig)
	if err != nil {
		resp.Errors = append(resp.Errors, err.Error())
	}

	resp.Results = s

	return c.JSON(http.StatusOK, resp)
}

func (a *App) apiWorkoutHandler(c echo.Context) error {
	resp := APIResponse{}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return a.renderAPIError(c, resp, err)
	}

	details := false
	if err = echo.QueryParamsBinder(c).Bool("details", &details).BindError(); err != nil {
		return a.renderAPIError(c, resp, err)
	}

	db := a.db
	if details {
		db = db.Preload("Data.Details")
	}

	w, err := a.getCurrentUser(c).GetWorkout(db, id)
	if err != nil {
		resp.Errors = append(resp.Errors, err.Error())
	}

	resp.Results = w

	return c.JSON(http.StatusOK, resp)
}

func (a *App) renderAPIError(c echo.Context, resp APIResponse, err error) error {
	resp.Errors = append(resp.Errors, err.Error())

	return c.JSON(http.StatusBadRequest, resp)
}
