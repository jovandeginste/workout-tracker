package app

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (a *App) redirectWithError(c echo.Context, target string, err error) error {
	a.setError(c, err.Error())

	return c.Redirect(http.StatusFound, target)
}

func (a *App) dashboardHandler(c echo.Context) error {
	data := a.defaultData(c)

	a.addWorkouts(data, c)
	a.addUserStatistics(data, c)

	return c.Render(http.StatusOK, "dashboard.html", data)
}

func (a *App) userLoginHandler(c echo.Context) error {
	data := a.defaultData(c)

	return c.Render(http.StatusOK, "user_login.html", data)
}
