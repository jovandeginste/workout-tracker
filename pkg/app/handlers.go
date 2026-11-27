package app

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"path"
	"time"

	"github.com/a-h/templ"
	"github.com/invopop/ctxi18n/i18n"
	"github.com/jovandeginste/workout-tracker/v2/pkg/database"
	"github.com/jovandeginste/workout-tracker/v2/pkg/geocoder"
	"github.com/jovandeginste/workout-tracker/v2/views/helpers"
	"github.com/jovandeginste/workout-tracker/v2/views/partials"
	"github.com/jovandeginste/workout-tracker/v2/views/user"
	"github.com/labstack/echo/v5"
	"github.com/spf13/cast"
)

var ErrUserNotFound = errors.New("user not found")

func (a *App) redirectWithError(c *echo.Context, target string, err error) error {
	if err != nil {
		a.addErrorT(c, "alerts.something-wrong", i18n.M{"message": err.Error()})
	}

	return c.Redirect(http.StatusFound, target)
}

// redirectHTMXOrFound lets HTMX perform a full page navigation, while normal
// requests get a regular redirect.
func processUploads(c *echo.Context, handle func(filename string, content []byte) (string, error)) ([]string, []string, error) {
	form, err := c.MultipartForm()
	if err != nil {
		return nil, nil, err
	}

	msg := []string{}
	errMsg := []string{}

	for _, file := range form.File["file"] {
		content, err := uploadedFile(file)
		if err != nil {
			errMsg = append(errMsg, err.Error())
			continue
		}

		result, err := handle(file.Filename, content)
		if err != nil {
			errMsg = append(errMsg, err.Error())
			continue
		}

		msg = append(msg, result)
	}

	return msg, errMsg, nil
}

func renderCompact(ctx context.Context, t templ.Component) (string, error) {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)

	if err := t.Render(ctx, buf); err != nil {
		return "", err
	}

	return htmlConcatenizer.ReplaceAllString(buf.String(), ""), nil
}

func streamAttachment(c *echo.Context, filename string, content []byte) error {
	c.Response().Header().Set(echo.HeaderContentDisposition, "attachment; filename=\""+path.Base(filename)+"\"")

	return c.Stream(http.StatusOK, "application/binary", bytes.NewReader(content))
}

func (a *App) redirectHTMXOrFound(c *echo.Context, target string) error {
	if isHtmx(c) {
		c.Response().Header().Set("Hx-Redirect", target)
		return c.String(http.StatusFound, "ok")
	}

	return c.Redirect(http.StatusFound, target)
}

func (a *App) statisticsHandler(c *echo.Context) error {
	u := a.getCurrentUser(c)

	statisticsParams := struct {
		Since     string   `query:"since"`
		Per       string   `query:"per"`
		Type      []string `query:"type"`
		Equipment []uint64 `query:"equipment"`
		Chart     string   `query:"chart"`
		Pause     string   `query:"pause"`
		Segment   uint64   `query:"segment"`
	}{
		Since: "1 year",
		Per:   "month",
		Chart: "line",
		Pause: "1",
	}

	if err := c.Bind(&statisticsParams); err != nil {
		return a.redirectWithError(c, a.Reverse("dashboard"), err)
	}

	if statisticsParams.Since == "" {
		statisticsParams.Since = "1 year"
	}

	if statisticsParams.Chart != "bar" {
		statisticsParams.Chart = "line"
	}

	// The "all" checkbox submits an empty value. Ignore it when specific
	// types are selected as well, so the selection always wins.
	selectedTypes := make([]string, 0, len(statisticsParams.Type))
	for _, t := range statisticsParams.Type {
		if t != "" {
			selectedTypes = append(selectedTypes, t)
		}
	}

	statisticsParams.Type = selectedTypes

	equipment := make([]uint64, 0, len(statisticsParams.Equipment))
	for _, id := range statisticsParams.Equipment {
		if id != 0 {
			equipment = append(equipment, id)
		}
	}
	statisticsParams.Equipment = equipment

	routeSegments, err := database.GetRouteSegments(a.db)
	if err != nil {
		return a.redirectWithError(c, a.Reverse("dashboard"), err)
	}

	typeCounts, err := database.GetWorkoutTypeCounts(a.db, u.ID, u)
	if err != nil {
		return a.redirectWithError(c, a.Reverse("dashboard"), err)
	}

	allEquipment, err := u.GetEquipmentForFilter(a.db)
	if err != nil {
		return a.redirectWithError(c, a.Reverse("dashboard"), err)
	}

	workoutTypes := make([]database.WorkoutType, 0, len(statisticsParams.Type))
	for _, t := range statisticsParams.Type {
		workoutTypes = append(workoutTypes, database.AsWorkoutType(t))
	}

	includePause := statisticsParams.Pause != "0"

	distancePacePoints, err := u.GetDistancePacePoints(database.StatConfig{
		Since:     statisticsParams.Since,
		Equipment: statisticsParams.Equipment,
	}, workoutTypes, includePause, u)
	if err != nil {
		return a.redirectWithError(c, a.Reverse("dashboard"), err)
	}

	return Render(c, http.StatusOK, user.Statistics(user.StatisticsData{
		Since:              statisticsParams.Since,
		Per:                statisticsParams.Per,
		WorkoutTypes:       statisticsParams.Type,
		SelectedEquipment:  statisticsParams.Equipment,
		ChartType:          statisticsParams.Chart,
		RouteSegments:      routeSegments,
		TypeCounts:         typeCounts,
		Equipment:          allEquipment,
		SelectedSegment:    statisticsParams.Segment,
		IncludePause:       includePause,
		DistancePacePoints: distancePacePoints,
	}))
}

