package app

import (
	"bytes"
	"net/http"
	"path"
	"strings"

	"github.com/invopop/ctxi18n/i18n"
	"github.com/jovandeginste/workout-tracker/v2/pkg/database"
	"github.com/jovandeginste/workout-tracker/v2/views/route_segments"
	"github.com/labstack/echo/v5"
	"github.com/spf13/cast"
)

func (a *App) addRoutesSegments(g *echo.Group) {
	routeSegmentsGroup := g.Group("/route_segments")
	a.GET(routeSegmentsGroup, "", a.routeSegmentsHandler, "route-segments")
	a.POST(routeSegmentsGroup, "", a.addRouteSegment, "route-segments-create")
	a.GET(routeSegmentsGroup, "/:id", a.routeSegmentsShowHandler, "route-segment-show")
	a.POST(routeSegmentsGroup, "/:id", a.routeSegmentsUpdateHandler, "route-segment-update")
	a.GET(routeSegmentsGroup, "/:id/download", a.routeSegmentsDownloadHandler, "route-segment-download")
	a.GET(routeSegmentsGroup, "/:id/edit", a.routeSegmentsEditHandler, "route-segment-edit")
	a.GET(routeSegmentsGroup, "/:id/delete", a.routeSegmentsDeleteConfirmHandler, "route-segment-delete-confirm")
	a.POST(routeSegmentsGroup, "/:id/delete", a.routeSegmentsDeleteHandler, "route-segment-delete")
	a.POST(routeSegmentsGroup, "/:id/refresh", a.routeSegmentsRefreshHandler, "route-segment-refresh")
	a.POST(routeSegmentsGroup, "/:id/matches", a.routeSegmentFindMatches, "route-segment-matches")
	a.GET(routeSegmentsGroup, "/add", a.routeSegmentsAddHandler, "route-segment-add")
}

func (a *App) routeSegmentsHandler(c *echo.Context) error {
	s, err := database.GetRouteSegments(a.db)
	if err != nil {
		return a.redirectWithError(c, a.Reverse("dashboard"), err)
	}

	for _, rs := range s {
		rs.FilterVisibleMatches(a.getCurrentUser(c))
	}

	return Render(c, http.StatusOK, route_segments.List(s))
}

func (a *App) routeSegmentsAddHandler(c *echo.Context) error {
	return Render(c, http.StatusOK, route_segments.Add())
}

func (a *App) addRouteSegment(c *echo.Context) error {
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
		a.addErrorN(c, "alerts.route-segments-added", len(errMsg), i18n.M{"count": len(errMsg), "list": strings.Join(errMsg, "; ")})
	}

	if len(msg) > 0 {
		a.addNoticeN(c, "notices.route-segments-added", len(msg), i18n.M{"count": len(msg), "list": strings.Join(msg, "; ")})
	}

	return c.Redirect(http.StatusFound, a.Reverse("route-segments"))
}

func (a *App) routeSegmentsShowHandler(c *echo.Context) error {
	rs, err := a.getRouteSegment(c)
	if err != nil {
		return a.redirectWithError(c, a.Reverse("route-segments"), err)
	}

	rs.FilterVisibleMatches(a.getCurrentUser(c))

	return Render(c, http.StatusOK, route_segments.Show(rs))
}

func (a *App) routeSegmentsDownloadHandler(c *echo.Context) error {
	rs, err := a.getRouteSegment(c)
	if err != nil {
		return a.redirectWithError(c, a.Reverse("route-segments"), err)
	}

	basename := path.Base(rs.Filename)

	c.Response().Header().Set(echo.HeaderContentDisposition, "attachment; filename=\""+basename+"\"")

	return c.Stream(http.StatusOK, "application/binary", bytes.NewReader(rs.Content))
}

func (a *App) routeSegmentsEditHandler(c *echo.Context) error {
	rs, err := a.getRouteSegment(c)
	if err != nil {
		return a.redirectWithError(c, a.Reverse("route-segments"), err)
	}

	return Render(c, http.StatusOK, route_segments.Edit(rs))
}

func (a *App) routeSegmentsRefreshHandler(c *echo.Context) error {
	rs, err := a.getRouteSegment(c)
	if err != nil {
		return a.redirectWithError(c, a.Reverse("route-segments"), err)
	}

	if err := rs.UpdateFromContent(); err != nil {
		return a.redirectWithError(c, a.Reverse("route-segment-show", c.Param("id")), err)
	}

	if err := rs.Save(a.db); err != nil {
		return a.redirectWithError(c, a.Reverse("route-segment-show", c.Param("id")), err)
	}

	a.addNoticeT(c, "route-segment.notice.refreshed", rs.Name)

	return c.Redirect(http.StatusFound, a.Reverse("route-segment-show", c.Param("id")))
}

func (a *App) routeSegmentsDeleteHandler(c *echo.Context) error { //nolint:dupl
	rs, err := a.getRouteSegment(c)
	if err != nil {
		return a.redirectWithError(c, a.Reverse("route-segments"), err)
	}

	if err := rs.Delete(a.db); err != nil {
		return a.redirectWithError(c, a.Reverse("route-segment-show", c.Param("id")), err)
	}

	a.addNoticeT(c, "route-segment.notice.deleted", rs.Name)

	if isHtmx(c) {
		c.Response().Header().Set("Hx-Redirect", a.Reverse("route-segments"))
		return c.String(http.StatusFound, "ok")
	}

	return c.Redirect(http.StatusFound, a.Reverse("route-segments"))
}

func (a *App) routeSegmentsUpdateHandler(c *echo.Context) error {
	rs, err := a.getRouteSegment(c)
	if err != nil {
		return a.redirectWithError(c, a.Reverse("route-segments"), err)
	}

	rs.Name = c.FormValue("name")
	rs.Notes = c.FormValue("notes")
	rs.Bidirectional = isChecked(c.FormValue("bidirectional"))
	rs.Circular = isChecked(c.FormValue("circular"))
	rs.Dirty = true

	if err := rs.Save(a.db); err != nil {
		return a.redirectWithError(c, a.Reverse("route-segment-edit", c.Param("id")), err)
	}

	a.addNoticeT(c, "route-segment.notice.updated", rs.Name)

	return c.Redirect(http.StatusFound, a.Reverse("route-segment-show", c.Param("id")))
}

func (a *App) routeSegmentFindMatches(c *echo.Context) error {
	id, err := cast.ToUint64E(c.Param("id"))
	if err != nil {
		return a.redirectWithError(c, a.Reverse("route-segment-show", c.Param("id")), err)
	}

	rs, err := database.GetRouteSegment(a.db, id)
	if err != nil {
		return a.redirectWithError(c, a.Reverse("route-segment-show", c.Param("id")), err)
	}

	rs.Dirty = true
	if err := rs.Save(a.db); err != nil {
		return a.redirectWithError(c, a.Reverse("route-segment-show", c.Param("id")), err)
	}

	a.addNoticeT(c, "route-segment.notice.searching", rs.Name)

	return c.Redirect(http.StatusFound, a.Reverse("route-segment-show", c.Param("id")))
}

func (a *App) routeSegmentsDeleteConfirmHandler(c *echo.Context) error {
	rs, err := a.getRouteSegment(c)
	if err != nil {
		return a.redirectWithError(c, a.Reverse("route-segments"), err)
	}

	return Render(c, http.StatusOK, route_segments.DeleteModal(rs))
}
