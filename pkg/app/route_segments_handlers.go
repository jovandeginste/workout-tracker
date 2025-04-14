package app

import (
	"bytes"
	"net/http"
	"path"
	"strconv"
	"strings"

	"github.com/invopop/ctxi18n/i18n"
	"github.com/jovandeginste/workout-tracker/v2/pkg/database"
	"github.com/jovandeginste/workout-tracker/v2/views/route_segments"
	"github.com/labstack/echo/v4"
	"github.com/stackus/hxgo/hxecho"
)

func (a *App) addRoutesSegments(e *echo.Group) {
	routeSegmentsGroup := e.Group("/route_segments")
	routeSegmentsGroup.GET("", a.routeSegmentsHandler).Name = "route-segments"
	routeSegmentsGroup.POST("", a.addRouteSegment).Name = "route-segments-create"
	routeSegmentsGroup.GET("/:id", a.routeSegmentsShowHandler).Name = "route-segment-show"
	routeSegmentsGroup.POST("/:id", a.routeSegmentsUpdateHandler).Name = "route-segment-update"
	routeSegmentsGroup.GET("/:id/download", a.routeSegmentsDownloadHandler).Name = "route-segment-download"
	routeSegmentsGroup.GET("/:id/edit", a.routeSegmentsEditHandler).Name = "route-segment-edit"
	routeSegmentsGroup.GET("/:id/delete", a.routeSegmentsDeleteConfirmHandler).Name = "route-segment-delete-confirm"
	routeSegmentsGroup.POST("/:id/delete", a.routeSegmentsDeleteHandler).Name = "route-segment-delete"
	routeSegmentsGroup.POST("/:id/refresh", a.routeSegmentsRefreshHandler).Name = "route-segment-refresh"
	routeSegmentsGroup.POST("/:id/matches", a.routeSegmentFindMatches).Name = "route-segment-matches"
	routeSegmentsGroup.GET("/add", a.routeSegmentsAddHandler).Name = "route-segment-add"
}

func (a *App) routeSegmentsHandler(c echo.Context) error {
	s, err := database.GetRouteSegments(a.db)
	if err != nil {
		return a.redirectWithError(c, a.echo.Reverse("dashboard"), err)
	}

	return Render(c, http.StatusOK, route_segments.List(s))
}

func (a *App) routeSegmentsAddHandler(c echo.Context) error {
	return Render(c, http.StatusOK, route_segments.Add())
}

func (a *App) addRouteSegment(c echo.Context) error {
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	files := form.File["file"]

	msg := []string{}
	errMsg := []string{}

	for _, file := range files {
		content, parseErr := uploadedFile(file)
		if parseErr != nil {
			errMsg = append(errMsg, parseErr.Error())
			continue
		}

		notes := c.FormValue("notes")

		w, addErr := database.AddRouteSegment(a.db, notes, file.Filename, content)
		if addErr != nil {
			errMsg = append(errMsg, addErr.Error())
			continue
		}

		msg = append(msg, w.Name)
	}

	if len(errMsg) > 0 {
		a.addErrorN(c, "alerts.route_segments_added", len(errMsg), i18n.M{"count": len(errMsg), "list": strings.Join(errMsg, "; ")})
	}

	if len(msg) > 0 {
		a.addNoticeN(c, "notices.route_segments_added", len(msg), i18n.M{"count": len(msg), "list": strings.Join(msg, "; ")})
	}

	return c.Redirect(http.StatusFound, a.echo.Reverse("route-segments"))
}

func (a *App) routeSegmentsShowHandler(c echo.Context) error {
	rs, err := a.getRouteSegment(c)
	if err != nil {
		return a.redirectWithError(c, a.echo.Reverse("route-segments"), err)
	}

	return Render(c, http.StatusOK, route_segments.Show(rs))
}

func (a *App) routeSegmentsDownloadHandler(c echo.Context) error {
	rs, err := a.getRouteSegment(c)
	if err != nil {
		return a.redirectWithError(c, a.echo.Reverse("route-segments"), err)
	}

	basename := path.Base(rs.Filename)

	c.Response().Header().Set(echo.HeaderContentDisposition, "attachment; filename=\""+basename+"\"")

	return c.Stream(http.StatusOK, "application/binary", bytes.NewReader(rs.Content))
}

