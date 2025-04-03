package app

import (
	"encoding/json"
	"errors"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/a-h/templ"
	"github.com/jovandeginste/workout-tracker/v2/pkg/database"
	"github.com/jovandeginste/workout-tracker/v2/pkg/importers"
	"github.com/jovandeginste/workout-tracker/v2/views/workouts"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	geojson "github.com/paulmach/orb/geojson"
)

var (
	ErrNotAuthorized = errors.New("not authorized")
	ErrInvalidAPIKey = errors.New("invalid API key")
	htmlConcatenizer = regexp.MustCompile(`\s*\n\s*`)
)

type APIResponse struct {
	Results any      `json:"results"`
	Errors  []string `json:"errors"`
}

func (ar *APIResponse) AddError(err ...error) {
	for _, e := range err {
		ar.Errors = append(ar.Errors, e.Error())
	}
}

// @title Workout Tracker
// @version 1.0
// @description A workout tracking web application for personal use (or family, friends), geared towards running and other GPX-based activities
// @contact.url https://github.com/jovandeginste/workout-tracker/issues
// @license.name MIT
// @license.url https://github.com/jovandeginste/workout-tracker?tab=License-1-ov-file

// @BasePath /api/v1
func (a *App) apiRoutes(e *echo.Group) {
	apiGroup := e.Group("/api/v1")
	apiGroup.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  a.jwtSecret(),
		TokenLookup: "cookie:token",
		ErrorHandler: func(c echo.Context, err error) error {
			log.Warn(err.Error())

			r := APIResponse{}
			r.AddError(err)
			r.AddError(ErrNotAuthorized)

			return c.JSON(http.StatusForbidden, r)
		},
		Skipper: func(ctx echo.Context) bool {
			if ctx.Request().Header.Get("Authorization") != "" {
				return true
			}

			if ctx.Request().URL.Query().Get("api-key") != "" {
				return true
			}

			return false
		},
		SuccessHandler: func(ctx echo.Context) {
			if err := a.setUserFromContext(ctx); err != nil {
				a.logger.Warn("error validating user", "error", err.Error())
				return
			}
		},
	}))

	apiGroup.Use(middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		Validator: a.ValidateAPIKeyMiddleware,
		KeyLookup: "query:api-key",
		Skipper: func(ctx echo.Context) bool {
			return ctx.Request().URL.Query().Get("api-key") == ""
		},
	}))
	apiGroup.Use(middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		Validator: a.ValidateAPIKeyMiddleware,
		Skipper: func(ctx echo.Context) bool {
			return ctx.Request().Header.Get("Authorization") == ""
		},
	}))

	apiGroup.GET("/whoami", a.apiWhoamiHandler).Name = "api-whoami"
	apiGroup.GET("/daily", a.apiDailyHandler).Name = "api-daily"
	apiGroup.POST("/daily", a.apiDailyUpdateHandler).Name = "api-daily-update"
	apiGroup.GET("/workouts", a.apiWorkoutsHandler).Name = "api-workouts"
	apiGroup.POST("/workouts", a.apiWorkoutAddHandler).Name = "api-workout-add"
	apiGroup.GET("/workouts/:id", a.apiWorkoutHandler).Name = "api-workout"
	apiGroup.GET("/workouts/:id/breakdown", a.apiWorkoutBreakdownHandler).Name = "api-workout-breakdown"
	apiGroup.GET("/workouts/coordinates", a.apiCoordinates).Name = "api-workouts-coordinates"
	apiGroup.GET("/workouts/centers", a.apiCenters).Name = "api-workouts-centers"
	apiGroup.GET("/workouts/calendar", a.apiCalendar).Name = "api-workouts-calendar"
	apiGroup.GET("/statistics", a.apiStatisticsHandler).Name = "api-statistics"
	apiGroup.GET("/totals", a.apiTotalsHandler).Name = "api-totals"
	apiGroup.GET("/records", a.apiRecordsHandler).Name = "api-records"
	apiGroup.POST("/import/:program", a.apiImportHandler).Name = "api-import"
}

