package helpers

import (
	"cmp"
	"context"
	"fmt"
	"slices"
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

	if rev, err := e.Router().Routes().Reverse(name, params...); err == nil && rev != "" {
		return rev
	}

	return "/invalid/route/#" + name
}

func WorkoutTypes() []database.WorkoutType {
	return database.WorkoutTypes()
}

// WorkoutTypesByCount returns the workout types the user actually has
// workouts for, ordered by count and then by name.
func WorkoutTypeStrings(types []database.WorkoutType) []string {
	result := make([]string, 0, len(types))
	for _, t := range types {
		result = append(result, t.String())
	}

	return result
}

func WorkoutTypesByCount(counts map[database.WorkoutType]int) []database.WorkoutType {
	types := make([]database.WorkoutType, 0, len(counts))

	for _, wt := range database.WorkoutTypes() {
		if counts[wt] > 0 {
			types = append(types, wt)
		}
	}

	slices.SortFunc(types, func(a, b database.WorkoutType) int {
		if c := cmp.Compare(counts[b], counts[a]); c != 0 {
			return c
		}

		return cmp.Compare(a, b)
	})

	return types
}

type Dataset struct {
	Label string
	Data  []any
}

type MapConfig struct {
	CenterLat     float64
	CenterLng     float64
	MinElevation  float64
	MaxElevation  float64
	MaxSpeed      float64
	SpeedName     string
	ElevationName string
	SlopeName     string
	LiftName      string
	ClimbsName    string
	Labels        map[string]string
	StreetsName   string
	AerialName    string
	ShowElevation bool
}

type TrendDataPoint struct {
	Date        string `json:"date"`
	Speed       string `json:"speed"`
	Type        string `json:"type"`
	SegmentID   uint64 `json:"segmentID"`
	SegmentName string `json:"segmentName"`
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
		{"asc", "common.ascending"},
		{"desc", "common.descending"},
	}
}

func OrderByOptions() []TranslatedKey {
	return []TranslatedKey{
		{"date", "common.date"},

		{"total_distance", "workout.total-distance"},
		{"total_duration", "workout.total-duration"},
		{"total_weight", "workout.total-weight"},
		{"total_repetitions", "workout.total-repetitions"},
		{"total_up", "workout.total-up"},
		{"total_down", "workout.total-down"},

		{"average_speed_no_pause", "workout.average-speed"},
		{"max_speed", "workout.max-speed"},
	}
}

func PreferredUnitsToJSON(units *database.UserPreferredUnits) map[string]string {
	return map[string]string{
		"distance":    units.Distance(nil),
		"speed":       units.Speed(nil),
		"elevation":   units.Elevation(),
		"heartRate":   units.HeartRate(),
		"cadence":     units.Cadence(),
		"temperature": units.Temperature(),
		"tempo":       units.Tempo(nil),
		"power":       units.Power(),
		"calories":    units.Calories(),
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

// IsIOS reports whether the request comes from an iOS device.
//
// iOS Safari ignores extensions it does not know, and disables every file when
// none of the listed extensions are known (WebKit bug 279606). So we leave the
// accept filter out there and rely on the server-side validation instead.
func IsIOS(ctx context.Context) bool {
	req := Request(ctx)
	if req == nil {
		return false
	}

	ua := req.UserAgent()

	return strings.Contains(ua, "iPhone") || strings.Contains(ua, "iPad") || strings.Contains(ua, "iPod")
}

func MarkdownToHTML(s string) string {
	s = strings.ReplaceAll(s, "\\n", "\n")
	doc := parser.NewWithExtensions(parser.CommonExtensions).Parse([]byte(s))
	renderer := html.NewRenderer(html.RendererOptions{Flags: html.CommonFlags})
	safeHTML := bluemonday.UGCPolicy().SanitizeBytes(markdown.Render(doc, renderer))

	return string(safeHTML)
}
