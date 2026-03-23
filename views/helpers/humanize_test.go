package helpers

import (
	"context"
	"testing"

	"github.com/jovandeginste/workout-tracker/v2/pkg/database"
	"github.com/stretchr/testify/assert"
)

func TestHumanDistanceForWorkout(t *testing.T) {
	ctx := context.Background()

	// Test nautical workout (Sailing)
	sailing := &database.Workout{
		Type: database.WorkoutTypeSailing,
	}

	// 1852 meters = 1 nm
	assert.Equal(t, "1", HumanDistanceForWorkout(ctx, sailing, 1852))
	assert.Equal(t, "10", HumanDistanceForWorkout(ctx, sailing, 18520))
}

func TestDistanceUnitForWorkout(t *testing.T) {
	ctx := context.Background()

	run := &database.Workout{Type: database.WorkoutTypeRunning}
	sailing := &database.Workout{Type: database.WorkoutTypeSailing}

	assert.Equal(t, "nm", DistanceUnitForWorkout(ctx, sailing))
	assert.NotEqual(t, "nm", DistanceUnitForWorkout(ctx, run))
}

func TestHumanSpeedForWorkout(t *testing.T) {
	ctx := context.Background()

	sailing := &database.Workout{Type: database.WorkoutTypeSailing}

	// 1.94384 knots = 1 m/s. Let's test 10 m/s ~ 19.4 knots
	assert.Equal(t, "19.44", HumanSpeedForWorkout(ctx, sailing, 10))
}

func TestSpeedUnitForWorkout(t *testing.T) {
	ctx := context.Background()

	sailing := &database.Workout{Type: database.WorkoutTypeSailing}
	run := &database.Workout{Type: database.WorkoutTypeRunning}

	assert.Equal(t, "kn", SpeedUnitForWorkout(ctx, sailing))
	assert.NotEqual(t, "kn", SpeedUnitForWorkout(ctx, run))
}

func TestTempoUnitForWorkout(t *testing.T) {
	ctx := context.Background()

	sailing := &database.Workout{Type: database.WorkoutTypeSailing}

	assert.Equal(t, "min/nm", TempoUnitForWorkout(ctx, sailing))
}
