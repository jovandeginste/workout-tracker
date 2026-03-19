package database

import (
	"maps"
	"slices"
)

type (
	WorkoutType string
)

const (
	WorkoutTypeUnknown        WorkoutType = "unknown"
	WorkoutTypeAutoDetect     WorkoutType = "auto"
	WorkoutTypeRunning        WorkoutType = "running"
	WorkoutTypeCycling        WorkoutType = "cycling"
	WorkoutTypeIndoorCycling  WorkoutType = "indoor-cycling"
	WorkoutTypeIndoorRunning  WorkoutType = "indoor-running"
	WorkoutTypeMountainBiking WorkoutType = "mountain-biking"
	WorkoutTypeAlpineClimbing WorkoutType = "alpine-climbing"
	WorkoutTypeMountaineering WorkoutType = "mountaineering"
	WorkoutTypeECycling       WorkoutType = "e-cycling"
	WorkoutTypeHorseRiding    WorkoutType = "horse-riding"
	WorkoutTypeInlineSkating  WorkoutType = "inline-skating"
	WorkoutTypeWalking        WorkoutType = "walking"
	WorkoutTypeSkiing         WorkoutType = "skiing"
	WorkoutTypeSnowboarding   WorkoutType = "snowboarding"
	WorkoutTypeSwimming       WorkoutType = "swimming"
	WorkoutTypeKayaking       WorkoutType = "kayaking"
	WorkoutTypeGolfing        WorkoutType = "golfing"
	WorkoutTypeHiking         WorkoutType = "hiking"
	WorkoutTypePushups        WorkoutType = "push-ups"
	WorkoutTypeSitups         WorkoutType = "sit-ups"
	WorkoutTypeSquats         WorkoutType = "squats"
	WorkoutTypeCore           WorkoutType = "core"
	WorkoutTypeChinUps        WorkoutType = "chin-ups"
	WorkoutTypePullUps        WorkoutType = "pull-ups"
	WorkoutTypePlanking       WorkoutType = "planking"
	WorkoutTypeWeightLifting  WorkoutType = "weight-lifting"
	WorkoutTypeRowing         WorkoutType = "rowing"
	WorkoutTypeTableTennis    WorkoutType = "table-tennis"
	WorkoutTypeTennis         WorkoutType = "tennis"
	WorkoutTypeIceSkating     WorkoutType = "ice-skating"
	WorkoutTypeBadminton      WorkoutType = "badminton"
	WorkoutTypeFootball       WorkoutType = "football"
	WorkoutTypeDancing        WorkoutType = "dancing"
	WorkoutTypeOther          WorkoutType = "other"

	WorkoutTypeClassLocation   = "location"
	WorkoutTypeClassDistance   = "distance"
	WorkoutTypeClassRepetition = "repetition"
	WorkoutTypeClassWeight     = "weight"
	WorkoutTypeClassDuration   = "duration"
)

type WorkoutTypeConfiguration struct {
	Location            bool
	Distance            bool
	Repetition          bool
	Weight              bool
	AreClimbsRelevant   bool
	AreDescentsRelevant bool
	MaxDeltaMeter       float64
}

const DefaultMaxDeltaMeter = 20.0

