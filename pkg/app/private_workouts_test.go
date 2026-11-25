package app

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/jovandeginste/workout-tracker/v2/pkg/database"
	"github.com/labstack/echo/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createUser(t *testing.T, a *App, username, name string, admin bool) *database.User {
	t.Helper()

	u := defaultUser(a.db)
	u.Username = username
	u.Name = name
	u.Admin = admin
	require.NoError(t, u.Save(a.db))

	return u
}

func renderAs(t *testing.T, a *App, u *database.User, target string, h echo.HandlerFunc, params ...echo.PathValue) (int, string) {
	t.Helper()

	req := httptest.NewRequest(http.MethodGet, target, nil)
	rec := httptest.NewRecorder()

	c := a.echo.NewContext(req, rec)
	a.setContext(c)
	c.Set("user_info", u)
	c.SetRequest(c.Request().WithContext(context.WithValue(c.Request().Context(), "user_info", u)))

	if len(params) > 0 {
		c.SetPathValues(params)
	}

	s := sessionLoadAndSave(a.sessionManager)
	require.NoError(t, s(h)(c))

	return rec.Code, rec.Body.String()
}

func makePrivate(t *testing.T, a *App, w *database.Workout) {
	t.Helper()

	require.NoError(t, a.db.Model(w).Update("private", true).Error)
	w.Private = true
}

func TestPrivateWorkouts_ShowIsHiddenFromOtherUsers(t *testing.T) {
	a := configuredApp(t)

	owner := createUser(t, a, "owner", "owner", false)
	other := createUser(t, a, "other", "other", false)
	admin := createUser(t, a, "boss", "boss", true)

	w := createAPIWorkout(t, a, owner, "secret workout", database.WorkoutTypeRunning, time.Date(2025, 1, 1, 10, 0, 0, 0, time.UTC))
	makePrivate(t, a, w)

	params := echo.PathValue{Name: "id", Value: strconv.FormatUint(w.ID, 10)}

	status, body := renderAs(t, a, owner, a.Reverse("workout-show", w.ID), a.workoutsShowHandler, params)
	assert.Equal(t, http.StatusOK, status)
	assert.Contains(t, body, "secret workout")

	status, body = renderAs(t, a, other, a.Reverse("workout-show", w.ID), a.workoutsShowHandler, params)
	assert.Equal(t, http.StatusFound, status)
	assert.NotContains(t, body, "secret workout")

	status, body = renderAs(t, a, admin, a.Reverse("workout-show", w.ID), a.workoutsShowHandler, params)
	assert.Equal(t, http.StatusOK, status)
	assert.Contains(t, body, "secret workout")
}

func TestPrivateWorkouts_AreHiddenFromDashboards(t *testing.T) {
	a := configuredApp(t)

	owner := createUser(t, a, "owner", "owner", false)
	other := createUser(t, a, "other", "other", false)
	admin := createUser(t, a, "boss", "boss", true)

	private := createAPIWorkout(t, a, owner, "secret workout", database.WorkoutTypeRunning, time.Date(2025, 1, 1, 10, 0, 0, 0, time.UTC))
	makePrivate(t, a, private)

	createAPIWorkout(t, a, owner, "public workout", database.WorkoutTypeRunning, time.Date(2025, 1, 2, 10, 0, 0, 0, time.UTC))

	t.Run("the owner sees everything on the dashboard", func(t *testing.T) {
		status, body := renderAs(t, a, owner, a.Reverse("dashboard"), a.dashboardHandler)

		require.Equal(t, http.StatusOK, status)
		assert.Contains(t, body, "secret workout")
		assert.Contains(t, body, "public workout")
	})

	t.Run("other users do not see private workouts", func(t *testing.T) {
		status, body := renderAs(t, a, other, a.Reverse("dashboard"), a.dashboardHandler)

		require.Equal(t, http.StatusOK, status)
		assert.NotContains(t, body, "secret workout")
		assert.Contains(t, body, "public workout")
	})

	t.Run("admins do see private workouts", func(t *testing.T) {
		status, body := renderAs(t, a, admin, a.Reverse("dashboard"), a.dashboardHandler)

		require.Equal(t, http.StatusOK, status)
		assert.Contains(t, body, "secret workout")
	})
}

