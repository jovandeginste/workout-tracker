package database

import (
	"context"
	"testing"
	"time"

	"github.com/invopop/ctxi18n"
	"github.com/jovandeginste/workout-tracker/v2/pkg/converters"
	apptranslations "github.com/jovandeginste/workout-tracker/v2/translations"
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

func TestUser_GetAllRecords_SortsByWorkoutCount(t *testing.T) {
	db := createMemoryDB(t)
	u := createStatsUser(t, db)

	for i := range 3 {
		createStatsWorkout(t, db, u, WorkoutTypeCycling, time.Date(2025, 1, i+1, 10, 0, 0, 0, time.UTC), 30000, 100)
	}

	for i := range 2 {
		createStatsWorkout(t, db, u, WorkoutTypeSwimming, time.Date(2025, 2, i+1, 10, 0, 0, 0, time.UTC), 1500, 10)
	}

	createStatsWorkout(t, db, u, WorkoutTypeRunning, time.Date(2025, 3, 1, 10, 0, 0, 0, time.UTC), 5000, 50)

	records, err := u.GetAllRecords(u)
	require.NoError(t, err)

	var order []WorkoutType
	for _, r := range records {
		if r.Active {
			order = append(order, r.WorkoutType)
		}
	}

	assert.Equal(t, []WorkoutType{WorkoutTypeCycling, WorkoutTypeSwimming, WorkoutTypeRunning}, order)
	assert.Equal(t, 3, records[0].Workouts)
	assert.Equal(t, 2, records[1].Workouts)
	assert.Equal(t, 1, records[2].Workouts)
}

func createDistancePaceWorkout(t *testing.T, db *gorm.DB, u *User, wt WorkoutType, date time.Time, distance, speed float64) *Workout {
	t.Helper()

	w := &Workout{
		UserID: u.ID,
		Date:   date,
		Name:   "distance pace workout",
		Type:   wt,
	}
	require.NoError(t, db.Create(w).Error)

	md := &MapData{
		WorkoutID: w.ID,
		WorkoutData: converters.WorkoutData{
			TotalDistance: distance,
			TotalDuration: time.Hour,
			AverageSpeed:  speed,
		},
	}
	require.NoError(t, db.Create(md).Error)

	return w
}

func TestUser_GetDistancePacePoints(t *testing.T) {
	db := createMemoryDB(t)
	u := createStatsUser(t, db)

	require.NoError(t, ctxi18n.LoadWithDefault(apptranslations.FS(), "en"))
	ctx, err := ctxi18n.WithLocale(context.Background(), "en")
	require.NoError(t, err)
	u.SetContext(ctx)

	recent := time.Now().AddDate(0, -1, 0)
	old := time.Now().AddDate(-2, 0, 0)

	createDistancePaceWorkout(t, db, u, WorkoutTypeRunning, recent, 5000, 3)
	createDistancePaceWorkout(t, db, u, WorkoutTypeCycling, recent.Add(time.Hour), 30000, 8)
	createDistancePaceWorkout(t, db, u, WorkoutTypeRunning, old, 4000, 2.5)

	// Workouts without distance or speed are not part of the chart.
	createStatsWorkout(t, db, u, WorkoutTypeWeightLifting, recent.Add(2*time.Hour), 0, 0)

	points, err := u.GetDistancePacePoints(StatConfig{Since: "1 year"}, nil, u)
	require.NoError(t, err)
	require.Len(t, points, 2)

	assert.Equal(t, WorkoutTypeRunning, points[0].WorkoutType)
	assert.Equal(t, 5000.0, points[0].Distance)
	assert.Equal(t, 3.0, points[0].AverageSpeed)
	assert.NotEmpty(t, points[0].LocalWorkoutType)

	assert.Equal(t, WorkoutTypeCycling, points[1].WorkoutType)
	assert.Equal(t, 30000.0, points[1].Distance)
	assert.Equal(t, 8.0, points[1].AverageSpeed)

	points, err = u.GetDistancePacePoints(StatConfig{Since: "1 year"}, []WorkoutType{WorkoutTypeCycling}, u)
	require.NoError(t, err)
	require.Len(t, points, 1)
	assert.Equal(t, WorkoutTypeCycling, points[0].WorkoutType)

	points, err = u.GetDistancePacePoints(StatConfig{Since: "forever"}, nil, u)
	require.NoError(t, err)
	require.Len(t, points, 3)
	assert.Equal(t, WorkoutTypeRunning, points[0].WorkoutType)
	assert.Equal(t, old.Year(), points[0].Date.Year())
}

func TestUser_GetDistancePacePointsFilterByEquipment(t *testing.T) {
	db := createMemoryDB(t)
	u := createStatsUser(t, db)

	require.NoError(t, ctxi18n.LoadWithDefault(apptranslations.FS(), "en"))
	ctx, err := ctxi18n.WithLocale(context.Background(), "en")
	require.NoError(t, err)
	u.SetContext(ctx)

	date := time.Now().AddDate(0, -1, 0)

	withEquipment := createDistancePaceWorkout(t, db, u, WorkoutTypeRunning, date, 5000, 3)
	createDistancePaceWorkout(t, db, u, WorkoutTypeRunning, date.Add(time.Hour), 6000, 3)

	e := &Equipment{Name: "distance pace filter shoes", UserID: u.ID}
	require.NoError(t, db.Create(e).Error)
	require.NoError(t, db.Model(withEquipment).Association("Equipment").Append(e))

	points, err := u.GetDistancePacePoints(StatConfig{Since: "forever", Equipment: []uint64{e.ID}}, nil, u)
	require.NoError(t, err)
	require.Len(t, points, 1)
	assert.Equal(t, 5000.0, points[0].Distance)
}
