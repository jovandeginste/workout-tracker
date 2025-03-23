package converters

import (
	"bytes"
	"encoding/xml"
	"errors"
	"math"
	"strconv"

	"github.com/muktihari/fit/decoder"
	"github.com/muktihari/fit/kit/semicircles"
	"github.com/muktihari/fit/profile/filedef"
	"github.com/tkrajina/gpxgo/gpx"
)

func ParseFit(content []byte) (*gpx.GPX, error) {
	// Decode the FIT file data
	dec := decoder.New(bytes.NewReader(content))

	f, err := dec.Decode()
	if err != nil {
		return nil, err
	}

	act := filedef.NewActivity(f.Messages...)
	if len(act.Sessions) == 0 {
		return nil, errors.New("no sessions found")
	}

	gpxFile := &gpx.GPX{
		Name:    act.FileId.TimeCreated.String(),
		Time:    &act.FileId.TimeCreated,
		Creator: act.FileId.Manufacturer.String(),
	}

	gpxFile.AppendTrack(&gpx.GPXTrack{
		Name: act.Sessions[0].SportProfileName,
		Type: act.Sessions[0].Sport.String(),
	})

	for _, r := range act.Records {
		p := &gpx.GPXPoint{
			Timestamp: r.Timestamp,
			Point: gpx.Point{
				Latitude:  semicircles.ToDegrees(r.PositionLat),
				Longitude: semicircles.ToDegrees(r.PositionLong),
			},
		}

		if math.IsNaN(p.Latitude) || math.IsNaN(p.Longitude) {
			continue
		}

		if a := r.EnhancedAltitudeScaled(); !math.IsNaN(a) {
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
