package app

import (
	"net/http"

	"github.com/jovandeginste/workout-tracker/pkg/database"
	"github.com/labstack/echo/v4"
)

func (a *App) userProfileHandler(c echo.Context) error {
	data := a.defaultData(c)
	return c.Render(http.StatusOK, "user_profile.html", data)
}

func (a *App) userProfileUpdateHandler(c echo.Context) error {
	u := a.getCurrentUser(c)
	p := &database.Profile{}

	if err := c.Bind(p); err != nil {
		return a.redirectWithError(c, a.echo.Reverse("user-profile"), err)
	}

	u.Profile.APIActive = p.APIActive
	u.Profile.Language = p.Language
	u.Profile.TotalsShow = p.TotalsShow
	u.Profile.Timezone = p.Timezone
	u.Profile.AutoImportDirectory = p.AutoImportDirectory

	if err := u.Profile.Save(a.db); err != nil {
		return a.redirectWithError(c, a.echo.Reverse("user-profile"), err)
	}

	if err := a.setUser(c); err != nil {
		return a.redirectWithError(c, a.echo.Reverse("user-profile"), err)
	}

	a.setNotice(c, "Profile updated")

	return c.Redirect(http.StatusFound, a.echo.Reverse("user-profile"))
}

func (a *App) userProfileResetAPIKeyHandler(c echo.Context) error {
	u := a.getCurrentUser(c)

	u.GenerateAPIKey(true)

	if err := u.Save(a.db); err != nil {
		return a.redirectWithError(c, a.echo.Reverse("user-profile"), err)
	}

	a.setNotice(c, "API key updated")

	return c.Redirect(http.StatusFound, a.echo.Reverse("user-profile"))
}

func (a *App) userRefreshHandler(c echo.Context) error {
	u := a.getCurrentUser(c)

	if err := u.MarkWorkoutsDirty(a.db); err != nil {
		return a.redirectWithError(c, a.echo.Reverse("user-profile"), err)
	}

	a.setNotice(c, "All workouts will be refreshed in the coming minutes.")

	return c.Redirect(http.StatusFound, a.echo.Reverse("user-profile"))
}