func (a *App) ValidateAPIKeyMiddleware(key string, c echo.Context) (bool, error) {
	u, err := database.GetUserByAPIKey(a.db, key)
	if err != nil {
		return false, ErrInvalidAPIKey
	}

	if !u.IsActive() || !u.Profile.APIActive {
		return false, ErrInvalidAPIKey
	}

	c.Set("user_info", u)
	c.Set("user_language", u.Profile.Language)

	return true, nil
}

// apiWhoamiHandler shows current user's information
// @Summary      Show the information of the owner of the API key
// @Produce      json
// @Success      200  {object}  APIResponse{results=database.User}
// @Failure      400  {object}  APIResponse
// @Failure      404  {object}  APIResponse
// @Failure      500  {object}  APIResponse
// @Router       /whoami [get]
func (a *App) apiWhoamiHandler(c echo.Context) error {
	user := a.getCurrentUser(c)
	return c.JSON(http.StatusOK, struct {
		database.UserData
		Profile database.Profile `json:"profile"`
	}{
		UserData: user.UserData,
		Profile:  user.Profile,
	})
}

// apiWorkoutsHandler lists current user's workouts
// @Summary      List all workouts of the current user
// @Produce      json
// @Success      200  {object}  APIResponse{results=[]database.Workout}
// @Failure      400  {object}  APIResponse
// @Failure      404  {object}  APIResponse
// @Failure      500  {object}  APIResponse
// @Router       /workouts [get]
func (a *App) apiWorkoutsHandler(c echo.Context) error {
	resp := APIResponse{}

	w, err := a.getCurrentUser(c).GetWorkouts(a.db)
	if err != nil {
		resp.AddError(err)
	}

	resp.Results = w

	return c.JSON(http.StatusOK, resp)
}

// apiCenters returns the center of all workouts of the current user
// @Summary      List the centers of all workouts of the current user
// @Produce      json
// @Success      200  {object}  APIResponse{results=geojson.FeatureCollection}
// @Failure      400  {object}  APIResponse
// @Failure      404  {object}  APIResponse
// @Failure      500  {object}  APIResponse
// @Router       /workouts/coordinates [get]
func (a *App) apiCenters(c echo.Context) error {
	resp := APIResponse{}
	coords := geojson.NewFeatureCollection()
	u := a.getCurrentUser(c)
	db := a.db.Preload("Data").Preload("Data.Details")

	wos, err := u.GetWorkouts(db)
	if err != nil {
		resp.AddError(err)
	}

	for _, w := range wos {
		if w.Data == nil {
			continue
		}

		p := w.Data.Center
		if p.IsZero() {
			continue
		}

		f := geojson.NewFeature(p.ToOrbPoint())
		a.fillGeoJSONProperties(c, w, f)

		coords.Append(f)
	}

	resp.Results = coords

	return c.JSON(http.StatusOK, resp)
}

// apiCoordinates returns all coordinates of all workouts of the current user
// @Summary      List all coordinates of all workouts of the current user
// @Produce      json
// @Success      200  {object}  APIResponse{results=geojson.FeatureCollection}
// @Failure      400  {object}  APIResponse
// @Failure      404  {object}  APIResponse
// @Failure      500  {object}  APIResponse
// @Router       /workouts/coordinates [get]
func (a *App) apiCoordinates(c echo.Context) error {
	resp := APIResponse{}
	coords := geojson.NewFeatureCollection()

	db := a.db.Preload("Data").Preload("Data.Details")
	u := a.getCurrentUser(c)

	wos, err := u.GetWorkouts(db)
	if err != nil {
		resp.AddError(err)
	}

	for _, w := range wos {
		if !w.HasTracks() {
			continue
		}

		for _, p := range w.Data.Details.Points {
			f := geojson.NewFeature(p.ToOrbPoint())

			coords.Append(f)
		}
	}

	resp.Results = coords

	return c.JSON(http.StatusOK, resp)
}

