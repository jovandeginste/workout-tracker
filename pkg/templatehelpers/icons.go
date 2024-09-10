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
	iconMap := map[string]string{
		"source":      "icon-solid icon-bookmark",
		"distance":    "icon-solid icon-road",
		"speed":       "icon-solid icon-gauge",
		"max-speed":   "icon-solid icon-gauge-high",
		"tempo":       "icon-solid icon-stopwatch",
		"duration":    "icon-regular icon-clock",
		"elevation":   "icon-solid icon-mountain",
		"location":    "icon-solid icon-map-location-dot",
		"repetitions": "icon-solid icon-calculator",
		"weight":      "icon-solid icon-weight-hanging",
		"heart-rate":  "icon-solid icon-heart-pulse",
		"cadence":     "icon-solid icon-stopwatch",
		"heading":     "icon-solid icon-compass",
		"date":        "icon-regular icon-calendar",
		"pause":       "icon-regular icon-hourglass",
		"calories":    "icon-solid icon-fire",
	}

	if icon, exists := iconMap[what]; exists {
		return iconDefaults + " " + icon
	}
	return ""
}

func miscIcon(what string) string {
	iconMap := map[string]string{
		"circular":      "icon-solid icon-circle-notch",
		"bidirectional": "icon-solid icon-arrow-right-arrow-left",
		"units":         "icon-solid icon-ruler",
		"file":          "icon-solid icon-file",
		"best":          "icon-solid icon-arrow-up-long",
		"worst":         "icon-solid icon-arrow-down-long",
		"up":            "icon-solid icon-chevron-up",
		"down":          "icon-solid icon-chevron-down",
		"metrics":       "icon-regular icon-rectangle-list",
		"translate":     "icon-solid icon-language",
		"expand":        "icon-solid icon-arrows-left-right",
	}

	if icon, exists := iconMap[what]; exists {
		return iconDefaults + " " + icon
	}
	return ""
}

func sportIcon(what string) string {
	iconMap := map[string]string{
		"cycling":        "icon-solid icon-person-biking",
		"running":        "icon-solid icon-person-running",
		"walking":        "icon-solid icon-person-walking",
		"swimming":       "icon-solid icon-person-swimming",
		"skiing":         "icon-solid icon-person-skiing",
		"snowboarding":   "icon-solid icon-person-snowboarding",
		"golfing":        "icon-solid icon-golf-ball-tee",
		"kayaking":       "icon-solid icon-sailboat",
		"hiking":         "icon-solid icon-person-hiking",
		"push-ups":       "icon-solid icon-dumbbell",
		"weight lifting": "icon-solid icon-dumbbell",
	}

	if icon, exists := iconMap[what]; exists {
		return iconDefaults + " " + icon
	}
	return ""
}

func pageIcon(what string) string {
	iconMap := map[string]string{
		"dashboard":         "icon-solid icon-chart-line",
		"statistics":        "icon-solid icon-chart-simple",
		"admin":             "icon-solid icon-gear",
		"actions":           "icon-solid icon-gear",
		"user-profile":      "icon-solid icon-user-circle",
		"user-add":          "icon-solid icon-user-plus",
		"workout":           "icon-solid icon-dumbbell",
		"equipment":         "icon-solid icon-bicycle",
		"route-segment":     "icon-solid icon-route",
		"add":               "icon-solid icon-circle-plus",
		"workout-add":       "icon-solid icon-circle-plus",
		"equipment-add":     "icon-solid icon-circle-plus",
		"route-segment-add": "icon-solid icon-circle-plus",
	}

	if icon, exists := iconMap[what]; exists {
		return iconDefaults + " " + icon
	}
	return ""
}

func utilityIcon(what string) string {
	iconMap := map[string]string{
		"close":         "icon-solid icon-xmark",
		"edit":          "icon-solid icon-pen-to-square",
		"auto-update":   "icon-solid icon-arrows-rotate",
		"refresh":       "icon-solid icon-arrows-rotate",
		"delete":        "icon-solid icon-trash",
		"note":          "icon-solid icon-quote-left",
		"users":         "icon-solid icon-users",
		"user-signin":   "icon-solid icon-right-to-bracket",
		"user-signout":  "icon-solid icon-right-from-bracket",
		"user-register": "icon-solid icon-user-plus",
		"user":          "icon-solid icon-user",
		"show":          "icon-solid icon-eye",
		"hide":          "icon-solid icon-eye-slash",
		"copy":          "icon-solid icon-clipboard",
		"download":      "icon-solid icon-download",
		"attention":     "icon-solid icon-circle-exclamation",
		"check":         "icon-solid icon-square-check",
		"totals":        "icon-solid icon-calculator",
	}

	if icon, exists := iconMap[what]; exists {
		return iconDefaults + " " + icon
	}
	return ""
}
