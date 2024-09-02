package converters

import (
	"bytes"
	"encoding/xml"
	"errors"
	"math"
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

	gpxFile := &gpx.GPX{
		Name:    f.FileId.TimeCreated.String(),
		Time:    &f.FileId.TimeCreated,
		Creator: f.FileId.Manufacturer.String(),
	}

	m, err := f.Activity()
	if err != nil {
		return nil, err
	}

	if len(m.Sessions) == 0 {
		return nil, errors.New("no sessions found")
	}

	gpxFile.AppendTrack(&gpx.GPXTrack{
		Name: m.Sessions[0].SportProfileName,
		Type: m.Sessions[0].Sport.String(),
	})

	for _, r := range m.Records {
		if r.PositionLat.Invalid() ||
			r.PositionLong.Invalid() {
			continue
		}

		p := &gpx.GPXPoint{
			Timestamp: r.Timestamp,
			Point: gpx.Point{
				Latitude:  r.PositionLat.Degrees(),
				Longitude: r.PositionLong.Degrees(),
			},
		}

		if a := r.GetEnhancedAltitudeScaled(); !math.IsNaN(a) {
			p.Elevation = *gpx.NewNullableFloat64(a)
		}

		if r.HeartRate != 0xFF {
			p.Extensions.Nodes = append(p.Extensions.Nodes, gpx.ExtensionNode{
				XMLName: xml.Name{Local: "heart-rate"}, Data: strconv.FormatUint(uint64(r.HeartRate), 10),
			})
		}

		if r.Cadence != 0xFF {
			p.Extensions.Nodes = append(p.Extensions.Nodes, gpx.ExtensionNode{
				XMLName: xml.Name{Local: "cadence"}, Data: strconv.FormatUint(uint64(r.Cadence), 10),
			})
		}

		gpxFile.AppendPoint(p)
	}

	return gpxFile, nil
}
