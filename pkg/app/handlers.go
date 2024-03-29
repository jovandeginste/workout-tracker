package app

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

var ErrUserNotFound = errors.New("user not found")

func (a *App) redirectWithError(c echo.Context, target string, err error) error {
	if err != nil {
		a.setError(c, "Something went wrong: "+err.Error())
	}

	return c.Redirect(http.StatusFound, target)
}

func (a *App) statisticsHandler(c echo.Context) error {
	data := a.defaultData(c)
	u := a.getCurrentUser(c)

	if err := a.addWorkouts(u, data); err != nil {
		return a.redirectWithError(c, a.echo.Reverse("user-signout"), err)
	}

	if err := a.addUserStatistics(u, data); err != nil {
		return a.redirectWithError(c, a.echo.Reverse("user-signout"), err)
	}

	data["user"] = u

	return c.Render(http.StatusOK, "user_statistics.html", data)
}

func (a *App) dashboardHandler(c echo.Context) error {
	data := a.defaultData(c)

	u := a.getCurrentUser(c)
	if u == nil {
		return a.redirectWithError(c, a.echo.Reverse("user-signout"), ErrUserNotFound)
	}

	if err := a.addWorkouts(u, data); err != nil {
		return a.redirectWithError(c, a.echo.Reverse("user-signout"), err)
	}

	if err := a.addUserStatistics(u, data); err != nil {
		return a.redirectWithError(c, a.echo.Reverse("user-signout"), err)
	}

	if err := a.addUsers(data); err != nil {
		return a.redirectWithError(c, a.echo.Reverse("user-signout"), err)
	}

	if err := a.addRecentWorkouts(data); err != nil {
		return a.redirectWithError(c, a.echo.Reverse("user-signout"), err)
	}

	data["user"] = u

	return c.Render(http.StatusOK, "user_show.html", data)
}

func (a *App) userLoginHandler(c echo.Context) error {
	data := a.defaultData(c)

	return c.Render(http.StatusOK, "user_login.html", data)
}
