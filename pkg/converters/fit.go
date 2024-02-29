package converters

import (
	"bytes"
	"encoding/xml"
	"strconv"

	"github.com/tkrajina/gpxgo/gpx"
	"github.com/tormoder/fit"
)

func ParseFit(fitFile []byte) (*gpx.GPX, error) {
	// Decode the FIT file data
	f, err := fit.Decode(bytes.NewReader(fitFile))
	if err != nil {
		return nil, err
	}

	m, err := f.Activity()
	if err != nil {
		return nil, err
	}

	gpxFile := &gpx.GPX{
		Creator:  "Garmin Connect",
		Link:     "connect.garmin.com",
		LinkText: "Garmin Connect",
		Time:     &m.Sessions[0].StartTime,
	}

	gpxFile.AppendTrack(&gpx.GPXTrack{
		Name: m.Sessions[0].SportProfileName,
		Type: m.Sessions[0].Sport.String(),
	})

	for _, r := range m.Records {
		p := &gpx.GPXPoint{
			Point: gpx.Point{
				Latitude:  r.PositionLat.Degrees(),
				Longitude: r.PositionLong.Degrees(),
				Elevation: *gpx.NewNullableFloat64(r.GetEnhancedAltitudeScaled()),
			},
			Timestamp: r.Timestamp,
			Extensions: gpx.Extension{
				Nodes: []gpx.ExtensionNode{
					{
						XMLName: xml.Name{Local: "ns3:TrackPointExtension"},
						Nodes: []gpx.ExtensionNode{
							{XMLName: xml.Name{Local: "ns3:hr"}, Data: strconv.Itoa(int(r.HeartRate))},
							{XMLName: xml.Name{Local: "ns3:cad"}, Data: strconv.Itoa(int(r.Cadence))},
						},
					},
				},
			},
		}

		gpxFile.AppendPoint(p)
	}

	return gpxFile, nil
}