// apiRecordsHandler lists current user's records for the specified workout type
// @Summary      List all records of the current user for the specified workout type
// @Param        type   query      string  true  "Workout type"
// @Produce      json
// @Success      200  {object}  APIResponse{results=database.WorkoutRecord}
// @Failure      400  {object}  APIResponse
// @Failure      404  {object}  APIResponse
// @Failure      500  {object}  APIResponse
// @Router       /records [get]
func (a *App) apiRecordsHandler(c echo.Context) error {
	resp := APIResponse{}

	var workoutType string

	if err := echo.QueryParamsBinder(c).String("type", &workoutType).BindError(); err != nil {
		return a.renderAPIError(c, resp, err)
	}

	s, err := a.getCurrentUser(c).GetRecords(database.AsWorkoutType(workoutType))
	if err != nil {
		resp.AddError(err)
	}

	resp.Results = s

	return c.JSON(http.StatusOK, resp)
}

// apiTotalsHandler lists current user's totals for the specified workout type
// @Summary      List all totals of the current user for the specified workout type
// @Param        type   query      string  false  "Workout type"
// @Produce      json
// @Success      200  {object}  APIResponse{results=database.Bucket}
// @Failure      400  {object}  APIResponse
// @Failure      404  {object}  APIResponse
// @Failure      500  {object}  APIResponse
// @Router       /totals [get]
func (a *App) apiTotalsHandler(c echo.Context) error {
	resp := APIResponse{}

	var workoutType string

	if err := echo.QueryParamsBinder(c).String("type", &workoutType).BindError(); err != nil {
		return a.renderAPIError(c, resp, err)
	}

	s, err := a.getCurrentUser(c).GetTotals(database.AsWorkoutType(workoutType))
	if err != nil {
		resp.AddError(err)
	}

	resp.Results = s

	return c.JSON(http.StatusOK, resp)
}

// apiStatisticsHandler returns a user's statistics for a given time range and bucket size
// @Summary      List all statistics of the current user
// @Param        since   query      string  false  "Start of time range"
// @Param        per     query      string  false  "Bucket size"
// @Produce      json
// @Success      200  {object}  APIResponse{results=database.Statistics}
// @Failure      400  {object}  APIResponse
// @Failure      404  {object}  APIResponse
// @Failure      500  {object}  APIResponse
// @Router       /statistics [get]
func (a *App) apiStatisticsHandler(c echo.Context) error {
	resp := APIResponse{}

	var statConfig database.StatConfig

	if err := c.Bind(&statConfig); err != nil {
		return a.renderAPIError(c, resp, err)
	}

	s, err := a.getCurrentUser(c).GetStatistics(statConfig)
	if err != nil {
		resp.AddError(err)
	}

	resp.Results = s

	return c.JSON(http.StatusOK, resp)
}

// apiWorkoutBreakdownHandler returns the breakdown per unit for a given workout
// @Summary      Break down a workdown per units
// @Param        id      path       int     true  "Workout ID"
// @Param        unit    query      string  false  "Unit"
// @Param        count   query      int     false  "Count"
// @Produce      json
// @Success      200  {object}  APIResponse{results=database.WorkoutBreakdown}
// @Failure      400  {object}  APIResponse
// @Failure      404  {object}  APIResponse
// @Failure      500  {object}  APIResponse
// @Router       /workouts/{id}/breakdown [get]
func (a *App) apiWorkoutBreakdownHandler(c echo.Context) error {
	resp := APIResponse{}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return a.renderAPIError(c, resp, err)
	}

	config := struct {
		Unit  string  `query:"unit"`
		Count float64 `query:"count"`
	}{
		Unit:  "km",
		Count: 1.0,
	}
	if err := c.Bind(&config); err != nil {
		return a.renderAPIError(c, resp, err)
	}

	w, err := a.getCurrentUser(c).GetWorkout(a.db, id)
	if err != nil {
		resp.AddError(err)
	}

	resp.Results, err = w.StatisticsPer(config.Count, config.Unit)
	if err != nil {
		return a.renderAPIError(c, resp, err)
	}

	return c.JSON(http.StatusOK, resp)
}

