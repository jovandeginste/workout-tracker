package database

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tkrajina/gpxgo/gpx"
)

func extensionNode(name, data string) gpx.ExtensionNode {
	return gpx.ExtensionNode{
		XMLName: xml.Name{Local: name},
		Data:    data,
	}
}

func TestExtraMetricsParseGPXExtensionsDropsMeaninglessZeroValues(t *testing.T) {
	extension := gpx.Extension{Nodes: []gpx.ExtensionNode{
		extensionNode("hr", "0"),
		extensionNode("heartrate", "0"),
		extensionNode("hr", "-1"),
		extensionNode("cad", "0"),
		extensionNode("temp", "0"),
		extensionNode("speed", "0"),
		extensionNode("TrackPointExtension", ""),
		extensionNode("hacc", "0"),
		extensionNode("accuracy", "0"),
		extensionNode("accuracy_horizontal", "-1"),
		extensionNode("hr", "142"),
	}}

	em := ExtraMetrics{}
	em.ParseGPXExtensions(extension)

	assert.Equal(t, 142.0, em.Get("heart-rate"))
	assert.Equal(t, 0.0, em.Get("cadence"))
	assert.Equal(t, 0.0, em.Get("temperature"))
	assert.Equal(t, 0.0, em.Get("speed"))
	assert.NotContains(t, em, "trackpointextension")
	assert.NotContains(t, em, "horizontal-accuracy")
	assert.NotContains(t, em, "accuracy")
	assert.NotContains(t, em, "accuracy_horizontal")
}

func TestUpdateExtraMetricsIgnoresAllZeroMetrics(t *testing.T) {
	points := []MapPoint{
		{ExtraMetrics: ExtraMetrics{
			"heart-rate":          0,
			"cadence":             0,
			"elevation":           0,
			"trackpointextension": 0,
		}},
		{ExtraMetrics: ExtraMetrics{"speed": 4.2}},
	}

	m := &MapData{Details: &MapDataDetails{Points: points}}
	m.UpdateExtraMetrics()

	assert.Equal(t, []string{"speed"}, m.ExtraMetrics)
	assert.NotContains(t, m.Details.Points[0].ExtraMetrics, "heart-rate")
}

func TestUpdateExtraMetricsKeepsMeaningfulHeartRate(t *testing.T) {
	points := []MapPoint{
		{ExtraMetrics: ExtraMetrics{"heart-rate": 0}},
		{ExtraMetrics: ExtraMetrics{"heart-rate": 142}},
	}

	m := &MapData{Details: &MapDataDetails{Points: points}}
	m.UpdateExtraMetrics()

	assert.Equal(t, []string{"heart-rate"}, m.ExtraMetrics)
	assert.NotContains(t, m.Details.Points[0].ExtraMetrics, "heart-rate")
	assert.Equal(t, 142.0, m.Details.Points[1].ExtraMetrics.Get("heart-rate"))
}

func TestUpdateExtraMetricsKeepsZeroCadenceWithRealReadings(t *testing.T) {
	points := []MapPoint{
		{ExtraMetrics: ExtraMetrics{"cadence": 0}},
		{ExtraMetrics: ExtraMetrics{"cadence": 80}},
	}

	m := &MapData{Details: &MapDataDetails{Points: points}}
	m.UpdateExtraMetrics()

	assert.Equal(t, []string{"cadence"}, m.ExtraMetrics)
	assert.Equal(t, 0.0, m.Details.Points[0].ExtraMetrics.Get("cadence"))
}
