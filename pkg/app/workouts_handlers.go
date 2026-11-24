package app

import (
	"bytes"
	"errors"
	"net/http"
	"path"
	"time"

	"github.com/google/uuid"
	"github.com/jovandeginste/workout-tracker/v2/pkg/database"
	"github.com/jovandeginste/workout-tracker/v2/views/workouts"
	"github.com/labstack/echo/v5"
	"github.com/spf13/cast"
)

func (a *App) addRoutesWorkouts(g *echo.Group) {
	workoutsGroup := g.Group("/workouts")
	a.GET(workoutsGroup, "", a.workoutsHandler, "workouts")
	a.POST(workoutsGroup, "", a.addWorkout, "workouts-create")
	a.GET(workoutsGroup, "/:id", a.workoutsShowHandler, "workout-show")
	a.POST(workoutsGroup, "/:id", a.workoutsUpdateHandler, "workout-update")
	a.GET(workoutsGroup, "/:id/download", a.workoutsDownloadHandler, "workout-download")
	a.GET(workoutsGroup, "/:id/edit", a.workoutsEditHandler, "workout-edit")
	a.GET(workoutsGroup, "/:id/delete", a.workoutsDeleteConfirmHandler, "workout-delete-confirm")
	a.POST(workoutsGroup, "/:id/delete", a.workoutsDeleteHandler, "workout-delete")
	a.POST(workoutsGroup, "/:id/refresh", a.workoutsRefreshHandler, "workout-refresh")
	a.POST(workoutsGroup, "/:id/share", a.workoutsShareHandler, "workout-share")
	a.DELETE(workoutsGroup, "/:id/share", a.workoutsShareDeleteHandler, "workout-share-delete")
	a.GET(workoutsGroup, "/:id/route-segment", a.workoutsCreateRouteSegmentHandler, "workout-route-segment")
	a.POST(workoutsGroup, "/:id/route-segment", a.workoutsCreateRouteSegmentFromWorkoutHandler, "workout-route-segment-create")
	a.POST(workoutsGroup, "/:id/toggle-lock", a.workoutsToggleLockHandler, "workout-toggle-lock")
	a.GET(workoutsGroup, "/add", a.workoutsAddHandler, "workout-add")
	a.GET(workoutsGroup, "/form", a.workoutsFormHandler, "workout-form")
}

func (a *App) workoutsHandler(c *echo.Context) error {
	u := a.getCurrentUser(c)
	if u.IsAnonymous() {
		return a.redirectWithError(c, a.Reverse("user-signout"), ErrUserNotFound)
	}

	filters, err := database.GetWorkoutsFilters(c)
	if err != nil {
		return a.redirectWithError(c, a.Reverse("dashboard"), err)
	}

	target, err := a.userForUsername(c, filters.Username)
	if err != nil {
		a.addErrorT(c, "translation.No_user_found_with_username_s", filters.Username)

		return Render(c, http.StatusOK, workouts.List(nil, filters))
	}

	w, err := target.GetWorkouts(database.VisibleWorkouts(filters.ToQuery(a.db), u))
	if err != nil {
		return a.redirectWithError(c, a.Reverse("dashboard"), err)
	}

	return Render(c, http.StatusOK, workouts.List(w, filters))
}

func (a *App) workoutsShowHandler(c *echo.Context) error {
	id, err := cast.ToUint64E(c.Param("id"))
	if err != nil {
		return a.redirectWithError(c, a.Reverse("workouts"), err)
	}

	w, err := database.GetWorkoutDetails(a.db, id)
	if err != nil {
		return a.redirectWithError(c, a.Reverse("workouts"), err)
	}

	if !w.IsVisibleTo(a.getCurrentUser(c)) {
		a.addErrorT(c, "translation.Workout_not_found")

		return c.Redirect(http.StatusFound, a.Reverse("workouts"))
	}

	return Render(c, http.StatusOK, workouts.Show(w))
}

func (a *App) workoutsAddHandler(c *echo.Context) error {
	return Render(c, http.StatusOK, workouts.Add())
}

