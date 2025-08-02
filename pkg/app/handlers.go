package app

import (
	"errors"
	"net/http"
	"time"

	"github.com/a-h/templ"
	"github.com/invopop/ctxi18n/i18n"
	"github.com/jovandeginste/workout-tracker/v2/pkg/database"
	"github.com/jovandeginste/workout-tracker/v2/pkg/geocoder"
	"github.com/jovandeginste/workout-tracker/v2/views/partials"
	"github.com/jovandeginste/workout-tracker/v2/views/user"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cast"
	"github.com/stackus/hxgo/hxecho"
)

var ErrUserNotFound = errors.New("user not found")

func (a *App) redirectWithError(c echo.Context, target string, err error) error {
	if err != nil {
		a.addErrorT(c, "alerts.something_wrong", i18n.M{"message": err.Error()})
	}

	return c.Redirect(http.StatusFound, target)
}

func (a *App) statisticsHandler(c echo.Context) error {
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

func (a *App) dailyDeleteHandler(c echo.Context) error {
	u := a.getCurrentUser(c)
	d := c.Param("date")

	t, err := time.Parse("2006-01-02", d)
	if err != nil {
		return a.redirectWithError(c, a.echo.Reverse("daily"), err)
	}

	m, err := u.GetMeasurementForDate(t)
	if err != nil {
		return a.redirectWithError(c, a.echo.Reverse("daily"), err)
	}

	if err := m.Delete(a.db); err != nil {
		return a.redirectWithError(c, a.echo.Reverse("daily"), err)
	}

	if hxecho.IsHtmx(c) {
		c.Response().Header().Set("Hx-Redirect", a.echo.Reverse("daily"))
		return c.String(http.StatusFound, "ok")
	}

	return c.Redirect(http.StatusFound, a.echo.Reverse("daily"))
}

func (a *App) dailyUpdateHandler(c echo.Context) error {
	d := &Measurement{units: a.getCurrentUser(c).PreferredUnits()}
	if err := c.Bind(d); err != nil {
		return a.redirectWithError(c, a.echo.Reverse("daily"), err)
	}

	m, err := a.getCurrentUser(c).GetMeasurementForDate(d.Time())
	if err != nil {
		return a.redirectWithError(c, a.echo.Reverse("daily"), err)
	}

	d.Update(m)

	if err := m.Save(a.db); err != nil {
		return a.redirectWithError(c, a.echo.Reverse("daily"), err)
	}

	return c.Redirect(http.StatusFound, a.echo.Reverse("daily"))
}

func (a *App) dailyHandler(c echo.Context) error {
	u := a.getCurrentUser(c)

	count := 20
	if cs := c.QueryParam("count"); cs != "" {
		if ci, err := cast.ToIntE(cs); err == nil {
			count = ci
		} else {
			return a.redirectWithError(c, a.echo.Reverse("daily"), err)
		}
	}

	return Render(c, http.StatusOK, user.Daily(u, count))
}

func (a *App) dashboardHandler(c echo.Context) error {
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
	return Render(c, http.StatusOK, user.Login())
}

func (a *App) lookupAddressHandler(c echo.Context) error {
	q := c.FormValue("location")

	results, err := geocoder.Search(q)
	if err != nil {
		a.addErrorT(c, "alerts.something_wrong", i18n.M{"message": err.Error()})
	}

	return Render(c, http.StatusOK, partials.AddressResults(results))
}

func (a *App) heatmapHandler(c echo.Context) error {
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
