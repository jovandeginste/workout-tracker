package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func init() { //nolint:gochecknoinits
	online = false
}

func defaultWorkout(t *testing.T) *Workout {
	u := defaultUser()
	f1, err := gpxFS.ReadFile("sample1.gpx")
	require.NoError(t, err)

	w := NewWorkout(
		u,
		WorkoutTypeAutoDetect,
		"some notes",
		f1,
	)

	require.NoError(t, err)

	return w
}

func TestWorkout_ParseWithType(t *testing.T) {
	u := defaultUser()
	f1, err := gpxFS.ReadFile("sample1.gpx")
	require.NoError(t, err)

	w := NewWorkout(
		u,
		WorkoutTypeWalking,
		"some notes",
		f1,
	)

	require.NoError(t, err)

	assert.NotNil(t, w)
	assert.Equal(t, WorkoutTypeWalking, w.Type)
}

func TestWorkout_Parse(t *testing.T) {
	w := defaultWorkout(t)

	assert.NotNil(t, w)
	assert.Equal(t, WorkoutTypeRunning, w.Type)
	assert.Equal(t, "Garmin Connect", w.Data.Creator)
	assert.Equal(t, "some notes", w.Notes)
	assert.InDelta(t, 39, w.Data.Center.Lat, 1)
	assert.InDelta(t, -77, w.Data.Center.Lng, 1)

	assert.Len(t, w.Data.Points, 206)
	assert.InDelta(t, 3125, w.Data.TotalDistance, 1)
	assert.InDelta(t, 3.297, w.Data.AverageSpeed(), 0.01)
	assert.InDelta(t, 3.297, w.Data.AverageSpeedNoPause(), 0.01)
	assert.Equal(t, "Untitled", w.Name)
	assert.Nil(t, w.Data.Address)
}

func TestWorkout_UpdateData(t *testing.T) {
	db := createMemoryDB(t)
	w := defaultWorkout(t)

	require.NoError(t, w.Save(db))

	ud := w.UpdatedAt
	d := w.Data

	w.Data = dummyMapData()
	require.NoError(t, w.Save(db))

	assert.NotEqual(t, d, w.Data)
	assert.NotEqual(t, ud, w.UpdatedAt)
	ud = w.UpdatedAt

	require.NoError(t, w.UpdateData(db))
	assert.Equal(t, d, w.Data)
	assert.NotEqual(t, ud, w.UpdatedAt)
}

func TestWorkout_SaveAndGet(t *testing.T) {
	db := createMemoryDB(t)
	w := defaultWorkout(t)

	assert.Zero(t, w.UpdatedAt)
	require.NoError(t, w.Save(db))
	assert.NotZero(t, w.UpdatedAt)

	newW, err := GetWorkout(db, int(w.ID))
	require.NoError(t, err)
	assert.Equal(t, w.ID, newW.ID)
	assert.Equal(t, w.Data, newW.Data)
}
