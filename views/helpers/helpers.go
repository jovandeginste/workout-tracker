package helpers

import (
	"context"
	"fmt"

	"github.com/jovandeginste/workout-tracker/pkg/database"
)

const timeFormat = "2006-01-02 15:04"

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

type TranslatedKey struct {
	Key         string
	Translation string
}

func OrderDirOptions() []TranslatedKey {
	return []TranslatedKey{
		{"asc", "ascending"},
		{"desc", "descending"},
	}
}

func OrderByOptions() []TranslatedKey {
	return []TranslatedKey{
		{"date", "Date"},

		{"total_distance", "Total distance"},
		{"total_duration", "Total duration"},
		{"total_weight", "Total weight"},
		{"total_repetitions", "Total repetitions"},
		{"total_up", "Total up"},
		{"total_down", "Total down"},

		{"average_speed_no_pause", "Average speed"},
		{"max_speed", "Max speed"},
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

func A2S(v any) string {
	switch e := v.(type) {
	case string:
		return e
	case bool:
		if e {
			return "true"
		}

		return "false"
	case int, uint:
		return fmt.Sprintf("%d", e)
	case float64:
		return fmt.Sprintf("%.2f", e)
	default:
		return fmt.Sprintf("%v", e)
	}
}