func (a *App) routeSegmentsEditHandler(c echo.Context) error {
	rs, err := a.getRouteSegment(c)
	if err != nil {
		return a.redirectWithError(c, a.echo.Reverse("route-segments"), err)
	}

	return Render(c, http.StatusOK, route_segments.Edit(rs))
}

func (a *App) routeSegmentsRefreshHandler(c echo.Context) error {
	rs, err := a.getRouteSegment(c)
	if err != nil {
		return a.redirectWithError(c, a.echo.Reverse("route-segments"), err)
	}

	if err := rs.UpdateFromContent(); err != nil {
		return a.redirectWithError(c, a.echo.Reverse("route-segment-show", c.Param("id")), err)
	}

	if err := rs.Save(a.db); err != nil {
		return a.redirectWithError(c, a.echo.Reverse("route-segment-show", c.Param("id")), err)
	}

	a.addNoticeT(c, "translation.The_workout_s_has_been_refreshed", rs.Name)

	return c.Redirect(http.StatusFound, a.echo.Reverse("route-segment-show", c.Param("id")))
}

func (a *App) routeSegmentsDeleteHandler(c echo.Context) error { //nolint:dupl
	rs, err := a.getRouteSegment(c)
	if err != nil {
		return a.redirectWithError(c, a.echo.Reverse("route-segments"), err)
	}

	if err := rs.Delete(a.db); err != nil {
		return a.redirectWithError(c, a.echo.Reverse("route-segment-show", c.Param("id")), err)
	}

	a.addNoticeT(c, "translation.The_workout_s_has_been_deleted", rs.Name)

	if hxecho.IsHtmx(c) {
		c.Response().Header().Set("Hx-Redirect", a.echo.Reverse("route-segments"))
		return c.String(http.StatusFound, "ok")
	}

	return c.Redirect(http.StatusFound, a.echo.Reverse("route-segments"))
}

func (a *App) routeSegmentsUpdateHandler(c echo.Context) error {
	rs, err := a.getRouteSegment(c)
	if err != nil {
		return a.redirectWithError(c, a.echo.Reverse("route-segments"), err)
	}

	rs.Name = c.FormValue("name")
	rs.Notes = c.FormValue("notes")
	rs.Bidirectional = isChecked(c.FormValue("bidirectional"))
	rs.Circular = isChecked(c.FormValue("circular"))
	rs.Dirty = true

	if err := rs.Save(a.db); err != nil {
		return a.redirectWithError(c, a.echo.Reverse("route-segment-edit", c.Param("id")), err)
	}

	a.addNoticeT(c, "translation.The_route_segment_s_has_been_updated", rs.Name)

	return c.Redirect(http.StatusFound, a.echo.Reverse("route-segment-show", c.Param("id")))
}

func (a *App) routeSegmentFindMatches(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return a.redirectWithError(c, a.echo.Reverse("route-segment-show", c.Param("id")), err)
	}

	rs, err := database.GetRouteSegment(a.db, id)
	if err != nil {
		return a.redirectWithError(c, a.echo.Reverse("route-segment-show", c.Param("id")), err)
	}

	rs.Dirty = true
	if err := rs.Save(a.db); err != nil {
		return a.redirectWithError(c, a.echo.Reverse("route-segment-show", c.Param("id")), err)
	}

	a.addNoticeT(c, "translation.Start_searching_in_the_background_for_matching_workouts_for_route_segment_s", rs.Name)

	return c.Redirect(http.StatusFound, a.echo.Reverse("route-segment-show", c.Param("id")))
}

func (a *App) routeSegmentsDeleteConfirmHandler(c echo.Context) error {
	rs, err := a.getRouteSegment(c)
	if err != nil {
		return a.redirectWithError(c, a.echo.Reverse("route-segments"), err)
	}

	return Render(c, http.StatusOK, route_segments.DeleteModal(rs))
}
