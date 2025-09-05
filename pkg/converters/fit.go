package converters

import (
	"bytes"
	"encoding/xml"
	"errors"
	"math"
	"time"

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

	name := act.Sessions[0].Sport.String() + " - " + act.Activity.LocalTimestamp.Format(time.DateTime)
	gpxFile := &gpx.GPX{
		Name:    name,
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

		gpxExtensionData := map[string]string{}
		if r.Cadence != math.MaxUint8 {
			gpxExtensionData["cadence"] = cast.ToString(r.Cadence)
		}

		if r.HeartRate != math.MaxUint8 {
			gpxExtensionData["heart-rate"] = cast.ToString(r.HeartRate)
		}

		if r.EnhancedSpeed != math.MaxUint32 {
			gpxExtensionData["speed"] = cast.ToString(r.EnhancedSpeedScaled())
		} else if r.Speed != math.MaxUint16 {
			gpxExtensionData["speed"] = cast.ToString(r.SpeedScaled())
		}

		if r.Temperature != math.MaxInt8 {
			gpxExtensionData["temperature"] = cast.ToString(r.Temperature)
		}

		for key, value := range gpxExtensionData {
			p.Extensions.Nodes = append(p.Extensions.Nodes, gpx.ExtensionNode{
				XMLName: xml.Name{Local: key}, Data: value,
			})
		}

		gpxFile.AppendPoint(p)
	}

	return gpxFile, nil
}
