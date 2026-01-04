package helpers

// We need every icon fully qualified here. Otherwise Tailwind will not pick
// it up and not add it to the generated CSS file.
var iconMap = map[string]string{
	// Category Icons
	"source":      "icon-[fa6-solid--bookmark]",
	"distance":    "icon-[fa6-solid--road]",
	"speed":       "icon-[fa6-solid--gauge]",
	"max-speed":   "icon-[fa6-solid--gauge-high]",
	"tempo":       "icon-[fa6-solid--stopwatch]",
	"duration":    "icon-[fa6-regular--clock]",
	"elevation":   "icon-[fa6-solid--mountain]",
	"gain":        "icon-[mdi--triangle-wave]",
	"slope":       "icon-[mdi--angle-acute]",
	"slope_cat":   "icon-[mdi--mountain-outline]",
	"location":    "icon-[fa6-solid--map-location-dot]",
	"repetitions": "icon-[fa6-solid--calculator]",
	"weight":      "icon-[fa6-solid--weight-hanging]",
	"heart-rate":  "icon-[fa6-solid--heart-pulse]",
	"cadence":     "icon-[fa6-solid--stopwatch]",
	"temperature": "icon-[fa6-solid--temperature-high]",
	"heading":     "icon-[fa6-solid--compass]",
	"accuracy":    "icon-[fa6-solid--crosshairs]",
	"date":        "icon-[fa6-regular--calendar]",
	"name":        "icon-[fa6-solid--tag]",
	"timezone":    "icon-[material-symbols--map]",
	"pause":       "icon-[fa6-regular--hourglass]",
	"calories":    "icon-[fa6-solid--fire]",
	"steps":       "icon-[ion--footsteps]",
	"scale":       "icon-[ion--scale]",
	"height":      "icon-[mdi--human-male-height]",

	// Misc Icons
	"welcome":        "icon-[fa6-solid--chevron-right]",
	"circular":       "icon-[fa6-solid--circle-notch]",
	"bidirectional":  "icon-[fa6-solid--arrow-right-arrow-left]",
	"units":          "icon-[fa6-solid--ruler]",
	"file":           "icon-[fa6-solid--file]",
	"best":           "icon-[fa6-solid--arrow-up-long]",
	"worst":          "icon-[fa6-solid--arrow-down-long]",
	"up":             "icon-[fa6-solid--chevron-up]",
	"down":           "icon-[fa6-solid--chevron-down]",
	"metrics":        "icon-[fa6-regular--rectangle-list]",
	"translate":      "icon-[fa6-solid--language]",
	"expand":         "icon-[fa6-solid--arrows-left-right]",
	"share":          "icon-[fa6-solid--share-from-square]",
	"generate-share": "icon-[fa6-solid--retweet]",

	// Sport Icons
	"cycling":        "icon-[material-symbols--pedal-bike-outline]",
	"e-cycling":      "icon-[material-symbols--electric-bike-outline]",
	"running":        "icon-[fa6-solid--person-running]",
	"walking":        "icon-[fa6-solid--person-walking]",
	"swimming":       "icon-[fa6-solid--person-swimming]",
	"inline-skating": "icon-[map--inline-skating]",
	"skiing":         "icon-[fa6-solid--person-skiing]",
	"snowboarding":   "icon-[fa6-solid--person-snowboarding]",
	"golfing":        "icon-[fa6-solid--golf-ball-tee]",
	"kayaking":       "icon-[fa6-solid--sailboat]",
	"hiking":         "icon-[fa6-solid--person-hiking]",
	"horse-riding":   "icon-[mdi--horse-human]",
	"push-ups":       "icon-[hugeicons--push-up-bar]",
	"weight-lifting": "icon-[fa6-solid--dumbbell]",
	"rowing":         "icon-[material-symbols--rowing]",
	"other":          "icon-[fa6-solid--question]",

	// Page Icons
	"dashboard":         "icon-[fa6-solid--chart-line]",
	"statistics":        "icon-[fa6-solid--chart-simple]",
	"admin":             "icon-[fa6-solid--gear]",
	"actions":           "icon-[fa6-solid--gear]",
	"user-profile":      "icon-[fa-solid--user-circle]",
	"user-add":          "icon-[fa6-solid--user-plus]",
	"workout":           "icon-[fa6-solid--dumbbell]",
	"equipment":         "icon-[fa6-solid--bicycle]",
	"route-segment":     "icon-[fa6-solid--route]",
	"add":               "icon-[fa6-solid--circle-plus]",
	"workout-add":       "icon-[fa6-solid--circle-plus]",
	"equipment-add":     "icon-[fa6-solid--circle-plus]",
	"route-segment-add": "icon-[fa6-solid--circle-plus]",
	"heatmap":           "icon-[fa6-solid--fire]",
	"changelog":         "icon-[mdi--clipboard-check]",

	// Utility Icons
	"close":         "icon-[fa6-solid--xmark]",
	"edit":          "icon-[fa6-solid--pen-to-square]",
	"auto-update":   "icon-[fa6-solid--arrows-rotate]",
	"refresh":       "icon-[fa6-solid--arrows-rotate]",
	"delete":        "icon-[fa6-solid--trash]",
	"note":          "icon-[fa6-solid--quote-left]",
	"breakdown":     "icon-[fa6-solid--table-cells]",
	"users":         "icon-[fa6-solid--users]",
	"user-signin":   "icon-[fa6-solid--right-to-bracket]",
	"user-signout":  "icon-[fa6-solid--right-from-bracket]",
	"user-register": "icon-[fa6-solid--user-plus]",
	"user":          "icon-[fa6-solid--user]",
	"show":          "icon-[fa6-solid--eye]",
	"hide":          "icon-[fa6-solid--eye-slash]",
	"copy":          "icon-[fa6-solid--clipboard]",
	"download":      "icon-[fa6-solid--download]",
	"attention":     "icon-[fa6-solid--circle-exclamation]",
	"check":         "icon-[fa6-solid--square-check]",
	"totals":        "icon-[fa6-solid--calculator]",
	"missing":       "icon-[material-symbols--block]",
	"locked":        "icon-[fa6-solid--user-lock]",
	"unlocked":      "icon-[fa6-solid--lock-open]",

	// Brand Icons
	"github": "icon-[octicon--mark-github-16]",
}

func iconFor(what string) string {
	if icon, exists := iconMap[what]; exists {
		return `<span class="icon-decoration ` + icon + `"></span>`
	}

	return `<span class="icon-decoration icon-[fa6-solid--question]"></span>`
}
