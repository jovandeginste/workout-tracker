package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

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

func TestParseWorkoutWithType(t *testing.T) {
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

func TestParseWorkout(t *testing.T) {
	w := defaultWorkout(t)

	assert.NotNil(t, w)
	assert.Equal(t, WorkoutTypeRunning, w.Type)

	assert.Len(t, w.Data.Points, 206)
	assert.InDelta(t, 3125, w.Data.TotalDistance, 1)
	assert.InDelta(t, 3.297, w.Data.AverageSpeed(), 0.01)
	assert.InDelta(t, 3.297, w.Data.AverageSpeedNoPause(), 0.01)
	assert.Equal(t, "Untitled", w.Name)
	assert.Equal(t, "US", w.Data.Address.CountryCode)
	assert.Equal(t, "Washington", w.Data.Address.City)
}
