package database

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWorkoutFilters_SetDefaultsEquipment(t *testing.T) {
	filters := WorkoutFilters{Equipment: []uint64{0, 3, 0, 5}}

	filters.setDefaults()

	assert.Equal(t, []uint64{3, 5}, filters.Equipment)
}

func TestWorkoutFilters_Equipment(t *testing.T) {
	db := createMemoryDB(t)
	u := createStatsUser(t, db)

	workout := func(name string, date time.Time) *Workout {
		w := &Workout{
			UserID: u.ID,
			Date:   date,
			Name:   name,
			Type:   WorkoutTypeRunning,
		}
		require.NoError(t, db.Create(w).Error)

		return w
	}

	w1 := workout("first", time.Date(2025, 1, 1, 10, 0, 0, 0, time.UTC))
	w2 := workout("second", time.Date(2025, 1, 2, 10, 0, 0, 0, time.UTC))

	e1 := &Equipment{Name: "filter test bike", UserID: u.ID}
	e2 := &Equipment{Name: "filter test shoes", UserID: u.ID}
	require.NoError(t, db.Create(e1).Error)
	require.NoError(t, db.Create(e2).Error)

	require.NoError(t, db.Model(w1).Association("Equipment").Append(e1))
	require.NoError(t, db.Model(w2).Association("Equipment").Append(e2))

	tests := []struct {
		name      string
		equipment []uint64
		want      []uint64
	}{
		{name: "no filter", want: []uint64{w2.ID, w1.ID}},
		{name: "one equipment", equipment: []uint64{e1.ID}, want: []uint64{w1.ID}},
		{name: "multiple equipment", equipment: []uint64{e1.ID, e2.ID}, want: []uint64{w2.ID, w1.ID}},
		{name: "unknown equipment", equipment: []uint64{e1.ID + e2.ID + 1000}, want: []uint64{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filters := WorkoutFilters{Equipment: tt.equipment}

			workouts, err := GetWorkouts(filters.ToQuery(db))
			require.NoError(t, err)

			ids := make([]uint64, len(workouts))
			for i, w := range workouts {
				ids[i] = w.ID
			}

			assert.Equal(t, tt.want, ids)
		})
	}
}
