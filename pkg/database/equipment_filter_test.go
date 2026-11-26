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

func createFilterWorkout(t *testing.T, db *gorm.DB, u *User, name string, date time.Time, distance float64) (*Workout, *MapData) {
	t.Helper()

	w := &Workout{
		UserID: u.ID,
		Date:   date,
		Name:   name,
		Type:   WorkoutTypeRunning,
	}
	require.NoError(t, db.Create(w).Error)

	md := &MapData{
		WorkoutID: w.ID,
		WorkoutData: converters.WorkoutData{
			TotalDistance: distance,
			TotalDuration: time.Hour,
		},
	}
	require.NoError(t, db.Create(md).Error)

	return w, md
}

func TestStatisticsFilterByEquipment(t *testing.T) {
	db := createMemoryDB(t)
	u := createStatsUser(t, db)

	require.NoError(t, ctxi18n.LoadWithDefault(apptranslations.FS(), "en"))
	ctx, err := ctxi18n.WithLocale(context.Background(), "en")
	require.NoError(t, err)
	u.SetContext(ctx)

	withEquipment, _ := createFilterWorkout(t, db, u, "with equipment", time.Date(2025, 1, 1, 10, 0, 0, 0, time.UTC), 1000)
	createFilterWorkout(t, db, u, "without equipment", time.Date(2025, 2, 1, 10, 0, 0, 0, time.UTC), 2000)

	e := &Equipment{Name: "statistics filter bike", UserID: u.ID}
	require.NoError(t, db.Create(e).Error)
	require.NoError(t, db.Model(withEquipment).Association("Equipment").Append(e))

	stats, err := u.GetStatistics(StatConfig{Since: "forever", Per: "month", Equipment: []uint64{e.ID}}, u)
	require.NoError(t, err)
	require.Len(t, stats.Buckets, 1)

	buckets, ok := stats.Buckets[WorkoutTypeRunning]
	require.True(t, ok)
	require.Len(t, buckets.Buckets, 1)

	for _, bucket := range buckets.Buckets {
		assert.Equal(t, 1, bucket.Workouts)
		assert.Equal(t, 1000.0, bucket.Distance)
	}
}

func TestHeatmapPointsFilterByEquipment(t *testing.T) {
	db := createMemoryDB(t)
	u := createStatsUser(t, db)

	withEquipment, md1 := createFilterWorkout(t, db, u, "with equipment", time.Date(2025, 1, 1, 10, 0, 0, 0, time.UTC), 1000)
	_, md2 := createFilterWorkout(t, db, u, "without equipment", time.Date(2025, 2, 1, 10, 0, 0, 0, time.UTC), 2000)

	e := &Equipment{Name: "heatmap filter bike", UserID: u.ID}
	require.NoError(t, db.Create(e).Error)
	require.NoError(t, db.Model(withEquipment).Association("Equipment").Append(e))

	require.NoError(t, db.Create(&MapDataDetails{MapDataID: md1.ID, HeatmapPoints: []HeatmapPoint{{1, 2}}}).Error)
	require.NoError(t, db.Create(&MapDataDetails{MapDataID: md2.ID, HeatmapPoints: []HeatmapPoint{{3, 4}}}).Error)

	points, err := u.GetHeatmapPointsForTypes(db, nil, []uint64{e.ID})
	require.NoError(t, err)
	assert.Equal(t, []HeatmapPoint{{1, 2}}, points)
}
