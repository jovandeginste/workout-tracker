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
	"github.com/muktihari/fit/profile/mesgdef"
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

	activityTime := act.Activity.LocalTimestamp
	if activityTime.IsZero() {
		activityTime = act.Sessions[0].StartTime.Local()
	}

	name := act.Sessions[0].Sport.String() + " - " + activityTime.Format(time.DateTime)
	gpxFile := &gpx.GPX{
		Name:    name,
		Time:    &act.FileId.TimeCreated,
		Creator: act.FileId.Manufacturer.String(),
	}

	if act.Sessions[0].TotalCalories != math.MaxUint16 {
		gpxFile.Extensions.Nodes = append(gpxFile.Extensions.Nodes, gpx.ExtensionNode{
			XMLName: xml.Name{Local: "total-calories"}, Data: cast.ToString(act.Sessions[0].TotalCalories),
		})
	}

	gpxFile.AppendTrack(&gpx.GPXTrack{
		Name: act.Sessions[0].SportProfileName,
		Type: act.Sessions[0].Sport.String(),
	})

	for _, r := range act.Records {
		p := &gpx.GPXPoint{
			Timestamp: r.Timestamp,
		}

		lat := semicircles.ToDegrees(r.PositionLat)
		lon := semicircles.ToDegrees(r.PositionLong)

		if !math.IsNaN(lat) && !math.IsNaN(lon) {
			p.Point = gpx.Point{
				Latitude:  lat,
				Longitude: lon,
			}
		}

		if r.EnhancedAltitude != math.MaxUint32 {
			p.Elevation = *gpx.NewNullableFloat64(r.EnhancedAltitudeScaled())
		}

		gpxExtensionData := getGPXExtensionData(r)
		for key, value := range gpxExtensionData {
			p.Extensions.Nodes = append(p.Extensions.Nodes, gpx.ExtensionNode{
				XMLName: xml.Name{Local: key}, Data: value,
			})
		}

		gpxFile.AppendPoint(p)
	}

	return gpxFile, nil
}

func getGPXExtensionData(r *mesgdef.Record) map[string]string {
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

	if r.Distance != math.MaxUint32 {
		gpxExtensionData["distance"] = cast.ToString(r.DistanceScaled())
	}

	if r.Calories != math.MaxUint16 {
		gpxExtensionData["calories"] = cast.ToString(r.Calories)
	}

	return gpxExtensionData
}