// apiWorkoutAddHandler creates a new workout
// @Summary      Create a new workout
// @Accept       json
// @Param        workout body      ManualWorkout     true "Workout data"
// @Produce      json
// @Success      200  {object}  APIResponse{results=ManualWorkout}
// @Failure      400  {object}  APIResponse
// @Failure      404  {object}  APIResponse
// @Failure      500  {object}  APIResponse
// @Router       /workouts/ [post]
func (a *App) apiWorkoutAddHandler(c echo.Context) error {
	resp := APIResponse{}

	d := &ManualWorkout{units: a.getCurrentUser(c).PreferredUnits()}
	if err := json.NewDecoder(c.Request().Body).Decode(d); err != nil {
		return a.renderAPIError(c, resp, err)
	}

	workout := &database.Workout{}
	resp.Results = workout

	d.Update(workout)

	workout.User = a.getCurrentUser(c)
	workout.UserID = a.getCurrentUser(c).ID
	workout.Data.Creator = "api"

	if err := workout.Save(a.db); err != nil {
		return a.renderAPIError(c, resp, err)
	}

	return c.JSON(http.StatusOK, resp)
}

// apiWorkoutHandler returns all information about a workout
// @Summary      Get all information about a workout
// @Param        id      path       int     true  "Workout ID"
// @Param        details query      bool    false "Include details"
// @Produce      json
// @Success      200  {object}  APIResponse{results=database.Workout}
// @Failure      400  {object}  APIResponse
// @Failure      404  {object}  APIResponse
// @Failure      500  {object}  APIResponse
// @Router       /workouts/{id} [get]
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
		resp.AddError(err)
	}

	resp.Results = w

	return c.JSON(http.StatusOK, resp)
}

// apiImportHandler imports a workout
// @Summary      Import a workout
// @Param        program path  string true "Program that generates the workout file"
// @Param        name query    string false "Name of the imported workout"
// @Param        type query    string false "Type of the imported workout"
// @Produce      json
// @Success      200  {object}  APIResponse{results=database.Workout}
// @Failure      400  {object}  APIResponse
// @Failure      404  {object}  APIResponse
// @Failure      500  {object}  APIResponse
// @Router       /import/{program} [post]
func (a *App) apiImportHandler(c echo.Context) error {
	resp := APIResponse{}

	program := c.Param("program")
	a.logger.Info("Importing with program: " + program)

	file, err := importers.Import(program, c, c.Request().Body)
	if err != nil {
		return a.renderAPIError(c, resp, err)
	}

	w, addErr := a.getCurrentUser(c).AddWorkout(a.db, database.WorkoutType(file.Type), file.Notes, file.Filename, file.Content)
	if len(addErr) > 0 {
		return a.renderAPIError(c, resp, addErr...)
	}

	resp.Results = w

	return c.JSON(http.StatusOK, resp)
}

type Event struct {
	Title string    `json:"title"`
	Start time.Time `json:"start"`
	URL   string    `json:"url"`
}

