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
	"github.com/muktihari/fit/profile/typedef"
	"github.com/spf13/cast"
	"github.com/tkrajina/gpxgo/gpx"
)

type gpxMap map[string]string

type fitConverter struct {
	isSwimming   bool
	poolLength   float64
	lastLengthID int
	lengths      []*mesgdef.Length
}

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

	fc := newFitConverter(act)

	for _, r := range act.Records {
		p := fc.fitRecToGPXPt(r)

		gpxFile.AppendPoint(p)
	}

	return gpxFile, nil
}

func newFitConverter(act *filedef.Activity) *fitConverter {
	fc := &fitConverter{
		lastLengthID: -1,
		poolLength:   act.Sessions[0].PoolLengthScaled(),
	}

	if fc.poolLength == 0 {
		return fc
	}

	if act.Sessions[0].Sport != typedef.SportSwimming {
		return fc
	}

	for _, l := range act.Lengths {
		if l.Event != typedef.EventLength {
			continue
		}

		if l.LengthType != typedef.LengthTypeActive {
			continue
		}

		fc.lengths = append(fc.lengths, l)
	}

	fc.isSwimming = len(fc.lengths) > 0

	return fc
}

func (fc *fitConverter) fitRecToGPXPt(r *mesgdef.Record) *gpx.GPXPoint {
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
	if fc.isSwimming {
		fc.addLengthInfo(gpxExtensionData, r.Timestamp)
	}

	for key, value := range gpxExtensionData {
		p.Extensions.Nodes = append(p.Extensions.Nodes, gpx.ExtensionNode{
			XMLName: xml.Name{Local: key}, Data: value,
		})
	}

	return p
}

func (fc *fitConverter) addLengthInfo(gpxExtensionData gpxMap, timestamp time.Time) {
	d := 0.0
	id := 0

	for i, lp := range fc.lengths {
		endTime := lp.StartTime.Add(time.Duration(lp.TotalTimerTimeScaled() * float64(time.Second)))
		if endTime.After(timestamp) {
			break
		}

		d += fc.poolLength
		id = i + 1
	}

	if id >= len(fc.lengths) {
		gpxExtensionData["distance"] = cast.ToString(d)
		return
	}

	l := fc.lengths[id]
	gpxExtensionData["strokes"] = cast.ToString(l.TotalStrokes)

	if _, ok := gpxExtensionData["distance"]; ok {
		return
	}

	fraction := timestamp.Sub(l.StartTime).Seconds() / l.TotalTimerTimeScaled()
	gpxExtensionData["distance"] = cast.ToString(d + (fraction * fc.poolLength))
}

func getGPXExtensionData(r *mesgdef.Record) gpxMap {
	gpxExtensionData := gpxMap{}

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

	if r.Power != math.MaxUint16 {
		gpxExtensionData["power"] = cast.ToString(r.Power)
	}

	return gpxExtensionData
}