func TestPrivateWorkouts_StatisticsAreViewerAware(t *testing.T) {
	a := configuredApp(t)

	owner := createUser(t, a, "owner", "owner", false)
	other := createUser(t, a, "other", "other", false)
	admin := createUser(t, a, "boss", "boss", true)

	private := createAPIWorkout(t, a, owner, "secret workout", database.WorkoutTypeRunning, time.Date(2025, 1, 1, 10, 0, 0, 0, time.UTC))
	makePrivate(t, a, private)

	// Give the private workout a distinctive distance, so we can tell whether it
	// is included in the statistics.
	require.NoError(t, a.db.Model(&database.MapData{}).
		Where("workout_id = ?", private.ID).
		Update("total_distance", 42000).Error)

	createAPIWorkout(t, a, owner, "public workout", database.WorkoutTypeRunning, time.Date(2025, 1, 2, 10, 0, 0, 0, time.UTC))

	params := echo.PathValue{Name: "id", Value: strconv.FormatUint(owner.ID, 10)}

	// Statistics are calculated over the workouts the viewer may see: another
	// user only sees the public workouts, the owner and admins see everything.
	status, body := renderAs(t, a, other, a.Reverse("user-show", owner.ID), a.userShowHandler, params)
	require.Equal(t, http.StatusOK, status)
	assert.Contains(t, body, "1 km")
	assert.NotContains(t, body, "42 km")

	status, body = renderAs(t, a, admin, a.Reverse("user-show", owner.ID), a.userShowHandler, params)
	require.Equal(t, http.StatusOK, status)
	assert.Contains(t, body, "42 km")

	status, body = renderAs(t, a, owner, a.Reverse("user-show", owner.ID), a.userShowHandler, params)
	require.Equal(t, http.StatusOK, status)
	assert.Contains(t, body, "42 km")
}

func TestPrivateWorkouts_AreHiddenFromEquipmentPages(t *testing.T) {
	a := configuredApp(t)

	owner := createUser(t, a, "owner", "owner", false)
	other := createUser(t, a, "other", "other", false)

	e := &database.Equipment{UserID: owner.ID, Name: "my bike"}
	require.NoError(t, a.db.Create(e).Error)

	private := createAPIWorkout(t, a, owner, "secret workout", database.WorkoutTypeCycling, time.Date(2025, 1, 1, 10, 0, 0, 0, time.UTC))
	makePrivate(t, a, private)

	public := createAPIWorkout(t, a, owner, "public workout", database.WorkoutTypeCycling, time.Date(2025, 1, 2, 10, 0, 0, 0, time.UTC))

	require.NoError(t, a.db.Model(e).Association("Workouts").Append(private, public))

	params := echo.PathValue{Name: "id", Value: strconv.FormatUint(e.ID, 10)}

	status, body := renderAs(t, a, other, a.Reverse("equipment-show", e.ID), a.equipmentShowHandler, params)
	require.Equal(t, http.StatusOK, status)
	assert.NotContains(t, body, "secret workout")
	assert.Contains(t, body, "public workout")
}

func TestPrivateWorkouts_CanBeSetThroughTheAPI(t *testing.T) {
	a := configuredApp(t)

	u := defaultAPIUser(a.db)
	w := createAPIWorkout(t, a, u, "workout", database.WorkoutTypeRunning, time.Date(2025, 1, 1, 10, 0, 0, 0, time.UTC))

	ts := httptest.NewServer(a.echo)
	defer ts.Close()

	status, _ := apiWorkoutRequest(t, ts, http.MethodPatch, a.Reverse("api-workout-update", w.ID), `{"private":true}`)
	require.Equal(t, http.StatusOK, status)

	var stored database.Workout
	require.NoError(t, a.db.First(&stored, w.ID).Error)
	assert.True(t, stored.Private)

	status, _ = apiWorkoutRequest(t, ts, http.MethodPatch, a.Reverse("api-workout-update", w.ID), `{"private":false}`)
	require.Equal(t, http.StatusOK, status)

	require.NoError(t, a.db.First(&stored, w.ID).Error)
	assert.False(t, stored.Private)
}

func TestPrivateWorkouts_CanBeSetThroughTheWebForm(t *testing.T) {
	a := configuredApp(t)

	owner := createUser(t, a, "owner", "owner", false)
	w := createAPIWorkout(t, a, owner, "workout", database.WorkoutTypeRunning, time.Date(2025, 1, 1, 10, 0, 0, 0, time.UTC))

	update := func(private bool) {
		form := url.Values{
			"name": {"workout"},
			"type": {database.WorkoutTypeRunning.String()},
		}
		if private {
			form.Set("private", "true")
		}

		req := httptest.NewRequest(http.MethodPost, a.Reverse("workout-update", w.ID), strings.NewReader(form.Encode()))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
		rec := httptest.NewRecorder()

		c := a.echo.NewContext(req, rec)
		a.setContext(c)
		c.Set("user_info", owner)
		c.SetRequest(c.Request().WithContext(context.WithValue(c.Request().Context(), "user_info", owner)))
		c.SetPathValues(echo.PathValues{{Name: "id", Value: strconv.FormatUint(w.ID, 10)}})

		s := sessionLoadAndSave(a.sessionManager)
		require.NoError(t, s(a.workoutsUpdateHandler)(c))

		var stored database.Workout
		require.NoError(t, a.db.First(&stored, w.ID).Error)
		assert.Equal(t, private, stored.Private)
	}

	update(true)
	update(false)
}

