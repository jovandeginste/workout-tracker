package database

import "slices"

type WorkoutType string

const (
	// We need to add each of these types to the "messages.html" partial view.
	// Then it gets picked up by the i18n system, added to the list of translatable
	// strings, etc.
	WorkoutTypeAutoDetect    WorkoutType = "auto"
	WorkoutTypeRunning       WorkoutType = "running"
	WorkoutTypeCycling       WorkoutType = "cycling"
	WorkoutTypeWalking       WorkoutType = "walking"
	WorkoutTypeSkiing        WorkoutType = "skiing"
	WorkoutTypeSnowboarding  WorkoutType = "snowboarding"
	WorkoutTypeSwimming      WorkoutType = "swimming"
	WorkoutTypeKayaking      WorkoutType = "kayaking"
	WorkoutTypeGolfing       WorkoutType = "golfing"
	WorkoutTypeHiking        WorkoutType = "hiking"
	WorkoutTypePushups       WorkoutType = "push-ups"
	WorkoutTypeWeightLifting WorkoutType = "weight lifting"
)

type WorkoutTypeConfiguration struct {
	Location   bool
	Distance   bool
	Repetition bool
	Weight     bool
}

var workoutTypeConfigs = map[WorkoutType]WorkoutTypeConfiguration{
	WorkoutTypeRunning:      {Location: true, Distance: true, Repetition: false, Weight: false},
	WorkoutTypeCycling:      {Location: true, Distance: true, Repetition: false, Weight: false},
	WorkoutTypeWalking:      {Location: true, Distance: true, Repetition: false, Weight: false},
	WorkoutTypeSkiing:       {Location: true, Distance: true, Repetition: false, Weight: false},
	WorkoutTypeSnowboarding: {Location: true, Distance: true, Repetition: false, Weight: false},
	WorkoutTypeSwimming:     {Location: true, Distance: true, Repetition: false, Weight: false},
	WorkoutTypeKayaking:     {Location: true, Distance: true, Repetition: false, Weight: false},
	WorkoutTypeGolfing:      {Location: true, Distance: true, Repetition: false, Weight: false},
	WorkoutTypeHiking:       {Location: true, Distance: true, Repetition: false, Weight: false},

	WorkoutTypePushups:       {Location: false, Distance: false, Repetition: true, Weight: false},
	WorkoutTypeWeightLifting: {Location: false, Distance: false, Repetition: true, Weight: true},
}

func WorkoutTypes() []WorkoutType {
	keys := []WorkoutType{}

	for k := range workoutTypeConfigs {
		keys = append(keys, k)
	}

	slices.Sort(keys)

	return keys
}

func DistanceWorkoutTypes() []WorkoutType {
	keys := []WorkoutType{}

	for k, c := range workoutTypeConfigs {
		if !c.Distance {
			continue
		}

		keys = append(keys, k)
	}

	slices.Sort(keys)

	return keys
}

func LocationWorkoutTypes() []WorkoutType {
	keys := []WorkoutType{}

	for k, c := range workoutTypeConfigs {
		if !c.Location {
			continue
		}

		keys = append(keys, k)
	}

	slices.Sort(keys)

	return keys
}

func (wt WorkoutType) String() string {
	return string(wt)
}

func (wt WorkoutType) IsDistance() bool {
	return workoutTypeConfigs[wt].Distance
}

func (wt WorkoutType) IsRepetition() bool {
	return workoutTypeConfigs[wt].Repetition
}

func (wt WorkoutType) IsDuration() bool {
	_, ok := workoutTypeConfigs[wt]
	return ok
}

func (wt WorkoutType) IsWeight() bool {
	return workoutTypeConfigs[wt].Weight
}

func (wt WorkoutType) IsLocation() bool {
	return workoutTypeConfigs[wt].Location
}

func AsWorkoutType(s string) WorkoutType {
	return WorkoutType(s)
}
