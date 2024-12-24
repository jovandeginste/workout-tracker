package converters

import (
	"encoding/xml"
	"errors"
	"strconv"

	"github.com/galeone/tcx"
	"github.com/tkrajina/gpxgo/gpx"
)

func ParseTCX(tcxFile []byte) (*gpx.GPX, error) {
	var t tcx.TCXDB

	if err := xml.Unmarshal(tcxFile, &t); err != nil {
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

	for _, a := range t.Acts.Act {
		for _, l := range a.Laps {
			for _, p := range l.Trk.Pt {
				gpxP := tcxPtToGPXPt(&p)
				if gpxP == nil {
					continue
				}

				g.AppendPoint(gpxP)
			}
		}
	}

	return g, nil
}

func tcxPtToGPXPt(t *tcx.Trackpoint) *gpx.GPXPoint {
	if t == nil {
		return nil
	}

	if t.Time.IsZero() ||
		(t.Lat == 0 && t.Long == 0) {
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

	if t.HR != 0 {
		p.Extensions.Nodes = append(p.Extensions.Nodes, gpx.ExtensionNode{
			XMLName: xml.Name{Local: "heart-rate"}, Data: strconv.FormatUint(uint64(t.HR), 10),
		})
	}

	if t.Cad != 0 {
		p.Extensions.Nodes = append(p.Extensions.Nodes, gpx.ExtensionNode{
			XMLName: xml.Name{Local: "cadence"}, Data: strconv.FormatUint(uint64(t.Cad), 10),
		})
	}

	return p
}
