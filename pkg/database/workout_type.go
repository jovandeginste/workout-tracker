package database

import "slices"

type WorkoutType string

const (
	// We need to add each of these types to the "messages.html" partial view.
	// Then it gets picked up by the i18n system, added to the list of translatable
	// strings, etc.
	WorkoutTypeAutoDetect   WorkoutType = "auto"
	WorkoutTypeRunning      WorkoutType = "running"
	WorkoutTypeCycling      WorkoutType = "cycling"
	WorkoutTypeWalking      WorkoutType = "walking"
	WorkoutTypeSkiing       WorkoutType = "skiing"
	WorkoutTypeSnowboarding WorkoutType = "snowboarding"
	WorkoutTypeSwimming     WorkoutType = "swimming"
	WorkoutTypeKayaking     WorkoutType = "kayaking"
)

func WorkoutTypes() []WorkoutType {
	return []WorkoutType{WorkoutTypeRunning, WorkoutTypeCycling, WorkoutTypeWalking, WorkoutTypeSkiing, WorkoutTypeSnowboarding, WorkoutTypeSwimming, WorkoutTypeKayaking}
}

func DistanceWorkoutTypes() []WorkoutType {
	return []WorkoutType{WorkoutTypeRunning, WorkoutTypeCycling, WorkoutTypeWalking, WorkoutTypeSkiing, WorkoutTypeSnowboarding, WorkoutTypeSwimming, WorkoutTypeKayaking}
}

func (wt WorkoutType) String() string {
	return string(wt)
}

func (wt WorkoutType) IsDistance() bool {
	return slices.Contains(DistanceWorkoutTypes(), wt)
}
