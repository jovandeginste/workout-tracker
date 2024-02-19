package templatehelpers

import (
	"fmt"
	"html/template"
	"time"

	"github.com/dustin/go-humanize"
)

func LocalDate(t time.Time) string {
	return t.Local().Format("2006-01-02 15:04") //nolint:gosmopolitan
}

func HumanDistance(d float64) string {
	value, prefix := humanize.ComputeSI(d)

	return fmt.Sprintf("%.2f %sm", value, prefix)
}

func HumanSpeed(mps float64) string {
	mph := mps * 3600
	value, prefix := humanize.ComputeSI(mph)

	return fmt.Sprintf("%.2f %sm/h", value, prefix)
}

func HumanTempo(mps float64) string {
	mpk := 1000000 / (mps * 60)
	value, prefix := humanize.ComputeSI(mpk)

	return fmt.Sprintf("%.2f min/%sm", value, prefix)
}

func IconFor(what string) string {
	switch what {
	case "distance":
		return "icon-solid icon-baseline icon-space-sm icon-before icon-road"
	case "speed":
		return "icon-solid icon-baseline icon-space-sm icon-before icon-gauge"
	case "max-speed":
		return "icon-solid icon-baseline icon-space-sm icon-before icon-gauge-high"
	case "tempo":
		return "icon-solid icon-baseline icon-space-sm icon-before icon-stopwatch"
	case "duration":
		return "icon-regular icon-baseline icon-space-sm icon-before icon-clock"
	case "elevation":
		return "icon-solid icon-baseline icon-space-sm icon-before icon-mountain"
	case "location":
		return "icon-regular icon-baseline icon-space-sm icon-before icon-map-location-dot"
	case "date":
		return "icon-regular icon-baseline icon-space-sm icon-before icon-calendar"
	case "pauze":
		return "icon-regular icon-baseline icon-space-sm icon-before icon-hourglass"
	case "up":
		return "icon-solid icon-baseline icon-space-sm icon-before icon-chevron-up"
	case "down":
		return "icon-solid icon-baseline icon-space-sm icon-before icon-chevron-down"

	case "best":
		return "icon-solid icon-baseline icon-space-sm icon-before icon-arrow-up-long"
	case "worst":
		return "icon-solid icon-baseline icon-space-sm icon-before icon-arrow-down-long"

	case "running":
		return "icon-solid icon-baseline icon-space-sm icon-before icon-person-running"

	case "dashboard":
		return "icon-solid icon-baseline icon-space-sm icon-before icon-chart-line"
	case "admin":
		return "icon-solid icon-baseline icon-space-sm icon-before icon-gear"
	case "user-profile":
		return "icon-solid icon-baseline icon-space-sm icon-before icon-user-circle"
	case "user-add":
		return "icon-solid icon-baseline icon-space-sm icon-before icon-user-plus"
	case "workout":
		return "icon-solid icon-baseline icon-space-sm icon-before icon-dumbbell"
	case "workout-add":
		return "icon-solid icon-baseline icon-space-sm icon-before icon-circle-plus"

	case "close":
		return "icon-solid icon-baseline icon-space-sm icon-before icon-xmark"
	case "edit":
		return "icon-solid icon-baseline icon-space-sm icon-before icon-pen-to-square"
	case "auto-update", "refresh":
		return "icon-solid icon-baseline icon-space-sm icon-before icon-arrows-rotate"
	case "delete":
		return "icon-solid icon-baseline icon-space-sm icon-before icon-trash"
	case "note":
		return "icon-solid icon-baseline icon-space-sm icon-before icon-quote-left"
	case "users":
		return "icon-solid icon-baseline icon-space-sm icon-before icon-users"
	case "user-signin":
		return "icon-solid icon-baseline icon-space-sm icon-before icon-right-to-bracket"
	case "user-signout":
		return "icon-solid icon-baseline icon-space-sm icon-before icon-right-from-bracket"
	case "user-register":
		return "icon-solid icon-baseline icon-space-sm icon-before icon-user-plus"

	}

	return "icon-solid icon-baseline icon-space-sm icon-before icon-question"
}

func BoolToHTML(b bool) template.HTML {
	if b {
		return `<i class="text-green-500 fas fa-check"></i>`
	}

	return `<i class="text-rose-500 fas fa-times"></i>`
}

func BoolToCheckbox(b bool) template.HTML {
	if b {
		return "checked"
	}

	return ""
}

func BuildDecoratedAttribute(icon, name string, value interface{}) interface{} {
	return struct {
		Icon  string
		Name  string
		Value interface{}
	}{
		Icon:  icon,
		Name:  name,
		Value: value,
	}
}
