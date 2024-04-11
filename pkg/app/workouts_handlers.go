package app

import (
	"bytes"
	"errors"
	"net/http"
	"path"
	"strconv"

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
		return a.redirectWithError(c, "/workouts", err)
	}

	w, err := database.GetWorkoutDetails(a.db, id)
	if err != nil {
		return a.redirectWithError(c, "/workouts", err)
	}

	data["workout"] = w

	return c.Render(http.StatusOK, "workouts_show.html", data)
}

func (a *App) workoutsAddHandler(c echo.Context) error {
	data := a.defaultData(c)
	return c.Render(http.StatusOK, "workouts_add.html", data)
}

func (a *App) workoutsDeleteHandler(c echo.Context) error {
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

func (a *App) workoutsUpdateHandler(c echo.Context) error {
	workout, err := a.getWorkout(c)
	if err != nil {
		return a.redirectWithError(c, a.echo.Reverse("workout-show", c.Param("id")), err)
	}

	workout.Name = c.FormValue("name")
	workout.Notes = c.FormValue("notes")
	workout.Type = database.WorkoutType(c.FormValue("type"))

	if err := workout.Save(a.db); err != nil {
		return a.redirectWithError(c, a.echo.Reverse("workout-show", c.Param("id")), err)
	}

	a.setNotice(c, "The workout '%s' has been updated.", workout.Name)

	return c.Redirect(http.StatusFound, a.echo.Reverse("workout-show", c.Param("id")))
}

func (a *App) workoutsDownloadHandler(c echo.Context) error {
	workout, err := a.getWorkout(c)
	if err != nil {
		return a.redirectWithError(c, "/workouts", err)
	}

	if workout.GPX == nil ||
		workout.GPX.Filename == "" ||
		workout.GPX.Content == nil {
		return a.redirectWithError(c, "/workouts", errors.New("workout has no content"))
	}

	basename := path.Base(workout.GPX.Filename)

	c.Response().Header().Set(echo.HeaderContentDisposition, "attachment; filename=\""+basename+"\"")

	return c.Stream(http.StatusOK, "application/binary", bytes.NewReader(workout.GPX.Content))
}

func (a *App) workoutsEditHandler(c echo.Context) error {
	data := a.defaultData(c)

	workout, err := a.getWorkout(c)
	if err != nil {
		return a.redirectWithError(c, "/workouts", err)
	}

	data["workout"] = workout

	return c.Render(http.StatusOK, "workouts_edit.html", data)
}
