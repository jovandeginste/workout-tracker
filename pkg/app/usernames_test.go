package app

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/jovandeginste/workout-tracker/v2/pkg/database"
	"github.com/labstack/echo/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createUser(t *testing.T, a *App, username, name string) *database.User {
	t.Helper()

	u := defaultUser(a.db)
	u.Username = username
	u.Name = name
	require.NoError(t, u.Save(a.db))

	return u
}

func renderForUser(t *testing.T, a *App, u *database.User, target string, h echo.HandlerFunc, params ...echo.PathValue) string {
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
	require.Equal(t, http.StatusOK, rec.Code)

	return rec.Body.String()
}

func TestUsernames_WorkoutListShowsUsername(t *testing.T) {
	a := configuredApp(t)

	u := defaultUser(a.db)
	require.NoError(t, u.Save(a.db))

	createAPIWorkout(t, a, u, "my workout", database.WorkoutTypeRunning, time.Date(2025, 1, 1, 10, 0, 0, 0, time.UTC))

	body := renderForUser(t, a, u, a.Reverse("workouts"), a.workoutsHandler)

	assert.Contains(t, body, u.Username)
	assert.Contains(t, body, "my workout")
	assert.Contains(t, body, `name="username"`)
}

func TestUsernames_WorkoutShowShowsUsernameAndName(t *testing.T) {
	a := configuredApp(t)

	u := defaultUser(a.db)
	require.NoError(t, u.Save(a.db))

	w := createAPIWorkout(t, a, u, "my workout", database.WorkoutTypeRunning, time.Date(2025, 1, 1, 10, 0, 0, 0, time.UTC))

	body := renderForUser(t, a, u, a.Reverse("workout-show", w.ID), a.workoutsShowHandler,
		echo.PathValue{Name: "id", Value: strconv.FormatUint(w.ID, 10)})

	assert.Contains(t, body, u.Username)
	assert.Contains(t, body, u.Name)
}

func TestUsernames_WorkoutListFiltersByUsername(t *testing.T) {
	a := configuredApp(t)

	me := defaultUser(a.db)
	require.NoError(t, me.Save(a.db))

	other := createUser(t, a, "other-user", "Other Name")

	createAPIWorkout(t, a, me, "my workout", database.WorkoutTypeRunning, time.Date(2025, 1, 1, 10, 0, 0, 0, time.UTC))
	createAPIWorkout(t, a, other, "other workout", database.WorkoutTypeRunning, time.Date(2025, 1, 2, 10, 0, 0, 0, time.UTC))

	t.Run("without a username it shows my own workouts", func(t *testing.T) {
		body := renderForUser(t, a, me, a.Reverse("workouts"), a.workoutsHandler)

		assert.Contains(t, body, "my workout")
		assert.NotContains(t, body, "other workout")
	})

	t.Run("with a username it shows the workouts of that user", func(t *testing.T) {
		target := a.Reverse("workouts") + "?username=" + other.Username
		body := renderForUser(t, a, me, target, a.workoutsHandler)

		assert.Contains(t, body, "other workout")
		assert.NotContains(t, body, "my workout")
	})

	t.Run("with an unknown username it shows an empty list and a message", func(t *testing.T) {
		target := a.Reverse("workouts") + "?username=does-not-exist"
		body := renderForUser(t, a, me, target, a.workoutsHandler)

		assert.Contains(t, body, "No user found with username does-not-exist")
		assert.NotContains(t, body, "my workout")
		assert.NotContains(t, body, "other workout")
	})
}

func TestUsernames_UserSuggestions(t *testing.T) {
	a := configuredApp(t)

	me := defaultUser(a.db)
	require.NoError(t, me.Save(a.db))

	other := createUser(t, a, "other-user", "Other Name")

	t.Run("without a query it suggests all users", func(t *testing.T) {
		body := renderForUser(t, a, me, a.Reverse("user-suggest"), a.userSuggestHandler)

		assert.Contains(t, body, `value="`+me.Username+`"`)
		assert.Contains(t, body, `value="`+other.Username+`"`)
	})

	t.Run("it matches on the username", func(t *testing.T) {
		target := a.Reverse("user-suggest") + "?username=other"
		body := renderForUser(t, a, me, target, a.userSuggestHandler)

		assert.Contains(t, body, `value="`+other.Username+`"`)
		assert.NotContains(t, body, `value="`+me.Username+`"`)
	})

	t.Run("it matches on the name", func(t *testing.T) {
		target := a.Reverse("user-suggest") + "?username=Other+Name"
		body := renderForUser(t, a, me, target, a.userSuggestHandler)

		assert.Contains(t, body, `value="`+other.Username+`"`)
		assert.NotContains(t, body, `value="`+me.Username+`"`)
	})

	t.Run("it filters out unknown users", func(t *testing.T) {
		target := a.Reverse("user-suggest") + "?username=does-not-exist"
		body := renderForUser(t, a, me, target, a.userSuggestHandler)

		assert.NotContains(t, body, "<option")
	})

	t.Run("it is reachable through the router", func(t *testing.T) {
		tokenReq := httptest.NewRequest(http.MethodGet, "/", nil)
		tokenRec := httptest.NewRecorder()
		require.NoError(t, a.createToken(me, a.echo.NewContext(tokenReq, tokenRec)))

		cookies := tokenRec.Result().Cookies()
		require.Len(t, cookies, 1)

		req := httptest.NewRequest(http.MethodGet, a.Reverse("user-suggest")+"?username=other", nil)
		req.AddCookie(cookies[0])
		rec := httptest.NewRecorder()

		a.echo.ServeHTTP(rec, req)

		require.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), `value="`+other.Username+`"`)
	})
}

func TestUsernames_EquipmentListShowsUsername(t *testing.T) {
	a := configuredApp(t)

	u := defaultUser(a.db)
	require.NoError(t, u.Save(a.db))

	e := &database.Equipment{UserID: u.ID, Name: "my bike"}
	require.NoError(t, a.db.Create(e).Error)

	body := renderForUser(t, a, u, a.Reverse("equipment"), a.equipmentHandler)

	assert.Contains(t, body, u.Username)
	assert.Contains(t, body, "my bike")
	assert.Contains(t, body, `name="username"`)
}
