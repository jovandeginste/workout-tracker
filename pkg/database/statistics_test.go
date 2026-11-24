package database

import (
	"testing"
	"time"

	"github.com/jovandeginste/workout-tracker/v2/pkg/converters"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func createStatsUser(t *testing.T, db *gorm.DB) *User {
	t.Helper()

	u := defaultUser()
	u.SetDB(db)

	require.NoError(t, u.Create(db))

	u.Profile.UserID = u.ID
	require.NoError(t, u.Profile.Save(db))

	return u
}

func createStatsWorkout(t *testing.T, db *gorm.DB, u *User, wt WorkoutType, date time.Time, distance, up float64) {
	t.Helper()

	w := &Workout{
		UserID: u.ID,
		Date:   date,
		Name:   "workout",
		Type:   wt,
	}
	require.NoError(t, db.Create(w).Error)

	md := &MapData{
		WorkoutID: w.ID,
		WorkoutData: converters.WorkoutData{
			TotalDistance: distance,
			TotalUp:       up,
			TotalDuration: time.Hour,
		},
	}
	require.NoError(t, db.Create(md).Error)
}

func TestUser_GetDefaultTotals_MultipleTypes(t *testing.T) {
	db := createMemoryDB(t)
	u := createStatsUser(t, db)

	createStatsWorkout(t, db, u, WorkoutTypeRunning, time.Date(2025, 1, 1, 10, 0, 0, 0, time.UTC), 1000, 100)
	createStatsWorkout(t, db, u, WorkoutTypeRunning, time.Date(2025, 2, 1, 10, 0, 0, 0, time.UTC), 2000, 200)
	createStatsWorkout(t, db, u, WorkoutTypeCycling, time.Date(2025, 3, 1, 10, 0, 0, 0, time.UTC), 30000, 300)

	u.Profile.TotalsShow = WorkoutTypeList{WorkoutTypeCycling, WorkoutTypeRunning}

	totals, err := u.GetDefaultTotals(u)
	require.NoError(t, err)
	require.Len(t, totals, 2)

	assert.Equal(t, WorkoutTypeCycling, totals[0].WorkoutType)
	assert.Equal(t, 1, totals[0].Workouts)
	assert.Equal(t, 30000.0, totals[0].Distance)
	assert.Equal(t, 300.0, totals[0].Up)
	assert.Equal(t, time.Hour, totals[0].Duration)

	assert.Equal(t, WorkoutTypeRunning, totals[1].WorkoutType)
	assert.Equal(t, 2, totals[1].Workouts)
	assert.Equal(t, 3000.0, totals[1].Distance)
	assert.Equal(t, 300.0, totals[1].Up)
	assert.Equal(t, 2*time.Hour, totals[1].Duration)
}

func TestUser_GetDefaultTotals_AutoDetect(t *testing.T) {
	db := createMemoryDB(t)
	u := createStatsUser(t, db)

	createStatsWorkout(t, db, u, WorkoutTypeRunning, time.Date(2025, 1, 1, 10, 0, 0, 0, time.UTC), 1000, 100)
	createStatsWorkout(t, db, u, WorkoutTypeCycling, time.Date(2025, 2, 1, 10, 0, 0, 0, time.UTC), 30000, 300)
	createStatsWorkout(t, db, u, WorkoutTypeCycling, time.Date(2025, 3, 1, 10, 0, 0, 0, time.UTC), 30000, 300)

	u.Profile.TotalsShow = WorkoutTypeList{WorkoutTypeAutoDetect, WorkoutTypeCycling}

	totals, err := u.GetDefaultTotals(u)
	require.NoError(t, err)
	require.Len(t, totals, 1, "auto-detect should resolve to the most used type and dedupe")

	assert.Equal(t, WorkoutTypeCycling, totals[0].WorkoutType)
	assert.Equal(t, 2, totals[0].Workouts)
}

func TestUser_GetDefaultTotals_EmptyDefaultsToAutoDetect(t *testing.T) {
	db := createMemoryDB(t)
	u := createStatsUser(t, db)

	createStatsWorkout(t, db, u, WorkoutTypeSwimming, time.Date(2025, 1, 1, 10, 0, 0, 0, time.UTC), 1500, 10)

	u.Profile.TotalsShow = nil

	totals, err := u.GetDefaultTotals(u)
	require.NoError(t, err)
	require.Len(t, totals, 1)

	assert.Equal(t, WorkoutTypeSwimming, totals[0].WorkoutType)
}

func TestUser_GetDefaultTotals_WithoutWorkouts(t *testing.T) {
	db := createMemoryDB(t)
	u := createStatsUser(t, db)

	u.Profile.TotalsShow = WorkoutTypeList{WorkoutTypeSwimming}

	totals, err := u.GetDefaultTotals(u)
	require.NoError(t, err)
	require.Len(t, totals, 1)

	assert.Equal(t, WorkoutTypeSwimming, totals[0].WorkoutType)
	assert.Equal(t, 0, totals[0].Workouts)
	assert.Zero(t, totals[0].Distance)
	assert.Zero(t, totals[0].Up)
}

func TestUser_GetDefaultTotals_AutoDetectWithoutWorkouts(t *testing.T) {
	db := createMemoryDB(t)
	u := createStatsUser(t, db)

	u.Profile.TotalsShow = WorkoutTypeList{WorkoutTypeAutoDetect}

	totals, err := u.GetDefaultTotals(u)
	require.NoError(t, err)
	require.Len(t, totals, 1)

	assert.Equal(t, WorkoutTypeAutoDetect, totals[0].WorkoutType)
	assert.Equal(t, 0, totals[0].Workouts)
}

func TestUser_GetDefaultTotals_Anonymous(t *testing.T) {
	totals, err := AnonymousUser().GetDefaultTotals(nil)

	require.ErrorIs(t, err, ErrAnonymousUser)
	assert.Nil(t, totals)
}

func TestUser_GetTotalsFor_Empty(t *testing.T) {
	totals, err := (&User{}).GetTotalsFor(nil, nil)

	require.NoError(t, err)
	assert.Nil(t, totals)
}
