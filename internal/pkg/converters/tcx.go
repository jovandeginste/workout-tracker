package converters

import (
	"encoding/xml"
	"errors"

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
				if p.Time.IsZero() ||
					(p.Lat == 0 && p.Long == 0) {
					continue
				}

				gpxP := &gpx.GPXPoint{
					Point: gpx.Point{
						Latitude:  p.Lat,
						Longitude: p.Long,
						Elevation: *gpx.NewNullableFloat64(p.Alt),
					},
					Timestamp: p.Time,
				}

				g.AppendPoint(gpxP)
			}
		}
	}

	return g, nil
}
