package workouts

import (
	"context"
	"testing"
	"time"

	"github.com/jovandeginste/workout-tracker/v2/pkg/database"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddDetailsMovingDuration(t *testing.T) {
	w := &database.Workout{
		Data: &database.MapData{
			Details: &database.MapDataDetails{
				Points: []database.MapPoint{
					{Duration: 0, TotalDuration: 0},
					{Duration: time.Minute, TotalDuration: time.Minute},                     // standing still
					{Duration: time.Minute, TotalDuration: 2 * time.Minute},                 // standing still
					{Duration: time.Minute, TotalDuration: 3 * time.Minute, Distance: 1000}, // moving
					{Duration: time.Minute, TotalDuration: 4 * time.Minute, Distance: 2000}, // moving
				},
			},
		},
	}

	data := map[string]*dataset{
		"position":        {Data: []any{}},
		"time":            {Data: []any{}},
		"distance":        {Data: []any{}},
		"speed":           {Data: []any{}},
		"duration":        {Data: []any{}},
		"moving-duration": {Data: []any{}},
		"slope":           {Data: []any{}},
	}

	addDetails(context.Background(), data, w)

	require.Len(t, data["moving-duration"].Data, len(w.Data.Details.Points))

	// The pauses are not added to the moving duration.
	assert.Equal(t, 0.0, data["moving-duration"].Data[1])
	assert.Equal(t, 0.0, data["moving-duration"].Data[2])
	assert.Equal(t, 60.0, data["moving-duration"].Data[3])
	assert.Equal(t, 120.0, data["moving-duration"].Data[4])

	// The total duration keeps counting.
	assert.Equal(t, 240.0, data["duration"].Data[4])
}
