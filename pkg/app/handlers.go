package app

import (
	"errors"
	"net/http"

	"github.com/a-h/templ"
	"github.com/jovandeginste/workout-tracker/pkg/database"
	"github.com/jovandeginste/workout-tracker/pkg/geocoder"
	"github.com/jovandeginste/workout-tracker/views/partials"
	"github.com/jovandeginste/workout-tracker/views/user"
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
	a.setContext(c)

	u := a.getCurrentUser(c)
	if u.IsAnonymous() {
		return a.redirectWithError(c, a.echo.Reverse("user-signout"), ErrUserNotFound)
	}

	statisticsParams := struct {
		Since string `query:"since"`
		Per   string `query:"per"`
	}{
		Since: "1 year",
		Per:   "month",
	}

	if err := c.Bind(&statisticsParams); err != nil {
		return a.redirectWithError(c, a.echo.Reverse("dashboard"), err)
	}

	return Render(c, http.StatusOK, user.Statistics(u, statisticsParams.Since, statisticsParams.Per))
}

func (a *App) dashboardHandler(c echo.Context) error {
	a.setContext(c)

	u := a.getCurrentUser(c)
	if u.IsAnonymous() {
		return a.redirectWithError(c, a.echo.Reverse("user-signout"), ErrUserNotFound)
	}

	w, err := u.GetWorkouts(a.db)
	if err != nil {
		return a.redirectWithError(c, a.echo.Reverse("user-signout"), ErrUserNotFound)
	}

	users, err := database.GetUsers(a.db)
	if err != nil {
		return a.redirectWithError(c, a.echo.Reverse("user-signout"), ErrUserNotFound)
	}

	recent, err := database.GetRecentWorkouts(a.db, 20)
	if err != nil {
		return a.redirectWithError(c, a.echo.Reverse("user-signout"), ErrUserNotFound)
	}

	return Render(c, http.StatusOK, user.Show(u, users, w, recent))
}

func (a *App) userLoginHandler(c echo.Context) error {
	a.setContext(c)

	return Render(c, http.StatusOK, user.Login())
}

func (a *App) lookupAddressHandler(c echo.Context) error {
	a.setContext(c)

	q := c.FormValue("location")

	results, err := geocoder.Search(q)
	if err != nil {
		a.setError(c, "Something went wrong: "+err.Error())
	}

	return Render(c, http.StatusOK, partials.AddressResults(results))
}

func (a *App) heatmapHandler(c echo.Context) error {
	a.setContext(c)

	u := a.getCurrentUser(c)
	if u.IsAnonymous() {
		return a.redirectWithError(c, a.echo.Reverse("user-signout"), ErrUserNotFound)
	}

	w, err := u.GetWorkouts(a.db)
	if err != nil {
		return a.redirectWithError(c, a.echo.Reverse("user-signout"), err)
	}

	return Render(c, http.StatusOK, user.Heatmap(w))
}

func Render(ctx echo.Context, statusCode int, t templ.Component) error {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)

	if err := t.Render(ctx.Request().Context(), buf); err != nil {
		return err
	}

	return ctx.HTML(statusCode, buf.String())
}
