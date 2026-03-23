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
	assert.True(t, WorkoutTypeSailing.IsDistance())

	assert.False(t, WorkoutTypeAutoDetect.IsDistance())
}

func TestWorkoutType_Collections(t *testing.T) {
	for _, wt := range []WorkoutType{WorkoutTypeRunning, WorkoutTypeCycling, WorkoutTypeWalking} {
		assert.Contains(t, WorkoutTypes(), wt)
		assert.Contains(t, DistanceWorkoutTypes(), wt)
	}
}

func TestWorkoutType_SnowboardingSkiing(t *testing.T) {
	assert.False(t, WorkoutTypeSnowboarding.AreClimbsRelevant())
	assert.False(t, WorkoutTypeSkiing.AreClimbsRelevant())
	assert.True(t, WorkoutTypeSnowboarding.AreDescentsRelevant())
	assert.True(t, WorkoutTypeSkiing.AreDescentsRelevant())
	assert.True(t, WorkoutTypeCycling.AreClimbsRelevant() && WorkoutTypeCycling.AreDescentsRelevant())
	assert.True(t, WorkoutTypeSnowboarding.MaxDeltaMeter() > 100)
	assert.True(t, WorkoutTypeSkiing.MaxDeltaMeter() > 100)
	assert.True(t, WorkoutTypeCycling.MaxDeltaMeter() < 100)
}

func TestWorkoutType_IsNautical(t *testing.T) {
	assert.True(t, WorkoutTypeSailing.IsNautical())
	assert.False(t, WorkoutTypeRunning.IsNautical())
	assert.False(t, WorkoutTypeCycling.IsNautical())
}
