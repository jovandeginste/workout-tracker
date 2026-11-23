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
	assert.Equal(t, 120.0, WorkoutTypeCycling.MaxDeltaMeter())
}

func TestWorkoutType_IsNautical(t *testing.T) {
	assert.True(t, WorkoutTypeSailing.IsNautical())
	assert.False(t, WorkoutTypeRunning.IsNautical())
	assert.False(t, WorkoutTypeCycling.IsNautical())
}

func TestWorkoutTypeList_Sanitize(t *testing.T) {
	types := WorkoutTypeList{WorkoutTypeRunning, WorkoutTypeCycling, WorkoutTypeRunning, "invalid", WorkoutTypeAutoDetect, WorkoutTypeAutoDetect}

	types.Sanitize()

	assert.Equal(t, WorkoutTypeList{WorkoutTypeRunning, WorkoutTypeCycling, WorkoutTypeAutoDetect}, types)
}

func TestWorkoutTypeList_SanitizeEmpty(t *testing.T) {
	var types WorkoutTypeList

	types.Sanitize()

	assert.Empty(t, types)
}

func TestWorkoutTypeList_Has(t *testing.T) {
	types := WorkoutTypeList{WorkoutTypeRunning, WorkoutTypeCycling}

	assert.True(t, types.Has(WorkoutTypeRunning))
	assert.False(t, types.Has(WorkoutTypeSwimming))
}
