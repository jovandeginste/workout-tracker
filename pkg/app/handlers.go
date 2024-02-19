package app

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/jovandeginste/workouts/pkg/database"
	"github.com/labstack/echo/v4"
)

func (a *App) redirectWithError(c echo.Context, err error) error {
	a.setError(c, err.Error())

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func (a *App) dashboardHandler(c echo.Context) error {
	data := a.defaultData(c)

	a.addWorkouts(data, c)
	a.addUserStatistics(data, c)

	return c.Render(http.StatusOK, "dashboard.html", data)
}

func (a *App) loginHandler(c echo.Context) error {
	data := a.defaultData(c)

	return c.Render(http.StatusOK, "user_login.html", data)
}

func (a *App) workoutsHandler(c echo.Context) error {
	data := a.defaultData(c)

	a.addWorkouts(data, c)

	return c.Render(http.StatusOK, "workouts_list.html", data)
}

func (a *App) workoutsShowHandler(c echo.Context) error {
	data := a.defaultData(c)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return a.redirectWithError(c, err)
	}

	w, err := a.getUser(c).GetWorkout(a.db, id)
	if err != nil {
		return a.redirectWithError(c, err)
	}

	data["workout"] = w
	data["workoutStatistics"] = w.StatisticsPerKilometer()

	return c.Render(http.StatusOK, "workouts_show.html", data)
}

func (a *App) workoutsAddHandler(c echo.Context) error {
	data := a.defaultData(c)
	return c.Render(http.StatusOK, "workouts_add.html", data)
}

func (a *App) workoutsDeleteHandler(c echo.Context) error {
	workout, ok := c.Get("workout").(*database.Workout)
	if !ok {
		return c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("/workouts/%d", c.Param("id")))
	}

	return c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("/workouts/%d", workout.ID))
}

func (a *App) workoutsRefreshHandler(c echo.Context) error {
	workout, ok := c.Get("workout").(*database.Workout)
	if !ok {
		return c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("/workouts/%d", c.Param("id")))
	}

	return c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("/workouts/%d", workout.ID))
}

func (a *App) workoutsUpdateHandler(c echo.Context) error {
	if err := a.addWorkoutToContext(c); err != nil {
		return a.redirectWithError(c, err)
	}

	action := c.FormValue("action")
	switch action {
	case "delete":
		return a.workoutsDeleteHandler(c)
	case "refresh":
		return a.workoutsRefreshHandler(c)
	default:
		data := a.defaultData(c)
		return c.Render(http.StatusOK, "workouts_update.html", data)
	}
}

func (a *App) addWorkoutToContext(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return a.redirectWithError(c, err)
	}

	w, err := a.getUser(c).GetWorkout(a.db, id)
	if err != nil {
		return a.redirectWithError(c, err)
	}

	c.Set("workout", w)

	return nil
}

func (a *App) workoutsEditHandler(c echo.Context) error {
	data := a.defaultData(c)

	if err := a.addWorkoutToContext(c); err != nil {
		return a.redirectWithError(c, err)
	}

	data["workout"] = c.Get("workout")

	return c.Render(http.StatusOK, "workouts_edit.html", data)
}

func (a *App) userProfileHandler(c echo.Context) error {
	data := a.defaultData(c)
	return c.Render(http.StatusOK, "user_profile.html", data)
}
