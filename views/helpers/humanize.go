package helpers

import (
	"context"
	"time"

	"github.com/jovandeginste/workout-tracker/v2/pkg/templatehelpers"
)

func HumanDuration(d time.Duration) string {
	return templatehelpers.HumanDuration(d)
}

func HumanElevation(ctx context.Context, d float64) string {
	return templatehelpers.HumanElevationFor(CurrentUser(ctx).PreferredUnits().Elevation())(d)
}

func HumanDistance(ctx context.Context, d float64) string {
	return templatehelpers.HumanDistanceFor(CurrentUser(ctx).PreferredUnits().Distance())(d)
}

func HumanTempo(ctx context.Context, d float64) string {
	return templatehelpers.HumanTempoFor(CurrentUser(ctx).PreferredUnits().Distance())(d)
}

func HumanSpeed(ctx context.Context, d float64) string {
	return templatehelpers.HumanSpeedFor(CurrentUser(ctx).PreferredUnits().Speed())(d)
}

func HumanCalories(d float64) string {
	return templatehelpers.HumanCaloriesKcal(d)
}