var workoutTypeConfigs = map[WorkoutType]WorkoutTypeConfiguration{
	WorkoutTypeAlpineClimbing: {Location: true, Distance: true, Repetition: false, Weight: false, AreDescentsRelevant: false, AreClimbsRelevant: true, MaxDeltaMeter: DefaultMaxDeltaMeter},
	WorkoutTypeBadminton:      {Location: false, Distance: false, Repetition: false, Weight: false, AreDescentsRelevant: false, AreClimbsRelevant: false, MaxDeltaMeter: DefaultMaxDeltaMeter},
	WorkoutTypeChinUps:        {Location: false, Distance: false, Repetition: true, Weight: true, AreDescentsRelevant: false, AreClimbsRelevant: false, MaxDeltaMeter: DefaultMaxDeltaMeter},
	WorkoutTypeCore:           {Location: false, Distance: false, Repetition: true, Weight: false, AreDescentsRelevant: false, AreClimbsRelevant: false, MaxDeltaMeter: DefaultMaxDeltaMeter},
	WorkoutTypeCycling:        {Location: true, Distance: true, Repetition: false, Weight: false, AreDescentsRelevant: true, AreClimbsRelevant: true, MaxDeltaMeter: DefaultMaxDeltaMeter},
	WorkoutTypeDancing:        {Location: false, Distance: false, Repetition: false, Weight: false, AreDescentsRelevant: false, AreClimbsRelevant: false, MaxDeltaMeter: DefaultMaxDeltaMeter},
	WorkoutTypeECycling:       {Location: true, Distance: true, Repetition: false, Weight: false, AreDescentsRelevant: true, AreClimbsRelevant: true, MaxDeltaMeter: DefaultMaxDeltaMeter},
	WorkoutTypeFootball:       {Location: true, Distance: true, Repetition: false, Weight: false, AreDescentsRelevant: false, AreClimbsRelevant: false, MaxDeltaMeter: DefaultMaxDeltaMeter},
	WorkoutTypeGolfing:        {Location: true, Distance: true, Repetition: false, Weight: false, AreDescentsRelevant: true, AreClimbsRelevant: true, MaxDeltaMeter: DefaultMaxDeltaMeter},
	WorkoutTypeHiking:         {Location: true, Distance: true, Repetition: false, Weight: false, AreDescentsRelevant: true, AreClimbsRelevant: true, MaxDeltaMeter: DefaultMaxDeltaMeter},
	WorkoutTypeHorseRiding:    {Location: true, Distance: true, Repetition: false, Weight: false, AreDescentsRelevant: true, AreClimbsRelevant: true, MaxDeltaMeter: DefaultMaxDeltaMeter},
	WorkoutTypeIceSkating:     {Location: true, Distance: true, Repetition: false, Weight: false, AreDescentsRelevant: false, AreClimbsRelevant: false, MaxDeltaMeter: DefaultMaxDeltaMeter},
	WorkoutTypeIndoorCycling:  {Location: false, Distance: true, Repetition: false, Weight: false, AreDescentsRelevant: false, AreClimbsRelevant: false, MaxDeltaMeter: DefaultMaxDeltaMeter},
	WorkoutTypeIndoorRunning:  {Location: false, Distance: true, Repetition: false, Weight: false, AreDescentsRelevant: false, AreClimbsRelevant: false, MaxDeltaMeter: DefaultMaxDeltaMeter},
	WorkoutTypeInlineSkating:  {Location: true, Distance: true, Repetition: false, Weight: false, AreDescentsRelevant: true, AreClimbsRelevant: true, MaxDeltaMeter: DefaultMaxDeltaMeter},
	WorkoutTypeKayaking:       {Location: true, Distance: true, Repetition: false, Weight: false, AreDescentsRelevant: false, AreClimbsRelevant: false, MaxDeltaMeter: DefaultMaxDeltaMeter},
	WorkoutTypeMountainBiking: {Location: true, Distance: true, Repetition: false, Weight: false, AreDescentsRelevant: true, AreClimbsRelevant: true, MaxDeltaMeter: DefaultMaxDeltaMeter},
	WorkoutTypeMountaineering: {Location: true, Distance: true, Repetition: false, Weight: true, AreDescentsRelevant: true, AreClimbsRelevant: true, MaxDeltaMeter: DefaultMaxDeltaMeter},
	WorkoutTypeOther:          {Location: true, Distance: true, Repetition: true, Weight: true, AreDescentsRelevant: true, AreClimbsRelevant: true, MaxDeltaMeter: DefaultMaxDeltaMeter},
	WorkoutTypePlanking:       {Location: false, Distance: false, Repetition: false, Weight: false, AreDescentsRelevant: false, AreClimbsRelevant: false, MaxDeltaMeter: DefaultMaxDeltaMeter},
	WorkoutTypePullUps:        {Location: false, Distance: false, Repetition: true, Weight: true, AreDescentsRelevant: false, AreClimbsRelevant: false, MaxDeltaMeter: DefaultMaxDeltaMeter},
	WorkoutTypePushups:        {Location: false, Distance: false, Repetition: true, Weight: false, AreDescentsRelevant: false, AreClimbsRelevant: true, MaxDeltaMeter: DefaultMaxDeltaMeter},
	WorkoutTypeRowing:         {Location: true, Distance: true, Repetition: false, Weight: false, AreDescentsRelevant: false, AreClimbsRelevant: false, MaxDeltaMeter: DefaultMaxDeltaMeter},
	WorkoutTypeRunning:        {Location: true, Distance: true, Repetition: false, Weight: false, AreDescentsRelevant: true, AreClimbsRelevant: true, MaxDeltaMeter: DefaultMaxDeltaMeter},
	WorkoutTypeSitups:         {Location: false, Distance: false, Repetition: true, Weight: false, AreDescentsRelevant: false, AreClimbsRelevant: false, MaxDeltaMeter: DefaultMaxDeltaMeter},
	WorkoutTypeSkiing:         {Location: true, Distance: true, Repetition: false, Weight: false, AreDescentsRelevant: true, AreClimbsRelevant: false, MaxDeltaMeter: 120},
	WorkoutTypeSnowboarding:   {Location: true, Distance: true, Repetition: false, Weight: false, AreDescentsRelevant: true, AreClimbsRelevant: false, MaxDeltaMeter: 120},
	WorkoutTypeSquats:         {Location: false, Distance: false, Repetition: true, Weight: true, AreDescentsRelevant: false, AreClimbsRelevant: false, MaxDeltaMeter: DefaultMaxDeltaMeter},
	WorkoutTypeSwimming:       {Location: true, Distance: true, Repetition: false, Weight: false, AreDescentsRelevant: false, AreClimbsRelevant: false, MaxDeltaMeter: DefaultMaxDeltaMeter},
	WorkoutTypeTableTennis:    {Location: false, Distance: false, Repetition: false, Weight: false, AreDescentsRelevant: false, AreClimbsRelevant: false, MaxDeltaMeter: DefaultMaxDeltaMeter},
	WorkoutTypeTennis:         {Location: true, Distance: true, Repetition: false, Weight: false, AreDescentsRelevant: false, AreClimbsRelevant: false, MaxDeltaMeter: DefaultMaxDeltaMeter},
	WorkoutTypeWalking:        {Location: true, Distance: true, Repetition: false, Weight: false, AreDescentsRelevant: true, AreClimbsRelevant: true, MaxDeltaMeter: DefaultMaxDeltaMeter},
	WorkoutTypeWeightLifting:  {Location: false, Distance: false, Repetition: true, Weight: true, AreDescentsRelevant: false, AreClimbsRelevant: true, MaxDeltaMeter: DefaultMaxDeltaMeter},
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
	if wt == "" {
		return string(WorkoutTypeUnknown)
	}

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

// MaxDeltaMeter is the maximum distance in meters that a point can be away from
// the route segment for route segment matching
func (wt WorkoutType) MaxDeltaMeter() float64 {
	return workoutTypeConfigs[wt].MaxDeltaMeter
}

func (wt WorkoutType) AreClimbsRelevant() bool {
	return workoutTypeConfigs[wt].AreClimbsRelevant
}

func (wt WorkoutType) AreDescentsRelevant() bool {
	return workoutTypeConfigs[wt].AreDescentsRelevant
}
