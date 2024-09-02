package app

import (
	"net/http"

	"github.com/jovandeginste/workout-tracker/internal/database"
	"github.com/labstack/echo/v4"
)

func (a *App) addRoutesSelf(e *echo.Group) {
	selfGroup := e.Group("/user")
	selfGroup.GET("/profile", a.userProfileHandler).Name = "user-profile"
	selfGroup.POST("/profile", a.userProfileUpdateHandler).Name = "user-profile-update"
	selfGroup.POST("/profile/preferred-units", a.userProfilePreferredUnitsUpdateHandler).Name = "user-profile-preferred-units-update"
	selfGroup.POST("/refresh", a.userRefreshHandler).Name = "user-refresh"
	selfGroup.POST("/reset-api-key", a.userProfileResetAPIKeyHandler).Name = "user-profile-reset-api-key"
	selfGroup.POST("/update-version", a.userUpdateVersion).Name = "user-update-version"
}

func (a *App) userProfileHandler(c echo.Context) error {
	data := a.defaultData(c)
	return c.Render(http.StatusOK, "user_profile.html", data)
}

func (a *App) userProfilePreferredUnitsUpdateHandler(c echo.Context) error {
	u := a.getCurrentUser(c)
	p := database.UserPreferredUnits{}

	if err := c.Bind(&p); err != nil {
		return a.redirectWithError(c, a.echo.Reverse("user-profile"), err)
	}

	u.Profile.PreferredUnits = p

	if err := u.Profile.Save(a.db); err != nil {
		return a.redirectWithError(c, a.echo.Reverse("user-profile"), err)
	}

	if err := a.setUser(c); err != nil {
		return a.redirectWithError(c, a.echo.Reverse("user-profile"), err)
	}

	a.setNotice(c, "Preferred units updated")

	return c.Redirect(http.StatusFound, a.echo.Reverse("user-profile"))
}

func (a *App) userProfileUpdateHandler(c echo.Context) error {
	u := a.getCurrentUser(c)
	p := &u.Profile

	p.ResetBools()

	if err := c.Bind(p); err != nil {
		return a.redirectWithError(c, a.echo.Reverse("user-profile"), err)
	}

	p.UserID = u.ID

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

func (a *App) userUpdateVersion(c echo.Context) error {
	data := a.defaultData(c)
	u := a.getCurrentUser(c)

	u.LastVersion = a.Version.Sha
	if err := u.Save(a.db); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.Render(http.StatusOK, "version_updated.html", data)
}
