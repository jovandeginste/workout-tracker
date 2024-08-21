package app

import (
	"bytes"
	"errors"
	"net/http"
	"path"
	"strconv"
	"time"

	"github.com/jovandeginste/workout-tracker/pkg/database"
	"github.com/labstack/echo/v4"
)

func (a *App) workoutsHandler(c echo.Context) error {
	data := a.defaultData(c)

	if err := a.addWorkouts(a.getCurrentUser(c), data); err != nil {
		return a.redirectWithError(c, a.echo.Reverse("dashboard"), err)
	}

	return c.Render(http.StatusOK, "workouts_list.html", data)
}

func (a *App) workoutsShowHandler(c echo.Context) error {
	data := a.defaultData(c)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return a.redirectWithError(c, a.echo.Reverse("workouts"), err)
	}

	w, err := database.GetWorkoutDetails(a.db, id)
	if err != nil {
		return a.redirectWithError(c, a.echo.Reverse("workouts"), err)
	}

	data["workout"] = w

	return c.Render(http.StatusOK, "workouts_show.html", data)
}

func (a *App) workoutsAddHandler(c echo.Context) error {
	data := a.defaultData(c)
	return c.Render(http.StatusOK, "workouts_add.html", data)
}

func (a *App) workoutsFormHandler(c echo.Context) error {
	w := &database.Workout{}

	if c.FormValue("id") != "" {
		id, err := strconv.Atoi(c.FormValue("id"))
		if err != nil {
			return c.Render(http.StatusOK, "workout_form.html", w)
		}

		w, err = a.getCurrentUser(c).GetWorkout(a.db, id)
		if err != nil {
			return c.Render(http.StatusOK, "workout_form.html", w)
		}
	}

	if w.Type == "" || c.FormValue("type") != "" {
		w.Type = database.WorkoutType(c.FormValue("type"))
	}

	if w.Date == nil {
		t := time.Now()
		w.Date = &t
	}

	if w.Name == "" {
		w.Name = w.Type.String() + " - " + w.Date.Format(time.RFC3339)
	}

	return c.Render(http.StatusOK, "workout_form.html", w)
}

func (a *App) workoutsDeleteHandler(c echo.Context) error { //nolint:dupl
	workout, err := a.getWorkout(c)
	if err != nil {
		return a.redirectWithError(c, a.echo.Reverse("workout-show", c.Param("id")), err)
	}

	if err := workout.Delete(a.db); err != nil {
		return a.redirectWithError(c, a.echo.Reverse("workout-show", c.Param("id")), err)
	}

	a.setNotice(c, "The workout '%s' has been deleted.", workout.Name)

	return c.Redirect(http.StatusFound, a.echo.Reverse("workouts"))
}

func (a *App) workoutsRefreshHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return a.redirectWithError(c, a.echo.Reverse("workout-show", c.Param("id")), err)
	}

	workout, err := a.getCurrentUser(c).GetWorkout(a.db.Preload("GPX"), id)
	if err != nil {
		return a.redirectWithError(c, a.echo.Reverse("workout-show", c.Param("id")), err)
	}

	if err := workout.UpdateData(a.db); err != nil {
		return a.redirectWithError(c, a.echo.Reverse("workout-show", c.Param("id")), err)
	}

	a.setNotice(c, "The workout '%s' has been refreshed.", workout.Name)

	return c.Redirect(http.StatusFound, a.echo.Reverse("workout-show", c.Param("id")))
}

func (a *App) workoutsDownloadHandler(c echo.Context) error {
	workout, err := a.getWorkout(c)
	if err != nil {
		return a.redirectWithError(c, a.echo.Reverse("workouts"), err)
	}

	if !workout.HasFile() {
		return a.redirectWithError(c, a.echo.Reverse("workouts"), errors.New("workout has no content"))
	}

	basename := path.Base(workout.GPX.Filename)

	c.Response().Header().Set(echo.HeaderContentDisposition, "attachment; filename=\""+basename+"\"")

	return c.Stream(http.StatusOK, "application/binary", bytes.NewReader(workout.GPX.Content))
}

func (a *App) workoutsEditHandler(c echo.Context) error {
	data := a.defaultData(c)

	workout, err := a.getWorkout(c)
	if err != nil {
		return a.redirectWithError(c, a.echo.Reverse("workouts"), err)
	}

	data["workout"] = workout

	return c.Render(http.StatusOK, "workouts_edit.html", data)
}

func (a *App) workoutsCreateRouteSegmentFromWorkoutHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return a.redirectWithError(c, a.echo.Reverse("workouts"), err)
	}

	workout, err := database.GetWorkoutDetails(a.db, id)
	if err != nil {
		return a.redirectWithError(c, a.echo.Reverse("workouts"), err)
	}

	var params database.RoutSegmentCreationParams

	if err := c.Bind(&params); err != nil {
		return a.redirectWithError(c, a.echo.Reverse("workouts"), err)
	}

	content, err := database.RouteSegmentFromPoints(workout, &params)
	if err != nil {
		return a.redirectWithError(c, a.echo.Reverse("workouts"), err)
	}

	rs, err := database.AddRouteSegment(a.db, "", params.Filename(), content)
	if err != nil {
		return a.redirectWithError(c, a.echo.Reverse("workouts"), err)
	}

	return c.Redirect(http.StatusFound, a.echo.Reverse("route-segment-show", rs.ID))
}

func (a *App) workoutsCreateRouteSegmentHandler(c echo.Context) error {
	data := a.defaultData(c)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return a.redirectWithError(c, a.echo.Reverse("workouts"), err)
	}

	workout, err := database.GetWorkoutDetails(a.db, id)
	if err != nil {
		return a.redirectWithError(c, a.echo.Reverse("workouts"), err)
	}

	data["workout"] = workout

	return c.Render(http.StatusOK, "workouts_route_segment.html", data)
}
