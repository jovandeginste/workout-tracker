package app

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/jovandeginste/workout-tracker/v2/pkg/database"
	"github.com/jovandeginste/workout-tracker/v2/views/user"
	"github.com/labstack/echo/v5"
)

func (a *App) addRoutesSelf(g *echo.Group) {
	selfGroup := g.Group("/user")
	a.GET(selfGroup, "/profile", a.userProfileHandler, "user-profile")
	a.POST(selfGroup, "/profile", a.userProfileUpdateHandler, "user-profile-update")
	a.POST(selfGroup, "/profile/preferred-units", a.userProfilePreferredUnitsUpdateHandler, "user-profile-preferred-units-update")
	a.POST(selfGroup, "/profile/route-segment-config", a.userProfileRouteSegmentConfigUpdateHandler, "user-profile-route-segment-config-update")
	a.POST(selfGroup, "/refresh", a.userRefreshHandler, "user-refresh")
	a.POST(selfGroup, "/reset-api-key", a.userProfileResetAPIKeyHandler, "user-profile-reset-api-key")
	a.POST(selfGroup, "/update-version", a.userUpdateVersion, "user-update-version")
	a.GET(selfGroup, "/export", a.userExportAllHandler, "user-export-all")
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

func (a *App) userProfileRouteSegmentConfigUpdateHandler(c *echo.Context) error {
	u := a.getCurrentUser(c)

	if err := c.Request().ParseForm(); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	values := c.Request().PostForm

	if period := values.Get("route_segment_trend_period"); period != "" {
		u.Profile.RouteSegmentConfig.TrendPeriod = period
		u.Profile.RouteSegmentConfig.Validate()
	}

	if _, ok := values["route_segment_trend_break_detection"]; ok {
		u.Profile.RouteSegmentConfig.TrendBreakDetection = values.Get("route_segment_trend_break_detection") == "true"
	}

	if err := u.Profile.Save(a.db); err != nil {
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

func (a *App) userExportAllHandler(c *echo.Context) error {
	u := a.getCurrentUser(c)

	var wIDs []uint64

	if err := a.db.
		Model(&database.Workout{}).
		Where(&database.Workout{UserID: u.ID}).
		Pluck("ID", &wIDs).Error; err != nil {
		return a.redirectWithError(c, a.Reverse("user-profile"), err)
	}

	tmpFile, err := os.CreateTemp("", "workouts-export-*.zip")
	if err != nil {
		return a.redirectWithError(c, a.Reverse("user-profile"), err)
	}

	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()

	archive := zip.NewWriter(tmpFile)

	for _, w := range wIDs {
		if err := a.writeWorkoutToArchive(archive, w); err != nil {
			return a.redirectWithError(c, a.Reverse("user-profile"), err)
		}
	}

	if err := archive.Close(); err != nil {
		return a.redirectWithError(c, a.Reverse("user-profile"), err)
	}

	if _, err := tmpFile.Seek(0, io.SeekStart); err != nil {
		return a.redirectWithError(c, a.Reverse("user-profile"), err)
	}

	c.Response().Header().Set(echo.HeaderContentDisposition, "attachment; filename=\"workouts-export.zip\"")
	return c.Stream(http.StatusOK, "application/zip", tmpFile)
}

func (a *App) writeWorkoutToArchive(archive *zip.Writer, wID uint64) error {
	w, err := database.GetWorkoutDetails(a.db, wID)
	if err != nil {
		return err
	}

	if w.GPX == nil || len(w.GPX.Content) == 0 {
		return nil
	}

	filename := w.GPX.Filename
	if filename == "" {
		filename = "workout.gpx"
	}

	filename = strings.ReplaceAll(filename, "\\", "_")
	filename = strings.ReplaceAll(filename, "/", "_")

	filename = fmt.Sprintf("%d_%s", w.ID, filename)

	writer, err := archive.Create(filename)
	if err != nil {
		return err
	}

	if _, err := writer.Write(w.GPX.Content); err != nil {
		return err
	}

	return nil
}
