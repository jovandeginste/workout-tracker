package database

import "slices"

type WorkoutType string

const (
	WorkoutTypeAutoDetect WorkoutType = "auto"
	WorkoutTypeRunning    WorkoutType = "running"
	WorkoutTypeCycling    WorkoutType = "cycling"
	WorkoutTypeWalking    WorkoutType = "walking"
)

func WorkoutTypes() []WorkoutType {
	return []WorkoutType{WorkoutTypeRunning, WorkoutTypeCycling, WorkoutTypeWalking}
}

func DistanceWorkoutTypes() []WorkoutType {
	return []WorkoutType{WorkoutTypeRunning, WorkoutTypeCycling, WorkoutTypeWalking}
}

func (wt WorkoutType) String() string {
	return string(wt)
}

func (wt WorkoutType) IsDistance() bool {
	return slices.Contains(DistanceWorkoutTypes(), wt)
}
