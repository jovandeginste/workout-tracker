package database

import "slices"

type WorkoutType string

const (
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
