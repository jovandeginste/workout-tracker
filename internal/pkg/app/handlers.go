package app

import (
	"cmp"
	"errors"
	"net/http"

	"github.com/jovandeginste/workout-tracker/internal/pkg/geocoder"
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
	data["since"] = cmp.Or(c.QueryParam("since"), "1 year")
	data["per"] = cmp.Or(c.QueryParam("per"), "month")

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

func (a *App) lookupAddressHandler(c echo.Context) error {
	data := a.defaultData(c)

	q := c.FormValue("location")

	results, err := geocoder.Search(q)
	if err != nil {
		a.setError(c, "Something went wrong: "+err.Error())
	}

	data["addresses"] = results

	return c.Render(http.StatusOK, "address_results.html", data)
}
