package converters

import (
	"encoding/xml"
	"math"

	"github.com/muktihari/fit/profile/mesgdef"
	"github.com/spf13/cast"
	"github.com/tkrajina/gpxgo/gpx"
)

// Extensions that hold session or lap summaries as reported by the recording
// device. They are stored on the GPX root and take precedence over values that
// are calculated from the track points.
const (
	ExtensionTotalDistance  = "total-distance"
	ExtensionTotalAscent    = "total-ascent"
	ExtensionTotalDescent   = "total-descent"
	ExtensionTotalTimerTime = "total-timer-time"
	ExtensionMaxSpeed       = "max-speed"
	ExtensionAverageCadence = "avg-cadence"
	ExtensionMaxCadence     = "max-cadence"
	ExtensionAveragePower   = "avg-power"
	ExtensionMaxPower       = "max-power"
)

// AddGPXExtension appends a numeric value to the GPX root extensions.
func AddGPXExtension(g *gpx.GPX, key string, value float64) {
	g.Extensions.Nodes = append(g.Extensions.Nodes, gpx.ExtensionNode{
		XMLName: xml.Name{Local: key},
		Data:    cast.ToString(value),
	})
}

// addSessionExtensions copies the summary values of a FIT session that the
// parser would otherwise discard. Values that are not recorded are skipped.
func addSessionExtensions(g *gpx.GPX, s *mesgdef.Session) {
	if s == nil {
		return
	}

	if s.TotalDistance != math.MaxUint32 {
		AddGPXExtension(g, ExtensionTotalDistance, s.TotalDistanceScaled())
	}

	if s.TotalAscent != math.MaxUint16 {
		AddGPXExtension(g, ExtensionTotalAscent, float64(s.TotalAscent))
	}

	if s.TotalDescent != math.MaxUint16 {
		AddGPXExtension(g, ExtensionTotalDescent, float64(s.TotalDescent))
	}

	if s.TotalTimerTime != math.MaxUint32 {
		AddGPXExtension(g, ExtensionTotalTimerTime, s.TotalTimerTimeScaled())
	}

	if s.MaxSpeed != math.MaxUint16 {
		AddGPXExtension(g, ExtensionMaxSpeed, s.MaxSpeedScaled())
	}

	if s.AvgCadence != math.MaxUint8 {
		AddGPXExtension(g, ExtensionAverageCadence, float64(s.AvgCadence))
	}

	if s.MaxCadence != math.MaxUint8 {
		AddGPXExtension(g, ExtensionMaxCadence, float64(s.MaxCadence))
	}

	if s.AvgPower != math.MaxUint16 {
		AddGPXExtension(g, ExtensionAveragePower, float64(s.AvgPower))
	}

	if s.MaxPower != math.MaxUint16 {
		AddGPXExtension(g, ExtensionMaxPower, float64(s.MaxPower))
	}
}
