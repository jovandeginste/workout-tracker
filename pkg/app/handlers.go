package app

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (a *App) redirectWithError(c echo.Context, err error) error {
	a.setError(c, err.Error())

	return c.Redirect(http.StatusMovedPermanently, "/")
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
	data["workout_map_data"] = w.MapData()

	return c.Render(http.StatusOK, "workouts_show.html", data)
}

func (a *App) dashboardHandler(c echo.Context) error {
	data := a.defaultData(c)
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

func (a *App) workoutsAddHandler(c echo.Context) error {
	data := a.defaultData(c)
	return c.Render(http.StatusOK, "workouts_add.html", data)
}

func (a *App) workoutsStatisticsHandler(c echo.Context) error {
	data := a.defaultData(c)

	a.addWorkouts(data, c)

	return c.Render(http.StatusOK, "workouts_statistics.html", data)
}
