package app

import (
	"strconv"
	"testing"
	"time"

	"github.com/jovandeginste/workout-tracker/v2/pkg/database"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWorkoutListFiltersByEquipment(t *testing.T) {
	a := configuredApp(t)

	u := defaultUser(a.db)
	require.NoError(t, u.Save(a.db))

	withEquipment := createAPIWorkout(t, a, u, "with equipment", database.WorkoutTypeRunning, time.Date(2025, 1, 1, 10, 0, 0, 0, time.UTC))
	withoutEquipment := createAPIWorkout(t, a, u, "without equipment", database.WorkoutTypeRunning, time.Date(2025, 1, 2, 10, 0, 0, 0, time.UTC))

	e := &database.Equipment{Name: "filter test bike", UserID: u.ID}
	require.NoError(t, a.db.Create(e).Error)
	require.NoError(t, a.db.Model(withEquipment).Association("Equipment").Append(e))

	_, body := renderAs(t, a, u, a.Reverse("workouts"), a.workoutsHandler)
	assert.Contains(t, body, `name="equipment"`)
	assert.Contains(t, body, e.Name)

	target := a.Reverse("workouts") + "?equipment=" + strconv.FormatUint(e.ID, 10)
	_, body = renderAs(t, a, u, target, a.workoutsHandler)

	assert.Contains(t, body, withEquipment.Name)
	assert.NotContains(t, body, withoutEquipment.Name)
}

func TestStatisticsAndHeatmapShowEquipmentFilter(t *testing.T) {
	a := configuredApp(t)

	u := defaultUser(a.db)
	require.NoError(t, u.Save(a.db))

	e := &database.Equipment{Name: "filter test shoes", UserID: u.ID}
	require.NoError(t, a.db.Create(e).Error)

	_, body := renderAs(t, a, u, a.Reverse("statistics"), a.statisticsHandler)
	assert.Contains(t, body, `name="equipment"`)
	assert.Contains(t, body, e.Name)

	_, body = renderAs(t, a, u, a.Reverse("heatmap"), a.heatmapHandler)
	assert.Contains(t, body, `name="equipment"`)
	assert.Contains(t, body, e.Name)
}