func shareViaRouter(t *testing.T, a *App, u *database.User, uuid string) (int, string) {
	t.Helper()

	req := httptest.NewRequest(http.MethodGet, a.Reverse("share", uuid), nil)

	if u != nil {
		tokenReq := httptest.NewRequest(http.MethodGet, "/", nil)
		tokenRec := httptest.NewRecorder()
		require.NoError(t, a.createToken(u, a.echo.NewContext(tokenReq, tokenRec)))

		cookies := tokenRec.Result().Cookies()
		require.NotEmpty(t, cookies)

		req.AddCookie(cookies[0])
	}

	rec := httptest.NewRecorder()

	a.echo.ServeHTTP(rec, req)

	return rec.Code, rec.Body.String()
}

func TestPrivateWorkouts_SharedLinkIsHidden(t *testing.T) {
	a := configuredApp(t)

	owner := createUser(t, a, "owner", "owner", false)
	other := createUser(t, a, "other", "other", false)
	admin := createUser(t, a, "boss", "boss", true)

	privateUUID := "0e8a9a9a-3f0d-4d4b-9d0e-0a2a1a1a1a1a"
	publicUUID := "1f9b0b0b-4a1e-4e5c-8e1f-1b3b2b2b2b2b"

	privateWorkout := createAPIWorkout(t, a, owner, "secret workout", database.WorkoutTypeRunning, time.Date(2025, 1, 1, 10, 0, 0, 0, time.UTC))
	makePrivate(t, a, privateWorkout)
	require.NoError(t, a.db.Model(privateWorkout).Update("public_uuid", privateUUID).Error)

	publicWorkout := createAPIWorkout(t, a, owner, "public workout", database.WorkoutTypeRunning, time.Date(2025, 1, 2, 10, 0, 0, 0, time.UTC))
	require.NoError(t, a.db.Model(publicWorkout).Update("public_uuid", publicUUID).Error)

	t.Run("a public workout is still visible to anonymous visitors", func(t *testing.T) {
		status, body := shareViaRouter(t, a, nil, publicUUID)

		assert.Equal(t, http.StatusOK, status)
		assert.Contains(t, body, "public workout")
	})

	t.Run("a private workout is hidden from anonymous visitors", func(t *testing.T) {
		status, body := shareViaRouter(t, a, nil, privateUUID)

		assert.Equal(t, http.StatusFound, status)
		assert.NotContains(t, body, "secret workout")
	})

	t.Run("a private workout is hidden from other users", func(t *testing.T) {
		status, body := shareViaRouter(t, a, other, privateUUID)

		assert.Equal(t, http.StatusFound, status)
		assert.NotContains(t, body, "secret workout")
	})

	t.Run("the owner sees the shared private workout", func(t *testing.T) {
		status, body := shareViaRouter(t, a, owner, privateUUID)

		assert.Equal(t, http.StatusOK, status)
		assert.Contains(t, body, "secret workout")
	})

	t.Run("an admin sees the shared private workout", func(t *testing.T) {
		status, body := shareViaRouter(t, a, admin, privateUUID)

		assert.Equal(t, http.StatusOK, status)
		assert.Contains(t, body, "secret workout")
	})
}

func TestPrivateWorkouts_RouteSegmentCreationIsOwnerScoped(t *testing.T) {
	a := configuredApp(t)

	owner := createUser(t, a, "owner", "owner", false)
	other := createUser(t, a, "other", "other", false)

	w := createAPIWorkout(t, a, owner, "secret workout", database.WorkoutTypeRunning, time.Date(2025, 1, 1, 10, 0, 0, 0, time.UTC))
	makePrivate(t, a, w)

	var md database.MapData

	require.NoError(t, a.db.Where("workout_id = ?", w.ID).First(&md).Error)
	require.NoError(t, a.db.Create(&database.MapDataDetails{
		MapDataID: md.ID,
		Points: []database.MapPoint{
			{Lat: 50, Lng: 4},
			{Lat: 50.001, Lng: 4},
			{Lat: 50.002, Lng: 4},
		},
	}).Error)

	createSegment := func(u *database.User) int {
		form := url.Values{"name": {"segment"}, "start": {"1"}, "end": {"3"}}

		req := httptest.NewRequest(http.MethodPost, a.Reverse("workout-route-segment-create", w.ID), strings.NewReader(form.Encode()))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
		rec := httptest.NewRecorder()

		c := a.echo.NewContext(req, rec)
		a.setContext(c)
		c.Set("user_info", u)
		c.SetRequest(c.Request().WithContext(context.WithValue(c.Request().Context(), "user_info", u)))
		c.SetPathValues(echo.PathValues{{Name: "id", Value: strconv.FormatUint(w.ID, 10)}})

		s := sessionLoadAndSave(a.sessionManager)
		require.NoError(t, s(a.workoutsCreateRouteSegmentFromWorkoutHandler)(c))

		return rec.Code
	}

	segments := func() int64 {
		t.Helper()

		var count int64
		require.NoError(t, a.db.Model(&database.RouteSegment{}).Count(&count).Error)

		return count
	}

	assert.Equal(t, http.StatusFound, createSegment(other))
	assert.Zero(t, segments())

	assert.Equal(t, http.StatusFound, createSegment(owner))
	assert.EqualValues(t, 1, segments())
}
