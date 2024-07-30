package database

import (
	"strconv"

	"github.com/tkrajina/gpxgo/gpx"
)

type ExtraMetrics map[string]float64

func (em ExtraMetrics) Set(key string, value float64) {
	em[key] = value
}

func (em ExtraMetrics) Get(key string) float64 {
	return em[key]
}

func (em ExtraMetrics) ParseGPXExtensions(extension gpx.Extension) {
	for _, n := range extension.Nodes {
		if key, value := getGPXExtensionKeyValue(&n); key != "" {
			em.Set(key, value)
		}

		for _, subN := range n.Nodes {
			if key, value := getGPXExtensionKeyValue(&subN); key != "" {
				em.Set(key, value)
			}
		}
	}
}

func getGPXExtensionKeyValue(n *gpx.ExtensionNode) (string, float64) {
	name := standardExtensionName(n.XMLName.Local)

	if data, err := strconv.ParseFloat(n.Data, 64); err == nil {
		return name, data
	}

	return "", 0
}

func standardExtensionName(name string) string {
	switch name {
	case "course":
		return "heading"
	case "hAcc": // horizontal accuracy estimate [mm]
		return "horizontal-accuracy"
	case "vAcc": // vertical accuracy estimate [mm]
		return "vertical-accuracy"
	case "ns3:hr", "hr":
		return "heart-rate"
	case "ns3:cad", "cad":
		return "cadence"
	default:
		return name
	}
}
