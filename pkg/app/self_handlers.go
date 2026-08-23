package app

import (
	"net/http"

	"github.com/jovandeginste/workout-tracker/v2/pkg/database"
	"github.com/jovandeginste/workout-tracker/v2/views/user"
	"github.com/labstack/echo/v5"
)

func (a *App) addRoutesSelf(g *echo.Group) {
	selfGroup := g.Group("/user")
	a.GET(selfGroup, "/profile", a.userProfileHandler, "user-profile")
	a.POST(selfGroup, "/profile", a.userProfileUpdateHandler, "user-profile-update")
	a.POST(selfGroup, "/profile/preferred-units", a.userProfilePreferredUnitsUpdateHandler, "user-profile-preferred-units-update")
	a.POST(selfGroup, "/profile/route-segment-trend-period", a.userProfileRouteSegmentTrendPeriodUpdateHandler, "user-profile-route-segment-trend-period-update")
	a.POST(selfGroup, "/profile/route-segment-trend-break-detection", a.userProfileRouteSegmentTrendBreakDetectionUpdateHandler, "user-profile-route-segment-trend-break-detection-update")
	a.POST(selfGroup, "/refresh", a.userRefreshHandler, "user-refresh")
	a.POST(selfGroup, "/reset-api-key", a.userProfileResetAPIKeyHandler, "user-profile-reset-api-key")
	a.POST(selfGroup, "/update-version", a.userUpdateVersion, "user-update-version")
}

func (a *App) userProfileHandler(c *echo.Context) error {
	u := a.getCurrentUser(c)

	return Render(c, http.StatusOK, user.Profile(u))
}

func (a *App) userProfilePreferredUnitsUpdateHandler(c *echo.Context) error {
	u := a.getCurrentUser(c)
	p := database.UserPreferredUnits{}

	if err := c.Bind(&p); err != nil {
		return a.redirectWithError(c, a.Reverse("user-profile"), err)
	}

	u.Profile.PreferredUnits = p

	if err := u.Profile.Save(a.db); err != nil {
		return a.redirectWithError(c, a.Reverse("user-profile"), err)
	}

	if err := a.setUser(c); err != nil {
		return a.redirectWithError(c, a.Reverse("user-profile"), err)
	}

	a.addNoticeT(c, "Preferred units updated")

	return c.Redirect(http.StatusFound, a.Reverse("user-profile"))
}

func (a *App) userProfileUpdateHandler(c *echo.Context) error {
	u := a.getCurrentUser(c)
	p := &u.Profile

	p.ResetBools()

	if err := c.Bind(p); err != nil {
		return a.redirectWithError(c, a.Reverse("user-profile"), err)
	}

	p.UserID = u.ID

	if err := u.Profile.Save(a.db); err != nil {
		return a.redirectWithError(c, a.Reverse("user-profile"), err)
	}

	if err := a.setUser(c); err != nil {
		return a.redirectWithError(c, a.Reverse("user-profile"), err)
	}

	a.addNoticeT(c, "translation.Profile_updated")

	return c.Redirect(http.StatusFound, a.Reverse("user-profile"))
}

func (a *App) userProfileRouteSegmentTrendPeriodUpdateHandler(c *echo.Context) error {
	u := a.getCurrentUser(c)
	u.Profile.RouteSegmentConfig.TrendPeriod = c.FormValue("route_segment_trend_period")
	u.Profile.RouteSegmentConfig.Validate()

	if err := u.Profile.Save(a.db); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	if err := a.setUser(c); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}

func (a *App) userProfileRouteSegmentTrendBreakDetectionUpdateHandler(c *echo.Context) error {
	u := a.getCurrentUser(c)

	u.Profile.RouteSegmentConfig.TrendBreakDetection = c.FormValue("route_segment_trend_break_detection") == "true"

	if err := u.Profile.Save(a.db); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	if err := a.setUser(c); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}

func (a *App) userProfileResetAPIKeyHandler(c *echo.Context) error {
	u := a.getCurrentUser(c)

	u.GenerateAPIKey(true)

	if err := u.Save(a.db); err != nil {
		return a.redirectWithError(c, a.Reverse("user-profile"), err)
	}

	a.addNoticeT(c, "translation.API_key_updated")

	if isHtmx(c) {
		c.Response().Header().Set("Hx-Redirect", a.Reverse("user-profile"))
		return c.String(http.StatusFound, "ok")
	}

	return c.Redirect(http.StatusFound, a.Reverse("user-profile"))
}

func (a *App) userRefreshHandler(c *echo.Context) error {
	u := a.getCurrentUser(c)

	if err := u.MarkWorkoutsDirty(a.db); err != nil {
		return a.redirectWithError(c, a.Reverse("user-profile"), err)
	}

	a.addNoticeT(c, "translation.All_workouts_will_be_refreshed_in_the_coming_minutes")

	return c.Redirect(http.StatusFound, a.Reverse("user-profile"))
}

func (a *App) userUpdateVersion(c *echo.Context) error {
	u := a.getCurrentUser(c)

	u.LastVersion = a.Version.Sha
	if err := u.Save(a.db); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return Render(c, http.StatusOK, user.VersionUpdated())
}