func (a *App) workoutsFormHandler(c *echo.Context) error {
	w := &database.Workout{}

	if c.FormValue("id") != "" {
		id, err := cast.ToUint64E(c.FormValue("id"))
		if err != nil {
			return Render(c, http.StatusOK, workouts.Form(w))
		}

		w, err = a.getCurrentUser(c).GetWorkout(a.db, id)
		if err != nil {
			return Render(c, http.StatusOK, workouts.Form(w))
		}
	}

	if w.Type == "" || c.FormValue("type") != "" {
		w.Type = database.WorkoutType(c.FormValue("type"))
	}

	if w.Date.IsZero() {
		w.Date = time.Now()
	}

	if w.Name == "" {
		w.Name = a.i18nT(c, w.Type.StringT()) + " - " + w.Date.Format(time.RFC3339)
	}

	return Render(c, http.StatusOK, workouts.Form(w))
}

func (a *App) workoutsDeleteHandler(c *echo.Context) error { //nolint:dupl
	workout, err := a.getWorkout(c)
	if err != nil {
		return a.redirectWithError(c, a.Reverse("workout-show", c.Param("id")), err)
	}

	if err := workout.Delete(a.db); err != nil {
		return a.redirectWithError(c, a.Reverse("workout-show", c.Param("id")), err)
	}

	a.addNoticeT(c, "translation.The_workout_s_has_been_deleted", workout.Name)

	if isHtmx(c) {
		c.Response().Header().Set("Hx-Redirect", a.Reverse("workouts"))
		return c.String(http.StatusFound, "ok")
	}

	return c.Redirect(http.StatusFound, a.Reverse("workouts"))
}

func (a *App) workoutShowShared(c *echo.Context) error {
	u, err := uuid.Parse(c.Param("uuid"))
	if err != nil {
		return a.redirectWithError(c, a.Reverse("workouts"), err)
	}

	w, err := database.GetWorkoutDetailsByUUID(a.db, u)
	if err != nil {
		return a.redirectWithError(c, a.Reverse("workouts"), err)
	}

	if !w.IsVisibleTo(a.getCurrentUser(c)) {
		a.addErrorT(c, "translation.Workout_not_found")

		return c.Redirect(http.StatusFound, a.Reverse("dashboard"))
	}

	return Render(c, http.StatusOK, workouts.Show(w))
}

func (a *App) workoutsShareHandler(c *echo.Context) error {
	id, err := cast.ToUint64E(c.Param("id"))
	if err != nil {
		return a.redirectWithError(c, a.Reverse("workout-show", c.Param("id")), err)
	}

	workout, err := a.getCurrentUser(c).GetWorkout(a.db, id)
	if err != nil {
		return a.redirectWithError(c, a.Reverse("workout-show", c.Param("id")), err)
	}

	u := uuid.New()
	workout.PublicUUID = &u

	if err := workout.Save(a.db); err != nil {
		return a.redirectWithError(c, a.Reverse("workout-show", c.Param("id")), err)
	}

	a.addNoticeT(c, "translation.The_workout_s_now_has_a_shareable_link", workout.Name)

	return c.Redirect(http.StatusFound, a.Reverse("workout-show", c.Param("id")))
}

func (a *App) workoutsShareDeleteHandler(c *echo.Context) error {
	id, err := cast.ToUint64E(c.Param("id"))
	if err != nil {
		return a.redirectWithError(c, a.Reverse("workout-show", c.Param("id")), err)
	}

	workout, err := a.getCurrentUser(c).GetWorkout(a.db, id)
	if err != nil {
		return a.redirectWithError(c, a.Reverse("workout-show", c.Param("id")), err)
	}

	workout.PublicUUID = nil

	if err := workout.Save(a.db); err != nil {
		return a.redirectWithError(c, a.Reverse("workout-show", c.Param("id")), err)
	}

	a.addNoticeT(c, "translation.The_shareable_link_is_deleted_for_the_workout_s", workout.Name)

	return c.Redirect(http.StatusFound, a.Reverse("workout-show", c.Param("id")))
}