// apiCalendar returns the calendar events of all workouts of the current user
// @Summary      List the calendar events of all workouts of the current user
// @Param        start query    string false "Start date of the calendar view"
// @Param        end query    string false "End date of the calendar view"
// @Produce      json
// @Success      200  {object}  APIResponse{results=[]Event}
// @Failure      400  {object}  APIResponse
// @Failure      404  {object}  APIResponse
// @Failure      500  {object}  APIResponse
// @Router       /workouts/coordinates [get]
func (a *App) apiCalendar(c echo.Context) error {
	resp := APIResponse{}
	events := []Event{}

	queryParams := struct {
		Start *string `query:"start"`
		End   *string `query:"end"`
	}{}
	if err := c.Bind(&queryParams); err != nil {
		return a.renderAPIError(c, resp, err)
	}

	u := a.getCurrentUser(c)
	db := a.db.Preload("Data").Preload("Data.Details")

	if queryParams.Start != nil {
		db = db.Where("workouts.date >= ?", queryParams.Start)
	}

	if queryParams.End != nil {
		db = db.Where("workouts.date < ?", queryParams.End)
	}

	wos, err := u.GetWorkouts(db)
	if err != nil {
		return a.renderAPIError(c, resp, err)
	}

	for _, w := range wos {
		buf := templ.GetBuffer()
		defer templ.ReleaseBuffer(buf)

		t := workouts.EventTitle(w, u.PreferredUnits())
		if err := t.Render(c.Request().Context(), buf); err != nil {
			return a.renderAPIError(c, resp, err)
		}

		d := buf.String()
		// Remove all newlines and surrounding whitespace
		d = htmlConcatenizer.ReplaceAllString(d, "")

		events = append(events, Event{
			Title: d,
			Start: w.Date,
			URL:   a.echo.Reverse("workout-show", w.ID),
		})
	}

	resp.Results = events

	return c.JSON(http.StatusOK, resp)
}

func (a *App) renderAPIError(c echo.Context, resp APIResponse, err ...error) error {
	resp.AddError(err...)

	return c.JSON(http.StatusBadRequest, resp)
}

func (a *App) fillGeoJSONProperties(c echo.Context, w *database.Workout, f *geojson.Feature) {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)

	t := workouts.Details(w)
	if err := t.Render(c.Request().Context(), buf); err != nil {
		return
	}

	d := buf.String()
	// Remove all newlines and surrounding whitespace
	d = htmlConcatenizer.ReplaceAllString(d, "")

	f.Properties["details"] = d
}

// apiDailyHandler returns the daily measurements for the user
// @Summary      List the daily measurements of the current user
// @Param        limit query    int false "Number of measurements to return; default 50; -1 is no limit"
// @Produce      json
// @Success      200  {object}  APIResponse{results=[]database.Measurement}
// @Failure      400  {object}  APIResponse
// @Failure      404  {object}  APIResponse
// @Failure      500  {object}  APIResponse
// @Router       /daily [get]
func (a *App) apiDailyHandler(c echo.Context) error {
	resp := APIResponse{}
	u := a.getCurrentUser(c)

	limit := 50

	if l := c.QueryParam("limit"); l != "" {
		if nl, err := strconv.Atoi(l); err == nil {
			limit = nl
		} else {
			return a.renderAPIError(c, resp, err)
		}
	}

	m, err := u.GetLatestMeasurements(limit)
	if err != nil {
		return a.renderAPIError(c, resp, err)
	}

	resp.Results = m

	return c.JSON(http.StatusOK, resp)
}

// apiDailyUpdateHandler updates the daily measurement for the user
// @Summary      Update the daily measurement of the current user
// @Accept       json
// @Param        measurement body      Measurement     true "Measurement data"
// @Produce      json
// @Success      200  {object}  APIResponse{results=database.Measurement}
// @Failure      400  {object}  APIResponse
// @Failure      404  {object}  APIResponse
// @Failure      500  {object}  APIResponse
// @Router       /daily [post]
func (a *App) apiDailyUpdateHandler(c echo.Context) error {
	resp := APIResponse{}

	d := &Measurement{units: a.getCurrentUser(c).PreferredUnits()}
	if err := json.NewDecoder(c.Request().Body).Decode(d); err != nil {
		return a.renderAPIError(c, resp, err)
	}

	m, err := a.getCurrentUser(c).GetMeasurementForDate(d.Time())
	if err != nil {
		return a.renderAPIError(c, resp, err)
	}

	d.Update(m)

	if err := m.Save(a.db); err != nil {
		return a.renderAPIError(c, resp, err)
	}

	resp.Results = m

	return c.JSON(http.StatusOK, resp)
}
