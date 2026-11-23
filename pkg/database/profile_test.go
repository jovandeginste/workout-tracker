package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProfile_TotalsShowRoundTrip(t *testing.T) {
	db := createMemoryDB(t)
	u := createStatsUser(t, db)

	u.Profile.TotalsShow = WorkoutTypeList{WorkoutTypeRunning, WorkoutTypeCycling, WorkoutTypeRunning, "invalid"}

	require.NoError(t, u.Profile.Save(db))

	var p Profile
	require.NoError(t, db.First(&p, u.Profile.ID).Error)

	assert.Equal(t, WorkoutTypeList{WorkoutTypeRunning, WorkoutTypeCycling}, p.TotalsShow)
}

func TestMigrateTotalsShow(t *testing.T) {
	tests := []struct {
		name     string
		stored   any
		expected WorkoutTypeList
	}{
		{name: "legacy single type", stored: "running", expected: WorkoutTypeList{WorkoutTypeRunning}},
		{name: "legacy auto", stored: "auto", expected: WorkoutTypeList{WorkoutTypeAutoDetect}},
		{name: "json list", stored: `["running","cycling"]`, expected: WorkoutTypeList{WorkoutTypeRunning, WorkoutTypeCycling}},
		{name: "empty", stored: "", expected: nil},
		{name: "null", stored: nil, expected: nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := createMemoryDB(t)
			u := createStatsUser(t, db)

			require.NoError(t, db.Table("profiles").Where("id = ?", u.Profile.ID).Update("totals_show", tt.stored).Error)
			require.NoError(t, migrateTotalsShow(db))

			var p Profile
			require.NoError(t, db.First(&p, u.Profile.ID).Error)

			assert.Equal(t, tt.expected, p.TotalsShow)
		})
	}
}

func TestMigrateTotalsShow_Idempotent(t *testing.T) {
	db := createMemoryDB(t)
	u := createStatsUser(t, db)

	require.NoError(t, db.Table("profiles").Where("id = ?", u.Profile.ID).Update("totals_show", "cycling").Error)
	require.NoError(t, migrateTotalsShow(db))
	require.NoError(t, migrateTotalsShow(db))

	var p Profile
	require.NoError(t, db.First(&p, u.Profile.ID).Error)

	assert.Equal(t, WorkoutTypeList{WorkoutTypeCycling}, p.TotalsShow)
}
