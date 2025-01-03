package helpers

import (
	"context"

	"github.com/jovandeginste/workout-tracker/pkg/database"
	"github.com/jovandeginste/workout-tracker/pkg/templatehelpers"
)

const timeFormat = "2006-01-02 15:04"

func iconFor(name string) string {
	return string(templatehelpers.IconFor(name))
}

func RouteFor(ctx context.Context, name string, params ...any) string {
	e := appEcho(ctx)
	if e == nil {
		return "/invalid/route/#" + name
	}

	if rev := e.Reverse(name, params...); rev != "" {
		return rev
	}

	return "/invalid/route/#" + name
}

func UserPreferredUnits(ctx context.Context) *database.UserPreferredUnits {
	return CurrentUser(ctx).PreferredUnits()
}

func WorkoutTypes() []database.WorkoutType {
	return database.WorkoutTypes()
}

func FilterOptions() []string {
	return StatisticSinceOptions()
}

func OrderDirOptions() map[string]string {
	return map[string]string{
		"asc":  "ascending",
		"desc": "descending",
	}
}

func OrderByOptions() map[string]string {
	return map[string]string{
		"date": "Date",

		"total_distance":    "Total distance",
		"total_duration":    "Total duration",
		"total_weight":      "Total weight",
		"total_repetitions": "Total repetitions",
		"total_up":          "Total up",
		"total_down":        "Total down",

		"average_speed_no_pause": "Average speed",
		"max_speed":              "Max speed",
	}
}

func PreferredUnitsToJSON(units *database.UserPreferredUnits) map[string]string {
	return map[string]string{
		"Distance":  units.Distance(),
		"Speed":     units.Speed(),
		"Elevation": units.Elevation(),
		"HeartRate": units.HeartRate(),
		"Cadence":   units.Cadence(),
	}
}

func BoolToHTML(b bool) string {
	if b {
		return `<i class="text-green-500 icon-[fa-solid--check]"></i>`
	}

	return `<i class="text-rose-500 icon-[fa-solid--times]"></i>`
}
