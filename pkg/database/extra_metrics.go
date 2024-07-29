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
	for i := range extension.Nodes {
		key, value := getGPXExtensionKeyValue(&extension.Nodes[i])
		if key == "" {
			continue
		}

		em.Set(key, value)
	}
}

func getGPXExtensionKeyValue(n *gpx.ExtensionNode) (string, float64) {
	name := n.XMLName.Local

	if data, err := strconv.ParseFloat(n.Data, 64); err == nil {
		return name, data
	}

	return "", 0
}
