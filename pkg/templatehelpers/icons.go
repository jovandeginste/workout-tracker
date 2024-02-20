package templatehelpers

import "html/template"

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

	return template.HTML("icon-solid icon-baseline icon-space-sm icon-before icon-question")
}

func categoryIcon(what string) string {
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
	case "pause":
		return "icon-regular icon-baseline icon-space-sm icon-before icon-hourglass"
	case "up":
		return "icon-solid icon-baseline icon-space-sm icon-before icon-chevron-up"
	case "down":
		return "icon-solid icon-baseline icon-space-sm icon-before icon-chevron-down"
	default:
		return ""
	}
}

func miscIcon(what string) string {
	switch what {
	case "best":
		return "icon-solid icon-baseline icon-space-sm icon-before icon-arrow-up-long"
	case "worst":
		return "icon-solid icon-baseline icon-space-sm icon-before icon-arrow-down-long"
	default:
		return ""
	}
}

func sportIcon(what string) string {
	switch what {
	case "running":
		return "icon-solid icon-baseline icon-space-sm icon-before icon-person-running"
	default:
		return ""
	}
}

func pageIcon(what string) string {
	switch what {
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
	default:
		return ""
	}
}

func utilityIcon(what string) string {
	switch what {
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
	case "user":
		return "icon-solid icon-baseline icon-space-sm icon-before icon-user"
	default:
		return ""
	}
}
