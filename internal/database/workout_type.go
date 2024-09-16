package database

import (
	"html"
	"html/template"
	"slices"

	"github.com/jovandeginste/workout-tracker/internal/pkg/templatehelpers"
)

type (
	WorkoutType string
)

const (
	// We need to add each of these types to the "messages.html" partial view.
	// Then it gets picked up by the i18n system, added to the list of translatable
	// strings, etc.
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
	WorkoutTypeWeightLifting WorkoutType = "weight lifting"

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

	for k := range workoutTypeConfigs {
		workoutTypes = append(workoutTypes, k)
	}

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

func (wt *WorkoutType) PreferredSpeedMetric(v float64, preferredUnits *UserPreferredUnits) template.HTML {
	speedUnit := preferredUnits.Speed()
	speedFormatter := templatehelpers.HumanSpeedFor(speedUnit)
	primaryText := html.EscapeString(speedFormatter(v) + " " + speedUnit)
	tempoUnit := preferredUnits.Tempo()
	tempoFormatter := templatehelpers.HumanTempoFor(tempoUnit)
	secondaryText := html.EscapeString(tempoFormatter(v) + " " + tempoUnit)

	if *wt == WorkoutTypeRunning {
		// Swap tempo and speed, so that tempo is primarily shown for running
		primaryText, secondaryText = secondaryText, primaryText
	}

	//nolint:gosec // We escaped all unsafe HTML
	return template.HTML("<abbr title='" + secondaryText + "'>" + primaryText + "</abbr>")
}
