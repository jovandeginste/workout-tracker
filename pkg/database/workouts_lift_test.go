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

// liftRunPoints builds a monotonic run. When stopDuration is set, even points
// are stops; the odd points move moveDistance in moveDuration and change the
// elevation by elevationStep. The elevation of a stop does not change.
func liftRunPoints(startElevation, elevationStep, moveDistance float64, moveDuration, stopDuration time.Duration, count int) []MapPoint {
	points := make([]MapPoint, 0, count)

	var totalDistance float64
	var totalDuration time.Duration

	elevation := startElevation

	for i := range count {
		point := MapPoint{
			Lat:             45.0 + float64(i)*0.0002,
			Lng:             6.0,
			Elevation:       elevation,
			TotalDistance:   totalDistance,
			TotalDistance2D: totalDistance,
			TotalDuration:   totalDuration,
		}

		stop := stopDuration > 0 && i > 0 && i%2 == 0

		if i > 0 {
			if stop {
				point.Duration = stopDuration
			} else {
				elevation += elevationStep
				point.Elevation = elevation
				point.Distance = moveDistance
				point.Distance2D = moveDistance
				point.Duration = moveDuration
			}

			totalDistance += point.Distance
			totalDuration += point.Duration
			point.TotalDistance = totalDistance
			point.TotalDistance2D = totalDistance
			point.TotalDuration = totalDuration
		}

		points = append(points, point)
	}

	return points
}

func TestDetectLiftSegmentsIgnoresStoppedDescent(t *testing.T) {
	// A ski run: fast descent, but stopping every few seconds.
	points := liftRunPoints(1500, -10, 10, time.Second, 2*time.Second, 62)

	assert.Empty(t, DetectLiftSegments(points))
}

func TestDetectLiftSegmentsIgnoresUnevenDescent(t *testing.T) {
	// A ski run without stops, but with strongly varying speed.
	points := make([]MapPoint, 0, 120)

	var totalDistance float64
	var totalDuration time.Duration
	elevation := 1500.0

	for i := range 120 {
		point := MapPoint{
			Lat:             45.0 + float64(i)*0.0002,
			Lng:             6.0,
			Elevation:       elevation,
			TotalDistance:   totalDistance,
			TotalDistance2D: totalDistance,
			TotalDuration:   totalDuration,
		}

		if i > 0 {
			distance := 1.0
			if i%2 == 0 {
				distance = 15
			}

			elevation -= 5
			point.Distance = distance
			point.Distance2D = distance
			point.Duration = time.Second

			totalDistance += distance
			totalDuration += time.Second
		}

		points = append(points, point)
	}

	assert.Empty(t, DetectLiftSegments(points))
}

func TestDetectLiftSegmentsIgnoresElevationJump(t *testing.T) {
	// A steady horizontal movement where the altitude sticks and then jumps:
	// this is bad barometer data, not a lift.
	points := make([]MapPoint, 0, 40)

	var totalDistance float64
	var totalDuration time.Duration
	elevation := 1000.0

	for i := range 40 {
		if i == 20 {
			elevation -= 100
		}

		point := MapPoint{
			Lat:             45.0 + float64(i)*0.0002,
			Lng:             6.0,
			Elevation:       elevation,
			TotalDistance:   totalDistance,
			TotalDistance2D: totalDistance,
			TotalDuration:   totalDuration,
		}

		if i > 0 {
			point.Distance = 20
			point.Distance2D = 20
			point.Duration = 5 * time.Second

			totalDistance += point.Distance
			totalDuration += point.Duration
			point.TotalDistance = totalDistance
			point.TotalDistance2D = totalDistance
			point.TotalDuration = totalDuration
		}

		points = append(points, point)
	}

	assert.Empty(t, DetectLiftSegments(points))
}

func TestDetectLiftSegmentsFindsDescendingLift(t *testing.T) {
	// A cable car going down, e.g. across a valley.
	points := liftRunPoints(1500, -10, 20, 5*time.Second, 0, 40)

	lifts := DetectLiftSegments(points)

	require.Len(t, lifts, 1)
	assert.Equal(t, SlopeKindLift, lifts[0].Type)
	assert.InDelta(t, 390, lifts[0].Gain, 1)
}

func TestApplyLiftExclusions(t *testing.T) {
	points := liftTestPoints()
	totalDistance := points[len(points)-1].TotalDistance2D
	totalDuration := points[len(points)-1].TotalDuration

	data := &MapData{
		WorkoutData: converters.WorkoutData{
			TotalDistance:   totalDistance,
			TotalDistance2D: totalDistance,
			TotalDuration:   totalDuration,
			TotalUp:         200,
		},
		Details: &MapDataDetails{Points: points},
	}

	data.Lifts = DetectLiftSegments(points)
	applyLiftExclusions(data, WorkoutTypeSkiing)

	require.Len(t, data.Lifts, 1)
	assert.Less(t, data.TotalDistance, totalDistance)
	assert.Less(t, data.TotalDuration, totalDuration)
	assert.Zero(t, data.TotalUp)
}
