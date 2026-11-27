package database

import (
	"time"

	"github.com/jovandeginste/workout-tracker/v2/pkg/converters"
	"github.com/spf13/cast"
	"github.com/tkrajina/gpxgo/gpx"
)

// extensionValue returns the numeric value of the first matching extension
// node, including one level of nested nodes.
func extensionValue(ext gpx.Extension, key string) (float64, bool) {
	for i := range ext.Nodes {
		node := &ext.Nodes[i]

		if name, value := getGPXExtensionKeyValue(node); name == key {
			return value, true
		}

		for j := range node.Nodes {
			if name, value := getGPXExtensionKeyValue(&node.Nodes[j]); name == key {
				return value, true
			}
		}
	}

	return 0, false
}

// gpxExtensionValue returns a value from the GPX root extensions.
func gpxExtensionValue(g *gpx.GPX, key string) (float64, bool) {
	if g == nil {
		return 0, false
	}

	node, ok := g.Extensions.GetNode(gpx.AnyNamespace, key)
	if !ok {
		return 0, false
	}

	value, err := cast.ToFloat64E(node.Data)
	if err != nil {
		return 0, false
	}

	return value, true
}

// deviceDistance returns the cumulative distance recorded by the device.
func deviceDistance(p *gpx.GPXPoint) (float64, bool) {
	return extensionValue(p.Extensions, "distance")
}

// applyExtensionValue overwrites dst with the value of a GPX extension when
// the extension is present and not negative.
func applyExtensionValue(g *gpx.GPX, key string, dst *float64) {
	if value, ok := gpxExtensionValue(g, key); ok && value >= 0 {
		*dst = value
	}
}

// applyDevicePause derives the pause duration from the device timer time.
func applyDevicePause(g *gpx.GPX, data *MapData) {
	timer, ok := gpxExtensionValue(g, converters.ExtensionTotalTimerTime)
	if !ok || timer <= 0 {
		return
	}

	pause := data.TotalDuration - time.Duration(timer*float64(time.Second))
	pause = max(0, min(pause, data.TotalDuration))

	// Ignore sub-second differences between the timer and the track points.
	if pause < time.Second {
		pause = 0
	}

	data.PauseDuration = pause
}

// applyDeviceValues replaces the values that were derived from the track
// points with the summary values that the recording device reported.
func applyDeviceValues(g *gpx.GPX, data *MapData, ignoreElevation, ignoreDeviceValues bool) {
	if ignoreDeviceValues {
		return
	}

	if distance, ok := gpxExtensionValue(g, converters.ExtensionTotalDistance); ok && distance > 0 {
		data.TotalDistance = distance
		data.TotalDistance2D = distance
	}

	if !ignoreElevation {
		applyExtensionValue(g, converters.ExtensionTotalAscent, &data.TotalUp)
		applyExtensionValue(g, converters.ExtensionTotalDescent, &data.TotalDown)
	}

	applyExtensionValue(g, converters.ExtensionAverageCadence, &data.AverageCadence)
	applyExtensionValue(g, converters.ExtensionMaxCadence, &data.MaxCadence)
	applyExtensionValue(g, converters.ExtensionAveragePower, &data.AveragePower)
	applyExtensionValue(g, converters.ExtensionMaxPower, &data.MaxPower)

	applyDevicePause(g, data)
}

// updateMaxSpeed sets the maximum speed. Speeds recorded by the device take
// precedence; otherwise the speed is derived from the track geometry, because
// GPS glitches and fallback distances produce values that never happened.
func (m *MapData) updateMaxSpeed(g *gpx.GPX, workoutType WorkoutType, ignoreDeviceValues bool) {
	if !ignoreDeviceValues {
		if deviceMaxSpeed, ok := maxDeviceSpeed(g, m.Details.Points); ok {
			m.MaxSpeed = deviceMaxSpeed

			return
		}
	}

	m.MaxSpeed = max(m.MaxSpeed, maxSpeedForPoints(m.Details.Points, workoutType.MaxDeltaMeter()))
}

// maxDeviceSpeed returns the maximum speed that the device recorded, either
// per track point or as a session summary. The boolean is false when the
// workout contains no device speed at all.
func maxDeviceSpeed(g *gpx.GPX, points []MapPoint) (float64, bool) {
	deviceMax, hasDeviceMax := gpxExtensionValue(g, converters.ExtensionMaxSpeed)
	if !hasDeviceMax {
		deviceMax = 0
	}

	pointMax := 0.0

	for _, p := range points {
		pointMax = max(pointMax, p.ExtraMetrics.Get("speed"))
	}

	if pointMax > 0 || hasDeviceMax {
		return max(pointMax, deviceMax), true
	}

	return 0, false
}
