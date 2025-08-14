package helpers

import (
	"context"
	"fmt"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	"github.com/jovandeginste/workout-tracker/v2/pkg/converters"
	"github.com/jovandeginste/workout-tracker/v2/pkg/database"
	"github.com/jovandeginste/workout-tracker/v2/pkg/templatehelpers"
	"github.com/microcosm-cc/bluemonday"
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

func FilterOptions() []TranslatedKey {
	return StatisticSinceOptions()
}

type TranslatedKey struct {
	Key         string
	Translation string
}

func FindTranslationForKey(tks []TranslatedKey, key string) string {
	for _, tk := range tks {
		if tk.Key == key {
			return tk.Translation
		}
	}

	return ""
}

func OrderDirOptions() []TranslatedKey {
	return []TranslatedKey{
		{"asc", "translation.ascending"},
		{"desc", "translation.descending"},
	}
}

func OrderByOptions() []TranslatedKey {
	return []TranslatedKey{
		{"date", "translation.Date"},

		{"total_distance", "translation.Total_distance"},
		{"total_duration", "translation.Total_duration"},
		{"total_weight", "translation.Total_weight"},
		{"total_repetitions", "translation.Total_repetitions"},
		{"total_up", "translation.Total_up"},
		{"total_down", "translation.Total_down"},

		{"average_speed_no_pause", "translation.Average_speed"},
		{"max_speed", "translation.Max_speed"},
	}
}

func PreferredUnitsToJSON(units *database.UserPreferredUnits) map[string]string {
	return map[string]string{
		"distance":    units.Distance(),
		"speed":       units.Speed(),
		"elevation":   units.Elevation(),
		"heartRate":   units.HeartRate(),
		"cadence":     units.Cadence(),
		"temperature": units.Temperature(),
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
	case int, uint64:
		return fmt.Sprintf("%d", e)
	case float64:
		return templatehelpers.RoundFloat64(e)
	default:
		return fmt.Sprintf("%v", e)
	}
}

func SupportedFileTypes() string {
	return strings.Join(converters.SupportedFileTypes, ", ")
}

func MarkdownToHTML(s string) string {
	s = strings.ReplaceAll(s, "\\n", "\n")
	doc := parser.NewWithExtensions(parser.CommonExtensions).Parse([]byte(s))
	renderer := html.NewRenderer(html.RendererOptions{Flags: html.CommonFlags})
	safeHTML := bluemonday.UGCPolicy().SanitizeBytes(markdown.Render(doc, renderer))

	return string(safeHTML)
}
