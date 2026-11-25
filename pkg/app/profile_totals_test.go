package app

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/jovandeginste/workout-tracker/v2/pkg/database"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func updateUserProfile(t *testing.T, a *App, u *database.User, form url.Values) {
	t.Helper()

	req := httptest.NewRequest(http.MethodPost, a.Reverse("user-profile-update"), strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rec := httptest.NewRecorder()

	c := a.echo.NewContext(req, rec)
	a.setContext(c)
	a.setCurrentUser(c, u)

	s := sessionLoadAndSave(a.sessionManager)
	h := s(a.userProfileUpdateHandler)

	require.NoError(t, h(c))
	assert.Equal(t, http.StatusFound, rec.Code)
}

func TestRoute_UserProfileTotalsShowFoldable(t *testing.T) {
	a := configuredApp(t)

	u := defaultUser(a.db)
	u.Profile.TotalsShow = database.WorkoutTypeList{database.WorkoutTypeRunning, database.WorkoutTypeCycling}
	require.NoError(t, u.Save(a.db))

	c, rec := newTestContext(t, a, a.Reverse("user-profile"), u)

	s := sessionLoadAndSave(a.sessionManager)
	h := s(a.userProfileHandler)

	require.NoError(t, h(c))
	require.Equal(t, http.StatusOK, rec.Code)

	body := rec.Body.String()
	assert.Contains(t, body, "<details>")
	assert.Contains(t, body, "running, cycling")
	assert.Contains(t, body, `class="sr-only peer"`)
	assert.NotContains(t, body, `class="hidden peer"`)
	assert.Contains(t, body, `name="totals_show" value="running" checked`)
	assert.Contains(t, body, `name="totals_show" value="cycling" checked`)
}

func TestRoute_UserProfileUpdateTotalsShow(t *testing.T) {
	t.Run("stores multiple workout types", func(t *testing.T) {
		a := configuredApp(t)

		u := defaultUser(a.db)
		require.NoError(t, u.Save(a.db))

		form := url.Values{}
		form.Add("totals_show", database.WorkoutTypeRunning.String())
		form.Add("totals_show", database.WorkoutTypeCycling.String())

		updateUserProfile(t, a, u, form)

		var p database.Profile
		require.NoError(t, a.db.First(&p, u.Profile.ID).Error)

		assert.Equal(t, database.WorkoutTypeList{database.WorkoutTypeRunning, database.WorkoutTypeCycling}, p.TotalsShow)
	})

	t.Run("defaults to auto-detect when none are selected", func(t *testing.T) {
		a := configuredApp(t)

		u := defaultUser(a.db)
		u.Profile.TotalsShow = database.WorkoutTypeList{database.WorkoutTypeRunning, database.WorkoutTypeCycling}
		require.NoError(t, u.Save(a.db))

		updateUserProfile(t, a, u, url.Values{})

		var p database.Profile
		require.NoError(t, a.db.First(&p, u.Profile.ID).Error)

		assert.Equal(t, database.WorkoutTypeList{database.WorkoutTypeAutoDetect}, p.TotalsShow)
	})
}
