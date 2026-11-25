package database

import (
	"testing"
	"time"

	"github.com/jovandeginste/workout-tracker/v2/pkg/converters"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func liftTestPoints() []MapPoint {
	points := make([]MapPoint, 0, 25)
	var totalDistance float64
	var totalDuration time.Duration

	for i := range 25 {
		elevation := 1000.0
		if i >= 5 {
			elevation += float64(i-4) * 10
		}

		points = append(points, MapPoint{
			Lat:             45.0 + float64(i)*0.00027,
			Lng:             6.0,
			Elevation:       elevation,
			Distance:        30,
			Distance2D:      30,
			Duration:        15 * time.Second,
			TotalDistance:   totalDistance,
			TotalDistance2D: totalDistance,
			TotalDuration:   totalDuration,
		})

		totalDistance += 30
		totalDuration += 15 * time.Second
	}

	return points
}

func TestDetectLiftSegments(t *testing.T) {
	lifts := DetectLiftSegments(liftTestPoints())

	require.Len(t, lifts, 1)
	assert.Equal(t, SlopeKindLift, lifts[0].Type)
	assert.InDelta(t, 200, lifts[0].Gain, 1)
	assert.GreaterOrEqual(t, lifts[0].Duration, liftMinDuration)
}

func TestAutoDetectLiftsOnlyForSkiing(t *testing.T) {
	assert.True(t, WorkoutTypeSkiing.AutoDetectLifts())
	assert.True(t, WorkoutTypeSnowboarding.AutoDetectLifts())
	assert.False(t, WorkoutTypeCycling.AutoDetectLifts())
	assert.False(t, WorkoutTypeRunning.AutoDetectLifts())
}

func TestApplyLiftExclusions(t *testing.T) {
	points := liftTestPoints()
	data := &MapData{
		WorkoutData: converters.WorkoutData{TotalUp: 200},
		Details:     &MapDataDetails{Points: points},
	}

	totalDistance := points[len(points)-1].TotalDistance2D
	totalDuration := points[len(points)-1].TotalDuration

	distance, _, duration := applyLiftExclusions(data, WorkoutTypeSkiing, totalDistance, totalDistance, totalDuration)

	require.Len(t, data.Lifts, 1)
	assert.Less(t, distance, totalDistance)
	assert.Less(t, duration, totalDuration)
	assert.Zero(t, data.TotalUp)
}
