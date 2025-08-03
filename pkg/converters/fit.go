package converters

import (
	"bytes"
	"encoding/xml"
	"errors"
	"math"

	"github.com/muktihari/fit/decoder"
	"github.com/muktihari/fit/kit/semicircles"
	"github.com/muktihari/fit/profile/filedef"
	"github.com/spf13/cast"
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

		if r.EnhancedAltitude != math.MaxUint32 {
			p.Elevation = *gpx.NewNullableFloat64(r.EnhancedAltitudeScaled())
		}

		if r.EnhancedSpeed != math.MaxUint32 {
			p.Extensions.Nodes = append(p.Extensions.Nodes, gpx.ExtensionNode{
				XMLName: xml.Name{Local: "enhanced-speed"}, Data: cast.ToString(r.EnhancedSpeedScaled()),
			})
		}

		if r.HeartRate != math.MaxUint8 {
			p.Extensions.Nodes = append(p.Extensions.Nodes, gpx.ExtensionNode{
				XMLName: xml.Name{Local: "heart-rate"}, Data: cast.ToString(r.HeartRate),
			})
		}

		if r.Cadence != math.MaxUint8 {
			p.Extensions.Nodes = append(p.Extensions.Nodes, gpx.ExtensionNode{
				XMLName: xml.Name{Local: "cadence"}, Data: cast.ToString(r.Cadence),
			})
		}

		if r.Temperature != math.MaxInt8 {
			p.Extensions.Nodes = append(p.Extensions.Nodes, gpx.ExtensionNode{
				XMLName: xml.Name{Local: "temperature"}, Data: cast.ToString(r.Temperature),
			})
		}

		gpxFile.AppendPoint(p)
	}

	return gpxFile, nil
}
