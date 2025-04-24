package database

import (
	"maps"
	"slices"
)

type (
	WorkoutType string
)

const (
	WorkoutTypeAutoDetect    WorkoutType = "auto"
	WorkoutTypeRunning       WorkoutType = "running"
	WorkoutTypeCycling       WorkoutType = "cycling"
	WorkoutTypeECycling      WorkoutType = "e-cycling"
	WorkoutTypeWalking       WorkoutType = "walking"
	WorkoutTypeSkiing        WorkoutType = "skiing"
	WorkoutTypeSnowboarding  WorkoutType = "snowboarding"
	WorkoutTypeSwimming      WorkoutType = "swimming"
	WorkoutTypeKayaking      WorkoutType = "kayaking"
	WorkoutTypeGolfing       WorkoutType = "golfing"
	WorkoutTypeHiking        WorkoutType = "hiking"
	WorkoutTypePushups       WorkoutType = "push-ups"
	WorkoutTypeWeightLifting WorkoutType = "weight-lifting"
	WorkoutTypeRowing        WorkoutType = "rowing"

	WorkoutTypeClassLocation   = "location"
	WorkoutTypeClassDistance   = "distance"
	WorkoutTypeClassRepetition = "repetition"
	WorkoutTypeClassWeight     = "weight"
	WorkoutTypeClassDuration   = "duration"
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
	WorkoutTypeECycling:     {Location: true, Distance: true, Repetition: false, Weight: false},
	WorkoutTypeWalking:      {Location: true, Distance: true, Repetition: false, Weight: false},
	WorkoutTypeSkiing:       {Location: true, Distance: true, Repetition: false, Weight: false},
	WorkoutTypeSnowboarding: {Location: true, Distance: true, Repetition: false, Weight: false},
	WorkoutTypeSwimming:     {Location: true, Distance: true, Repetition: false, Weight: false},
	WorkoutTypeKayaking:     {Location: true, Distance: true, Repetition: false, Weight: false},
	WorkoutTypeRowing:       {Location: true, Distance: true, Repetition: false, Weight: false},
	WorkoutTypeGolfing:      {Location: true, Distance: true, Repetition: false, Weight: false},
	WorkoutTypeHiking:       {Location: true, Distance: true, Repetition: false, Weight: false},

	WorkoutTypePushups:       {Location: false, Distance: false, Repetition: true, Weight: false},
	WorkoutTypeWeightLifting: {Location: false, Distance: false, Repetition: true, Weight: true},
}

var (
	workoutTypes        []WorkoutType
	workoutTypesByClass map[string][]WorkoutType
)

func WorkoutTypes() []WorkoutType {
	if len(workoutTypes) > 0 {
		return workoutTypes
	}

	workoutTypes = slices.Collect(maps.Keys(workoutTypeConfigs))

	slices.Sort(workoutTypes)

	return workoutTypes
}

func getOrSetByClass(class string, fn func(c WorkoutTypeConfiguration) bool) []WorkoutType {
	if workoutTypesByClass == nil {
		workoutTypesByClass = make(map[string][]WorkoutType)
	}

	if wt, ok := workoutTypesByClass[class]; ok {
		return wt
	}

	keys := []WorkoutType{}

	for k, c := range workoutTypeConfigs {
		if !fn(c) {
			continue
		}

		keys = append(keys, k)
	}

	slices.Sort(keys)
	workoutTypesByClass[WorkoutTypeClassDistance] = keys

	return keys
}

func DistanceWorkoutTypes() []WorkoutType {
	return getOrSetByClass(WorkoutTypeClassDistance, func(c WorkoutTypeConfiguration) bool {
		return c.Distance
	})
}

func WeightWorkoutTypes() []WorkoutType {
	return getOrSetByClass(WorkoutTypeClassWeight, func(c WorkoutTypeConfiguration) bool {
		return c.Weight
	})
}

func RepetitionWorkoutTypes() []WorkoutType {
	return getOrSetByClass(WorkoutTypeClassRepetition, func(c WorkoutTypeConfiguration) bool {
		return c.Repetition
	})
}

func LocationWorkoutTypes() []WorkoutType {
	return getOrSetByClass(WorkoutTypeClassLocation, func(c WorkoutTypeConfiguration) bool {
		return c.Location
	})
}

func DurationWorkoutTypes() []WorkoutType {
	return getOrSetByClass(WorkoutTypeClassDuration, func(c WorkoutTypeConfiguration) bool {
		return true // All workout types store duration
	})
}

func (wt WorkoutType) StringT() string {
	return "sports." + wt.String()
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
