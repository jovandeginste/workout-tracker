package database

import (
	"maps"
	"slices"
)

type (
	WorkoutType string

	// WorkoutTypeList is a list of workout types, stored as JSON in the database
	WorkoutTypeList []WorkoutType
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
	WorkoutTypeSailing        WorkoutType = "sailing"
	WorkoutTypeIceSkating     WorkoutType = "ice-skating"
	WorkoutTypeBadminton      WorkoutType = "badminton"
	WorkoutTypeFootball       WorkoutType = "football"
	WorkoutTypeDancing        WorkoutType = "dancing"
	WorkoutTypeOther          WorkoutType = "other"
)

type WorkoutTypeConfiguration struct {
	Location            bool
	Distance            bool
	Repetition          bool
	Weight              bool
	AreClimbsRelevant   bool
	AreDescentsRelevant bool
	MaxDeltaMeter       float64
	Nautical            bool
}

const DefaultMaxDeltaMeter = 20.0

func workoutTypeConfig(location, distance, repetition, weight bool, opts ...func(*WorkoutTypeConfiguration)) WorkoutTypeConfiguration {
	c := WorkoutTypeConfiguration{
		Location:      location,
		Distance:      distance,
		Repetition:    repetition,
		Weight:        weight,
		MaxDeltaMeter: DefaultMaxDeltaMeter,
	}

	for _, opt := range opts {
		opt(&c)
	}

	return c
}

func withClimbs(c *WorkoutTypeConfiguration) {
	c.AreClimbsRelevant = true
}

func withDescents(c *WorkoutTypeConfiguration) {
	c.AreDescentsRelevant = true
}

func withHighMaxDelta(c *WorkoutTypeConfiguration) {
	c.MaxDeltaMeter = 120
}

func withNautical(c *WorkoutTypeConfiguration) {
	c.Nautical = true
}

var workoutTypeConfigs = map[WorkoutType]WorkoutTypeConfiguration{
	WorkoutTypeAlpineClimbing: workoutTypeConfig(true, true, false, false, withClimbs),
	WorkoutTypeBadminton:      {},
	WorkoutTypeChinUps:        workoutTypeConfig(false, false, true, true),
	WorkoutTypeCore:           workoutTypeConfig(false, false, true, false),
	WorkoutTypeCycling:        workoutTypeConfig(true, true, false, false, withClimbs, withDescents, withHighMaxDelta),
	WorkoutTypeDancing:        {},
	WorkoutTypeECycling:       workoutTypeConfig(true, true, false, false, withClimbs, withDescents, withHighMaxDelta),
	WorkoutTypeFootball:       workoutTypeConfig(true, true, false, false),
	WorkoutTypeGolfing:        workoutTypeConfig(true, true, false, false, withClimbs, withDescents),
	WorkoutTypeHiking:         workoutTypeConfig(true, true, false, false, withClimbs, withDescents),
	WorkoutTypeHorseRiding:    workoutTypeConfig(true, true, false, false, withClimbs, withDescents),
	WorkoutTypeIceSkating:     workoutTypeConfig(true, true, false, false),
	WorkoutTypeIndoorCycling:  workoutTypeConfig(false, true, false, false, withHighMaxDelta),
	WorkoutTypeIndoorRunning:  workoutTypeConfig(false, true, false, false),
	WorkoutTypeInlineSkating:  workoutTypeConfig(true, true, false, false, withClimbs, withDescents),
	WorkoutTypeKayaking:       workoutTypeConfig(true, true, false, false),
	WorkoutTypeMountainBiking: workoutTypeConfig(true, true, false, false, withClimbs, withDescents),
	WorkoutTypeMountaineering: workoutTypeConfig(true, true, false, true, withClimbs, withDescents),
	WorkoutTypeOther:          workoutTypeConfig(true, true, true, true, withClimbs, withDescents),
	WorkoutTypePlanking:       {},
	WorkoutTypePullUps:        workoutTypeConfig(false, false, true, true),
	WorkoutTypePushups:        workoutTypeConfig(false, false, true, false, withClimbs),
	WorkoutTypeRowing:         workoutTypeConfig(true, true, false, false),
	WorkoutTypeRunning:        workoutTypeConfig(true, true, false, false, withClimbs, withDescents),
	WorkoutTypeSailing:        workoutTypeConfig(true, true, false, false, withNautical),
	WorkoutTypeSitups:         workoutTypeConfig(false, false, true, false),
	WorkoutTypeSkiing:         workoutTypeConfig(true, true, false, false, withDescents, withHighMaxDelta),
	WorkoutTypeSnowboarding:   workoutTypeConfig(true, true, false, false, withDescents, withHighMaxDelta),
	WorkoutTypeSquats:         workoutTypeConfig(false, false, true, true),
	WorkoutTypeSwimming:       workoutTypeConfig(true, true, false, false),
	WorkoutTypeTableTennis:    {},
	WorkoutTypeTennis:         workoutTypeConfig(true, true, false, false),
	WorkoutTypeWalking:        workoutTypeConfig(true, true, false, false, withClimbs, withDescents),
	WorkoutTypeWeightLifting:  workoutTypeConfig(false, false, true, true, withClimbs),
}

var workoutTypes []WorkoutType

func WorkoutTypes() []WorkoutType {
	if len(workoutTypes) > 0 {
		return workoutTypes
	}

	workoutTypes = slices.Collect(maps.Keys(workoutTypeConfigs))

	slices.Sort(workoutTypes)

	return workoutTypes
}

func workoutTypesByClass(fn func(c WorkoutTypeConfiguration) bool) []WorkoutType {
	keys := []WorkoutType{}

	for k, c := range workoutTypeConfigs {
		if !fn(c) {
			continue
		}

		keys = append(keys, k)
	}

	slices.Sort(keys)

	return keys
}

func DistanceWorkoutTypes() []WorkoutType {
	return workoutTypesByClass(func(c WorkoutTypeConfiguration) bool {
		return c.Distance
	})
}

func WeightWorkoutTypes() []WorkoutType {
	return workoutTypesByClass(func(c WorkoutTypeConfiguration) bool {
		return c.Weight
	})
}

func RepetitionWorkoutTypes() []WorkoutType {
	return workoutTypesByClass(func(c WorkoutTypeConfiguration) bool {
		return c.Repetition
	})
}

func LocationWorkoutTypes() []WorkoutType {
	return workoutTypesByClass(func(c WorkoutTypeConfiguration) bool {
		return c.Location
	})
}

func DurationWorkoutTypes() []WorkoutType {
	return workoutTypesByClass(func(c WorkoutTypeConfiguration) bool {
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

// Has returns true if the list contains the given workout type
func (l WorkoutTypeList) Has(t WorkoutType) bool {
	return slices.Contains(l, t)
}

// Sanitize removes unknown and duplicate workout types from the list
func (l *WorkoutTypeList) Sanitize() {
	if l == nil {
		return
	}

	seen := map[WorkoutType]bool{}
	result := make(WorkoutTypeList, 0, len(*l))

	for _, t := range *l {
		if t != WorkoutTypeAutoDetect && !t.IsDuration() {
			continue
		}

		if seen[t] {
			continue
		}

		seen[t] = true
		result = append(result, t)
	}

	*l = result
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

func (wt WorkoutType) IsNautical() bool {
	return workoutTypeConfigs[wt].Nautical
}
