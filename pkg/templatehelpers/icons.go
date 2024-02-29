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
	case "file":
		return iconDefaults + " icon-solid icon-file"
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
	case "date":
		return iconDefaults + " icon-regular icon-calendar"
	case "pause":
		return iconDefaults + " icon-regular icon-hourglass"
	case "up":
		return iconDefaults + " icon-solid icon-chevron-up"
	case "down":
		return iconDefaults + " icon-solid icon-chevron-down"
	default:
		return ""
	}
}

func miscIcon(what string) string {
	switch what {
	case "best":
		return iconDefaults + " icon-solid icon-arrow-up-long"
	case "worst":
		return iconDefaults + " icon-solid icon-arrow-down-long"
	default:
		return ""
	}
}

func sportIcon(what string) string {
	switch what {
	case "running":
		return iconDefaults + " icon-solid icon-person-running"
	case "cycling":
		return iconDefaults + " icon-solid icon-person-biking"
	case "walking":
		return iconDefaults + " icon-solid icon-person-walking"
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
	case "admin":
		return iconDefaults + " icon-solid icon-gear"
	case "user-profile":
		return iconDefaults + " icon-solid icon-user-circle"
	case "user-add":
		return iconDefaults + " icon-solid icon-user-plus"
	case "workout":
		return iconDefaults + " icon-solid icon-dumbbell"
	case "workout-add":
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
	default:
		return ""
	}
}