func (a *App) workoutsRefreshHandler(c *echo.Context) error {
	id, err := cast.ToUint64E(c.Param("id"))
	if err != nil {
		return a.redirectWithError(c, a.Reverse("workout-show", c.Param("id")), err)
	}

	workout, err := a.getCurrentUser(c).GetWorkout(a.db, id)
	if err != nil {
		return a.redirectWithError(c, a.Reverse("workout-show", c.Param("id")), err)
	}

	workout.Dirty = true
	if err := workout.Save(a.db); err != nil {
		return a.redirectWithError(c, a.Reverse("workout-show", c.Param("id")), err)
	}

	a.addNoticeT(c, "translation.The_workout_s_will_be_refreshed_soon", workout.Name)

	return c.Redirect(http.StatusFound, a.Reverse("workout-show", c.Param("id")))
}

func (a *App) workoutsDownloadHandler(c *echo.Context) error {
	workout, err := a.getWorkout(c)
	if err != nil {
		return a.redirectWithError(c, a.Reverse("workouts"), err)
	}

	if !workout.HasFile() {
		return a.redirectWithError(c, a.Reverse("workouts"), errors.New("workout has no content"))
	}

	basename := path.Base(workout.GPX.Filename)

	c.Response().Header().Set(echo.HeaderContentDisposition, "attachment; filename=\""+basename+"\"")

	return c.Stream(http.StatusOK, "application/binary", bytes.NewReader(workout.GPX.Content))
}

func (a *App) workoutsEditHandler(c *echo.Context) error {
	w, err := a.getWorkout(c)
	if err != nil {
		return a.redirectWithError(c, a.Reverse("workouts"), err)
	}

	return Render(c, http.StatusOK, workouts.Edit(w))
}

func (a *App) workoutsToggleLockHandler(c *echo.Context) error {
	workout, err := a.getWorkout(c)
	if err != nil {
		return a.redirectWithError(c, a.Reverse("workouts"), err)
	}

	workout.Locked = !workout.Locked

	if err := workout.Save(a.db); err != nil {
		return a.redirectWithError(c, a.Reverse("workout-show", c.Param("id")), err)
	}

	if workout.Locked {
		a.addNoticeT(c, "translation.WorkoutLocked", workout.Name)
	} else {
		a.addNoticeT(c, "translation.WorkoutUnLocked", workout.Name)
	}

	return c.Redirect(http.StatusFound, a.Reverse("workout-show", c.Param("id")))
}

func (a *App) workoutsCreateRouteSegmentFromWorkoutHandler(c *echo.Context) error {
	id, err := cast.ToUint64E(c.Param("id"))
	if err != nil {
		return a.redirectWithError(c, a.Reverse("workouts"), err)
	}

	workout, err := database.GetWorkoutDetails(a.db, id)
	if err != nil {
		return a.redirectWithError(c, a.Reverse("workouts"), err)
	}

	var params database.RoutSegmentCreationParams

	if err := c.Bind(&params); err != nil {
		return a.redirectWithError(c, a.Reverse("workouts"), err)
	}

	content, err := database.RouteSegmentFromPoints(workout, &params)
	if err != nil {
		return a.redirectWithError(c, a.Reverse("workouts"), err)
	}

	rs, err := database.AddRouteSegment(a.db, "", params.Filename(), content)
	if err != nil {
		return a.redirectWithError(c, a.Reverse("workouts"), err)
	}

	a.addNoticeT(c, "translation.The_route_segment_s_has_been_created_we_search_for_matches_in_the_background", rs.Name)

	return c.Redirect(http.StatusFound, a.Reverse("route-segment-show", rs.ID))
}

func (a *App) workoutsCreateRouteSegmentHandler(c *echo.Context) error {
	w, err := a.getWorkout(c)
	if err != nil {
		return a.redirectWithError(c, a.Reverse("workouts"), err)
	}

	return Render(c, http.StatusOK, workouts.CreateRouteSegment(w))
}

func (a *App) workoutsDeleteConfirmHandler(c *echo.Context) error {
	w, err := a.getWorkout(c)
	if err != nil {
		return a.redirectWithError(c, a.Reverse("workouts"), err)
	}

	return Render(c, http.StatusOK, workouts.DeleteModal(w))
}
