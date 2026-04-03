package helpers

import (
	"context"
	"time"

	"github.com/jovandeginste/workout-tracker/v2/pkg/database"
	"github.com/jovandeginste/workout-tracker/v2/pkg/templatehelpers"
)

func HumanDuration(d time.Duration) string {
	return templatehelpers.HumanDuration(d)
}

func HumanElevation(ctx context.Context, d float64) string {
	return templatehelpers.HumanElevationFor(CurrentUser(ctx).PreferredUnits().Elevation())(d)
}

func HumanDistance(ctx context.Context, d float64) string {
	return templatehelpers.HumanDistanceFor(CurrentUser(ctx).PreferredUnits().Distance(nil))(d)
}

func HumanDistanceForWorkout(ctx context.Context, w *database.Workout, d float64) string {
	if w != nil {
		return templatehelpers.HumanDistanceFor(CurrentUser(ctx).PreferredUnits().Distance(&w.Type))(d)
	}
	return HumanDistance(ctx, d)
}

func DistanceUnitForWorkout(ctx context.Context, w *database.Workout) string {
	if w != nil {
		return CurrentUser(ctx).PreferredUnits().Distance(&w.Type)
	}
	return CurrentUser(ctx).PreferredUnits().Distance(nil)
}

func HumanTempo(ctx context.Context, d float64) string {
	return templatehelpers.HumanTempoFor(CurrentUser(ctx).PreferredUnits().Distance(nil))(d)
}

func HumanTempoForWorkout(ctx context.Context, w *database.Workout, d float64) string {
	if w != nil {
		return templatehelpers.HumanTempoFor(CurrentUser(ctx).PreferredUnits().Tempo(&w.Type))(d)
	}
	return HumanTempo(ctx, d)
}

func TempoUnitForWorkout(ctx context.Context, w *database.Workout) string {
	if w != nil {
		return CurrentUser(ctx).PreferredUnits().Tempo(&w.Type)
	}
	return CurrentUser(ctx).PreferredUnits().Tempo(nil)
}

func HumanSpeed(ctx context.Context, d float64) string {
	return templatehelpers.HumanSpeedFor(CurrentUser(ctx).PreferredUnits().Speed(nil))(d)
}

func HumanSpeedForWorkout(ctx context.Context, w *database.Workout, d float64) string {
	if w != nil {
		return templatehelpers.HumanSpeedFor(CurrentUser(ctx).PreferredUnits().Speed(&w.Type))(d)
	}
	return HumanSpeed(ctx, d)
}

func SpeedUnitForWorkout(ctx context.Context, w *database.Workout) string {
	if w != nil {
		return CurrentUser(ctx).PreferredUnits().Speed(&w.Type)
	}
	return CurrentUser(ctx).PreferredUnits().Speed(nil)
}

func HumanCadence(d float64) string {
	return templatehelpers.HumanCadence(d)
}

func HumanPower(d float64) string {
	return templatehelpers.HumanPower(d)
}

func HumanCalories(d float64) string {
	return templatehelpers.HumanCalories(d)
}

func HumanWeight(ctx context.Context, d float64) string {
	return templatehelpers.HumanWeightFor(CurrentUser(ctx).PreferredUnits().Weight())(d)
}

func HumanHeight(ctx context.Context, d float64) string {
	return templatehelpers.HumanHeightFor(CurrentUser(ctx).PreferredUnits().Height())(d)
}

func HumanHeightSingle(ctx context.Context, d float64) string {
	return templatehelpers.HumanHeightSingleFor(CurrentUser(ctx).PreferredUnits().Height())(d)
}
