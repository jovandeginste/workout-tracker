package app

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (a *App) mapHandler(c echo.Context) error {
	gpxContent, err := uploadedGPXFile(c)
	if err != nil {
		return err
	}

	data := gpxAsMapData(gpxContent)

	return c.Render(http.StatusOK, "map.html", data)
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
