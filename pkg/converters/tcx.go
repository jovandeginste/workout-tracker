package converters

import (
	"encoding/xml"
	"errors"

	"github.com/galeone/tcx"
	"github.com/spf13/cast"
	"github.com/tkrajina/gpxgo/gpx"
)

type tcxConverter struct {
	totalDistance float64
}

func ParseTCX(content []byte) (*gpx.GPX, error) {
	var t tcx.TCXDB

	if err := xml.Unmarshal(content, &t); err != nil {
		return nil, err
	}

	g := &gpx.GPX{}

	if t.Acts == nil || len(t.Acts.Act) == 0 {
		return nil, errors.New("no activities found")
	}

	g.Name = t.Acts.Act[0].Id.String()
	g.Time = &t.Acts.Act[0].Id

	if t.Auth != nil {
		g.Creator = t.Auth.Name
	} else {
		g.Creator = "TCX importer"
	}

	g.AppendTrack(&gpx.GPXTrack{
		Name: t.Acts.Act[0].Id.String(),
		Type: t.Acts.Act[0].Sport,
	})

	tc := tcxConverter{}

	var (
		totalDistance float64
		totalCalories float64
		maxSpeed      float64
	)

	for _, a := range t.Acts.Act {
		for _, l := range a.Laps {
			totalDistance += l.Dist
			totalCalories += l.Calories
			maxSpeed = max(maxSpeed, l.MaxSpeed)

			if l.Trk == nil {
				continue
			}

			for _, p := range l.Trk.Pt {
				gpxP := tc.tcxPtToGPXPt(&p)
				if gpxP == nil {
					continue
				}

				g.AppendPoint(gpxP)
			}
		}
	}

	if totalDistance > 0 {
		AddGPXExtension(g, ExtensionTotalDistance, totalDistance)
	}

	if totalCalories > 0 {
		AddGPXExtension(g, "total-calories", totalCalories)
	}

	if maxSpeed > 0 {
		AddGPXExtension(g, ExtensionMaxSpeed, maxSpeed)
	}

	return g, nil
}

func (tc *tcxConverter) tcxPtToGPXPt(t *tcx.Trackpoint) *gpx.GPXPoint {
	if t == nil {
		return nil
	}

	if t.Time.IsZero() {
		return nil
	}

	p := &gpx.GPXPoint{
		Point: gpx.Point{
			Latitude:  t.Lat,
			Longitude: t.Long,
			Elevation: *gpx.NewNullableFloat64(t.Alt),
		},
		Timestamp: t.Time,
	}

	setIfNotZero(p, "cadence", t.Cad)
	setIfNotZero(p, "distance", t.Dist)
	setIfNotZero(p, "heart-rate", t.HR)
	setIfNotZero(p, "power", t.Power)
	setIfNotZero(p, "speed", t.Speed)
	setIfNotZero(p, "strokes", t.StrokeCount)

	if t.Dist == 0 && t.SwimDistance > 0 {
		tc.totalDistance += t.SwimDistance
		setIfNotZero(p, "distance", tc.totalDistance)
	}

	return p
}

func setIfNotZero(p *gpx.GPXPoint, key string, value float64) {
	if value == 0 {
		return
	}

	p.Extensions.Nodes = append(p.Extensions.Nodes, gpx.ExtensionNode{
		XMLName: xml.Name{Local: key}, Data: cast.ToString(value),
	})
}
