package database

import (
	"time"

	"github.com/codingsince1985/geo-golang"
	"github.com/codingsince1985/geo-golang/openstreetmap"
	"github.com/tkrajina/gpxgo/gpx"
	"github.com/westphae/geomag/pkg/egm96"
)

var online = true

func correctAltitude(lat, long, alt float64) float64 {
	loc := egm96.NewLocationGeodetic(lat, long, alt)

	h, err := loc.HeightAboveMSL()
	if err != nil {
		return alt
	}

	return h
}

// parseGPX parses a GPX file, returns GPX.
func parseGPX(gpxBytes []byte) (*gpx.GPX, error) {
	gpxContent, err := gpx.ParseBytes(gpxBytes)
	if err != nil {
		return nil, err
	}

	return gpxContent, nil
}

type MapData struct {
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
	gpxContent.ReduceGpxToSingleTrack()

	if len(gpxContent.Tracks) == 0 {
		return nil
	}

	if len(gpxContent.Tracks[0].Segments) == 0 {
		return nil
	}

	return gpxContent.Tracks[0].Segments[0].Points
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

func distanceBetween(p1 gpx.GPXPoint, p2 gpx.GPXPoint) float64 {
	return gpx.HaversineDistance(p1.Latitude, p1.Longitude, p2.Latitude, p2.Longitude)
}

func createMapData(gpxContent *gpx.GPX) *MapData {
	gpxContent.ReduceGpxToSingleTrack()

	if len(gpxContent.Tracks) == 0 {
		return nil
	}

	if len(gpxContent.Tracks[0].Segments) == 0 {
		return nil
	}

	mapCenter := center(gpxContent)

	totalDistance := gpxContent.Tracks[0].Segments[0].Length3D()
	totalDuration := time.Duration(gpxContent.Tracks[0].Segments[0].Duration()) * time.Second
	pauseDuration := time.Duration(gpxContent.Tracks[0].Segments[0].MovingData().StoppedTime) * time.Second

	updown := gpxContent.Tracks[0].Segments[0].UphillDownhill()

	data := &MapData{
		Name:          gpxName(gpxContent),
		Center:        mapCenter,
		Address:       mapCenter.Address(),
		TotalDistance: totalDistance,
		TotalDuration: totalDuration,
		MaxSpeed:      gpxContent.Tracks[0].Segments[0].MovingData().MaxSpeed,
		PauseDuration: pauseDuration,
		MinElevation:  correctAltitude(mapCenter.Lat, mapCenter.Lng, gpxContent.Tracks[0].Segments[0].ElevationBounds().MinElevation),
		MaxElevation:  correctAltitude(mapCenter.Lat, mapCenter.Lng, gpxContent.Tracks[0].Segments[0].ElevationBounds().MaxElevation),
		TotalUp:       updown.Uphill,
		TotalDown:     updown.Downhill,
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
			Elevation:     correctAltitude(pt.Point.Latitude, pt.Point.Longitude, pt.Elevation.Value()),
		})
	}

	return data
}
