package database

import (
	"strings"

	"github.com/spf13/cast"
	"github.com/tkrajina/gpxgo/gpx"
)

type ExtraMetrics map[string]float64

func (em ExtraMetrics) Set(key string, value float64) {
	em[key] = value
}

// SetIfMeaningful stores a metric value, unless zero is not a real reading for
// the metric. Devices often report zero for metrics they did not measure; a
// heart rate of zero, for example, is not a heart rate.
func (em ExtraMetrics) SetIfMeaningful(key string, value float64) {
	if extraMetricValueIsMeaningful(key, value) {
		em.Set(key, value)
	}
}

func (em ExtraMetrics) Get(key string) float64 {
	return em[key]
}

func (em ExtraMetrics) ParseGPXExtensions(extension gpx.Extension) {
	for _, n := range extension.Nodes {
		if key, value := getGPXExtensionKeyValue(&n); key != "" {
			em.SetIfMeaningful(key, value)
		}

		for _, subN := range n.Nodes {
			if key, value := getGPXExtensionKeyValue(&subN); key != "" {
				em.SetIfMeaningful(key, value)
			}
		}
	}
}

func extraMetricValueIsMeaningful(key string, value float64) bool {
	if key == "heart-rate" || strings.Contains(key, "accuracy") {
		return value > 0
	}

	return true
}

func getGPXExtensionKeyValue(n *gpx.ExtensionNode) (string, float64) {
	name := standardExtensionName(n.XMLName.Local)

	// An empty extension node (a container without a value) is not a metric.
	if strings.TrimSpace(n.Data) == "" {
		return "", 0
	}

	if data, err := cast.ToFloat64E(n.Data); err == nil {
		return name, data
	}

	return "", 0
}

func standardExtensionName(name string) string {
	name = strings.ToLower(name)
	name = strings.TrimPrefix(name, "gpxdata:")
	name = strings.TrimPrefix(name, "ns3:")

	switch name {
	case "course":
		return "heading"
	case "hacc": // horizontal accuracy estimate [mm]
		return "horizontal-accuracy"
	case "vacc": // vertical accuracy estimate [mm]
		return "vertical-accuracy"
	case "hr", "heartrate":
		return "heart-rate"
	case "cad":
		return "cadence"
	case "atemp", "temp":
		return "temperature"
	default:
		return name
	}
}
