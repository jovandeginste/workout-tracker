package app

import (
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/jovandeginste/workout-tracker/v2/pkg/converters"
	"github.com/jovandeginste/workout-tracker/v2/pkg/database"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createDashboardWorkout(t *testing.T, a *App, u *database.User, wt database.WorkoutType, date time.Time) {
	t.Helper()

	w := &database.Workout{
		UserID: u.ID,
		Date:   date,
		Name:   "workout",
		Type:   wt,
	}
	require.NoError(t, a.db.Create(w).Error)

	md := &database.MapData{
		WorkoutID: w.ID,
		WorkoutData: converters.WorkoutData{
			TotalDistance: 1000,
			TotalUp:       100,
			TotalDuration: time.Hour,
		},
	}
	require.NoError(t, a.db.Create(md).Error)
}

func TestRoute_DashboardMultipleTotals(t *testing.T) {
	a := configuredApp(t)

	u := defaultUser(a.db)
	u.Profile.TotalsShow = database.WorkoutTypeList{database.WorkoutTypeRunning, database.WorkoutTypeCycling}
	require.NoError(t, u.Save(a.db))

	createDashboardWorkout(t, a, u, database.WorkoutTypeRunning, time.Date(2025, 1, 1, 10, 0, 0, 0, time.UTC))
	createDashboardWorkout(t, a, u, database.WorkoutTypeCycling, time.Date(2025, 2, 1, 10, 0, 0, 0, time.UTC))

	c, rec := newTestContext(t, a, a.Reverse("dashboard"), u)

	s := sessionLoadAndSave(a.sessionManager)
	h := s(a.dashboardHandler)

	require.NoError(t, h(c))
	require.Equal(t, http.StatusOK, rec.Code)

	// Each selected workout type renders one row of four statistic tiles.
	assert.Equal(t, 8, strings.Count(rec.Body.String(), `class="md:basis-1/2 lg:basis-1/4"`))
}
