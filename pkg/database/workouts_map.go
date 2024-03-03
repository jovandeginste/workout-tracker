package database

import (
	"time"

	"github.com/codingsince1985/geo-golang"
	"github.com/codingsince1985/geo-golang/openstreetmap"
	"github.com/tkrajina/gpxgo/gpx"
	"github.com/westphae/geomag/pkg/egm96"
)

var online = true

func creatorNeedsCorrection(creator string) bool {
	return creator != "Garmin Connect" && creator != "Apple Watch" && creator != "StravaGPX iPhone" && creator != "StravaGPX"
}

func correctAltitude(creator string, lat, long, alt float64) float64 {
	if !creatorNeedsCorrection(creator) {
		return alt
	}

	loc := egm96.NewLocationGeodetic(lat, long, alt)

	h, err := loc.HeightAboveMSL()
	if err != nil {
		return alt
	}

	return h
}

type MapData struct {
	Creator       string
	Name          string
	Date          string
	Center        MapCenter
	Address       *geo.Address
	TotalDistance float64
	TotalDuration time.Duration
	MaxSpeed      float64
	PauseDuration time.Duration
	MinElevation  float64
	MaxElevation  float64
	TotalUp       float64
	TotalDown     float64
	Points        []MapPoint
}
type MapCenter struct {
	Lat float64
	Lng float64
}
type MapPoint struct {
	Lat           float64
	Lng           float64
	Distance      float64
	TotalDistance float64
	Duration      time.Duration
	TotalDuration time.Duration
	Time          time.Time
	Elevation     float64
}

func (m *MapData) AverageSpeed() float64 {
	return m.TotalDistance / m.TotalDuration.Seconds()
}

func (m *MapData) AverageSpeedNoPause() float64 {
	return m.TotalDistance / (m.TotalDuration - m.PauseDuration).Seconds()
}

func (m *MapPoint) AverageSpeed() float64 {
	return m.Distance / m.Duration.Seconds()
}

// center returns the center point (lat, lng) of gpx points
func center(gpxContent *gpx.GPX) MapCenter {
	lat, lng := 0.0, 0.0
	points := allGPXPoints(gpxContent)

	for _, pt := range points {
		lat += pt.Point.Latitude
		lng += pt.Point.Longitude
	}

	size := float64(len(points))

	return MapCenter{
		Lat: lat / size,
		Lng: lng / size,
	}
}

func (m *MapCenter) Address() *geo.Address {
	if !online {
		return nil
	}

	geocoder := openstreetmap.Geocoder()

	address, err := geocoder.ReverseGeocode(m.Lat, m.Lng)
	if err != nil {
		return nil
	}

	return address
}

// allGPXPoints returns the first track segment's points
func allGPXPoints(gpxContent *gpx.GPX) []gpx.GPXPoint {
	var points []gpx.GPXPoint

	for _, track := range gpxContent.Tracks {
		for _, segment := range track.Segments {
			points = append(points, segment.Points...)
		}
	}

	return points
}

func gpxName(gpxContent *gpx.GPX) string {
	if gpxContent.Name != "" {
		return gpxContent.Name
	}

	if len(gpxContent.Tracks) == 0 {
		return "(no name)"
	}

	return gpxContent.Tracks[0].Name
}

// Determines the date to use for the workout
func gpxDate(gpxContent *gpx.GPX) *time.Time {
	// If a date is specified in the metadata, use that (not all apps have this, notably Workoutdoors doesn't)
	if gpxContent.Time != nil {
		return gpxContent.Time
	}

	// Otherwise use the first track's first segment's timestamp
	if len(gpxContent.Tracks) > 0 {
		if len(gpxContent.Tracks[0].Segments) > 0 {
			if len(gpxContent.Tracks[0].Segments[0].Points) > 0 {
				return &gpxContent.Tracks[0].Segments[0].Points[0].Timestamp
			}
		}
	}

	// This is not good as the database requires date to be set, but we don't know any suitable date...
	return nil
}

func distanceBetween(p1 gpx.GPXPoint, p2 gpx.GPXPoint) float64 {
	return gpx.HaversineDistance(p1.Latitude, p1.Longitude, p2.Latitude, p2.Longitude)
}

func createMapData(gpxContent *gpx.GPX) *MapData {
	if len(gpxContent.Tracks) == 0 {
		return nil
	}

	var (
		totalDistance, maxElevation, minElevation, uphill, downhill, maxSpeed float64
		totalDuration, pauseDuration                                          time.Duration
	)

	for _, track := range gpxContent.Tracks {
		for _, segment := range track.Segments {
			if len(segment.Points) > 0 {
				totalDistance += segment.Length3D()
				totalDuration += time.Duration(segment.Duration()) * time.Second
				pauseDuration += (time.Duration(segment.MovingData().StoppedTime)) * time.Second
				minElevation = min(minElevation, segment.ElevationBounds().MinElevation)
				maxElevation = max(maxElevation, segment.ElevationBounds().MaxElevation)
				uphill += segment.UphillDownhill().Uphill
				downhill += segment.UphillDownhill().Downhill
				maxSpeed = max(maxSpeed, segment.MovingData().MaxSpeed)
			}
		}
	}

	// Now reduce the whole GPX to a single track to calculate the center
	gpxContent.ReduceGpxToSingleTrack()
	mapCenter := center(gpxContent)

	data := &MapData{
		Creator:       gpxContent.Creator,
		Name:          gpxName(gpxContent),
		Center:        mapCenter,
		Address:       mapCenter.Address(),
		TotalDistance: totalDistance,
		TotalDuration: totalDuration,
		MaxSpeed:      maxSpeed,
		PauseDuration: pauseDuration,
		MinElevation:  correctAltitude(gpxContent.Creator, mapCenter.Lat, mapCenter.Lng, minElevation),
		MaxElevation:  correctAltitude(gpxContent.Creator, mapCenter.Lat, mapCenter.Lng, maxElevation),
		TotalUp:       uphill,
		TotalDown:     downhill,
	}

	return data
}

func gpxAsMapData(gpxContent *gpx.GPX) *MapData {
	data := createMapData(gpxContent)

	points := allGPXPoints(gpxContent)
	if len(points) == 0 {
		return data
	}

	totalDist := 0.0
	totalTime := 0.0
	prevPoint := points[0]

	for i, pt := range points {
		dist := 0.0
		t := 0.0

		if i > 0 {
			dist = distanceBetween(prevPoint, pt)
			t = pt.TimeDiff(&prevPoint)

			prevPoint = pt
		}

		totalDist += dist
		totalTime += t

		data.Points = append(data.Points, MapPoint{
			Lat:           pt.Point.Latitude,
			Lng:           pt.Point.Longitude,
			Time:          pt.Timestamp,
			Distance:      dist,
			TotalDistance: totalDist,
			Duration:      time.Duration(t) * time.Second,
			TotalDuration: time.Duration(totalTime) * time.Second,
			Elevation:     correctAltitude(gpxContent.Creator, pt.Point.Latitude, pt.Point.Longitude, pt.Elevation.Value()),
		})
	}

	return data
}