func (a *App) dailyDeleteHandler(c *echo.Context) error {
	u := a.getCurrentUser(c)
	d := c.Param("date")

	t, err := time.Parse("2006-01-02", d)
	if err != nil {
		return a.redirectWithError(c, a.Reverse("daily"), err)
	}

	m, err := u.GetMeasurementForDate(t)
	if err != nil {
		return a.redirectWithError(c, a.Reverse("daily"), err)
	}

	if err := m.Delete(a.db); err != nil {
		return a.redirectWithError(c, a.Reverse("daily"), err)
	}

	return a.redirectHTMXOrFound(c, a.Reverse("daily"))
}

func (a *App) dailyUpdateHandler(c *echo.Context) error {
	d := &Measurement{units: a.getCurrentUser(c).PreferredUnits()}
	if err := c.Bind(d); err != nil {
		return a.redirectWithError(c, a.Reverse("daily"), err)
	}

	m, err := a.getCurrentUser(c).GetMeasurementForDate(d.Time())
	if err != nil {
		return a.redirectWithError(c, a.Reverse("daily"), err)
	}

	d.Update(m)

	if err := m.Save(a.db); err != nil {
		return a.redirectWithError(c, a.Reverse("daily"), err)
	}

	return c.Redirect(http.StatusFound, a.Reverse("daily"))
}

func (a *App) dailyHandler(c *echo.Context) error {
	u := a.getCurrentUser(c)

	count := 20
	if cs := c.QueryParam("count"); cs != "" {
		if ci, err := cast.ToIntE(cs); err == nil {
			count = ci
		} else {
			return a.redirectWithError(c, a.Reverse("daily"), err)
		}
	}

	return Render(c, http.StatusOK, user.Daily(u, count))
}

func (a *App) dashboardHandler(c *echo.Context) error {
	u := a.getCurrentUser(c)

	users, err := database.GetUsers(a.db)
	if err != nil {
		return a.redirectWithError(c, a.Reverse("user-signout"), ErrUserNotFound)
	}

	recent, err := database.GetRecentWorkouts(a.db, u, 20)
	if err != nil {
		return a.redirectWithError(c, a.Reverse("user-signout"), ErrUserNotFound)
	}

	return Render(c, http.StatusOK, user.Show(u, users, recent, u))
}

func (a *App) userLoginHandler(c *echo.Context) error {
	return Render(c, http.StatusOK, user.Login())
}

func (a *App) lookupAddressHandler(c *echo.Context) error {
	q := c.FormValue("location")

	results, err := geocoder.Search(q)
	if err != nil {
		a.addErrorT(c, "alerts.something-wrong", i18n.M{"message": err.Error()})
	}

	return Render(c, http.StatusOK, partials.AddressResults(results))
}

func workoutTypesFromQuery(c *echo.Context) []database.WorkoutType {
	raw := c.QueryParams()["type"]
	types := make([]database.WorkoutType, 0, len(raw))

	for _, t := range raw {
		if t != "" {
			types = append(types, database.AsWorkoutType(t))
		}
	}

	return types
}

func equipmentIDsFromQuery(c *echo.Context) []uint64 {
	raw := c.QueryParams()["equipment"]
	ids := make([]uint64, 0, len(raw))

	for _, value := range raw {
		if id := cast.ToUint64(value); id != 0 {
			ids = append(ids, id)
		}
	}

	return ids
}

func (a *App) heatmapHandler(c *echo.Context) error {
	u := a.getCurrentUser(c)
	selected := workoutTypesFromQuery(c)
	selectedEquipment := equipmentIDsFromQuery(c)

	counts, err := database.GetWorkoutTypeCounts(a.db, u.ID, u)
	if err != nil {
		return a.redirectWithError(c, a.Reverse("dashboard"), err)
	}

	equipment, err := u.GetEquipmentForFilter(a.db)
	if err != nil {
		return a.redirectWithError(c, a.Reverse("dashboard"), err)
	}

	return Render(c, http.StatusOK, user.Heatmap(helpers.WorkoutTypesByCount(counts), helpers.WorkoutTypeStrings(selected), selected, equipment, selectedEquipment))
}

func Render(ctx *echo.Context, statusCode int, t templ.Component) error {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)

	if err := t.Render(ctx.Request().Context(), buf); err != nil {
		return err
	}

	return ctx.HTML(statusCode, buf.String())
}
