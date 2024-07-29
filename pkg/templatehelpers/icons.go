package templatehelpers

import "html/template"

const iconDefaults = "icon-baseline icon-space-sm icon-before"

func IconFor(what string) template.HTML {
	iconFunctions := []func(string) string{
		categoryIcon,
		sportIcon,
		pageIcon,
		utilityIcon,
		miscIcon,
	}

	for _, f := range iconFunctions {
		if icon := f(what); icon != "" {
			return template.HTML(icon) //nolint:gosec
		}
	}

	return template.HTML(iconDefaults + " icon-solid icon-question") //nolint:gosec
}

func categoryIcon(what string) string {
	switch what {
	case "source":
		return iconDefaults + " icon-solid icon-bookmark"
	case "distance":
		return iconDefaults + " icon-solid icon-road"
	case "speed":
		return iconDefaults + " icon-solid icon-gauge"
	case "max-speed":
		return iconDefaults + " icon-solid icon-gauge-high"
	case "tempo":
		return iconDefaults + " icon-solid icon-stopwatch"
	case "duration":
		return iconDefaults + " icon-regular icon-clock"
	case "elevation":
		return iconDefaults + " icon-solid icon-mountain"
	case "location":
		return iconDefaults + " icon-solid icon-map-location-dot"
	case "repetitions":
		return iconDefaults + " icon-solid icon-calculator"
	case "weight":
		return iconDefaults + " icon-solid icon-weight-hanging"
	case "heartrate":
		return iconDefaults + " icon-solid icon-heart-pulse"
	case "cadence":
		return iconDefaults + " icon-solid icon-stopwatch"
	case "date":
		return iconDefaults + " icon-regular icon-calendar"
	case "pause":
		return iconDefaults + " icon-regular icon-hourglass"
	default:
		return ""
	}
}

func miscIcon(what string) string {
	switch what {
	case "units":
		return iconDefaults + " icon-solid icon-ruler"
	case "file":
		return iconDefaults + " icon-solid icon-file"
	case "best":
		return iconDefaults + " icon-solid icon-arrow-up-long"
	case "worst":
		return iconDefaults + " icon-solid icon-arrow-down-long"
	case "up":
		return iconDefaults + " icon-solid icon-chevron-up"
	case "down":
		return iconDefaults + " icon-solid icon-chevron-down"
	case "metrics":
		return iconDefaults + " icon-regular icon-rectangle-list"
	default:
		return ""
	}
}

func sportIcon(what string) string {
	// We need every icon fully qualified here. Otherwise Tailwind will not pick
	// it up and not add it to the generated CSS file.
	switch what {
	case "cycling":
		return iconDefaults + " icon-solid icon-person-biking"
	case "running":
		return iconDefaults + " icon-solid icon-person-running"
	case "walking":
		return iconDefaults + " icon-solid icon-person-walking"
	case "swimming":
		return iconDefaults + " icon-solid icon-person-swimming"
	case "skiing":
		return iconDefaults + " icon-solid icon-person-skiing"
	case "snowboarding":
		return iconDefaults + " icon-solid icon-person-snowboarding"
	case "golfing":
		return iconDefaults + " icon-solid icon-golf-ball-tee"
	case "kayaking":
		return iconDefaults + " icon-solid icon-sailboat"
	case "hiking":
		return iconDefaults + " icon-solid icon-person-hiking"
	case "push-ups":
		return iconDefaults + " icon-solid icon-dumbbell"
	case "weight lifting":
		return iconDefaults + " icon-solid icon-dumbbell"
	default:
		return ""
	}
}

func pageIcon(what string) string {
	switch what {
	case "dashboard":
		return iconDefaults + " icon-solid icon-chart-line"
	case "statistics":
		return iconDefaults + " icon-solid icon-chart-simple"
	case "admin", "actions":
		return iconDefaults + " icon-solid icon-gear"
	case "user-profile":
		return iconDefaults + " icon-solid icon-user-circle"
	case "user-add":
		return iconDefaults + " icon-solid icon-user-plus"
	case "workout":
		return iconDefaults + " icon-solid icon-dumbbell"
	case "equipment":
		return iconDefaults + " icon-solid icon-bicycle"
	case "add", "workout-add", "equipment-add":
		return iconDefaults + " icon-solid icon-circle-plus"
	default:
		return ""
	}
}

func utilityIcon(what string) string {
	switch what {
	case "close":
		return iconDefaults + " icon-solid icon-xmark"
	case "edit":
		return iconDefaults + " icon-solid icon-pen-to-square"
	case "auto-update", "refresh":
		return iconDefaults + " icon-solid icon-arrows-rotate"
	case "delete":
		return iconDefaults + " icon-solid icon-trash"
	case "note":
		return iconDefaults + " icon-solid icon-quote-left"
	case "users":
		return iconDefaults + " icon-solid icon-users"
	case "user-signin":
		return iconDefaults + " icon-solid icon-right-to-bracket"
	case "user-signout":
		return iconDefaults + " icon-solid icon-right-from-bracket"
	case "user-register":
		return iconDefaults + " icon-solid icon-user-plus"
	case "user":
		return iconDefaults + " icon-solid icon-user"
	case "show":
		return iconDefaults + " icon-solid icon-eye"
	case "hide":
		return iconDefaults + " icon-solid icon-eye-slash"
	case "copy":
		return iconDefaults + " icon-solid icon-clipboard"
	case "download":
		return iconDefaults + " icon-solid icon-download"
	case "attention":
		return iconDefaults + " icon-solid icon-circle-exclamation"
	case "check":
		return iconDefaults + " icon-solid icon-square-check"
	case "totals":
		return iconDefaults + " icon-solid icon-calculator"
	default:
		return ""
	}
}
