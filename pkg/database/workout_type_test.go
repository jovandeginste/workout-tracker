package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWorkoutType_String(t *testing.T) {
	assert.Equal(t, "running", WorkoutTypeRunning.String())
	assert.Equal(t, "cycling", WorkoutTypeCycling.String())
	assert.Equal(t, "walking", WorkoutTypeWalking.String())
}

func TestWorkoutType_IsDistance(t *testing.T) {
	assert.True(t, WorkoutTypeRunning.IsDistance())
	assert.True(t, WorkoutTypeCycling.IsDistance())
	assert.True(t, WorkoutTypeWalking.IsDistance())

	assert.False(t, WorkoutTypeAutoDetect.IsDistance())
}

func TestWorkoutType_Collections(t *testing.T) {
	for _, wt := range []WorkoutType{WorkoutTypeRunning, WorkoutTypeCycling, WorkoutTypeWalking} {
		assert.Contains(t, WorkoutTypes(), wt)
		assert.Contains(t, DistanceWorkoutTypes(), wt)
	}
}
